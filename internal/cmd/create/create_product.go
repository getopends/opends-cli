package create

import "github.com/spf13/cobra"

func NewCmdCreateProduct() *cobra.Command {
	return &cobra.Command{
		Use:     "product",
		Short:   "Create product",
		Aliases: []string{"products", "prod"},
	}
}
