package create

import "github.com/spf13/cobra"

var flags = &CreateFlags{}

type CreateFlags struct {
	Force  bool
	File   string
	Output string
}

func (c *CreateFlags) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&c.Force, "force", "F", false, "force")
	cmd.PersistentFlags().StringVarP(&c.File, "filename", "f", "", "file")
	cmd.PersistentFlags().StringVarP(&c.Output, "output", "o", "", "file")
}

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create --force --file [FILENAME]",
	}

	flags.AddFlags(cmd)

	cmd.AddCommand(
		NewCmdCreateTransaction(),
		NewCmdCreateProvider(),
		NewCmdCreateService(),
		NewCmdCreateProduct(),
		NewCmdCreateOperator(),
	)

	return cmd
}
