module github.com/getcouragenow/bootstrap/tool/protofig/protoc-gen-msgnames

go 1.14

require (
	github.com/golang/protobuf v1.4.0
	github.com/pseudomuto/protokit v0.2.0
	github.com/stretchr/testify v1.5.1 // indirect
	google.golang.org/protobuf v1.21.0
)

replace github.com/getcouragenow/bootstrap/tool/protofig/protoc-gen-method-names => ./

replace github.com/getcouragenow/bootstrap/tool/protofig/protoc-gen-method-names/parser => ./parser

replace github.com/getcouragenow/bootstrap/tool/protofig/protoc-gen-method-names/generator => ./generator/
