package create

import "github.com/spf13/cobra"

type CreateFlags struct {
	Force *bool
}

func (c *CreateFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(c.Force, "force", false, "force")
}

func NewCmdCreate() {}
