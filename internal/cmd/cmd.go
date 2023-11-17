package cmd

import (
	"github.com/opends/opendsctl/internal/cmd/create"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use: "opendsctl",
	}

	newCmd.AddCommand(create.NewCmdCreate())

	return newCmd
}
