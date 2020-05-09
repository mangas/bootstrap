package oses

/*
Package oses is for getting os, user, and git information
*/

import (
	"fmt"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/termutil"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// osProperties is the current environment for user's OS
type osProperties struct {
	name   string        // os username
	root   string        // root is the rootdir or homedir of the current user
	groups []*user.Group // groups user belongs to
	osInfo OSInfoGetter  // Os properties
}

func initOSProperties() (*osProperties, error) {
	shellUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	gids, err := shellUser.GroupIds()
	if err != nil {
		return nil, err
	}
	var groups []*user.Group
	for _, gid := range gids {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	osInfo, err := getOsInfoGetter()
	if err != nil {
		return nil, err
	}
	return &osProperties{
		name:   shellUser.Username,
		root:   shellUser.HomeDir,
		groups: groups,
		osInfo: osInfo,
	}, nil
}
func (o *osProperties) GetName() string    { return o.name }
func (o *osProperties) GetRoot() string    { return o.root }
func (o *osProperties) GetAccount() string { return o.name }
func (o *osProperties) GetGroups() []*user.Group {
	return o.groups
}
func (o *osProperties) GetOsInfo() OSInfoGetter { return o.osInfo }
func (o *osProperties) ToMapString() map[string]string {
	ms := map[string]string{}
	var s strings.Builder
	for i := 0; i < len(o.GetGroups()); i++ {
		s.WriteString(o.GetGroups()[i].Name)
		if i < (len(o.GetGroups()) - 1) {
			s.WriteRune(',')
			s.WriteRune(' ')
		}
	}
	ms["Username"] = o.GetName()
	ms["User Homedir"] = o.GetRoot()
	ms["User Groups"] = s.String()
	return ms
}

// gitConfig is the current environment for user's git configuration
// users of this tool must have their own
type gitConfig struct {
	name    string       // git config --global user.name
	root    string       // ex: github.com/winwiselyxx
	account string       // git config --global user.email
	osInfo  OSInfoGetter // just added it for the sake of ease
}

func initGitConfig() (*gitConfig, error) {
	userName, err := runUnixCmd("git", "config", "user.name")
	if err != nil {
		return nil, err
	}
	root := fmt.Sprintf("github.com/%s", *userName)
	account, err := runUnixCmd("git", "config", "user.email")
	if err != nil {
		return nil, err
	}
	return &gitConfig{
		name:    *userName,
		root:    root,
		account: *account,
		osInfo:  nil,
	}, nil
}

func (g *gitConfig) GetName() string    { return g.name }
func (g *gitConfig) GetRoot() string    { return g.root }
func (g *gitConfig) GetAccount() string { return g.account }
func (g *gitConfig) ToMapString() map[string]string {
	ms := map[string]string{}
	ms["Git Username"] = g.GetName()
	ms["Git Email"] = g.GetAccount()
	ms["Git URL"] = g.GetRoot()
	return ms
}

type goConfig struct {
	goRoot string
	goPath string
}

func initGoConfig() *goConfig {
	return &goConfig{
		runtime.GOROOT(),
		os.Getenv("GOPATH"),
	}
}
func (g *goConfig) GoRoot() string { return g.goRoot }
func (g *goConfig) GoPath() string { return g.goPath }
func (g *goConfig) ToMapString() map[string]string {
	ms := map[string]string{}
	ms["GOROOT"] = g.GoRoot()
	ms["GOPATH"] = g.GoPath()
	return ms
}

type UserOsEnv struct {
	osProperties *osProperties
	goEnv        *goConfig
	gitUser      *gitConfig
}

func InitUserOsEnv() (*UserOsEnv, error) {
	osProp, err := initOSProperties()
	if err != nil {
		return nil, err
	}
	gitUser, err := initGitConfig()
	if err != nil {
		return nil, err
	}
	goenv := initGoConfig()
	return &UserOsEnv{
		osProperties: osProp,
		gitUser:      gitUser,
		goEnv:        goenv,
	}, nil
}

func (u *UserOsEnv) GetGoPath() string              { return u.goEnv.GoPath() }
func (u *UserOsEnv) GetGitUser() *gitConfig         { return u.gitUser }
func (u *UserOsEnv) GetGoRoot() string              { return u.goEnv.GoRoot() }
func (u *UserOsEnv) GetGoEnv() *goConfig            { return u.goEnv }
func (u *UserOsEnv) GetOsProperties() *osProperties { return u.osProperties }

func (u *UserOsEnv) PrintUserOsEnv() {
	termutil.CreateTable(u.GetOsProperties().GetOsInfo().ToMapString(), "OS Env")
	termutil.CreateTable(u.GetOsProperties().ToMapString(), "User Env")
	termutil.CreateTable(u.GetGitUser().ToMapString(), "Git Env")
	termutil.CreateTable(u.GetGoEnv().ToMapString(), "Go Env")
}
