package cmd

import (
	e "github.com/cloudposse/atmos/internal/exec"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

// terraformGenerateBackendCmd generates backend config for a terraform components
var terraformGenerateBackendCmd = &cobra.Command{
	Use:                "backend",
	Short:              "generate backend",
	Long:               `This command generates the backend config for a terraform component`,
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
	Run: func(cmd *cobra.Command, args []string) {
		err := e.ExecuteTerraformGenerateBackend(cmd, args)
		if err != nil {
			color.Red("%s\n\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	terraformGenerateBackendCmd.DisableFlagParsing = false
	terraformGenerateBackendCmd.PersistentFlags().StringP("stack", "s", "", "")

	err := terraformGenerateBackendCmd.MarkPersistentFlagRequired("stack")
	if err != nil {
		color.Red("%s\n\n", err)
		os.Exit(1)
	}

	terraformGenerateCmd.AddCommand(terraformGenerateBackendCmd)
}
