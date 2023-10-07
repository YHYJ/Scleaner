/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-17 16:08:21

Description: 程序子命令'package'时执行
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/scleaner/function"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Clear Orphaned Dependencies",
	Long:  `Purge orphaned packages that were installed as dependencies but are no longer depended on.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		noLogoFlag, _ := cmd.Flags().GetBool("nologo")
		function.PackageCleaner(noLogoFlag)
	},
}

func init() {
	packageCmd.Flags().BoolP("nologo", "n", false, "Do not display the logo")

	packageCmd.Flags().BoolP("help", "h", false, "help for package command")
	rootCmd.AddCommand(packageCmd)
}
