package main

import (
	"bytes"
	"fmt"
	"github.com/getcouragenow/bootstrap/tool/protokit"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	plugin "google.golang.org/protobuf/types/pluginpb"

	"log"
	"strings"
	"text/template"
	"unicode"
)

type (
	protoapi struct {
		params   commandLineParams
		output   *bytes.Buffer
		Messages []*Message
		mnames   []string
		comments *protokit.Comment
		Fname    string
		Name     string
	}

	Message struct {
		Name   string
		Fields []field
		Label  descriptorpb.FieldDescriptorProto_Label
	}

	field struct {
		JsonKey string
		Key     string
		Type    string
		Label   descriptorpb.FieldDescriptorProto_Label
	}
)

func newGenerator(params commandLineParams) *protoapi {
	return &protoapi{
		params:   params,
		Messages: []*Message{},
		output:   bytes.NewBuffer(nil),
	}
}

func (p *protoapi) Generate(in *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	resp := new(plugin.CodeGeneratorResponse)

	p.scanAllMessages(in, resp)
	p.GenerateGoDefault(in, resp)

	return resp
}

func (p *protoapi) scanAllMessages(req *plugin.CodeGeneratorRequest, _ *plugin.CodeGeneratorResponse) {
	descriptors := protokit.ParseCodeGenRequest(req)
	for _, d := range descriptors {
		p.scanMessages(d)
	}
}

func (p *protoapi) scanMessages(d *protokit.FileDescriptor) {
	for _, md := range d.GetMessages() {
		p.scanMessage(md)
	}
}

func (p *protoapi) scanMessage(md *protokit.Descriptor) {
	for _, smd := range md.GetMessages() {
		p.scanMessage(smd)
	}
	{
		fields := make([]field, len(md.GetMessageFields()))
		maps := make(map[string]*descriptorpb.DescriptorProto)
		for _, nt := range md.NestedType {
			if nt.Options.GetMapEntry() {
				name := nt.GetName()
				log.Println(name)
				maps[name] = nt
			}
		}

		for i, fd := range md.GetMessageFields() {
			typeName := fd.GetTypeName()
			if typeName == "" {
				typeName = fd.GetType().String()
			}

			f := field{
				Type:    typeName,
				Key:     fd.GetName(),
				JsonKey: fd.GetJsonName(),
			}
			fields[i] = f
		}
		p.Messages = append(p.Messages, &Message{
			Name:   md.GetName(),
			Fields: fields,
		})
	}
}

func (f field) isRepeated() bool {
	return f.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED
}

func (p *protoapi) GenerateGoDefault(req *plugin.CodeGeneratorRequest, resp *plugin.CodeGeneratorResponse) {
	descriptors := protokit.ParseCodeGenRequest(req)

	for _, d := range descriptors {
		p.Fname = d.GetName()
		p.Name = d.GetPackage()
		p.generateGoDefault()
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s.default.go", p.Name)),
			Content: proto.String(p.output.String()),
		})
	}
}

func (p *protoapi) generateGoDefault() {
	fmap := template.FuncMap{
		"toTitle":      strings.Title,
		"toPascalCase": toPascalCase,
		"yamlTpl":      getYamlTpl,
		"shellTpl":     getShellTpl,
	}
	t, err := template.New(p.Fname).Funcs(fmap).Parse(gotpl)
	if err != nil {
		panic(fmt.Sprintf("Error parsing default template: %v", err))
	}
	if err := t.Execute(p.output, p); err != nil {
		panic(fmt.Sprintf("Should be able to marshal input to go output: %v", err))
	}
}

func isWhitespace(r rune) bool {
	return r == '\t' || r == '\n' || r == ' ' || r == '\r'
}

func toPascalCase(s string) string {
	prev := ' '
	result := strings.Map(
		func(r rune) rune {
			if isWhitespace(prev) || '_' == prev || '-' == prev {
				prev = r
				return unicode.ToTitle(r)
			} else if isWhitespace(r) || '_' == r || '-' == r {
				prev = r
				return -1
			} else {
				prev = r
				return unicode.ToLower(r)

			}
		},
		s)
	return result
}

func getYamlTpl(_ string) string {
	return yamlTmpl
}

func getShellTpl(_ string) string {
	return shellTpl
}
