# oc-console

[![GoDoc](https://godoc.org/github.com/cblecker/oc-console?status.svg)](https://godoc.org/github.com/cblecker/oc-console)
[![Go](https://github.com/cblecker/oc-console/workflows/Go/badge.svg)](https://github.com/cblecker/oc-console/actions?query=workflow%3AGo)

A [cli plugin] that allows you to open the OpenShift 4 web console in your
default web browser.

## Installation

```
go install github.com/cblecker/oc-console
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
