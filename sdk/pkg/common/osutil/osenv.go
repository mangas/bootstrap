package osutil

/*
Package osutil is the common utils for getting os information
*/

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
	"text/tabwriter"
)

type GlobalEnvGetter interface {
	GetName() string
	GetRoot() string
	GetAccount() string
	GetGroups() []*user.Group
	GetOsInfo() OSInfoGetter
}

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
	root := fmt.Sprintf("github.com/%s", userName)
	account, err := runUnixCmd("git", "")
	if err != nil {
		return nil, err
	}
	osInfo, err := getOsInfoGetter()
	if err != nil {
		return nil, err
	}
	return &gitConfig{
		name:    *userName,
		root:    root,
		account: *account,
		osInfo:  osInfo,
	}, nil
}

func (g *gitConfig) GetName() string          { return g.name }
func (g *gitConfig) GetRoot() string          { return g.root }
func (g *gitConfig) GetAccount() string       { return g.account }
func (g *gitConfig) GetGroups() []*user.Group { return nil }
func (g *gitConfig) GetOsInfo() OSInfoGetter  { return g.osInfo }

type UserOsEnv struct {
	osProperties GlobalEnvGetter
	goRoot       string
	goPath       string
	gitUser      GlobalEnvGetter
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
	goroot := runtime.GOROOT()
	gopath := os.Getenv("GOPATH")
	return &UserOsEnv{
		osProperties: osProp,
		gitUser:      gitUser,
		goRoot:       goroot,
		goPath:       gopath,
	}, nil
}

func (u *UserOsEnv) GetGoPath() string                { return u.goPath }
func (u *UserOsEnv) GetGitUser() GlobalEnvGetter      { return u.gitUser }
func (u *UserOsEnv) GetGoRoot() string                { return u.goRoot }
func (u *UserOsEnv) GetOsProperties() GlobalEnvGetter { return u.osProperties }

func (u *UserOsEnv) PrintUserOsEnv() error {
	out := getTabWriterOutput()
	PrintDelimiters(out)
	if _, err := fmt.Fprintf(out, "%s\n", titleColor("OS & USER INFO")); err != nil {
		return err
	}
	PrintDelimiters(out)
	u.prow(out, "OS", u.GetOsProperties().GetOsInfo().GetOsName())
	u.prow(out, "Kernel / Version", u.GetOsProperties().GetOsInfo().GetKernel())
	u.prow(out, "Platform", u.GetOsProperties().GetOsInfo().GetPlatform())
	u.prow(out, "#Cores", u.GetOsProperties().GetOsInfo().GetCores())
	u.prow(out, "#Memory", u.GetOsProperties().GetOsInfo().GetMemory())
	PrintDelimiters(out)
	u.prow(out, "Hostname", u.GetOsProperties().GetOsInfo().GetHostName())
	u.prow(out, "OS Username", u.GetOsProperties().GetAccount())
	var s strings.Builder
	for _, g := range u.GetOsProperties().GetGroups() {
		s.WriteString(g.Name)
		s.WriteRune(',')
		s.WriteRune(' ')
	}
	u.prow(out, "Groups", s.String())
	u.prow(out, "Home Dir", u.GetOsProperties().GetRoot())
	PrintDelimiters(out)
	u.prow(out, "Git User", u.gitUser.GetName())
	u.prow(out, "Git Email", u.gitUser.GetAccount())
	u.prow(out, "Git Root", u.gitUser.GetRoot())
	if err := out.Flush(); err != nil {
		return err
	}
	return nil
}

func (u *UserOsEnv) prow(out *tabwriter.Writer, key string, value interface{}) {
	printRow(out, "", key, value)
}
