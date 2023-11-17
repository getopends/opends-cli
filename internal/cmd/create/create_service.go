package create

import "github.com/spf13/cobra"

func NewCmdCreateService() *cobra.Command {
	return &cobra.Command{
		Use:     "service",
		Short:   "Create service",
		Aliases: []string{"services", "svc"},
	}
}
