package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type TypeMap map[string]map[string]interface{}
var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func main() {
	rand.Seed(time.Now().UnixNano())

	filename:= "../../../packages/mod-chat/server/pkg/api/service.pb.go"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fs := token.NewFileSet()
	astFile, err := parser.ParseFile(fs, "", bs, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	tm, err := handleFile(astFile)
	if err != nil {
		panic(err)
	}

	for k, v := range tm {
		f, err := os.Create(fmt.Sprintf("./output/data.%s.json", k))
		if err != nil {
			panic(err)
		}

		if err := json.NewEncoder(f).Encode(v); err != nil {
			panic(err)
		}

		_ = f.Close()
	}
}

func handleFile(f *ast.File) (TypeMap, error) {
	tm := make(TypeMap)

	for _, d := range f.Scope.Objects {
		ts, ok := d.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}

		var isProtoMessage bool
		for _, f := range st.Fields.List {
			if f.Tag == nil || f.Tag.Value == "" {
				continue
			}

			if strings.Contains(f.Tag.Value, "protobuf:") {
				isProtoMessage = true
				break
			}
		}
		if !isProtoMessage {
			continue
		}

		m, err := makeStruct(st)
		if err != nil {
			return nil, err
		}

		tm[d.Name] = m
	}

	return tm, nil
}

func makeStruct(st *ast.StructType) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		key := field.Names[0].Name
		if strings.HasPrefix(key, "XXX") {
			continue
		}

		if se, ok := field.Type.(*ast.StarExpr); ok {
			x, ok := se.X.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf("expected X to be Ident but was %T", se.X)
			}

			ts, ok := x.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				return nil, fmt.Errorf("expected ts to be Ident but was %T", x.Obj.Decl)
			}

			if st, ok := ts.Type.(*ast.StructType); ok {
				mm, err := makeStruct(st)
				if err != nil {
					return nil, fmt.Errorf("unable to parse inner struct %s, %w", key, err)
				}
				m[key] = mm
				continue
			}

		}

		id, ok := field.Type.(*ast.Ident)
		if !ok {
			fmt.Printf("expecting ident but got: %T\n", field.Type)
			continue
		}

		switch id.Name {
		case "int32":
			m[key] = rand.Int31n(1<<30)
		case "string":
			length := rand.Intn(100)
			var b strings.Builder
			for i := 0; i < length; i++ {
				b.WriteRune(chars[rand.Intn(len(chars))])
			}
			m[key] = b.String()
		case "bool":
			m[key] = rand.Int31n(9) >= 5
		default:
			return nil, fmt.Errorf("unimplemented type: %s", id.Name)
		}
	}

	return m, nil
}
