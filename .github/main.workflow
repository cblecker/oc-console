workflow "Release" {
  on = "push"
  resolves = ["goreleaser"]
}

action "is-tag" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "lint" {
  uses = "actions-contrib/golangci-lint@master"
  env = {
    GOPROXY = "https://proxy.golang.org"
  }
  args = "run"
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"
  secrets = [
    "GITHUB_TOKEN"
  ]
  args = "release"
  needs = ["is-tag", "lint"]
}

workflow "Lint" {
  on = "pull_request"
  resolves = ["lint"]
}
