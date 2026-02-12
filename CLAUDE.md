# oc-console

A kubectl/oc plugin that opens the OpenShift 4 web console in your default browser.

## Build Commands

```bash
# Build the binary
go build

# Run tests
go test ./...

# Run linter (requires golangci-lint v1.63.4 or later)
golangci-lint run

# Install to $GOPATH/bin
go install
```

## Architecture

The codebase consists of two main Go files following kubectl plugin conventions:

- **cmd/oc-console/oc-console.go** - Entry point and CLI setup using cobra
- **pkg/cmd/console.go** - Core plugin logic implementing the Complete/Validate/Run pattern

### Console URL Discovery

The plugin discovers the OpenShift console URL through:

1. **Primary method**: Queries the `console-public` ConfigMap in the `openshift-config-managed` namespace for the `consoleURL` field
2. **Fallback**: For OpenShift 3.x clusters, attempts to construct the URL from the API server address

Once discovered, the URL is opened using the `pkg/browser` package, which handles cross-platform browser launching.

## Key Dependencies

- **cobra** - CLI framework
- **cli-runtime** - Kubernetes CLI utilities for kubeconfig and client management
- **client-go** - Kubernetes API client
- **pkg/browser** - Cross-platform browser launching

## CI/CD

GitHub Actions workflow (`.github/workflows/go.yml`) runs:
- Go builds and tests on multiple platforms
- golangci-lint checks
- GoReleaser for tagged releases
