package baseproto
	
import (
	"google.golang.org/protobuf/reflect/protoreflect"
)	
	
var MessageNames = []string{
	"Config",
	"ConfigVal",
	
}
	
func HasMessageName(s string) bool {
	for _, mn := range MessageNames {
	    if s == mn {
			return true
		}
	}
	return false
}
	
func CreateMessage(s string) func() protoreflect.MessageType {
	switch s {
	case "Config":
		return createEmptyConfig
	case "ConfigVal":
		return createEmptyConfigVal
	
	default:
		return nil
	}
}

func createEmptyConfig() protoreflect.MessageType {
	comp := &Config{}
	return comp.ProtoReflect().Type()
}
func createEmptyConfigVal() protoreflect.MessageType {
	comp := &ConfigVal{}
	return comp.ProtoReflect().Type()
}

