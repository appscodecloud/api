# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pv.proto

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
  name='pv.proto',
  package='pv',
  syntax='proto3',
  serialized_pb=_b('\n\x08pv.proto\x12\x02pv\x1a\x1cgoogle/api/annotations.proto\x1a\x16\x61pi/dtypes/types.proto\"v\n\x11PVRegisterRequest\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nidentifier\x18\x03 \x01(\t\x12\x0e\n\x06plugin\x18\x04 \x01(\t\x12\x0c\n\x04size\x18\x05 \x01(\x03\x12\x10\n\x08\x65ndpoint\x18\x06 \x01(\t\"4\n\x13PVUnregisterRequest\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"2\n\x11PVDescribeRequest\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"x\n\x06PVInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04size\x18\x02 \x01(\x03\x12\x0e\n\x06status\x18\x03 \x01(\t\x12\x0e\n\x06volume\x18\x04 \x01(\t\x12\r\n\x05\x63laim\x18\x05 \x01(\t\x12\x0e\n\x06plugin\x18\x06 \x01(\t\x12\x13\n\x0b\x61\x63\x63\x65ssModes\x18\x07 \x03(\t\"L\n\x12PVDescribeResponse\x12\x1e\n\x06status\x18\x01 \x01(\x0b\x32\x0e.dtypes.Status\x12\x16\n\x02pv\x18\x02 \x01(\x0b\x32\n.pv.PVInfo2\xbc\x02\n\x11PersistentVolumes\x12`\n\x08\x44\x65scribe\x12\x15.pv.PVDescribeRequest\x1a\x16.pv.PVDescribeResponse\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/api/pv/v0.1/{cluster}/{name}\x12\x61\n\x08Register\x12\x15.pv.PVRegisterRequest\x1a\x14.dtypes.VoidResponse\"(\x82\xd3\xe4\x93\x02\"\x1a\x1d/api/pv/v0.1/{cluster}/{name}:\x01*\x12\x62\n\nUnregister\x12\x17.pv.PVUnregisterRequest\x1a\x14.dtypes.VoidResponse\"%\x82\xd3\xe4\x93\x02\x1f*\x1d/api/pv/v0.1/{cluster}/{name}b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,api_dot_dtypes_dot_types__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_PVREGISTERREQUEST = _descriptor.Descriptor(
  name='PVRegisterRequest',
  full_name='pv.PVRegisterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cluster', full_name='pv.PVRegisterRequest.cluster', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='pv.PVRegisterRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='identifier', full_name='pv.PVRegisterRequest.identifier', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='plugin', full_name='pv.PVRegisterRequest.plugin', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='size', full_name='pv.PVRegisterRequest.size', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='endpoint', full_name='pv.PVRegisterRequest.endpoint', index=5,
      number=6, type=9, cpp_type=9, label=1,
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
  serialized_start=70,
  serialized_end=188,
)


_PVUNREGISTERREQUEST = _descriptor.Descriptor(
  name='PVUnregisterRequest',
  full_name='pv.PVUnregisterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cluster', full_name='pv.PVUnregisterRequest.cluster', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='pv.PVUnregisterRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=190,
  serialized_end=242,
)


_PVDESCRIBEREQUEST = _descriptor.Descriptor(
  name='PVDescribeRequest',
  full_name='pv.PVDescribeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cluster', full_name='pv.PVDescribeRequest.cluster', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='pv.PVDescribeRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=244,
  serialized_end=294,
)


_PVINFO = _descriptor.Descriptor(
  name='PVInfo',
  full_name='pv.PVInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='pv.PVInfo.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='size', full_name='pv.PVInfo.size', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='pv.PVInfo.status', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume', full_name='pv.PVInfo.volume', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='claim', full_name='pv.PVInfo.claim', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='plugin', full_name='pv.PVInfo.plugin', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='accessModes', full_name='pv.PVInfo.accessModes', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=296,
  serialized_end=416,
)


_PVDESCRIBERESPONSE = _descriptor.Descriptor(
  name='PVDescribeResponse',
  full_name='pv.PVDescribeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='pv.PVDescribeResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='pv', full_name='pv.PVDescribeResponse.pv', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=418,
  serialized_end=494,
)

_PVDESCRIBERESPONSE.fields_by_name['status'].message_type = api_dot_dtypes_dot_types__pb2._STATUS
_PVDESCRIBERESPONSE.fields_by_name['pv'].message_type = _PVINFO
DESCRIPTOR.message_types_by_name['PVRegisterRequest'] = _PVREGISTERREQUEST
DESCRIPTOR.message_types_by_name['PVUnregisterRequest'] = _PVUNREGISTERREQUEST
DESCRIPTOR.message_types_by_name['PVDescribeRequest'] = _PVDESCRIBEREQUEST
DESCRIPTOR.message_types_by_name['PVInfo'] = _PVINFO
DESCRIPTOR.message_types_by_name['PVDescribeResponse'] = _PVDESCRIBERESPONSE

PVRegisterRequest = _reflection.GeneratedProtocolMessageType('PVRegisterRequest', (_message.Message,), dict(
  DESCRIPTOR = _PVREGISTERREQUEST,
  __module__ = 'pv_pb2'
  # @@protoc_insertion_point(class_scope:pv.PVRegisterRequest)
  ))
_sym_db.RegisterMessage(PVRegisterRequest)

PVUnregisterRequest = _reflection.GeneratedProtocolMessageType('PVUnregisterRequest', (_message.Message,), dict(
  DESCRIPTOR = _PVUNREGISTERREQUEST,
  __module__ = 'pv_pb2'
  # @@protoc_insertion_point(class_scope:pv.PVUnregisterRequest)
  ))
_sym_db.RegisterMessage(PVUnregisterRequest)

PVDescribeRequest = _reflection.GeneratedProtocolMessageType('PVDescribeRequest', (_message.Message,), dict(
  DESCRIPTOR = _PVDESCRIBEREQUEST,
  __module__ = 'pv_pb2'
  # @@protoc_insertion_point(class_scope:pv.PVDescribeRequest)
  ))
_sym_db.RegisterMessage(PVDescribeRequest)

PVInfo = _reflection.GeneratedProtocolMessageType('PVInfo', (_message.Message,), dict(
  DESCRIPTOR = _PVINFO,
  __module__ = 'pv_pb2'
  # @@protoc_insertion_point(class_scope:pv.PVInfo)
  ))
_sym_db.RegisterMessage(PVInfo)

PVDescribeResponse = _reflection.GeneratedProtocolMessageType('PVDescribeResponse', (_message.Message,), dict(
  DESCRIPTOR = _PVDESCRIBERESPONSE,
  __module__ = 'pv_pb2'
  # @@protoc_insertion_point(class_scope:pv.PVDescribeResponse)
  ))
_sym_db.RegisterMessage(PVDescribeResponse)


# @@protoc_insertion_point(module_scope)
