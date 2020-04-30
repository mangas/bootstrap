# protoc-gen-msgnames

Given a protofile, it will spit out the name of the messages
and puts that in a variable in golang.

### Install

1. Build it using go (`go build`)
2. Install it to your PATH (`install -Dm755 protoc-gen-msgnames ${GOPATH}/bin/`)

### Example Usage

Given a protofile in testoutput directory

Run this:

```
protoc --msgnames_out=./testoutput/ testoutput/test.proto
```

For test.proto like this
```
syntax = "proto3";

package wheelerdealin;

service PlayerTransfer {
  rpc Info(InfoCarReq) returns (InfoCarResp) {};
  rpc Buy(BuyCarReq) returns (BuyCarResp) {};
}

message Err {
  string reason = 1;
}

message InfoCarReq {
  string id = 1;
}

message BuyCarResp {
  bool success = 1;
  Err err_reason = 2;
}

message BuyCarReq {
  bool cash = 1;
  bool credit = 2;
}

message InfoCarResp {
  string id = 1;
  string manufacturer = 2;
  string year = 3;
  double mileage = 4;
  double price = 5;
  Err err_reason = 6;
}
```

It will produce a <PROTO_PACKAGE_NAME>.msgnames.go
which contains:

```
package wheelerdealin

var MessageNames = []string{
	"wheelerdealin.Err",
	"wheelerdealin.InfoCarReq",
	"wheelerdealin.BuyCarResp",
	"wheelerdealin.BuyCarReq",
	"wheelerdealin.InfoCarResp",
	
}
```

