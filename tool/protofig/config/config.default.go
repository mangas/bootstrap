package config
	
import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"google.golang.org/protobuf/encoding/protojson"
	nurl "net/url"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)	
	
type DefConfig struct {
	AppConfig []Component `json:"appConfig" yaml:"appConfig"`
}

type Component struct {
	Name   string                 `json:"componentName" yaml:"appConfig"`
	Config map[string]interface{} `json:"config" yaml:"config"` 
}	

var (	
	configMessageNames = []string{
		"MinioComponent",
		"MaintemplateComponent",
		"GcpComponent",
		"JwtComponent",
		"WorkflowComponent",
		
	}
	configShellTpl =  `
# {{ .Name }}	
{{ range $k, $v := .Config }}{{ $k | toUpperCaseSnake }} = "{{ $v }}"
{{ end }}
export {{ range $k, $v := .Config}}{{ $k | toUpperCaseSnake }} {{ end }}
` 
	configYamlTpl =  `apiVersion: v1
kind: Secret
metadata:
  name: {{ .Name | toCamelCase }}-secret
type: Opaque
data:
  {{ range $k, $v := .Config }}{{if eq ($v | typeOf ) "string"}}{{ $k }}: {{ $v | toB64 }}{{end}}
  {{ end }}` 
	configNumberSequence    = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
	configNumberReplacement = []byte("$1 $2 $3")
	configUppercaseAcronym = map[string]bool{
		"ID": true,
	}
)
	
func configIsUrl(str string) bool {
	u, err := nurl.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
	
func configIsEmail(str string) bool {
	rxEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return rxEmail.MatchString(str)
}
	
func ConfigHasMessageName(s string) bool {
	for _, mn := range configMessageNames {
	    if s == mn {
			return true
		}
	}
	return false
}
	
func (c *Component) ConfigToJSON() ([]byte, error) {
	switch c.Name {
	case "MinioComponent":
		msg, err := c.CreateMinioComponent()
		if err != nil { return nil, err }
		return msg.MarshalJSON()
	case "MaintemplateComponent":
		msg, err := c.CreateMaintemplateComponent()
		if err != nil { return nil, err }
		return msg.MarshalJSON()
	case "GcpComponent":
		msg, err := c.CreateGcpComponent()
		if err != nil { return nil, err }
		return msg.MarshalJSON()
	case "JwtComponent":
		msg, err := c.CreateJwtComponent()
		if err != nil { return nil, err }
		return msg.MarshalJSON()
	case "WorkflowComponent":
		msg, err := c.CreateWorkflowComponent()
		if err != nil { return nil, err }
		return msg.MarshalJSON()
	
	default:
		return nil, errors.New("component name is unknown")
	}
}
	
func (c *Component) ConfigToShellEnv() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	component, err := func() (*Component, error) {
		switch c.Name {
		case "MinioComponent":
			msg, err := c.CreateMinioComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "MaintemplateComponent":
			msg, err := c.CreateMaintemplateComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "GcpComponent":
			msg, err := c.CreateGcpComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "JwtComponent":
			msg, err := c.CreateJwtComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "WorkflowComponent":
			msg, err := c.CreateWorkflowComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		
		default:
			return nil, errors.New("component name is unknown")
		}
	}()
	if err != nil { return nil, err }
	shellTemplate := configShellTpl
	fmap := template.FuncMap{
		"toUpperCaseSnake": configToUpperCaseSnake,
	}
	t, err := template.New(component.Name).Funcs(fmap).Parse(shellTemplate)
	if err != nil {
			return nil, err
		}
	err = t.Execute(b, component)
	return b.Bytes(), err
}
	
func (c *Component) ConfigToK8sSecret() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	component, err := func() (*Component, error) {
		switch c.Name {
		case "MinioComponent":
			msg, err := c.CreateMinioComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "MaintemplateComponent":
			msg, err := c.CreateMaintemplateComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "GcpComponent":
			msg, err := c.CreateGcpComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "JwtComponent":
			msg, err := c.CreateJwtComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		case "WorkflowComponent":
			msg, err := c.CreateWorkflowComponent()
			if err != nil { return nil, err }
			return msg.ToComponent()
		
		default:
			return nil, errors.New("component name is unknown")
		}
	}()
	if err != nil { return nil, err }
	yamlTemplate := configYamlTpl
	fmap := template.FuncMap{
		"toB64": configToB64,
		"typeOf": configTypeOf,
		"toCamelCase": configToCamelCase,
	}
	t, err := template.New(component.Name).Funcs(fmap).Parse(yamlTemplate)
	if err != nil {
			return nil, err
		}
	err = t.Execute(b, component)
	return b.Bytes(), err
}
	
