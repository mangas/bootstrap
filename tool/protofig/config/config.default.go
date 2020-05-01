package config
	
import (
	b64 "encoding/base64"
	"errors"
	"google.golang.org/protobuf/encoding/protojson"
	nurl "net/url"
	"regexp"
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
	configYamlTpl = `apiVersion: v1
kind: Secret
metadata:
  name: {{ .Name }}
type: Opaque
data:
  {{ range .Fields }}{{ .Key }}: {{ .Value | toB64 }}
  {{ end }}	
`
)
	
func isUrl(str string) bool {
	u, err := nurl.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
	
func isEmail(str string) bool {
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
	
func (c *Component) ConfigCreateJSONMessage() ([]byte, error) {
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
	
func isSecret(s string) bool {
	rxSecret := regexp.MustCompile("[k|K]ey")
	return rxSecret.MatchString(s)
}
	
func toConfigVal(i interface{}) (cfg *ConfigVal, err error) {
	switch i.(type) {
	case string:
		if isUrl(i.(string)) {
			cfg = &ConfigVal{Val: &ConfigVal_CidrVal{CidrVal: []byte(i.(string))}}
			if err = cfg.Validate(); err != nil {
				return nil, err
			}
			return cfg, nil
		} else if isEmail(i.(string)) {
			cfg = &ConfigVal{Val: &ConfigVal_EmailVal{EmailVal: i.(string)}}
			if err = cfg.Validate(); err != nil {
				return nil, err
			}
			return cfg, nil
		} else {
			if isSecret(i.(string)) {
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


func (x *MinioComponent) MarshalYAML() ([]byte, error) {
	type KVal struct {
		Name string
		
	}
}

func (x *MaintemplateComponent) MarshalYAML() ([]byte, error) {
	return nil, false
}
func (x *GcpComponent) MarshalYAML() ([]byte, error) {
	return nil, false
}
func (x *JwtComponent) MarshalYAML() ([]byte, error) {
	return nil, false
}
func (x *WorkflowComponent) MarshalYAML() ([]byte, error) {
	return nil, false
}

	
func (x *MinioComponent) UnmarshalYAML(b []byte) error {
	return nil, false
}
func (x *MaintemplateComponent) UnmarshalYAML(b []byte) error {
	return nil, false
}
func (x *GcpComponent) UnmarshalYAML(b []byte) error {
	return nil, false
}
func (x *JwtComponent) UnmarshalYAML(b []byte) error {
	return nil, false
}
func (x *WorkflowComponent) UnmarshalYAML(b []byte) error {
	return nil, false
}

	
