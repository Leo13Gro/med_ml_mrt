# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ml_api.proto
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
    'ml_api.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cml_api.proto\x12\nyir.ml_api\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/api/annotations.proto\"1\n\x1fSegmentAndClassificationRequest\x12\x0e\n\x06mri_id\x18\x01 \x01(\t2}\n\x05MLAPI\x12t\n\x18SegmentAndClassification\x12+.yir.ml_api.SegmentAndClassificationRequest\x1a\x16.google.protobuf.Empty\"\x13\x82\xd3\xe4\x93\x02\r\"\x08/mlmagic:\x01*B\x18Z\x16yir/ml_service/api;apib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ml_api_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\026yir/ml_service/api;api'
  _globals['_MLAPI'].methods_by_name['SegmentAndClassification']._loaded_options = None
  _globals['_MLAPI'].methods_by_name['SegmentAndClassification']._serialized_options = b'\202\323\344\223\002\r\"\010/mlmagic:\001*'
  _globals['_SEGMENTANDCLASSIFICATIONREQUEST']._serialized_start=87
  _globals['_SEGMENTANDCLASSIFICATIONREQUEST']._serialized_end=136
  _globals['_MLAPI']._serialized_start=138
  _globals['_MLAPI']._serialized_end=263
# @@protoc_insertion_point(module_scope)
