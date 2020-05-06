module github.com/getcouragenow/bootstrap/tool/protofig/protoc-gen-msgnames

go 1.13

require (
	github.com/getcouragenow/bootstrap/tool/protokit v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.0
	google.golang.org/protobuf v1.21.0
)

replace github.com/getcouragenow/bootstrap/tool/protokit => ../../protokit/
