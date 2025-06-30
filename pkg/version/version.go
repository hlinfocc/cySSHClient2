package version

import "fmt"

const version = "2.2.1"
const SysTitle = "沁芳Linux远程主机管理系统"

func Simple() string {
	return version
}

func Full() string {
	return fmt.Sprintf("%s V%s", SysTitle, version)
}
