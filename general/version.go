/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 13:58:00

Description: 版本信息
*/

package general

import (
	"strconv"
	"time"
)

const (
	Name    string = "Scleaner"                 // 程序名
	Version string = "v0.7.6"                   // 程序版本
	Project string = "github.com/yhyj/scleaner" // 项目地址
)

var (
	GitCommitHash string = "Unknown" // Git 提交 Hash
	BuildTime     string = "Unknown" // 编译时间
	BuildBy       string = "Unknown" // 编译者
)

// ProgramInfo 返回程序信息
//
// 返回：
//   - 程序信息
func ProgramInfo() map[string]string {
	programInfo := make(map[string]string)

	// 解析构建时间
	BuildTimeConverted := "Unknown"
	timestamp, err := strconv.ParseInt(BuildTime, 10, 64)
	if err == nil {
		BuildTimeConverted = time.Unix(timestamp, 0).String()
	}

	programInfo["Name"] = Name
	programInfo["Version"] = Version
	programInfo["Project"] = Project
	programInfo["GitCommitHash"] = GitCommitHash
	programInfo["BuildTime"] = BuildTimeConverted
	programInfo["BuildBy"] = BuildBy

	return programInfo
}
