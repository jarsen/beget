package cmd

import (
	"fmt"
	"os"

	"github.com/jarsen/beget/core"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repository",
	Long:  `Initializes a repository in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll(core.ObjectsPath(), os.ModePerm)
		os.MkdirAll(core.RefsPath(), os.ModePerm)
		fmt.Printf("Initialized empty Jit repository in %s.", core.ProjectPath())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