func configIsSecret(s string) bool {
	rxSecret := regexp.MustCompile("[k|K]ey")
	return rxSecret.MatchString(s)
}
	
func(c *ConfigVal) FromConfigVal() (interface{}, error) {
	protoTyp := strings.Split(c.String(), ":")[0]
	switch protoTyp {
	case "string_val":
		return c.GetStringVal(), nil
	case "uint64_val":
		return c.GetUint64Val(), nil
	case "double_val":
		return c.GetDoubleVal(), nil
	case "bool_val":
		return c.GetBoolVal(), nil
	case "cidr_val":
		return c.GetCidrVal(), nil
	case "email_val":
		return c.GetEmailVal(), nil
	default:
		return nil, errors.New("unknown value")
	}
}

func toConfigVal(i interface{}) (cfg *ConfigVal, err error) {
	switch i.(type) {
	case string:
		if configIsUrl(i.(string)) {
			cfg = &ConfigVal{Val: &ConfigVal_CidrVal{CidrVal: []byte(i.(string))}}
			if err = cfg.Validate(); err != nil {
				return nil, err
			}
			return cfg, nil
		} else if configIsEmail(i.(string)) {
			cfg = &ConfigVal{Val: &ConfigVal_EmailVal{EmailVal: i.(string)}}
			if err = cfg.Validate(); err != nil {
				return nil, err
			}
			return cfg, nil
		} else {
			if configIsSecret(i.(string)) {
				cfg = &ConfigVal{Val: &ConfigVal_StringVal{StringVal: b64.StdEncoding.EncodeToString([]byte(i.(string)))}}
				if err = cfg.Validate(); err != nil {
					return nil, err
				}
			return cfg, nil	
			}
		    cfg = &ConfigVal{Val: &ConfigVal_StringVal{StringVal: i.(string)}}
			if err = cfg.Validate(); err != nil {
				return nil, err
			}
			return cfg, nil
		}

	case uint32, uint64, int, int32:
	    cfg := &ConfigVal{Val: &ConfigVal_Uint64Val{Uint64Val: i.(uint64)}}
		if err = cfg.Validate(); err != nil {
			return nil, err
		}
		return cfg, nil
	case bool:
	    cfg := &ConfigVal{Val: &ConfigVal_BoolVal{BoolVal: i.(bool)}}
		if err = cfg.Validate(); err != nil {
			return nil, err
		}
		return cfg, nil	
	case float64, float32:
		cfg := &ConfigVal{Val: &ConfigVal_DoubleVal{DoubleVal: i.(float64)}}
		if err = cfg.Validate(); err != nil {
			return nil, err
		}
		return cfg, nil
	default:
		return nil, errors.New("Unknown value")
	}
}
func (c *Component) CreateMinioComponent() (*MinioComponent, error) {
	minioAccesskey, err := toConfigVal(c.Config["minioAccesskey"])
	if err != nil {
		return nil, err
	}
	minioSecretkey, err := toConfigVal(c.Config["minioSecretkey"])
	if err != nil {
		return nil, err
	}
	
	return &MinioComponent{
	    MinioAccesskey: minioAccesskey,
	    MinioSecretkey: minioSecretkey,
	    
	}, nil
}
func (c *Component) CreateMaintemplateComponent() (*MaintemplateComponent, error) {
	minioAccesskey, err := toConfigVal(c.Config["minioAccesskey"])
	if err != nil {
		return nil, err
	}
	minioSecretkey, err := toConfigVal(c.Config["minioSecretkey"])
	if err != nil {
		return nil, err
	}
	minioLocation, err := toConfigVal(c.Config["minioLocation"])
	if err != nil {
		return nil, err
	}
	minioTimeout, err := toConfigVal(c.Config["minioTimeout"])
	if err != nil {
		return nil, err
	}
	minioSsl, err := toConfigVal(c.Config["minioSsl"])
	if err != nil {
		return nil, err
	}
	minioEnckey, err := toConfigVal(c.Config["minioEnckey"])
	if err != nil {
		return nil, err
	}
	minioEndpoint, err := toConfigVal(c.Config["minioEndpoint"])
	if err != nil {
		return nil, err
	}
	
	return &MaintemplateComponent{
	    MinioAccesskey: minioAccesskey,
	    MinioSecretkey: minioSecretkey,
	    MinioLocation: minioLocation,
	    MinioTimeout: minioTimeout,
	    MinioSsl: minioSsl,
	    MinioEnckey: minioEnckey,
	    MinioEndpoint: minioEndpoint,
	    
	}, nil
}
func (c *Component) CreateGcpComponent() (*GcpComponent, error) {
	gcpUser, err := toConfigVal(c.Config["gcpUser"])
	if err != nil {
		return nil, err
	}
	gcpProject, err := toConfigVal(c.Config["gcpProject"])
	if err != nil {
		return nil, err
	}
	gkeCluster, err := toConfigVal(c.Config["gkeCluster"])
	if err != nil {
		return nil, err
	}
	gkeZone, err := toConfigVal(c.Config["gkeZone"])
	if err != nil {
		return nil, err
	}
	gkeEmail, err := toConfigVal(c.Config["gkeEmail"])
	if err != nil {
		return nil, err
	}
	
	return &GcpComponent{
	    GcpUser: gcpUser,
	    GcpProject: gcpProject,
	    GkeCluster: gkeCluster,
	    GkeZone: gkeZone,
	    GkeEmail: gkeEmail,
	    
	}, nil
}
func (c *Component) CreateJwtComponent() (*JwtComponent, error) {
	privateKey, err := toConfigVal(c.Config["privateKey"])
	if err != nil {
		return nil, err
	}
	publicKey, err := toConfigVal(c.Config["publicKey"])
	if err != nil {
		return nil, err
	}
	
	return &JwtComponent{
	    PrivateKey: privateKey,
	    PublicKey: publicKey,
	    
	}, nil
}
func (c *Component) CreateWorkflowComponent() (*WorkflowComponent, error) {
	githubSha, err := toConfigVal(c.Config["githubSha"])
	if err != nil {
		return nil, err
	}
	githubRef, err := toConfigVal(c.Config["githubRef"])
	if err != nil {
		return nil, err
	}
	project, err := toConfigVal(c.Config["project"])
	if err != nil {
		return nil, err
	}
	registryHostname, err := toConfigVal(c.Config["registryHostname"])
	if err != nil {
		return nil, err
	}
	url, err := toConfigVal(c.Config["url"])
	if err != nil {
		return nil, err
	}
	locales, err := toConfigVal(c.Config["locales"])
	if err != nil {
		return nil, err
	}
	flutterChannel, err := toConfigVal(c.Config["flutterChannel"])
	if err != nil {
		return nil, err
	}
	releaseChannel, err := toConfigVal(c.Config["releaseChannel"])
	if err != nil {
		return nil, err
	}
	
	return &WorkflowComponent{
	    GithubSha: githubSha,
	    GithubRef: githubRef,
	    Project: project,
	    RegistryHostname: registryHostname,
	    Url: url,
	    Locales: locales,
	    FlutterChannel: flutterChannel,
	    ReleaseChannel: releaseChannel,
	    
	}, nil
}
	
	
func (x *MinioComponent) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		Multiline: true,
		AllowPartial: true,
	}
	return opt.Marshal(x)
}
func (x *MaintemplateComponent) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		Multiline: true,
		AllowPartial: true,
	}
	return opt.Marshal(x)
}
func (x *GcpComponent) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		Multiline: true,
		AllowPartial: true,
	}
	return opt.Marshal(x)
}
func (x *JwtComponent) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		Multiline: true,
		AllowPartial: true,
	}
	return opt.Marshal(x)
}
func (x *WorkflowComponent) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		Multiline: true,
		AllowPartial: true,
	}
	return opt.Marshal(x)
}

	
func (x *MinioComponent) UnmarshalJSON(b []byte) error {
	opt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return opt.Unmarshal(b, x)
}
func (x *MaintemplateComponent) UnmarshalJSON(b []byte) error {
	opt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return opt.Unmarshal(b, x)
}
func (x *GcpComponent) UnmarshalJSON(b []byte) error {
	opt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return opt.Unmarshal(b, x)
}
func (x *JwtComponent) UnmarshalJSON(b []byte) error {
	opt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return opt.Unmarshal(b, x)
}
func (x *WorkflowComponent) UnmarshalJSON(b []byte) error {
	opt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return opt.Unmarshal(b, x)
}


