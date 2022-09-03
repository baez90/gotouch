package root

import (
	"github.com/denizgursoy/gotouch/internal/commander"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	rootCommand := CreateRootCommand(commander.GetInstance())
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func CreateRootCommand(cmdr commander.Commander) *cobra.Command {
	createCommand := &cobra.Command{
		Use:   "gotouch",
		Short: "Helps you create new Go Projects",
		Long:  `Tag`,
		Run:   GetCreateCommandHandler(cmdr),
	}
	createCommand.Flags().StringP(FileFlagName, "f", "", "input file")

	createCommand.AddCommand(CreatePackageCommand(cmdr))

	return createCommand
}

func CreatePackageCommand(cmdr commander.Commander) *cobra.Command {
	packageCommand := &cobra.Command{
		Use:   "package",
		Short: "createYourZip",
		Long:  `Tag`,
		Run:   GetPackageCommandHandler(cmdr),
	}

	packageCommand.Flags().StringP(SourceDirectoryFlagName, "s", ".", "source directory")
	packageCommand.Flags().StringP(TargetDirectoryFlagName, "t", ".", "target directory")

	return packageCommand
}
