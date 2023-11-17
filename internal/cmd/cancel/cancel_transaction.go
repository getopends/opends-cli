package cancel

import "github.com/spf13/cobra"

var createTransactionFlags = &CancelTransactionFlags{}

type CancelTransactionFlags struct{}

func (c *CancelTransactionFlags) AddFlags(cmd *cobra.Command) {}

func NewCmdCancelTransaction() *cobra.Command {
	newCmd := &cobra.Command{
		Short:   "Cancel transaction",
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
