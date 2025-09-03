# oc-console

oc-console is a Go CLI application that provides a kubectl/oc plugin for opening the OpenShift 4 web console in your default browser. The plugin connects to an OpenShift cluster and retrieves the console URL from the cluster's configuration.

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

### Prerequisites and Dependencies
- Requires Go 1.24.0 or later (as specified in go.mod)
- Standard Go toolchain (no additional build tools needed)
- golangci-lint for code quality (install separately)

### Bootstrap, Build, and Test
1. **Download dependencies:**
   ```bash
   go mod download
   ```
   Takes ~10 seconds. NEVER CANCEL. Set timeout to 60+ seconds.

2. **Build the project:**
   ```bash
   go build -o oc-console .
   ```
   Takes ~5 seconds for binary build. NEVER CANCEL. Set timeout to 120+ seconds.
   
   For verbose build with all packages:
   ```bash
   go build -v ./...
   ```
   Takes ~55 seconds. NEVER CANCEL. Set timeout to 120+ seconds.

3. **Run tests:**
   ```bash
   go test ./...
   ```
   Takes ~9 seconds. Currently no test files exist but test infrastructure works. NEVER CANCEL. Set timeout to 60+ seconds.

4. **Run linting:**
   Install golangci-lint first:
   ```bash
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
   ```
   Then run:
   ```bash
   golangci-lint run
   ```
   Takes ~20 seconds. NEVER CANCEL. Set timeout to 60+ seconds.

### Run and Test the Application
- **Test basic functionality:**
  ```bash
  ./oc-console --help
  ```
  Should display usage information and available flags.

- **Test URL functionality (requires OpenShift cluster):**
  ```bash
  ./oc-console --url
  ```
  Will fail with connection error if not connected to OpenShift cluster - this is expected behavior.

## Validation

### Build Validation
- ALWAYS run the complete build sequence after making changes:
  1. `go mod download`
  2. `go build -o oc-console .`
  3. `./oc-console --help` (verify help output displays correctly)

### Code Quality Validation
- ALWAYS run `golangci-lint run` before committing changes or the CI (.github/workflows/go.yml) will fail
- The linter runs in CI and is required for all pull requests

### Functional Validation
- The application requires connection to an OpenShift cluster for full functionality
- Test with `./oc-console --help` to verify the binary works correctly
- Test with `./oc-console --url` to verify cluster connectivity (expected to fail in development environment)
- You cannot fully test the browser-opening functionality without an OpenShift cluster

### CI Validation
The GitHub Actions workflow (.github/workflows/go.yml) runs:
1. Build: `go build -v ./...`
2. Test: `go test ./...` 
3. Lint: golangci-lint

Ensure all these commands pass locally before committing.

## Common Tasks

The following are outputs from frequently run commands. Reference them instead of viewing, searching, or running bash commands to save time.

### Repository Structure
```
.
├── .github/
│   └── workflows/
│       ├── go.yml              # CI build, test, lint
│       └── go_release.yml      # Release with GoReleaser
├── pkg/
│   └── console/
│       └── console.go          # Main application logic
├── .goreleaser.yml             # Release configuration
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── oc-console.go               # Main entry point
├── README.md                   # Project documentation
└── LICENSE
```

### Key Files
- **oc-console.go**: Main entry point, minimal wrapper around pkg/console
- **pkg/console/console.go**: Core application logic using cobra CLI framework
- **go.mod**: Requires Go 1.24.0, uses Kubernetes client libraries and cobra
- **.github/workflows/go.yml**: CI pipeline for build, test, lint

### Dependencies
Main dependencies from go.mod:
- github.com/spf13/cobra - CLI framework
- github.com/pkg/browser - Browser opening functionality  
- k8s.io/client-go - Kubernetes client libraries
- k8s.io/cli-runtime - Kubernetes CLI utilities
- k8s.io/kubectl - kubectl integration

### Build Artifacts
- Binary: `oc-console` (created by `go build -o oc-console .`)
- No other build artifacts or generated files
- No dist/ or bin/ directories

### No Test Files
- Project currently has no unit tests
- `go test ./...` runs successfully but reports "[no test files]"
- Test infrastructure is available if tests need to be added

## Architecture Notes

### Application Flow
1. Uses kubectl/oc plugin architecture (checks KUBECTL_PLUGINS_CALLER env var)
2. Connects to Kubernetes/OpenShift cluster using standard kubeconfig
3. Fetches console URL from `openshift-config-managed/console-public` ConfigMap
4. Either prints URL (--url flag) or opens in default browser

### Error Handling
- Gracefully handles missing OpenShift cluster (falls back to using API server URL)
- Provides helpful error messages for connection issues
- Returns appropriate exit codes for automation

### Customization
- Supports all standard kubectl/oc flags for cluster connection
- Configurable via kubeconfig, environment variables, or command flags
- Works as both standalone binary and kubectl/oc plugin