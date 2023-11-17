package create

import "github.com/spf13/cobra"

func NewCmdCreateOperator() *cobra.Command {
	return &cobra.Command{
		Use:     "operator",
		Short:   "Create operator",
		Aliases: []string{"operators", "op"},
	}
}
