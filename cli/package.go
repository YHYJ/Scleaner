/*
File: package.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 14:01:12

Description: 子命令`package`功能函数
*/

package cli

import (
	"fmt"
	"strings"

	"github.com/yhyj/scleaner/general"
)

// PackageCleaner 清理孤立软件包
func PackageCleaner(noLogoFlag bool) {
	// 检查命令
	checkArgs := []string{"-Qtdq"}
	lonelyPackages, err := general.RunCommandGetResult("pacman", checkArgs)
	if err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
	}

	// 检查命令结果解析
	if lonelyPackages == "" {
		fmt.Println("\033[35m[✔]\033[0m 没有孤立依赖包")
		if !noLogoFlag {
			// Logo命令
			mascotArgs := []string{}
			mascot, err := general.RunCommandGetResult("repo-elephant", mascotArgs)
			if err != nil {
				fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
			}
			fmt.Println(mascot)
		}
	} else {
		// 卸载命令
		uninstallArgs := []string{"-Rn"}
		uninstallCmd := append(uninstallArgs, strings.Split(lonelyPackages, "\n")...)
		if err := general.RunCommand("pacman", uninstallCmd); err != nil {
			fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
		}
	}
}
