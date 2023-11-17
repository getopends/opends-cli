package create

import "github.com/spf13/cobra"

var createTransactionFlags = &CreateTransactionFlags{}

type CreateTransactionFlags struct {
	ExternalID                string
	CardNumber                string
	PhoneNumber               string
	AccountNumber             string
	AutoConfirm               bool
	RecevingCustomerFirstname string
	RecevingCustomerLastname  string
}

func (c *CreateTransactionFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.ExternalID, "external-id", "", "External id")
	cmd.Flags().StringVar(&c.CardNumber, "card-number", "", "Card number")
	cmd.Flags().StringVar(&c.PhoneNumber, "phone-number", "", "Phone number")
	cmd.Flags().StringVar(&c.AccountNumber, "account-number", "", "Account number")
	cmd.Flags().BoolVar(&c.AutoConfirm, "auto-confirm", false, "Confirm")
	cmd.Flags().StringVar(&c.RecevingCustomerFirstname, "receiving-customer-firstname", "", "First name")
	cmd.Flags().StringVar(&c.RecevingCustomerLastname, "receiving-customer-lastname", "", "Last name")
}

func NewCmdCreateTransaction() *cobra.Command {
	newCmd := &cobra.Command{
		Short:   "Create transaction",
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
