import 'dart:convert';
import 'baseproto.pb.dart';	
import 'package:fixnum/fixnum.dart';	
import 'config.pb.dart';	
	
class ConfigDefConfig {
  final List<ConfigAppConfig> appConfig;
  ConfigDefConfig({
    this.appConfig
  });
	
  factory ConfigDefConfig.fromJson(Map<String, dynamic> json) => ConfigDefConfig(
    appConfig: List<ConfigAppConfig>.from(json["appConfig"].map((x) => ConfigAppConfig.fromJson(x))),
  );

  Map<String, dynamic> toJson() => {
      "appConfig": List<dynamic>.from(appConfig.map((x) => x.toJson())),
  };
}
	
class ConfigAppConfig {
  String componentName;
  Map<String, dynamic> config;
	
  ConfigAppConfig({
    this.componentName,
    this.config,
  });
	
  factory ConfigAppConfig.fromJson(Map<String, dynamic> json) =>
    ConfigAppConfig(
      componentName: json["componentName"], config: json["config"],
	);
	
  factory ConfigAppConfig.fromEmpty() => ConfigAppConfig(
	  componentName: "", config: Map<String, dynamic>(),
	);
	
  Map<String, dynamic> toJson() => {
    "componentName": componentName,
    "config": config,
  };

  MinioComponent MinioComponentFromAppConfig() {
	MinioComponent msg = MinioComponent.create();
	msg.minioAccesskey = ConfigVal().fromDynamic(config["minioAccesskey"]);
	msg.minioSecretkey = ConfigVal().fromDynamic(config["minioSecretkey"]);
	
	return msg;
  }
  MaintemplateComponent MaintemplateComponentFromAppConfig() {
	MaintemplateComponent msg = MaintemplateComponent.create();
	msg.minioAccesskey = ConfigVal().fromDynamic(config["minioAccesskey"]);
	msg.minioSecretkey = ConfigVal().fromDynamic(config["minioSecretkey"]);
	msg.minioLocation = ConfigVal().fromDynamic(config["minioLocation"]);
	msg.minioTimeout = ConfigVal().fromDynamic(config["minioTimeout"]);
	msg.minioSsl = ConfigVal().fromDynamic(config["minioSsl"]);
	msg.minioEnckey = ConfigVal().fromDynamic(config["minioEnckey"]);
	msg.minioEndpoint = ConfigVal().fromDynamic(config["minioEndpoint"]);
	
	return msg;
  }
  GcpComponent GcpComponentFromAppConfig() {
	GcpComponent msg = GcpComponent.create();
	msg.gcpUser = ConfigVal().fromDynamic(config["gcpUser"]);
	msg.gcpProject = ConfigVal().fromDynamic(config["gcpProject"]);
	msg.gkeCluster = ConfigVal().fromDynamic(config["gkeCluster"]);
	msg.gkeZone = ConfigVal().fromDynamic(config["gkeZone"]);
	msg.gkeEmail = ConfigVal().fromDynamic(config["gkeEmail"]);
	
	return msg;
  }
  JwtComponent JwtComponentFromAppConfig() {
	JwtComponent msg = JwtComponent.create();
	msg.privateKey = ConfigVal().fromDynamic(config["privateKey"]);
	msg.publicKey = ConfigVal().fromDynamic(config["publicKey"]);
	
	return msg;
  }
  WorkflowComponent WorkflowComponentFromAppConfig() {
	WorkflowComponent msg = WorkflowComponent.create();
	msg.githubSha = ConfigVal().fromDynamic(config["githubSha"]);
	msg.githubRef = ConfigVal().fromDynamic(config["githubRef"]);
	msg.project = ConfigVal().fromDynamic(config["project"]);
	msg.registryHostname = ConfigVal().fromDynamic(config["registryHostname"]);
	msg.url = ConfigVal().fromDynamic(config["url"]);
	msg.locales = ConfigVal().fromDynamic(config["locales"]);
	msg.flutterChannel = ConfigVal().fromDynamic(config["flutterChannel"]);
	msg.releaseChannel = ConfigVal().fromDynamic(config["releaseChannel"]);
	
	return msg;
  }
  	
	
}

 extension ConfigMinioComponent on MinioComponent {
  ConfigAppConfig toAppConfig() {
	ConfigAppConfig component = ConfigAppConfig.fromEmpty();	
	component.componentName = "minioComponent";
	component.config["minioAccesskey"] = this.minioAccesskey.toDynamic(); 
	component.config["minioSecretkey"] = this.minioSecretkey.toDynamic(); 
	
	return component;
  }
}
 extension ConfigMaintemplateComponent on MaintemplateComponent {
  ConfigAppConfig toAppConfig() {
	ConfigAppConfig component = ConfigAppConfig.fromEmpty();	
	component.componentName = "maintemplateComponent";
	component.config["minioAccesskey"] = this.minioAccesskey.toDynamic(); 
	component.config["minioSecretkey"] = this.minioSecretkey.toDynamic(); 
	component.config["minioLocation"] = this.minioLocation.toDynamic(); 
	component.config["minioTimeout"] = this.minioTimeout.toDynamic(); 
	component.config["minioSsl"] = this.minioSsl.toDynamic(); 
	component.config["minioEnckey"] = this.minioEnckey.toDynamic(); 
	component.config["minioEndpoint"] = this.minioEndpoint.toDynamic(); 
	
	return component;
  }
}
 extension ConfigGcpComponent on GcpComponent {
  ConfigAppConfig toAppConfig() {
	ConfigAppConfig component = ConfigAppConfig.fromEmpty();	
	component.componentName = "gcpComponent";
	component.config["gcpUser"] = this.gcpUser.toDynamic(); 
	component.config["gcpProject"] = this.gcpProject.toDynamic(); 
	component.config["gkeCluster"] = this.gkeCluster.toDynamic(); 
	component.config["gkeZone"] = this.gkeZone.toDynamic(); 
	component.config["gkeEmail"] = this.gkeEmail.toDynamic(); 
	
	return component;
  }
}
 extension ConfigJwtComponent on JwtComponent {
  ConfigAppConfig toAppConfig() {
	ConfigAppConfig component = ConfigAppConfig.fromEmpty();	
	component.componentName = "jwtComponent";
	component.config["privateKey"] = this.privateKey.toDynamic(); 
	component.config["publicKey"] = this.publicKey.toDynamic(); 
	
	return component;
  }
}
 extension ConfigWorkflowComponent on WorkflowComponent {
  ConfigAppConfig toAppConfig() {
	ConfigAppConfig component = ConfigAppConfig.fromEmpty();	
	component.componentName = "workflowComponent";
	component.config["githubSha"] = this.githubSha.toDynamic(); 
	component.config["githubRef"] = this.githubRef.toDynamic(); 
	component.config["project"] = this.project.toDynamic(); 
	component.config["registryHostname"] = this.registryHostname.toDynamic(); 
	component.config["url"] = this.url.toDynamic(); 
	component.config["locales"] = this.locales.toDynamic(); 
	component.config["flutterChannel"] = this.flutterChannel.toDynamic(); 
	component.config["releaseChannel"] = this.releaseChannel.toDynamic(); 
	
	return component;
  }
}
	
	
extension ConfigConfigVal on ConfigVal {
  bool isEmail(String s) {
    return RegExp(
      r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+").hasMatch(s);
  }

  bool isCidr(String s) {
    RegExp ipv4 = RegExp(
      r"^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$");
    RegExp ipv6 = RegExp(
      r"^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(\/((1(1[0-9]|2[0-8]))|([0-9][0-9])|([0-9])))?$");

    return (ipv4.hasMatch(s) || ipv6.hasMatch(s));
  }

  ConfigVal fromDynamic(dynamic v) {
    ConfigVal cfg = ConfigVal.create();
    if (v is String) {
      if (isEmail(v)) {
        cfg.emailVal = v;
      } else if (isCidr(v)) {
        cfg.cidrVal = utf8.encode(v);
      }
      cfg.stringVal = v;
    }
    if (v is double) {
      cfg.doubleVal = v;
    }
    if (v is num) {
      cfg.uint64Val = v as Int64;
    }
    if (v is bool) {
      cfg.boolVal = v;
    }
    return cfg;
  }
	
  dynamic toDynamic() {
    if (this.hasStringVal()) {
      return this.stringVal;
    }
    else if (this.hasCidrVal()) {
	  return this.cidrVal;
    }
    else if (this.hasEmailVal()) {
	  return this.emailVal;
    }	
    else if (this.hasDoubleVal()) {
	  return this.doubleVal;
    }	
    else if (this.hasUint64Val()) {
	  return this.uint64Val;
    }
    else if (this.hasBoolVal()) {
	  return this.boolVal;
    }	
    return "";	
  }
}
	

