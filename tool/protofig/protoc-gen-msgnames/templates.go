package main

var (
	gotpl = `package {{ .Name }}
	
import (
	b64 "encoding/base64"
	"errors"
	"google.golang.org/protobuf/encoding/protojson"
	nurl "net/url"
	"regexp"
)	
	
type DefConfig struct {
	AppConfig []Component ` + "`" + `json:"appConfig" yaml:"appConfig"` + "`" + `
}

type Component struct {
	Name   string                 ` + "`" + `json:"componentName" yaml:"appConfig"` + "`" + `
	Config map[string]interface{} ` + "`" + `json:"config" yaml:"config"` + "`" + ` 
}	

var (	
	{{ .Name }}MessageNames = []string{
		{{ range .Messages }}"{{ .Name }}",
		{{ end }}
	}
	{{ .Name }}YamlTpl = ` + "`" + `{{ "" | yamlTpl }}` + "`" + `
)
	
func isUrl(str string) bool {
	u, err := nurl.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
	
func isEmail(str string) bool {
	rxEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return rxEmail.MatchString(str)
}
	
func {{ toTitle .Name }}HasMessageName(s string) bool {
	for _, mn := range {{ .Name }}MessageNames {
	    if s == mn {
			return true
		}
	}
	return false
}
	
func (c *Component) {{ toTitle .Name }}CreateJSONMessage() ([]byte, error) {
	switch c.Name {
	{{ range .Messages}}case "{{ .Name }}":
		msg, err := c.Create{{ .Name }}()
		if err != nil { return nil, err }
		return msg.MarshalJSON()
	{{ end }}
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
{{ range .Messages }}func (c *Component) Create{{ .Name }}() (*{{ .Name }}, error) {
	{{ range .Fields }}{{ .JsonKey }}, err := toConfigVal(c.Config["{{ .JsonKey }}"])
	if err != nil {
		return nil, err
	}
	{{ end }}
	return &{{ .Name }}{
	    {{ range .Fields }}{{ .Key | toPascalCase }}: {{ .JsonKey }},
	    {{ end }}
	}, nil
}
{{ end }}	
	
{{ range .Messages }}func (x *{{ .Name }}) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		Multiline: true,
		AllowPartial: true,
	}
	return opt.Marshal(x)
}
{{ end }}
	
{{ range .Messages }}func (x *{{ .Name }}) UnmarshalJSON(b []byte) error {
	opt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return opt.Unmarshal(b, x)
}
{{ end }}
	
{{ range .Messages }}func (x *{{ .Name }}) MarshalYAMLSecret() ([]byte, error) {
	
}
{{ end }}
`
	yamlTmpl = `apiVersion: v1
kind: Secret
metadata:
  name: {{ .Name }}
type: Opaque
data:
  {{ range .Fields }}{{ .Key }}: {{ .Value | toB64 }}
  {{ end }}	
`
)
