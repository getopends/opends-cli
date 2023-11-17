package create

import "github.com/spf13/cobra"

var createTransactionFlags = &ConfirmTransactionFlags{}

type ConfirmTransactionFlags struct{}

func (c *ConfirmTransactionFlags) AddFlags(cmd *cobra.Command) {}

func NewCmdConfirmTransaction() *cobra.Command {
	newCmd := &cobra.Command{
		Short:   "Confirm transaction",
		Use:     "transaction",
		Aliases: []string{"tra"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	createTransactionFlags.AddFlags(newCmd)

	return newCmd
}
