package main

import (
	"flag"
	"fmt"
	"runtime"

	u "github.com/getcouragenow/bootstrap/ci/dep/util"
)

func main() {
	install := flag.String("install", "", "install dependencies")
	flag.Parse()

	switch *install {
	case "flgo":
		installFlutterGO()
	}
}

func installFlutterGO() {
	switch runtime.GOOS {
	case "windows":
		u.Windows{}.InstallDependency()
	case "mac":
		u.Mac{}.InstallDependency()
	case "linux":
		u.Linux{}.InstallDependency()
	default:
		fmt.Println("No implemented")
		return
	}
}
