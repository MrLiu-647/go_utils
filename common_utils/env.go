package common_utils

import "runtime"

var OsThis = runtime.GOOS

const (
	OsMac     = "darwin"
	OsWindows = "windows"
	OsLinux   = "linux"
)

func IsMac() bool {
	return OsThis == OsMac
}

func IsLinux() bool {
	return OsThis == OsLinux
}

func IsWindows() bool {
	return OsThis == OsWindows
}

func IsProduct() bool {
	return IsLinux()
}