func (x *MinioComponent) ToComponent() (*Component, error) {
	var err error
	msgName := "MinioComponent"
	fields := make(map[string]interface{})
	
	fields["minioAccesskey"], err = x.MinioAccesskey.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioSecretkey"], err = x.MinioSecretkey.FromConfigVal()
	if err != nil { return nil, err }
	
	return &Component{
		Name: msgName,
		Config: fields,
	}, nil
}
func (x *MaintemplateComponent) ToComponent() (*Component, error) {
	var err error
	msgName := "MaintemplateComponent"
	fields := make(map[string]interface{})
	
	fields["minioAccesskey"], err = x.MinioAccesskey.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioSecretkey"], err = x.MinioSecretkey.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioLocation"], err = x.MinioLocation.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioTimeout"], err = x.MinioTimeout.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioSsl"], err = x.MinioSsl.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioEnckey"], err = x.MinioEnckey.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["minioEndpoint"], err = x.MinioEndpoint.FromConfigVal()
	if err != nil { return nil, err }
	
	return &Component{
		Name: msgName,
		Config: fields,
	}, nil
}
func (x *GcpComponent) ToComponent() (*Component, error) {
	var err error
	msgName := "GcpComponent"
	fields := make(map[string]interface{})
	
	fields["gcpUser"], err = x.GcpUser.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["gcpProject"], err = x.GcpProject.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["gkeCluster"], err = x.GkeCluster.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["gkeZone"], err = x.GkeZone.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["gkeEmail"], err = x.GkeEmail.FromConfigVal()
	if err != nil { return nil, err }
	
	return &Component{
		Name: msgName,
		Config: fields,
	}, nil
}
func (x *JwtComponent) ToComponent() (*Component, error) {
	var err error
	msgName := "JwtComponent"
	fields := make(map[string]interface{})
	
	fields["privateKey"], err = x.PrivateKey.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["publicKey"], err = x.PublicKey.FromConfigVal()
	if err != nil { return nil, err }
	
	return &Component{
		Name: msgName,
		Config: fields,
	}, nil
}
func (x *WorkflowComponent) ToComponent() (*Component, error) {
	var err error
	msgName := "WorkflowComponent"
	fields := make(map[string]interface{})
	
	fields["githubSha"], err = x.GithubSha.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["githubRef"], err = x.GithubRef.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["project"], err = x.Project.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["registryHostname"], err = x.RegistryHostname.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["url"], err = x.Url.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["locales"], err = x.Locales.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["flutterChannel"], err = x.FlutterChannel.FromConfigVal()
	if err != nil { return nil, err }
	
	fields["releaseChannel"], err = x.ReleaseChannel.FromConfigVal()
	if err != nil { return nil, err }
	
	return &Component{
		Name: msgName,
		Config: fields,
	}, nil
}
	
		
func configToB64(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}
	
