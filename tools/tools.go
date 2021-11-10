//go:build tools
// +build tools

package tools

import (
	_ "github.com/spf13/cobra/cobra"
	_ "go.mozilla.org/sops/v3/cmd/sops"
)
