# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: kafka.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'kafka.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bkafka.proto\x12\ryir.mri.kafka\"\xda\x02\n\x0cMriProcessed\x12/\n\x05nodes\x18\x64 \x03(\x0b\x32 .yir.mri.kafka.MriProcessed.Node\x12\x36\n\x08segments\x18\xc8\x01 \x03(\x0b\x32#.yir.mri.kafka.MriProcessed.Segment\x1a[\n\x04Node\x12\n\n\x02id\x18\x64 \x01(\t\x12\x0f\n\x06mri_id\x18\xc8\x01 \x01(\t\x12\x12\n\tknosp_012\x18\x90\x03 \x01(\x01\x12\x10\n\x07knosp_3\x18\xf4\x03 \x01(\x01\x12\x10\n\x07knosp_4\x18\xd8\x04 \x01(\x01\x1a\x83\x01\n\x07Segment\x12\n\n\x02id\x18\x64 \x01(\t\x12\x10\n\x07node_id\x18\xc8\x01 \x01(\t\x12\x11\n\x08image_id\x18\xac\x02 \x01(\t\x12\x0f\n\x06\x63ontor\x18\x90\x03 \x01(\t\x12\x12\n\tknosp_012\x18\xf4\x03 \x01(\x01\x12\x10\n\x07knosp_3\x18\xd8\x04 \x01(\x01\x12\x10\n\x07knosp_4\x18\xbc\x05 \x01(\x01\"0\n\x0bmriSplitted\x12\x0e\n\x06mri_id\x18\x64 \x01(\t\x12\x11\n\x08pages_id\x18\xc8\x01 \x03(\tB\x1bZ\x19yir/mri/api/broker;brokerb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kafka_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\031yir/mri/api/broker;broker'
  _globals['_MRIPROCESSED']._serialized_start=31
  _globals['_MRIPROCESSED']._serialized_end=377
  _globals['_MRIPROCESSED_NODE']._serialized_start=152
  _globals['_MRIPROCESSED_NODE']._serialized_end=243
  _globals['_MRIPROCESSED_SEGMENT']._serialized_start=246
  _globals['_MRIPROCESSED_SEGMENT']._serialized_end=377
  _globals['_MRISPLITTED']._serialized_start=379
  _globals['_MRISPLITTED']._serialized_end=427
# @@protoc_insertion_point(module_scope)
