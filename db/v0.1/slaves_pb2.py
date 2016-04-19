# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: slaves.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from api.dtypes import types_pb2 as api_dot_dtypes_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='slaves.proto',
  package='db',
  syntax='proto3',
  serialized_pb=_b('\n\x0cslaves.proto\x12\x02\x64\x62\x1a\x1cgoogle/api/annotations.proto\x1a\x16\x61pi/dtypes/types.proto\"@\n\x0fSlaveAddRequest\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06\x63onfig\x18\x03 \x01(\t2c\n\x06Slaves\x12Y\n\x03\x41\x64\x64\x12\x13.db.SlaveAddRequest\x1a\x14.dtypes.VoidResponse\"\'\x82\xd3\xe4\x93\x02!\x1a\x1c/api/db/v0.1/slave/{cluster}:\x01*b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,api_dot_dtypes_dot_types__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_SLAVEADDREQUEST = _descriptor.Descriptor(
  name='SlaveAddRequest',
  full_name='db.SlaveAddRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cluster', full_name='db.SlaveAddRequest.cluster', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='db.SlaveAddRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='config', full_name='db.SlaveAddRequest.config', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=74,
  serialized_end=138,
)

DESCRIPTOR.message_types_by_name['SlaveAddRequest'] = _SLAVEADDREQUEST

SlaveAddRequest = _reflection.GeneratedProtocolMessageType('SlaveAddRequest', (_message.Message,), dict(
  DESCRIPTOR = _SLAVEADDREQUEST,
  __module__ = 'slaves_pb2'
  # @@protoc_insertion_point(class_scope:db.SlaveAddRequest)
  ))
_sym_db.RegisterMessage(SlaveAddRequest)


# @@protoc_insertion_point(module_scope)
