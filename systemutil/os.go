package systemutil

import (
	"os"
	"runtime"
	"strings"
)

func GetArch() string {
	return runtime.GOARCH
}

func GetName() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name
}

func GetVersion() string {
	return runtime.Version()
}

func IsLinux() bool {
	return getOSMatches("Linux") || getOSMatches("LINUX")
}

func IsWindows() bool {
	return getOSMatches("Windows")
}

func IsMac() bool {
	return getOSMatches("Mac")
}

func getOSMatches(osNamePrefix string) bool {
	return strings.Contains(GetName(), osNamePrefix)
}

// 获取文本换行符
func GetFileSeparator() string {
	if IsWindows() {
		return `\r\n`
	}
	return `\n`
}

// 取得OS的文本文件换行符
func GetLineSeparator() string {
	return string(os.PathSeparator)
}

// 取得OS的搜索路径分隔符 Unix为:，Windows为;。
func GetPathSeparator() string {
	return string(os.PathListSeparator)
}
