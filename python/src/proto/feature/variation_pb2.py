# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/feature/variation.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dproto/feature/variation.proto\x12\x11\x62ucketeer.feature\"I\n\tVariation\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\tB1Z/github.com/bucketeer-io/bucketeer/proto/featureb\x06proto3')



_VARIATION = DESCRIPTOR.message_types_by_name['Variation']
Variation = _reflection.GeneratedProtocolMessageType('Variation', (_message.Message,), {
  'DESCRIPTOR' : _VARIATION,
  '__module__' : 'proto.feature.variation_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.feature.Variation)
  })
_sym_db.RegisterMessage(Variation)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z/github.com/bucketeer-io/bucketeer/proto/feature'
  _VARIATION._serialized_start=52
  _VARIATION._serialized_end=125
# @@protoc_insertion_point(module_scope)