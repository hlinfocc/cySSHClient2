package config

import (
	"os"
	"path/filepath"

	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

const (
	Separator     = string(os.PathSeparator)     // 路径分隔符（分隔路径元素）
	ListSeparator = string(os.PathListSeparator) // 路径列表分隔符（分隔多个路径）
)

func GetDbPath() string {
	exePath, err := os.Executable()
	errors.CheckError(err)
	path, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	errors.CheckError(err)

	return path + Separator + "cyssh.db"
}

func GetSocketPath() string {
	return "/tmp/cysshclient.sock"
}
