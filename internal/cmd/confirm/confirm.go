package confirm

import "github.com/spf13/cobra"

var flags = &ConfirmFlags{}

type ConfirmFlags struct {
	Force  bool
	File   string
	Output string
}

func (c *ConfirmFlags) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&c.Force, "force", "F", false, "force")
	cmd.PersistentFlags().StringVarP(&c.File, "filename", "f", "", "file")
	cmd.PersistentFlags().StringVarP(&c.Output, "output", "o", "", "file")
}

func NewCmdConfirm() *cobra.Command {
	cmd := &cobra.Command{
		Use: "confirm --force --file [FILENAME]",
	}

	flags.AddFlags(cmd)

	cmd.AddCommand(
		NewCmdConfirmTransaction(),
	)

	return cmd
}
