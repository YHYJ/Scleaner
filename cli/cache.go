/*
File: cache.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-22 13:35:18

Description: 子命令 'cache' 的实现
*/

package cli

import (
	"github.com/gookit/color"
	"github.com/yhyj/scleaner/general"
)

// CacheCleaner 清理缓存
func CacheCleaner() {
	// 清除 pip 缓存
	color.Printf("%s %s\n", general.NoticeText("-->"), general.LightText("Cleaning pip cache"))
	pipArgs := []string{"cache", "purge"}
	if err := general.RunCommand("pip", pipArgs); err != nil {
		color.Danger.Println(err)
	}
	color.Println()

	// 验证 npm 缓存文件夹
	color.Printf("%s %s\n", general.NoticeText("-->"), general.LightText("Verify the cache folder"))
	npmArgs := []string{"cache", "verify"}
	if err := general.RunCommand("npm", npmArgs); err != nil {
		color.Danger.Println(err)
	}
	color.Println()

	// 清除 yarn 缓存
	color.Printf("%s %s\n", general.NoticeText("-->"), general.LightText("Cleaning yarn cache"))
	yarnArgs := []string{"cache", "clean"}
	if err := general.RunCommand("yarn", yarnArgs); err != nil {
		color.Danger.Println(err)
	}
	color.Println()
}