func configTypeOf(i interface{}) string {
	return reflect.TypeOf(i).String()
}	
	
// Code below is
// taken from https://github.com/iancoleman/strcase

func configAddWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = configNumberSequence.ReplaceAll(b, configNumberReplacement)
	return string(b)
}


func configToUpperCaseSnake(s string) string {
	return configToUpperCaseDelimited(s, '_', 0, true)
}

func configToUpperCaseDelimited(s string, delimiter uint8, ignore uint8, upper bool) string {
	s = configAddWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			vIsCap := v >= 'A' && v <= 'Z'
			vIsLow := v >= 'a' && v <= 'z'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			if (vIsCap && nextIsLow) || (vIsLow && nextIsCap) {
				nextCaseIsChanged = true
			}
			if ignore > 0 && i-1 >= 0 && s[i-1] == ignore && nextCaseIsChanged {
				nextCaseIsChanged = false
			}
		}

		if i > 0 && n[len(n)-1] != delimiter && nextCaseIsChanged {
			// add underscore if next letter case type is changed
			if v >= 'A' && v <= 'Z' {
				n += string(delimiter) + string(v)
			} else if v >= 'a' && v <= 'z' {
				n += string(v) + string(delimiter)
			}
		} else if v == ' ' || v == '_' || v == '-' {
			// replace spaces/underscores with delimiters
			if uint8(v) == ignore {
				n += string(v)
			} else {
				n += string(delimiter)
			}
		} else {
			n = n + string(v)
		}
	}

	if upper {
		n = strings.ToUpper(n)
	} else {
		n = strings.ToLower(n)
	}
	return n
}
	

// Converts a string to CamelCase
func configToCamelInitCase(s string, initCase bool) string {
	s = configAddWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}	
	
// configToCamelCase converts a string to lowerCamelCase
func configToCamelCase(s string) string {
	if s == "" {
		return s
	}
	if configUppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return configToCamelInitCase(s, false)
}	
