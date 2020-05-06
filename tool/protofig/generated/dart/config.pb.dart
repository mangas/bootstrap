///
//  Generated code. Do not modify.
//  source: config.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'baseproto.pb.dart' as $0;

class MinioComponent extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MinioComponent', package: const $pb.PackageName('config'), createEmptyInstance: create)
    ..aOM<$0.ConfigVal>(1, 'minioAccesskey', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(2, 'minioSecretkey', subBuilder: $0.ConfigVal.create)
    ..hasRequiredFields = false
  ;

  MinioComponent._() : super();
  factory MinioComponent() => create();
  factory MinioComponent.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MinioComponent.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  MinioComponent clone() => MinioComponent()..mergeFromMessage(this);
  MinioComponent copyWith(void Function(MinioComponent) updates) => super.copyWith((message) => updates(message as MinioComponent));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MinioComponent create() => MinioComponent._();
  MinioComponent createEmptyInstance() => create();
  static $pb.PbList<MinioComponent> createRepeated() => $pb.PbList<MinioComponent>();
  @$core.pragma('dart2js:noInline')
  static MinioComponent getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MinioComponent>(create);
  static MinioComponent _defaultInstance;

  @$pb.TagNumber(1)
  $0.ConfigVal get minioAccesskey => $_getN(0);
  @$pb.TagNumber(1)
  set minioAccesskey($0.ConfigVal v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMinioAccesskey() => $_has(0);
  @$pb.TagNumber(1)
  void clearMinioAccesskey() => clearField(1);
  @$pb.TagNumber(1)
  $0.ConfigVal ensureMinioAccesskey() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.ConfigVal get minioSecretkey => $_getN(1);
  @$pb.TagNumber(2)
  set minioSecretkey($0.ConfigVal v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasMinioSecretkey() => $_has(1);
  @$pb.TagNumber(2)
  void clearMinioSecretkey() => clearField(2);
  @$pb.TagNumber(2)
  $0.ConfigVal ensureMinioSecretkey() => $_ensure(1);
}

class MaintemplateComponent extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MaintemplateComponent', package: const $pb.PackageName('config'), createEmptyInstance: create)
    ..aOM<$0.ConfigVal>(1, 'minioAccesskey', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(2, 'minioSecretkey', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(3, 'minioLocation', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(4, 'minioTimeout', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(5, 'minioSsl', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(6, 'minioEnckey', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(7, 'minioEndpoint', subBuilder: $0.ConfigVal.create)
    ..hasRequiredFields = false
  ;

  MaintemplateComponent._() : super();
  factory MaintemplateComponent() => create();
  factory MaintemplateComponent.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MaintemplateComponent.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  MaintemplateComponent clone() => MaintemplateComponent()..mergeFromMessage(this);
  MaintemplateComponent copyWith(void Function(MaintemplateComponent) updates) => super.copyWith((message) => updates(message as MaintemplateComponent));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MaintemplateComponent create() => MaintemplateComponent._();
  MaintemplateComponent createEmptyInstance() => create();
  static $pb.PbList<MaintemplateComponent> createRepeated() => $pb.PbList<MaintemplateComponent>();
  @$core.pragma('dart2js:noInline')
  static MaintemplateComponent getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MaintemplateComponent>(create);
  static MaintemplateComponent _defaultInstance;

  @$pb.TagNumber(1)
  $0.ConfigVal get minioAccesskey => $_getN(0);
  @$pb.TagNumber(1)
  set minioAccesskey($0.ConfigVal v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMinioAccesskey() => $_has(0);
  @$pb.TagNumber(1)
  void clearMinioAccesskey() => clearField(1);
  @$pb.TagNumber(1)
  $0.ConfigVal ensureMinioAccesskey() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.ConfigVal get minioSecretkey => $_getN(1);
  @$pb.TagNumber(2)
  set minioSecretkey($0.ConfigVal v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasMinioSecretkey() => $_has(1);
  @$pb.TagNumber(2)
  void clearMinioSecretkey() => clearField(2);
  @$pb.TagNumber(2)
  $0.ConfigVal ensureMinioSecretkey() => $_ensure(1);

  @$pb.TagNumber(3)
  $0.ConfigVal get minioLocation => $_getN(2);
  @$pb.TagNumber(3)
  set minioLocation($0.ConfigVal v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasMinioLocation() => $_has(2);
  @$pb.TagNumber(3)
  void clearMinioLocation() => clearField(3);
  @$pb.TagNumber(3)
  $0.ConfigVal ensureMinioLocation() => $_ensure(2);

  @$pb.TagNumber(4)
  $0.ConfigVal get minioTimeout => $_getN(3);
  @$pb.TagNumber(4)
  set minioTimeout($0.ConfigVal v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasMinioTimeout() => $_has(3);
  @$pb.TagNumber(4)
  void clearMinioTimeout() => clearField(4);
  @$pb.TagNumber(4)
  $0.ConfigVal ensureMinioTimeout() => $_ensure(3);

  @$pb.TagNumber(5)
  $0.ConfigVal get minioSsl => $_getN(4);
  @$pb.TagNumber(5)
  set minioSsl($0.ConfigVal v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasMinioSsl() => $_has(4);
  @$pb.TagNumber(5)
  void clearMinioSsl() => clearField(5);
  @$pb.TagNumber(5)
  $0.ConfigVal ensureMinioSsl() => $_ensure(4);

  @$pb.TagNumber(6)
  $0.ConfigVal get minioEnckey => $_getN(5);
  @$pb.TagNumber(6)
  set minioEnckey($0.ConfigVal v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasMinioEnckey() => $_has(5);
  @$pb.TagNumber(6)
  void clearMinioEnckey() => clearField(6);
  @$pb.TagNumber(6)
  $0.ConfigVal ensureMinioEnckey() => $_ensure(5);

  @$pb.TagNumber(7)
  $0.ConfigVal get minioEndpoint => $_getN(6);
  @$pb.TagNumber(7)
  set minioEndpoint($0.ConfigVal v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasMinioEndpoint() => $_has(6);
  @$pb.TagNumber(7)
  void clearMinioEndpoint() => clearField(7);
  @$pb.TagNumber(7)
  $0.ConfigVal ensureMinioEndpoint() => $_ensure(6);
}

class GcpComponent extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GcpComponent', package: const $pb.PackageName('config'), createEmptyInstance: create)
    ..aOM<$0.ConfigVal>(1, 'gcpUser', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(2, 'gcpProject', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(3, 'gkeCluster', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(4, 'gkeZone', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(5, 'gkeEmail', subBuilder: $0.ConfigVal.create)
    ..hasRequiredFields = false
  ;

  GcpComponent._() : super();
  factory GcpComponent() => create();
  factory GcpComponent.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GcpComponent.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GcpComponent clone() => GcpComponent()..mergeFromMessage(this);
  GcpComponent copyWith(void Function(GcpComponent) updates) => super.copyWith((message) => updates(message as GcpComponent));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GcpComponent create() => GcpComponent._();
  GcpComponent createEmptyInstance() => create();
  static $pb.PbList<GcpComponent> createRepeated() => $pb.PbList<GcpComponent>();
  @$core.pragma('dart2js:noInline')
  static GcpComponent getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GcpComponent>(create);
  static GcpComponent _defaultInstance;

  @$pb.TagNumber(1)
  $0.ConfigVal get gcpUser => $_getN(0);
  @$pb.TagNumber(1)
  set gcpUser($0.ConfigVal v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasGcpUser() => $_has(0);
  @$pb.TagNumber(1)
  void clearGcpUser() => clearField(1);
  @$pb.TagNumber(1)
  $0.ConfigVal ensureGcpUser() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.ConfigVal get gcpProject => $_getN(1);
  @$pb.TagNumber(2)
  set gcpProject($0.ConfigVal v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasGcpProject() => $_has(1);
  @$pb.TagNumber(2)
  void clearGcpProject() => clearField(2);
  @$pb.TagNumber(2)
  $0.ConfigVal ensureGcpProject() => $_ensure(1);

  @$pb.TagNumber(3)
  $0.ConfigVal get gkeCluster => $_getN(2);
  @$pb.TagNumber(3)
  set gkeCluster($0.ConfigVal v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasGkeCluster() => $_has(2);
  @$pb.TagNumber(3)
  void clearGkeCluster() => clearField(3);
  @$pb.TagNumber(3)
  $0.ConfigVal ensureGkeCluster() => $_ensure(2);

  @$pb.TagNumber(4)
  $0.ConfigVal get gkeZone => $_getN(3);
  @$pb.TagNumber(4)
  set gkeZone($0.ConfigVal v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasGkeZone() => $_has(3);
  @$pb.TagNumber(4)
  void clearGkeZone() => clearField(4);
  @$pb.TagNumber(4)
  $0.ConfigVal ensureGkeZone() => $_ensure(3);

  @$pb.TagNumber(5)
  $0.ConfigVal get gkeEmail => $_getN(4);
  @$pb.TagNumber(5)
  set gkeEmail($0.ConfigVal v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasGkeEmail() => $_has(4);
  @$pb.TagNumber(5)
  void clearGkeEmail() => clearField(5);
  @$pb.TagNumber(5)
  $0.ConfigVal ensureGkeEmail() => $_ensure(4);
}

class JwtComponent extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('JwtComponent', package: const $pb.PackageName('config'), createEmptyInstance: create)
    ..aOM<$0.ConfigVal>(1, 'privateKey', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(2, 'publicKey', subBuilder: $0.ConfigVal.create)
    ..hasRequiredFields = false
  ;

  JwtComponent._() : super();
  factory JwtComponent() => create();
  factory JwtComponent.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory JwtComponent.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  JwtComponent clone() => JwtComponent()..mergeFromMessage(this);
  JwtComponent copyWith(void Function(JwtComponent) updates) => super.copyWith((message) => updates(message as JwtComponent));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static JwtComponent create() => JwtComponent._();
  JwtComponent createEmptyInstance() => create();
  static $pb.PbList<JwtComponent> createRepeated() => $pb.PbList<JwtComponent>();
  @$core.pragma('dart2js:noInline')
  static JwtComponent getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<JwtComponent>(create);
  static JwtComponent _defaultInstance;

  @$pb.TagNumber(1)
  $0.ConfigVal get privateKey => $_getN(0);
  @$pb.TagNumber(1)
  set privateKey($0.ConfigVal v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPrivateKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearPrivateKey() => clearField(1);
  @$pb.TagNumber(1)
  $0.ConfigVal ensurePrivateKey() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.ConfigVal get publicKey => $_getN(1);
  @$pb.TagNumber(2)
  set publicKey($0.ConfigVal v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPublicKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearPublicKey() => clearField(2);
  @$pb.TagNumber(2)
  $0.ConfigVal ensurePublicKey() => $_ensure(1);
}

class WorkflowComponent extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('WorkflowComponent', package: const $pb.PackageName('config'), createEmptyInstance: create)
    ..aOM<$0.ConfigVal>(2, 'githubSha', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(3, 'githubRef', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(4, 'project', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(5, 'registryHostname', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(6, 'url', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(7, 'locales', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(8, 'flutterChannel', subBuilder: $0.ConfigVal.create)
    ..aOM<$0.ConfigVal>(9, 'releaseChannel', subBuilder: $0.ConfigVal.create)
    ..hasRequiredFields = false
  ;

  WorkflowComponent._() : super();
  factory WorkflowComponent() => create();
  factory WorkflowComponent.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WorkflowComponent.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  WorkflowComponent clone() => WorkflowComponent()..mergeFromMessage(this);
  WorkflowComponent copyWith(void Function(WorkflowComponent) updates) => super.copyWith((message) => updates(message as WorkflowComponent));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WorkflowComponent create() => WorkflowComponent._();
  WorkflowComponent createEmptyInstance() => create();
  static $pb.PbList<WorkflowComponent> createRepeated() => $pb.PbList<WorkflowComponent>();
  @$core.pragma('dart2js:noInline')
  static WorkflowComponent getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WorkflowComponent>(create);
  static WorkflowComponent _defaultInstance;

  @$pb.TagNumber(2)
  $0.ConfigVal get githubSha => $_getN(0);
  @$pb.TagNumber(2)
  set githubSha($0.ConfigVal v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasGithubSha() => $_has(0);
  @$pb.TagNumber(2)
  void clearGithubSha() => clearField(2);
  @$pb.TagNumber(2)
  $0.ConfigVal ensureGithubSha() => $_ensure(0);

  @$pb.TagNumber(3)
  $0.ConfigVal get githubRef => $_getN(1);
  @$pb.TagNumber(3)
  set githubRef($0.ConfigVal v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasGithubRef() => $_has(1);
  @$pb.TagNumber(3)
  void clearGithubRef() => clearField(3);
  @$pb.TagNumber(3)
  $0.ConfigVal ensureGithubRef() => $_ensure(1);

  @$pb.TagNumber(4)
  $0.ConfigVal get project => $_getN(2);
  @$pb.TagNumber(4)
  set project($0.ConfigVal v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasProject() => $_has(2);
  @$pb.TagNumber(4)
  void clearProject() => clearField(4);
  @$pb.TagNumber(4)
  $0.ConfigVal ensureProject() => $_ensure(2);

  @$pb.TagNumber(5)
  $0.ConfigVal get registryHostname => $_getN(3);
  @$pb.TagNumber(5)
  set registryHostname($0.ConfigVal v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasRegistryHostname() => $_has(3);
  @$pb.TagNumber(5)
  void clearRegistryHostname() => clearField(5);
  @$pb.TagNumber(5)
  $0.ConfigVal ensureRegistryHostname() => $_ensure(3);

  @$pb.TagNumber(6)
  $0.ConfigVal get url => $_getN(4);
  @$pb.TagNumber(6)
  set url($0.ConfigVal v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasUrl() => $_has(4);
  @$pb.TagNumber(6)
  void clearUrl() => clearField(6);
  @$pb.TagNumber(6)
  $0.ConfigVal ensureUrl() => $_ensure(4);

  @$pb.TagNumber(7)
  $0.ConfigVal get locales => $_getN(5);
  @$pb.TagNumber(7)
  set locales($0.ConfigVal v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasLocales() => $_has(5);
  @$pb.TagNumber(7)
  void clearLocales() => clearField(7);
  @$pb.TagNumber(7)
  $0.ConfigVal ensureLocales() => $_ensure(5);

  @$pb.TagNumber(8)
  $0.ConfigVal get flutterChannel => $_getN(6);
  @$pb.TagNumber(8)
  set flutterChannel($0.ConfigVal v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasFlutterChannel() => $_has(6);
  @$pb.TagNumber(8)
  void clearFlutterChannel() => clearField(8);
  @$pb.TagNumber(8)
  $0.ConfigVal ensureFlutterChannel() => $_ensure(6);

  @$pb.TagNumber(9)
  $0.ConfigVal get releaseChannel => $_getN(7);
  @$pb.TagNumber(9)
  set releaseChannel($0.ConfigVal v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasReleaseChannel() => $_has(7);
  @$pb.TagNumber(9)
  void clearReleaseChannel() => clearField(9);
  @$pb.TagNumber(9)
  $0.ConfigVal ensureReleaseChannel() => $_ensure(7);
}

