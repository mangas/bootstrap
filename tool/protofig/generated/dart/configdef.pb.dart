///
//  Generated code. Do not modify.
//  source: configdef.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Config extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Config', package: const $pb.PackageName('configdef'), createEmptyInstance: create)
    ..aOS(1, 'key')
    ..aOM<ConfigVal>(2, 'val', subBuilder: ConfigVal.create)
    ..hasRequiredFields = false
  ;

  Config._() : super();
  factory Config() => create();
  factory Config.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Config.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Config clone() => Config()..mergeFromMessage(this);
  Config copyWith(void Function(Config) updates) => super.copyWith((message) => updates(message as Config));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Config create() => Config._();
  Config createEmptyInstance() => create();
  static $pb.PbList<Config> createRepeated() => $pb.PbList<Config>();
  @$core.pragma('dart2js:noInline')
  static Config getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Config>(create);
  static Config _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => clearField(1);

  @$pb.TagNumber(2)
  ConfigVal get val => $_getN(1);
  @$pb.TagNumber(2)
  set val(ConfigVal v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasVal() => $_has(1);
  @$pb.TagNumber(2)
  void clearVal() => clearField(2);
  @$pb.TagNumber(2)
  ConfigVal ensureVal() => $_ensure(1);
}

enum ConfigVal_Val {
  stringVal, 
  uint64Val, 
  emailVal, 
  cidrVal, 
  boolVal, 
  doubleVal, 
  notSet
}

class ConfigVal extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, ConfigVal_Val> _ConfigVal_ValByTag = {
    1 : ConfigVal_Val.stringVal,
    2 : ConfigVal_Val.uint64Val,
    3 : ConfigVal_Val.emailVal,
    4 : ConfigVal_Val.cidrVal,
    5 : ConfigVal_Val.boolVal,
    6 : ConfigVal_Val.doubleVal,
    0 : ConfigVal_Val.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ConfigVal', package: const $pb.PackageName('configdef'), createEmptyInstance: create)
    ..oo(0, [1, 2, 3, 4, 5, 6])
    ..aOS(1, 'stringVal')
    ..a<$fixnum.Int64>(2, 'uint64Val', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, 'emailVal')
    ..a<$core.List<$core.int>>(4, 'cidrVal', $pb.PbFieldType.OY)
    ..aOB(5, 'boolVal')
    ..a<$core.double>(6, 'doubleVal', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  ConfigVal._() : super();
  factory ConfigVal() => create();
  factory ConfigVal.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigVal.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ConfigVal clone() => ConfigVal()..mergeFromMessage(this);
  ConfigVal copyWith(void Function(ConfigVal) updates) => super.copyWith((message) => updates(message as ConfigVal));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ConfigVal create() => ConfigVal._();
  ConfigVal createEmptyInstance() => create();
  static $pb.PbList<ConfigVal> createRepeated() => $pb.PbList<ConfigVal>();
  @$core.pragma('dart2js:noInline')
  static ConfigVal getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigVal>(create);
  static ConfigVal _defaultInstance;

  ConfigVal_Val whichVal() => _ConfigVal_ValByTag[$_whichOneof(0)];
  void clearVal() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get stringVal => $_getSZ(0);
  @$pb.TagNumber(1)
  set stringVal($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasStringVal() => $_has(0);
  @$pb.TagNumber(1)
  void clearStringVal() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get uint64Val => $_getI64(1);
  @$pb.TagNumber(2)
  set uint64Val($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUint64Val() => $_has(1);
  @$pb.TagNumber(2)
  void clearUint64Val() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get emailVal => $_getSZ(2);
  @$pb.TagNumber(3)
  set emailVal($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasEmailVal() => $_has(2);
  @$pb.TagNumber(3)
  void clearEmailVal() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.int> get cidrVal => $_getN(3);
  @$pb.TagNumber(4)
  set cidrVal($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCidrVal() => $_has(3);
  @$pb.TagNumber(4)
  void clearCidrVal() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get boolVal => $_getBF(4);
  @$pb.TagNumber(5)
  set boolVal($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasBoolVal() => $_has(4);
  @$pb.TagNumber(5)
  void clearBoolVal() => clearField(5);

  @$pb.TagNumber(6)
  $core.double get doubleVal => $_getN(5);
  @$pb.TagNumber(6)
  set doubleVal($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDoubleVal() => $_has(5);
  @$pb.TagNumber(6)
  void clearDoubleVal() => clearField(6);
}

