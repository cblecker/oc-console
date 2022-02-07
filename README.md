# oc-console

A [cli plugin] that allows you to open the OpenShift 4 web console in your
default web browser.

## Installation

```
go install github.com/cblecker/oc-console@master
```

[cli plugin]: https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/

## Use

```
Open the OpenShift console in your default browser.

Usage:
  oc console [flags]

Examples:
  # Open the OpenShift console in your default browser
  oc console

  # Display the URL for the OpenShift console
  oc console --url
```
