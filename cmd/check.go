package cmd

import (
	"fmt"
	"strings"

	"dupCheck/pkg/readYaml"

	"github.com/spf13/cobra"
)

const (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func verify(common, app map[string]string) {
	for k, _ := range common {
		if _, ok := app[k]; ok {
			issue := fmt.Sprintf("Potential issue with: %s\n", k)
			fmt.Println(string(colorRed), issue, string(colorReset))
		}
	}
}

func formatFileNames(baseFileName, check, correction string) string {
	fileName := baseFileName
	if strings.Contains(baseFileName, check) {
		fileName = strings.Replace(baseFileName, check, correction, -1)
	}
	return fileName
}

func checkFiles(filesToCheck []string) {
	for i := 0; i < len(filesToCheck); i++ {
		fileName := formatFileNames(filesToCheck[i], "application", application)
		fileName = formatFileNames(fileName, "environment", environment)
		secrets, nonsecrets := readYaml.GetEnvValues(fileName)
		fmt.Println(string(colorBlue), "Checking for duplicate values in "+fileName+"\n", string(colorReset))
		verify(secrets, nonsecrets)
		if i < len(filesToCheck) {
			for j := i + 1; j < len(filesToCheck); j++ {
				nextFile := formatFileNames(filesToCheck[j], "application", application)
				nextFile = formatFileNames(nextFile, "environment", environment)
				nextSecrets, nextNonsecrets := readYaml.GetEnvValues(nextFile)
				fmt.Println(string(colorBlue), "Checking "+fileName+" against "+nextFile+" for duplicate values\n", string(colorReset))
				verify(secrets, nextNonsecrets)
				verify(nonsecrets, nextSecrets)
			}
		}
	}

}

var check = &cobra.Command{
	Use:   "check",
	Short: "checks apps for potential duplicates",
	Long:  "checks apps for potential duplicates",
	Run: func(cmd *cobra.Command, args []string) {
		nonprodValuesFiles := []string{
			"../chime-cd/overrides/apps/common_nonprod_envs.yaml",
			"../chime-cd/overrides/apps/environment.yaml",
			"../chime-cd/overrides/apps/application/base.yaml",
			"../chime-cd/overrides/apps/application/common_nonprod_envs.yaml",
			"../chime-cd/overrides/apps/application/environment.yaml",
		}
		prodValuesFiles := []string{
			"../chime-cd/overrides/apps/prod.yaml",
			"../chime-cd/overrides/apps/application/base.yaml",
			"../chime-cd/overrides/apps/application/prod.yaml",
		}
		if environment == "production" || environment == "prod" {
			checkFiles(prodValuesFiles)
		} else {
			checkFiles(nonprodValuesFiles)
		}
	},
}
