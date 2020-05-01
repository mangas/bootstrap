package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getcouragenow/bootstrap/tool/protokit"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	plugin "google.golang.org/protobuf/types/pluginpb"
	"log"
	"strings"
)

type (
	protoapi struct {
		params   commandLineParams
		output   *bytes.Buffer
		messages []*message
		mnames   []string
		comments *protokit.Comment
		fname    string
		name     string
	}

	message struct {
		Name   string
		Fields []field
		Label  descriptorpb.FieldDescriptorProto_Label
	}

	DefConfig struct {
		AppConfig []Component `json:"appConfig" yaml:"appConfig"`
	}

	Component struct {
		Name   string                 `json:"componentName" yaml:"componentName"`
		Config map[string]interface{} `json:"config" yaml:"config"`
	}

	field struct {
		Key   string
		Value interface{}
		Type  string
		Label descriptorpb.FieldDescriptorProto_Label
	}
)

func newGenerator(params commandLineParams) *protoapi {
	return &protoapi{
		params:   params,
		messages: []*message{},
		mnames:   []string{},
		output:   bytes.NewBuffer(nil),
	}
}

func (p *protoapi) newDefConfig() *DefConfig {
	var cs []Component
	for _, m := range p.messages {
		fields := map[string]interface{}{}
		for _, f := range m.Fields {
			if f.Key != "" && f.Value != nil {
				fields[f.Key] = f.Value
			}
		}
		newComp := &Component{
			Name:   m.Name,
			Config: fields,
		}
		cs = append(cs, *newComp)
	}
	return &DefConfig{
		AppConfig: cs,
	}
}

func (p *protoapi) Generate(in *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	resp := new(plugin.CodeGeneratorResponse)

	p.scanAllMessages(in, resp)
	p.GenerateJsonDefault(in, resp)

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
				Type:  typeName,
				Key:   fd.GetJsonName(),
				Value: getFieldValue(fd.GetType().String(), fd.GetName(), fd.GetComments().Trailing),
			}
			fields[i] = f
		}
		p.messages = append(p.messages, &message{
			Name:   md.GetName(),
			Fields: fields,
		})
	}
}

func getFieldValue(s, kname, comment string) interface{} {
	switch s {
	case "TYPE_STRING":
		switch kname {
		case "email_val":
			return "test@example.com"
		default:
			return "REPLACE_THIS"
		}
	case "TYPE_BYTES":
		switch kname {
		case "cidr_val":
			return "IPv4_or_IPv6_here"
		default:
			return "bytes_value_here"
		}
	case "TYPE_DOUBLE", "TYPE_FLOAT":
		return 0.0
	case "TYPE_BOOL":
		return false
	case "TYPE_INT64", "TYPE_UINT64", "TYPE_INT32", "TYPE_UINT32":
		return 0
	case "TYPE_MESSAGE":
		return getNestedMessageValue(comment)
	default:
		return ""
	}
}

func getNestedMessageValue(s string) interface{} {
	typ := strings.Split(s, " ")
	if typ[0] == "repeated" {
		switch typ[len(typ)-1] {
		case "string":
			return `[REPLACE_THIS, REPLACE_THIS]`
		case "int", "uint64", "uint32", "int64", "int32":
			return [2]uint64{0, 0}
		case "double", "float":
			return [2]float64{0.0, 0.0}
		case "bool":
			return [2]bool{false, false}
		case "cidr":
			return [2]string{"127.0.0.1", "http://fe80::1ff:fe23:4567:890a"}
		case "email":
			return [2]string{"test@example.com", "winwisely268@example.com"}
		default:
			return `["", ""]`
		}
	}
	switch typ[len(typ)-1] {
	case "string":
		return `REPLACE_THIS`
	case "int", "uint64", "uint32", "int64", "int32":
		return 0
	case "double", "float":
		return 0.0
	case "bool":
		return false
	case "cidr":
		return "127.0.0.1"
	case "email":
		return "test@example.com"
	case "bytes":
		return "this is of type bytes"
	default:
		return `""`
	}
}

func (f field) isRepeated() bool {
	return f.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED
}

func (p *protoapi) GenerateJsonDefault(req *plugin.CodeGeneratorRequest, resp *plugin.CodeGeneratorResponse) {
	descriptors := protokit.ParseCodeGenRequest(req)

	for _, d := range descriptors {
		p.fname = d.GetName()
		p.name = d.GetPackage()
		defConfig := p.newDefConfig()
		defConfig.generateJsonDefault(p.output)
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s.%s.default.json", p.name, p.params.outfile)),
			Content: proto.String(p.output.String()),
		})
	}
}

func (dc *DefConfig) generateJsonDefault(output *bytes.Buffer) {
	jsonOut, err := json.MarshalIndent(dc, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = output.Write(jsonOut)
	if err != nil {
		log.Fatalf("fatal: should be able to write to output: %v", err)
	}
}
