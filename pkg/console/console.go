package console

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"k8s.io/kubectl/pkg/util/templates"
)

const (
	openShiftConfigManagedNamespaceName = "openshift-config-managed"
	consolePublicConfigMap              = "console-public"
)

var (
	consoleShort = templates.LongDesc(`
	  Open the OpenShift console in your default browser.`)

	consoleExample = templates.Examples(`
	  # Open the OpenShift console in your default browser
	  %[1]s %[2]s

	  # Display the URL for the OpenShift console
	  %[1]s %[2]s --url`)
)

// ConsoleCmdOptions are options supported by the console command.
type ConsoleCmdOptions struct { //nolint:golint
	configFlags *genericclioptions.ConfigFlags

	ClientConfig *rest.Config
	KubeClient   kubernetes.Interface

	// URL is true if the command should print the URL of the console instead of
	// opening the browser.
	URL bool

	// args is the slice of strings containing any arguments passed
	args []string

	context context.Context

	genericclioptions.IOStreams
}

// NewConsoleCmdOptions provides an instance of ConsoleCmdOptions with default values
func NewConsoleCmdOptions(streams genericclioptions.IOStreams) *ConsoleCmdOptions {
	return &ConsoleCmdOptions{
		configFlags: genericclioptions.NewConfigFlags(false),

		IOStreams: streams,

		context: context.Background(),
	}
}

// NewCmdConsoleConfig provides a cobra command wrapping ConsoleCmdOptions
func NewCmdConsoleConfig(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewConsoleCmdOptions(streams)

	callingBinary := getCallingBinary()

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s console", callingBinary),
		Short:   consoleShort,
		Example: fmt.Sprintf(consoleExample, callingBinary, "console"),
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&o.URL, "url", o.URL, "Print the console URL instead of opening it.")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

func getCallingBinary() string {
	if os.Getenv("KUBECTL_PLUGINS_CALLER") != "" {
		return os.Getenv("KUBECTL_PLUGINS_CALLER")
	}

	return "oc"
}

// Complete sets up the KubeClient
func (o *ConsoleCmdOptions) Complete(args []string) error {
	var err error

	o.args = args

	o.ClientConfig, err = o.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	kubeClient, err := kubernetes.NewForConfig(o.ClientConfig)
	if err != nil {
		return err
	}
	o.KubeClient = kubeClient

	return err
}

// Validate ensures that all required arguments and flag values are provided
func (o *ConsoleCmdOptions) Validate() error {
	if len(o.args) > 0 {
		return fmt.Errorf("no arguments are allowed")
	}

	return nil
}

// getWebConsoleURL retrieves the web console URL from the appropriate configmap
func (o *ConsoleCmdOptions) getWebConsoleURL() (string, error) {
	consolePublicConfig, err := o.KubeClient.CoreV1().
		ConfigMaps(openShiftConfigManagedNamespaceName).
		Get(o.context, consolePublicConfigMap, metav1.GetOptions{})

	// This means the command was run against 3.x server
	if errors.IsNotFound(err) {
		return o.ClientConfig.Host, nil
	}
	if err != nil {
		return "", fmt.Errorf("unable to determine console location: %v", err)
	}

	consoleURL, exists := consolePublicConfig.Data["consoleURL"]
	if !exists {
		return "", fmt.Errorf("unable to determine console location from the cluster")
	}
	return consoleURL, nil
}

// Run grabs the console URL, and either prints it to the terminal or opens it
// in your default web browser
func (o *ConsoleCmdOptions) Run() error {
	var err error

	consoleURL, err := o.getWebConsoleURL()
	if err != nil {
		return err
	}

	if o.URL {
		fmt.Fprintf(o.Out, "%s\n", consoleURL)
		return nil
	}

	err = browser.OpenURL(consoleURL)
	if err != nil {
		return err
	}

	return nil
}
