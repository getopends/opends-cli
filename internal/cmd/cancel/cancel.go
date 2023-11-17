package cancel

import "github.com/spf13/cobra"

var flags = &CancelFlags{}

type CancelFlags struct {
	Force  bool
	File   string
	Output string
}

func (c *CancelFlags) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&c.Force, "force", "F", false, "force")
	cmd.PersistentFlags().StringVarP(&c.File, "filename", "f", "", "file")
	cmd.PersistentFlags().StringVarP(&c.Output, "output", "o", "", "file")
}

func NewCmdCancel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel --force --file [FILENAME]",
		Short: "Cancel transaction",
	}

	flags.AddFlags(cmd)

	cmd.AddCommand(
		NewCmdCancelTransaction(),
	)

	return cmd
}
