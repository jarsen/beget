package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repository",
	Long:  `Initializes a repository in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		createObjectsDir(projectPath)
		createRefsDir(projectPath)
		fmt.Printf("Initialized empty Jit repository in %s.", projectPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createObjectsDir(projectPath string) {
	var refPath = path.Join(projectPath, ".git/objects")
	os.MkdirAll(refPath, os.ModePerm)
}

func createRefsDir(projectPath string) {
	var refPath = path.Join(projectPath, ".git/refs")
	os.MkdirAll(refPath, os.ModePerm)
}
