# protoc-gen-configdefault

Generate textual config from protobuf. The generated output will be of format JSON.
The produced JSON file will have default value assigned to it based on its primitive files.
If the protobuf message fields has message type, then its default value will be defined according to
the annotated comment in the field descriptor.

For example, see the declaration of baseproto.proto and config.proto in the protofig config dir.

## Install
```
make 
```

## Usage

```
protoc --configdefault_out=outfile=<NAME_OF_FILE>:<OUTPUT_DIR> <INPUT_FILE>

example:
protoc --configdefault_out=outfile=winwisely268:. baseproto.proto
```

For example given these declarations of proto file:

`FILENAME: baseproto.proto`
```
syntax = "proto3";

package config;

option go_package = ".;config";

import "validate/validate.proto";

// Config is basically a map<k,v>
// where k is always a string, and v is defined in the ConfigVal
message Config {
    string key = 1;
    ConfigVal val = 2;
}

message ConfigVal {
    oneof val {
        string string_val = 1;
        uint64 uint64_val = 2;
        string email_val = 3 [(validate.rules).string.email = true];
        bytes cidr_val = 4 [(validate.rules).bytes.ip = true];
        bool bool_val = 5;
    }
}
```

`FILENAME: config.proto`
```
syntax = "proto3";

package config;

option go_package = ".;config";

import "baseproto.proto";

message MinioComponent {
    config.ConfigVal minio_accesskey = 1; // string
    config.ConfigVal minio_secretkey = 2; // string
}

message MaintemplateComponent {
    MinioComponent minio_access_secret = 1; // string
    config.ConfigVal minio_location = 2; // string
    config.ConfigVal minio_timeout = 3; // uint64
    config.ConfigVal minio_ssl = 4; // bool
    config.ConfigVal minio_enckey = 5; // bytes
    config.ConfigVal minio_endpoint = 6; // cidr
}

message GcpComponent {
    config.ConfigVal gcp_user = 1; // string
    config.ConfigVal gcp_project = 2; // string
    config.ConfigVal gke_cluster = 3; // string
    config.ConfigVal gke_zone = 4; // string
    config.ConfigVal gke_email = 5; // email
}

message JwtComponent {
    config.ConfigVal private_key = 1; // string
    config.ConfigVal public_key = 2; // string
}

message WorkflowComponent {
    config.ConfigVal github_sha = 2; // string
    config.ConfigVal github_ref = 3; // string
    config.ConfigVal project = 4; // string
    config.ConfigVal registry_hostname = 5; // cidr
    config.ConfigVal url = 6; // string
    config.ConfigVal locales = 7; // repeated string
    config.ConfigVal flutter_channel = 8; // string
    config.ConfigVal release_channel = 9; // string
}
```

Will output this:
```
{
  "components": [
    {
      "componentName": "MinioComponent",
      "config": {
        "minio_accesskey": "REPLACE_THIS",
        "minio_secretkey": "REPLACE_THIS"
      }
    },
    {
      "componentName": "MaintemplateComponent",
      "config": {
        "minio_access_secret": "REPLACE_THIS",
        "minio_enckey": "[\"\", \"\"]",
        "minio_endpoint": "127.0.0.1",
        "minio_location": "REPLACE_THIS",
        "minio_ssl": false,
        "minio_timeout": 0
      }
    },
    {
      "componentName": "GcpComponent",
      "config": {
        "gcp_project": "REPLACE_THIS",
        "gcp_user": "REPLACE_THIS",
        "gke_cluster": "REPLACE_THIS",
        "gke_email": "test@example.com",
        "gke_zone": "REPLACE_THIS"
      }
    },
    {
      "componentName": "JwtComponent",
      "config": {
        "private_key": "REPLACE_THIS",
        "public_key": "REPLACE_THIS"
      }
    },
    {
      "componentName": "WorkflowComponent",
      "config": {
        "flutter_channel": "REPLACE_THIS",
        "github_ref": "REPLACE_THIS",
        "github_sha": "REPLACE_THIS",
        "locales": "[REPLACE_THIS, REPLACE_THIS]",
        "project": "REPLACE_THIS",
        "registry_hostname": "127.0.0.1",
        "release_channel": "REPLACE_THIS",
        "url": "REPLACE_THIS"
      }
    }
  ]
```