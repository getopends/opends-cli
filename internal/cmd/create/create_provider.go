package create

import "github.com/spf13/cobra"

func NewCmdCreateProvider() *cobra.Command {
	return &cobra.Command{
		Use:     "provider",
		Short:   "Create provider",
		Aliases: []string{"providers"},
	}
}
