package cmd

import (
	"github.com/opends/opendsctl/internal/cmd/cancel"
	"github.com/opends/opendsctl/internal/cmd/confirm"
	"github.com/opends/opendsctl/internal/cmd/create"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use: "opendsctl",
	}

	newCmd.AddCommand(create.NewCmdCreate())
	newCmd.AddCommand(confirm.NewCmdConfirm())
	newCmd.AddCommand(cancel.NewCmdCancel())

	return newCmd
}
