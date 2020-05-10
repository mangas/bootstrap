package build
//
// import (
// 	"github.com/magefile/mage/mg"
// 	"github.com/magefile/mage/sh"
// )
//
// // var projectName = "Bootstrap"
//
// // Build namespace
// type Build mg.Namespace
//
// // Windows bootstrap binary for windows(amd64)
// func (Build) Windows() {
// 	build("windows", "amd64", "bs.exe")
// }
//
// // Mac binary bootstrap for mac darwin (amd64)
// func (Build) Mac() {
// 	build("darwin", "amd64", "bs")
// }
//
// // Linux bootstrap binary for linux (amd64)
// func (Build) Linux() {
// 	build("linux", "amd64", "bs")
// }
//
// func build(goos, goarch, out string) {
// 	sh.RunV("mage", "-compile", out, "-goos", goos, "-goarch", goarch)
// }
