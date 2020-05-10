package dep
//
// import (
// 	"os"
//
// 	mageutil "github.com/getcouragenow/bootstrap/ci/os/util"
// 	"github.com/magefile/mage/mg"
// )
//
// var curDir = func() string {
// 	name, _ := os.Getwd()
// 	return name
// }()
//
// // OS namespace
// type OS mg.Namespace
//
// // Win install windows Dependencies.
// func (OS) Win() {
// 	mageutil.Windows{}.InstallDependency()
//
// }
//
// // Linux install linux Dependencies.
// func (OS) Linux() {
// 	mageutil.Linux{}.InstallDependency()
// 	// fmt.Println("Not yet ready")
// }
//
// // Mac install Mac Dependencies.
// func (OS) Mac() {
// 	mageutil.Mac{}.InstallDependency()
// }
