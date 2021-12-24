//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/spf13/cobra/cobra"
	_ "go.mozilla.org/sops/v3/cmd/sops"
)
