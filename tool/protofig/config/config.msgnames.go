package config
	
import (
	"google.golang.org/protobuf/reflect/protoreflect"
)	
	
var MessageNames = []string{
	"AppComponent",
	"MinioComponent",
	"MaintemplateComponent",
	"GcpComponent",
	"JwtComponent",
	"WorkflowComponent",
	
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
	case "AppComponent":
		return createEmptyAppComponent
	case "MinioComponent":
		return createEmptyMinioComponent
	case "MaintemplateComponent":
		return createEmptyMaintemplateComponent
	case "GcpComponent":
		return createEmptyGcpComponent
	case "JwtComponent":
		return createEmptyJwtComponent
	case "WorkflowComponent":
		return createEmptyWorkflowComponent
	
	default:
		return nil
	}
}

func createEmptyAppComponent() protoreflect.MessageType {
	comp := &AppComponent{}
	return comp.ProtoReflect().Type()
}
func createEmptyMinioComponent() protoreflect.MessageType {
	comp := &MinioComponent{}
	return comp.ProtoReflect().Type()
}
func createEmptyMaintemplateComponent() protoreflect.MessageType {
	comp := &MaintemplateComponent{}
	return comp.ProtoReflect().Type()
}
func createEmptyGcpComponent() protoreflect.MessageType {
	comp := &GcpComponent{}
	return comp.ProtoReflect().Type()
}
func createEmptyJwtComponent() protoreflect.MessageType {
	comp := &JwtComponent{}
	return comp.ProtoReflect().Type()
}
func createEmptyWorkflowComponent() protoreflect.MessageType {
	comp := &WorkflowComponent{}
	return comp.ProtoReflect().Type()
}

