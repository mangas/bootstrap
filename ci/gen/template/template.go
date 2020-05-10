package template
//
// import "fmt"
//
// // GetTemplate get template
// func GetTemplate(libName, lib, branch string) string {
// 	if libName == "" {
// 		libName = defaultLibName
// 	}
//
// 	if lib == "" {
// 		lib = defaultlib
// 	}
//
// 	if branch == "" {
// 		branch = defaultlibBranch
// 	}
// 	return fmt.Sprintf(template, libName, lib, branch)
// }
//
// var (
// 	defaultLibName   = "main"
// 	defaultlib       = "github.com/getcouragenow/"
// 	defaultlibBranch = "master"
// 	// Template basic template
// 	template = `// +build mage
//
// package main
//
// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"path"
// 	"runtime"
// )
//
// var (
// 	// Global
// 	home   = ""
// 	goos   = ""
// 	goarch = ""
// 	goPath = ""
// 	goBin  = ""
//
// 	// Repo
// 	libName   = "%s"
// 	lib       = "%s" + libName
// 	libBranch = "%s"
// 	libFsPath = path.Join(goPath, "src", lib)
//
// 	// Sample
// 	sample       = "keyboard_event"
// 	sampleName   = "client"
// 	sampleFsPath = path.Join(libFsPath, sampleName)
//
// 	// protobuf
// 	protoOutput = ""
//
// 	// Gir
// 	gitVersion = ""
//
// 	// Keystore
// 	andKeystoreAppName = "main"
// 	andKeystoreFsPath  = ""
//
// 	// Docker
// 	dockerTempFsPath = ""
//
// 	// Firebase
// 	// PROD
// 	fbProdProjID = "winwisely-app-prod"
//
// 	// DEV
// 	fbDevProjID = "winwisely-app-dev"
//
// 	// User ( the developer)
// 	fbUserProjID = "winwisely99-main-master"
//
// 	// Set the final variables used here !
// 	fbSessionProjectID         = fbUserProjID
// 	fbSessionProjectConsoleURL = "https://console.firebase.google.com/project/" + fbSessionProjectID + "/hosting/main"
// 	fbSessionProjectWebURL     = "https://" + fbSessionProjectID + ".firebaseapp.com/"
//
// 	gCloudSessionProjectConsoleURL = "https://console.cloud.google.com/home/dashboard?project=" + fbSessionProjectID
// 	gitHash                        = sExecute("git", "log", "-1", "--pretty=format:\"%H\"")
// )
//
// func init() {
// 	//set global
//
// 	goPath = os.Getenv("GOPATH")
// 	if goPath == "" {
// 		os.Setenv("GOPATH", sExecute("go", "env", "GOPATH"))
// 	}
// 	goBin = os.Getenv("GOBIN")
// 	goos = runtime.GOOS
// 	goarch = runtime.GOARCH
//
// 	// set git version
// 	gitVersion = sExecute("git", "describe", "--tags")
//
// 	// set keystoreFsPath
// 	pwd, _ := os.Getwd()
// 	andKeystoreFsPath = path.Join(pwd, "_keys", andKeystoreAppName)
//
// 	libFsPath = path.Join(goPath, "src", lib)
// 	sampleFsPath = path.Join(libFsPath, sampleName)
// 	home, _ = os.UserHomeDir()
// 	protoOutput = libFsPath
// }
//
// // Print help
// func Print() {
// 	fmt.Println()
// 	fmt.Println("GO_OS:", goos)
// 	fmt.Println("GO_AECH:", goarch)
// 	fmt.Println()
//
// 	fmt.Println()
// 	fmt.Println("GIT_VERSION:", gitVersion)
// 	fmt.Println("GIT_OTHER:", sExecute("git", "describe", "--all"))
// 	fmt.Println()
//
// 	fmt.Println()
// 	fmt.Println("AND_KEYSTORE_APP_NAME:", andKeystoreAppName)
// 	fmt.Println("AND_KEYSTORE_FSPATH:", andKeystoreFsPath)
// 	fmt.Println()
//
// 	fmt.Println()
// 	fmt.Println("LIB_NAME:", libName)
// 	fmt.Println("LIB:", lib)
// 	fmt.Println("LIB_BRANCH:", libBranch)
// 	fmt.Println("LIB_FSPATH:", libFsPath)
// 	fmt.Println("GOPATH: ", goPath)
// 	fmt.Println("HOME: ", home)
// 	fmt.Println()
//
// 	fmt.Println()
// 	fmt.Println("SAMPLE_NAME:", sampleName)
// 	fmt.Println("SAMPLE_FSPATH:", sampleFsPath)
// 	fmt.Println("PROTO_OUTPUT:", protoOutput)
// 	fmt.Println()
// }
//
// // GitClone git clone
// func GitClone() {
// 	execute("mkdir", "-p", libFsPath)
// 	os.Chdir(libFsPath)
// 	os.Chdir("..")
// 	execute("rm", "-rf", libName)
// 	execute("git", "clone", "ssh://git@"+lib+".git")
// }
//
// // GitPull git pull.
// func GitPull() {
// 	cdExecute(libFsPath, "git", "pull")
// }
//
// // GitClean git clean.
// func GitClean() {
// 	execute("rm", "-rf", libFsPath)
// }
//
// // Code open repo in vscode
// func Code() {
// 	execute("code", libFsPath)
// }
//
// // FluConfig flutter config
// func FluConfig() {
// 	execute("flutter", "config")
// 	execute("flutter", "doctor -v")
// 	execute("flutter", "devices")
// }
//
// // FluConfigDesk flutter config desk
// func FluConfigDesk() {
// 	execute("hover", "-h")
// }
//
// // FluClean flutter clean
// func FluClean() {
// 	cdExecute(sampleFsPath, "flutter", "clean")
// }
//
// // FluUpdate flutter update
// func FluUpdate() {
// 	cdExecute(sampleFsPath, "flutter", "packages", "get")
// 	GenIcon()
// 	GenHive()
// 	GenProto()
// 	cdExecute(sampleFsPath, "flutter", "analyze")
// }
//
// // GenIcon generate icon.
// func GenIcon() {
//
// 	// mobile and web
// 	fmt.Println()
// 	fmt.Println("Generating icons for Flutter")
// 	fmt.Println()
// 	cdExecute(sampleFsPath, "flutter", "pub", "run", "flutter_launcher_icons:main")
//
// 	// desktop
// 	fmt.Println()
// 	fmt.Println("Copying icon-png from flutter assets into go assets, so hover can use it")
// 	fmt.Println()
// 	execute("cp", sampleFsPath+"/assets/icon/icon.png", sampleFsPath+"/go/assets")
// }
//
// // GenHive generate hive.
// func GenHive() {
// 	cdExecute(sampleFsPath, "flutter", "packages", "pub", "run", "build_runner", "build", "--delete-conflicting-outputs")
// }
//
// // GenProto generate proto
// func GenProto() {
// 	execute("pub", "global", "activate", "protoc_plugin")
// 	execute("mkdir", "-p", sampleFsPath+"/lib/api/v1/google/protobuf")
//
// 	if runtime.GOOS == "windows" {
// 		fmt.Println("Windows detected")
// 		execute("protoc", "empty.proto", "timestamp.proto", "wrappers.proto",
// 			fmt.Sprintf("--proto_path=%s/server/third_party/google/protobuf/", libFsPath),
// 			fmt.Sprintf("--plugin=%s/AppData/Roaming/Pub/Cache/bin/protoc-gen-dart.bat", home),
// 			fmt.Sprintf("--dart_out=grpc:%s/client/lib/chat_view/api/v1/google/protobuf", protoOutput),
// 		)
// 		execute("protoc", "chat.proto",
// 			fmt.Sprintf("--proto_path=%s/server/api/proto/v1/", libFsPath),
// 			fmt.Sprintf("--plugin=%s/AppData/Roaming/Pub/Cache/bin/protoc-gen-dart.bat", home),
// 			fmt.Sprintf("--dart_out=grpc:%s/client/lib/chat_view/api/v1/", protoOutput),
// 		)
// 	} else {
// 		execute("protoc", "empty.proto", "timestamp.proto wrappers.proto",
// 			fmt.Sprintf("--proto_path=%s/server/third_party/google/protobuf", libFsPath),
// 			fmt.Sprintf("--plugin=protoc-gen-dart=%s/.pub-cache/bin/protoc-gen-dart", home),
// 			fmt.Sprintf("--dart_out=grpc:%s/lib/api/v1/google/protobuf", protoOutput),
// 		)
// 		execute("protoc", "chat.proto",
// 			fmt.Sprintf("--proto_path=%s/server/api/proto/v1/", libFsPath),
// 			fmt.Sprintf("--plugin=protoc-gen-dart=%s/.pub-cache/bin/protoc-gen-dart", home),
// 			fmt.Sprintf("--dart_out=grpc:%s/client/lib/chat_view/api/v1/", sampleFsPath),
// 		)
// 	}
// }
//
// // Flutter Desk
//
// // DFluLocalPrint flutter desk local print
// func DFluLocalPrint() {
// 	cdExecute(sampleFsPath, "hover", "doctor")
// }
//
// // DFluLocalRun flutter desk local run
// func DFluLocalRun() {
// 	cdExecute(sampleFsPath, "hover", "run")
// }
//
// // DFluLocalBuild flutter desk local build
// func DFluLocalBuild() {
//
// 	if goos == "linux" {
// 		fmt.Println("Linux detected")
// 		cdExecute(sampleFsPath, "hover", "build", goos+"-appimage", "--version-number", gitVersion)
// 		cdExecute(sampleFsPath, "hover", "build", goos+"-deb", "--version-number", gitVersion)
// 		cdExecute(sampleFsPath, "hover", "build", goos+"-rpm", "--version-number", gitVersion)
// 		cdExecute(sampleFsPath, "hover", "build", goos+"-snap", "--version-number", gitVersion)
//
// 	} else if goos == "windows" {
// 		fmt.Println("Windows detected")
// 		cdExecute(sampleFsPath, "hover", "build", goos, "--version-number", gitVersion)
// 		cdExecute(sampleFsPath, "hover", "build", goos+"-msi", "--version-number", gitVersion)
// 	} else if goos == "darwin" {
// 		fmt.Println("Mac detected")
// 		// cdExecute(sampleFsPath, "hover", "build", goos, "--version-number", gitVersion)
// 		// cdExecute(sampleFsPath, "hover", "build", goos, "--version-number", gitVersion)
// 		cdExecute(sampleFsPath, "hover", "build", goos+"-pkg", "--version-number", gitVersion)
// 		// cdExecute(sampleFsPath, "hover", "build", goos+"-dmg", "--version-number", gitVersion)
// 	}
//
// }
//
// // DFluBuildAll flutter desk build all
// func DFluBuildAll() {
// 	DFluBuildClean()
// 	DFluBuildInit()
// 	DFluBuild()
// }
//
// // DFluBuildInit flutter desk build init
// func DFluBuildInit() {
// 	cdExecute(sampleFsPath, "hover", "init", path.Join(lib, sampleName))
// }
//
// // DFluBuildClean flutter desk build clean
// func DFluBuildClean() {
// 	execute("rm", "-rf", sampleFsPath+"/go")
// 	execute("docker", "system", "prune")
// 	// execute("docker", "system", "prune", "--force")
// }
//
// // DFluBuild flutter desj build
// func DFluBuild() {
// 	fmt.Println("use docker to build")
// 	cdExecute(sampleFsPath, "hover", "build", "darwin", "--version-number", gitVersion)
// 	cdExecute(sampleFsPath, "hover", "build", "windows", "--version-number", gitVersion)
// 	cdExecute(sampleFsPath, "hover", "build", "linux", "--version-number", gitVersion)
// }
//
// // DFluBuildRun flutter desk build
// func DFluBuildRun() {
// 	DFluBuild()
// 	execute("open", path.Join(sampleFsPath, "go/build/outputs", goos))
// 	execute("open", path.Join(sampleFsPath, "go/build/outputs", goos, sampleName))
// }
//
// // DFluPackAll flutter desk pack all
// func DFluPackAll() {
// 	DFluPackClean()
// 	DFluPackInit()
// 	DFluPack()
// }
//
// // DFluPackClean flutter desk pack clean
// func DFluPackClean() {
// 	execute("rm", "-rf", path.Join(sampleFsPath, "go", "packaging"))
// 	execute("docker", "system", "prune")
// 	execute("docker", "system", "prune", "--force")
// }
//
// //DFluPackInit flutter desk pack init
// func DFluPackInit() {
// 	// hover init-packaging --help
//
// 	// darwin
// 	cdExecute(sampleFsPath, "hover", "init-packaging", "darwin-bundle")
// 	cdExecute(sampleFsPath, "hover", "init-packaging", "darwin-pkg")
//
// 	// 	# windows
// 	cdExecute(sampleFsPath, "hover", "init-packaging", "windows-msi")
//
// 	// 	# linux
// 	cdExecute(sampleFsPath, "hover", "init-packaging", "linux-appimage")
// 	cdExecute(sampleFsPath, "hover", "init-packaging", "linux-deb")
// 	cdExecute(sampleFsPath, "hover", "init-packaging", "linux-snap")
// }
//
// // DFluPack flutter desk pack
// func DFluPack() {
// 	// hover build --help
//
// 	// darwin (works)
// 	cdExecute(sampleFsPath, "hover", "build", "darwin-pkg", "--version-number", gitVersion)
//
// 	// windows (works)
// 	cdExecute(sampleFsPath, "hover", "build", "windows-msi", "--version-number", gitVersion)
//
// 	// linux (works)
// 	cdExecute(sampleFsPath, "hover", "build", "linux-deb", "--version-number", gitVersion)
//
// 	// broken: Issue: https://github.com/go-flutter-desktop/go-flutter/issues/287#issuecomment-544161253
// 	// marked as "will not fix" because its shitty Ubuntu SnapCraft error.
// 	// It seems that linux-appimage (https://appimage.org/) works on ubuntu so screw SnapCraft..
// 	// to install on ubuntu using a appimage is easy: https://askubuntu.com/questions/774490/what-is-an-appimage-how-do-i-install-it
// 	// We need to try and make sure it works with flutter ! Who has an ubuntu laptop ?
//
// 	// Update: works for basic go-flutter example
// 	// cdExecute(sampleFsPath, "hover", "build", "linux-snap")
//
// 	// 	# broken: Issue: https://github.com/go-flutter-desktop/go-flutter/issues/289
// 	// 	# Marked "as wont fix". So have to ask AppImage team.
//
// 	// 	# Update: works for basic go-flutter example
// 	// cdExecute(sampleFsPath, "hover", "build", "linux-appimage")
// }
//
// // DFluPackOpen flutter desk pack open
// func DFluPackOpen() {
// 	execute("open", path.Join(sampleFsPath, "go", "build", "outputs"))
// }
//
// // Pack debugging
//
// // # Will vary per run and machine.
// // DOCKERTEMP_FSPATH=/var/folders/wp/ff6sz9qs6g71jnm12nj2kbyw0000gp
//
// // DFluPackTmpList flutter desk pack tmp list
// func DFluPackTmpList() {
// 	execute("ls", "/var/folders/wp/")
// 	execute("stat", dockerTempFsPath)
// }
//
// // DFluPackTmpClean flutter desk pack tmp clean
// func DFluPackTmpClean() {
// 	// clean out docker tmp file dir
// 	execute("rm", "-rf", dockerTempFsPath)
// 	// execute("ls", "/var/folders/wp/")
// 	// cdExecute("/var/folders/wp/", "rm", "-rf", "*")
// 	// execute("ls", "/var/folders/wp/")
// }
//
// // DFluPackTmpZip flutter desk pack tmp zip
// func DFluPackTmpZip() {
// 	// to share whats in tmp
// 	cdExecute(dockerTempFsPath, "zip", "-r", "-X", path.Join(sampleFsPath, "go/build/outputs/dockertemp.zip"), "*")
// }
//
// // Desk Zipping
//
// // DFluDistAll flutter desk dist all
// func DFluDistAll() {
// 	DFluDistClean()
// 	DFluDistZip()
// 	DFluDistUnzip()
// }
//
// // DFluDistClean flutter desk dist clean
// func DFluDistClean() {
// 	execute("rm", "-rf", path.Join(sampleFsPath, "dist"))
// 	execute("mkdir", "-p", path.Join(sampleFsPath, "dist"))
// }
//
// // DFluDistZip flutter desk dist zip
// func DFluDistZip() {
// 	// zip build by OS.
// 	p := "go/build/outputs"
// 	// darwin
// 	cdExecute(path.Join(sampleFsPath, p, "darwin"), "zip", "-r", "-X", path.Join(sampleFsPath, "dist/darwin.zip"), "*")
// 	// windows
// 	cdExecute(path.Join(sampleFsPath, p, "windows"), "zip", "-r", "-X", path.Join(sampleFsPath, "dist/windows.zip"), "*")
// 	//linux
// 	cdExecute(path.Join(sampleFsPath, p, "linux"), "zip", "-r", "-X", path.Join(sampleFsPath, "dist/linux.zip"), "*")
// }
//
// // DFluDistUnzip flutter desk dist unzip
// func DFluDistUnzip() {
// 	p := "dist/out"
//
// 	execute("rm", "-rf", path.Join(sampleFsPath, p))
// 	execute("mkdir", "-p", path.Join(sampleFsPath, p))
//
// 	// darwin
// 	execute("unzip", "-p", path.Join(sampleFsPath, p, "darwin.zip"), "-d", path.Join(sampleFsPath, p, "darwin"))
// 	// windows
// 	execute("unzip", "-p", path.Join(sampleFsPath, p, "windows.zip"), "-d", path.Join(sampleFsPath, p, "windows"))
// 	// 	# linux
// 	execute("unzip", "-p", path.Join(sampleFsPath, p, "linux.zip"), "-d", path.Join(sampleFsPath, p, "linux"))
// }
//
// // Mob
//
// // FluMobRun flutter mob run
// func FluMobRun() {
// 	cdExecute(sampleFsPath, "flutter", "packages", "get")
// 	cdExecute(sampleFsPath, "flutter", "packages", "pub", "run", "build_runner", "build")
// 	cdExecute(sampleFsPath, "flutter", "run", "-d", "all")
// }
//
// // FluMobReleaseAndInit flu-mob-release-and-init
// func FluMobReleaseAndInit() {
// 	execute("mkdir", "-p", andKeystoreFsPath)
//
// 	// open keystore
// 	// execute("code", andKeystoreFsPath)
//
// 	// Create keystore
// 	// Warning: Keep the keystore file private; do not check it into public source control.
// 	// Cannot overwrite existing....
// 	// windows is a littel diff
// 	if goos == "windows" {
// 		execute("keytool", "-genkey", "-v", "-keystore", andKeystoreFsPath+"/key.jks", "-keyalg", "RSA", "-storetype", "JKS", "-keysize", "2048", "-validity", "10000", "-alias", "key")
// 	} else {
// 		execute("keytool", "-genkey", "-v", "-keystore", andKeystoreFsPath+"/key.jks", "-keyalg", "RSA", "-keysize", "2048", "-validity", "10000", "-alias", "key")
// 	}
// }
//
// // FluMobReleaseAndGen flu-mob-release-and-gen
// func FluMobReleaseAndGen() {
// 	// https://flutter.dev/docs/deployment/android
//
// 	execute("touch", sampleFsPath+"/android/key.properties")
// 	execute("code", sampleFsPath+"/android/key.properties")
// 	//  TODO gen this
// 	// - add the 4 lines
//
// 	execute("touch", sampleFsPath+"/android/app/proguard-rules.pro")
// 	execute("code", sampleFsPath+"/android/app/proguard-rules.pro")
// 	// TODO gen this
// 	// - add a few lines
//
// 	execute("touch", sampleFsPath+"/android/app/build.gradle")
// 	execute("code", sampleFsPath+"/android/app/build.gradle")
// 	// TODO: gen this
// 	// - add a few lines, etc etc
// 	// - add ref to proguard
//
// 	// Build it
// 	cdExecute(sampleFsPath, "flutter", "build", "appbundle")
//
// 	// Create an account on Google Play Store
// 	// https://play.google.com/apps/publish/?account=7105100900922845210
//
// 	// DO BETA RELEASE: https://play.google.com/apps/publish/?account=7105100900922845210#PrepareReleasePlace:p=tmp.05109091238716120360.1538517246276.1913796864&appid=4972552978345084231&releaseTrackId=4699993911531051870&releaseId=4701866095568168106
// 	// Just uplaod the "build/app/outputs/bundle/release/app.aab"
//
// 	// TODO: https://play.google.com/apps/publish/?account=7105100900922845210#MarketListingPlace:p=com.winwisely.whitelabel&appid=4972552978345084231
// 	// all the screenshots!!
// }
//
// // FluMobReleaseIOSGen flu-mob-release-ios-gen
// func FluMobReleaseIOSGen() {
// 	fmt.Println("Not implemented yet")
// 	// TODO: This needs to be done.
// }
//
// // WEB
//
// // FluWebConfig flutter web config
// func FluWebConfig() {
// 	// flutter channel ?
// 	// flutter channel dev
// 	// flutter upgrade
//
// 	execute("flutter", "config", "--enable-web")
// 	// turn off any desktop
// 	execute("flutter", "config", "--no-enable-linux-desktop")
// 	execute("flutter", "config", "--no-enable-macos-desktop")
// 	execute("flutter", "config", "--no-enable-windows-desktop")
// }
//
// // FluWebCreate flutter web create
// func FluWebCreate() {
// 	FluWebConfig()
// 	// works
// 	// make sure using new dir.
// 	execute("mkdir", "-p", sampleFsPath)
// 	cdExecute(sampleFsPath, "flutter", "create", "--web", ".")
// }
//
// // FluWebRun flutter web run
// func FluWebRun() {
// 	FluWebConfig()
// 	// works
// 	// Reload works too :)
// 	cdExecute(sampleFsPath, "flutter", "run", "-d", "chrome")
// }
//
// // FluWebBuild flutter web build
// func FluWebBuild() {
// 	FluWebConfig()
// 	// works :)
// 	cdExecute(sampleFsPath, "flutter", "build", "web", "--release")
// }
//
// // FluWebTest flutter web test
// func FluWebTest() {
// 	// works :)
// 	cdExecute(sampleFsPath, "flutter", "test", "--platform=chrome")
// }
//
// // FbDep firebase dep
// func FbDep() {
// 	// mac specific for now - sorry
// 	execute("brew", "install", "node")
// 	execute("npm", "i", "-g", "firebase-tools")
// }
//
// // FbPrint print
// func FbPrint() {
// 	fmt.Println()
// 	fmt.Println("FB_SESSION_PROJECT_ID:", fbSessionProjectID)
// 	fmt.Println("GIT_HASH:", gitHash)
// 	fmt.Println("FB_SESSION_PROJECT_CONSOLE_URL:", fbSessionProjectConsoleURL)
// 	fmt.Println("FB_SESSION_PROJECT_WEB_URL:", fbSessionProjectWebURL)
// 	fmt.Println("GCLOUD_SESSION_PROJECT_CONSOLE_URL:", gCloudSessionProjectConsoleURL)
// 	fmt.Println()
// }
//
// // FbAUTH fa=irebase auth
// func FbAUTH() {
// 	// firebase logout
// 	execute("firebase", "logout")
//
// 	// firebase login and pick which account to use via which browser you hit.
// 	execute("firebase", "login", "--no-localhost")
//
// 	// see what projects you have setup
// 	execute("firebase", "projects:list")
// }
//
// // FbInit firebase init
// func FbInit() {
// 	// This is a one time thing for a project
// 	execute("firebase", "init")
// }
//
// // FbConsole opens the web console.
// func FbConsole() {
// 	execute("open", gCloudSessionProjectConsoleURL)
// }
//
// // FbTest firebase test.
// func FbTest() {
// 	fmt.Println(libFsPath + "/client/build/web")
// }
//
// // FbServe firebase serve.
// func FbServe() {
// 	// runs locally
//
// 	// rebuilds flutter web
// 	FluWebBuild()
//
// 	// 	copy code from web build folder to public folder
// 	execute("rm", "-rf", "./public")
// 	execute("cp", "-r", libFsPath+"/client/build/web", "./public")
// 	execute("firebase", "serve", "--only", "hosting")
// }
//
// // FbDeploy firebase deploy
// func FbDeploy() {
// 	// rebuilds flutter web
// 	FluWebBuild()
//
// 	// copy code from web build folder to public folder
// 	execute("rm", "-rf", "./public")
// 	execute("cp", "-r", libFsPath+"/client/build/web", "./public")
// 	execute("firebase", "deploy", "--only", "hosting", "-m", "githash: "+gitHash)
// }
//
// // FbCIGet firebase gets secure keys used for CI automation.
// func FbCIGet() {
// 	execute("firebase", "login:ci")
// 	// PROJ: winwisely99-main-master
// 	// record the secret
// }
//
// // FbCIRevoke firebase revokes secure keys used for CI automation.
// func FbCIRevoke() {
// 	// revokes secure keys used for CI automation
// 	execute("firebase", "logout", "--token", "<token>")
// }
//
// // i18n
//
// // #SAMPLE=keyboard_event
// // # works
//
// // #SAMPLE=mousebuttons
// // # works
//
// // #SAMPLE=pointer_demo
// // # works and is amazing
// // # Is perfect sampler, as it has: licenses, setting ( dark mode, debugging)
// // # TODO: The i18n code is there and working, but there is no easy way to change it in the GUI.
// // # TODO: Looks like i can use my Google trans gocode to gen more if i can parse the ARBS
//
// // stocks-i18n-step1:
//
// // StocksI18nStep1 generate ui code. --> i18n code ( intl_messages.arb from lib/stock_strings.dart)
// func StocksI18nStep1() {
// 	cdExecute(path.Join(libFsPath, sample), "flutter", "pub", "pub", "run", "intl_translation:extract_to_arb", "--output-dir=lib/i18n", "lib/stock_strings.dart")
// 	cdExecute(path.Join(libFsPath, sample, "lib/i18n"), "ls", "-al")
// }
//
// // StocksI18nStep2 generate arb string. --> i18n code (a stock_messages_<locale>.dart for each stocks_<locale>.arb file and stock_messages_all.dart)
// func StocksI18nStep2() {
// 	cdExecute(path.Join(libFsPath, sample), "flutter", "pub", "pub", "run", "intl_translation:generate_from_arb", "--output-dir=lib/i18n", "--generated-file-prefix=stock_", "--no-use-deferred-loading lib/*.dart lib/i18n/stocks_*.arbs")
// 	// Now you have all the "stock_messages_all, en, es.dart regenerated
// 	// Its used the stocks_en,es.arb as sources
// 	// Now each of the locale i18n dart code is populated with the translated string.
// }
//
// // Utils
// func sExecute(cmd string, args ...string) string {
// 	c := exec.Command("git", "status")
// 	out, _ := c.Output()
// 	c.Run()
// 	return string(out)
// }
// func cdExecute(dir, cmd string, args ...string) {
// 	err := os.Chdir(dir)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	execute(cmd, args...)
// }
//
// func execute(cmd string, args ...string) {
// 	c := exec.Command(cmd, args...)
// 	c.Stderr = os.Stderr
// 	c.Stdout = os.Stdout
// 	err := c.Run()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
//
// // `
// )
