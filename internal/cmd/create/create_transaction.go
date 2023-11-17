package create

import "github.com/spf13/cobra"

var createTransactionFlags = &CreateTransactionFlags{}

type CreateTransactionFlags struct {
	ExternalID string
}

func (c *CreateTransactionFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.ExternalID, "external-id", "", "External id")
}

func NewCmdCreateTransaction() *cobra.Command {
	newCmd := &cobra.Command{
		Short:   "Create transaction",
		Use:     "transaction",
		Aliases: []string{"tra"},
	}

	createTransactionFlags.AddFlags(newCmd)

	return newCmd
}
