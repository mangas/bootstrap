package main

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
	"google.golang.org/protobuf/types/descriptorpb"
	"text/template"
)

type (
	prontomn struct {
		params   commandLineParams
		output   *bytes.Buffer
		messages map[string]*message
		mnames   []string
		comments *protokit.Comment
		// package name
		name string
	}

	message struct {
		Name  string
		Label descriptorpb.FieldDescriptorProto_Label
	}

	MessageNames struct {
		PackageName string
		Names       []string
	}
)

func newMessageNames(packageName string, names []string) *MessageNames {
	return &MessageNames{
		PackageName: packageName,
		Names:       names,
	}
}

func newGenerator(params commandLineParams) *prontomn {
	t := &prontomn{
		params:   params,
		messages: map[string]*message{},
		output:   bytes.NewBuffer(nil),
	}

	return t
}

func (p *prontomn) Generate(in *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	resp := new(plugin.CodeGeneratorResponse)

	p.scanAllMessages(in, resp)
	p.GenerateMessageNames(in, resp)

	return resp
}

// P forwards to g.gen.P, which prints output.
func (p *prontomn) P(args ...string) {
	for _, v := range args {
		p.output.WriteString(v)
	}
	p.output.WriteByte('\n')
}

func (p *prontomn) scanAllMessages(req *plugin.CodeGeneratorRequest, resp *plugin.CodeGeneratorResponse) {
	descriptors := protokit.ParseCodeGenRequest(req)
	for _, d := range descriptors {
		p.scanMessages(d)
	}
}

func (p *prontomn) GenerateMessageNames(req *plugin.CodeGeneratorRequest, resp *plugin.CodeGeneratorResponse) {
	descriptors := protokit.ParseCodeGenRequest(req)

	for _, d := range descriptors {
		p.name = d.GetPackage()
		for _, m := range d.GetMessages() {
			p.mnames = append(p.mnames, m.GetName())
		}
		msgNames := newMessageNames(p.name, p.mnames)
		p.generateGoFile(msgNames)
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s.msgnames.go", p.name)),
			Content: proto.String(p.output.String()),
		})
	}
}

func (p *prontomn) scanMessages(d *protokit.FileDescriptor) {
	for _, md := range d.GetMessages() {
		p.scanMessage(md)
	}
}

func (p *prontomn) scanMessage(md *protokit.Descriptor) {
	for _, smd := range md.GetMessages() {
		p.scanMessage(smd)
	}
	{
		maps := make(map[string]*descriptorpb.DescriptorProto)
		for _, t := range md.NestedType {
			if t.Options.GetMapEntry() {
				pkg := md.GetPackage()
				name := fmt.Sprintf(".%s.%s.%s", pkg, md.GetName(),
					t.GetName())
				maps[name] = t
			}
		}
		p.messages[md.GetFullName()] = &message{
			Name: md.GetName(),
		}
	}
}

func (p *prontomn) generateGoFile(names *MessageNames) {
	gotpl := `package {{ .PackageName }}
	
import (
	"google.golang.org/protobuf/reflect/protoreflect"
)	
	
var MessageNames = []string{
	{{ range .Names }}"{{ . }}",
	{{ end }}
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
	{{ range .Names }}case "{{ . }}":
		return createEmpty{{ . }}
	{{ end }}
	default:
		return nil
	}
}

{{ range .Names }}func createEmpty{{ . }}() protoreflect.MessageType {
	comp := &{{ . }}{}
	return comp.ProtoReflect().Type()
}
{{ end }}
`
	t := template.New(p.name)
	t, err := t.Parse(gotpl)
	if err != nil {
		panic(fmt.Sprintf("Error parsing default template: %v", err))
	}
	if err := execTemplate(names, t, p.output); err != nil {
		panic(fmt.Sprintf("Should be able to marshal input to go output: %v", err))
	}
}

func execTemplate(names *MessageNames, tpl *template.Template, tout *bytes.Buffer) error {
	return tpl.Execute(tout, names)
}
