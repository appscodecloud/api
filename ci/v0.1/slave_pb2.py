# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: slave.proto

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
  name='slave.proto',
  package='ci',
  syntax='proto3',
  serialized_pb=_b('\n\x0bslave.proto\x12\x02\x63i\x1a\x1cgoogle/api/annotations.proto\x1a\x16\x61pi/dtypes/types.proto\"N\n\x11SlaveListResponse\x12\x1e\n\x06status\x18\x01 \x01(\x0b\x32\x0e.dtypes.Status\x12\x19\n\x06slaves\x18\x02 \x03(\x0b\x32\t.ci.Slave\"%\n\x05Slave\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\"\x8a\x01\n\x12SlaveCreateRequest\x12\x11\n\texecutors\x18\x01 \x01(\x05\x12\x0e\n\x06labels\x18\x02 \x01(\t\x12\x1b\n\x13user_startup_script\x18\x03 \x01(\t\x12\x18\n\x10saltbase_version\x18\x04 \x01(\t\x12\x1a\n\x12\x63i_starter_version\x18\x05 \x01(\t\"$\n\x14SlaveDescribeRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\xdf\x01\n\x15SlaveDescribeResponse\x12\x1e\n\x06status\x18\x01 \x01(\x0b\x32\x0e.dtypes.Status\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x11\n\texecutors\x18\x03 \x01(\x03\x12\x13\n\x0bnode_status\x18\x04 \x01(\t\x12\x16\n\x0eoffline_reason\x18\x05 \x01(\t\x12\r\n\x05label\x18\x06 \x01(\t\x12\x10\n\x08provider\x18\x07 \x01(\t\x12\x0b\n\x03sku\x18\x08 \x01(\t\x12\x16\n\x0estartup_script\x18\t \x01(\t\x12\x12\n\ncreated_at\x18\n \x01(\t\"\"\n\x12SlaveDeleteRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"#\n\x13SlaveRestartRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"6\n\x14SlaveRestartResponse\x12\x1e\n\x06status\x18\x01 \x01(\x0b\x32\x0e.dtypes.Status2\xe9\x03\n\x06Slaves\x12O\n\x04List\x12\x13.dtypes.VoidRequest\x1a\x15.ci.SlaveListResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/api/ci/v0.1/slaves\x12\x63\n\x08\x44\x65scribe\x12\x18.ci.SlaveDescribeRequest\x1a\x19.ci.SlaveDescribeResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/api/ci/v0.1/slaves/{name}\x12]\n\x06\x43reate\x12\x16.ci.SlaveCreateRequest\x1a\x1b.dtypes.LongRunningResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x1a\x13/api/ci/v0.1/slaves:\x01*\x12\x61\n\x06\x44\x65lete\x12\x16.ci.SlaveDeleteRequest\x1a\x1b.dtypes.LongRunningResponse\"\"\x82\xd3\xe4\x93\x02\x1c*\x1a/api/ci/v0.1/slaves/{name}\x12g\n\x07Restart\x12\x17.ci.SlaveRestartRequest\x1a\x18.ci.SlaveRestartResponse\")\x82\xd3\xe4\x93\x02#\x12!/api/ci/v0.1/slaves/{name}/rebootb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,api_dot_dtypes_dot_types__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_SLAVELISTRESPONSE = _descriptor.Descriptor(
  name='SlaveListResponse',
  full_name='ci.SlaveListResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ci.SlaveListResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='slaves', full_name='ci.SlaveListResponse.slaves', index=1,
      number=2, type=11, cpp_type=10, label=3,
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
  serialized_start=73,
  serialized_end=151,
)


_SLAVE = _descriptor.Descriptor(
  name='Slave',
  full_name='ci.Slave',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ci.Slave.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='ci.Slave.status', index=1,
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
  serialized_start=153,
  serialized_end=190,
)


_SLAVECREATEREQUEST = _descriptor.Descriptor(
  name='SlaveCreateRequest',
  full_name='ci.SlaveCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='executors', full_name='ci.SlaveCreateRequest.executors', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='labels', full_name='ci.SlaveCreateRequest.labels', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='user_startup_script', full_name='ci.SlaveCreateRequest.user_startup_script', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='saltbase_version', full_name='ci.SlaveCreateRequest.saltbase_version', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ci_starter_version', full_name='ci.SlaveCreateRequest.ci_starter_version', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=193,
  serialized_end=331,
)


_SLAVEDESCRIBEREQUEST = _descriptor.Descriptor(
  name='SlaveDescribeRequest',
  full_name='ci.SlaveDescribeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ci.SlaveDescribeRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=333,
  serialized_end=369,
)


_SLAVEDESCRIBERESPONSE = _descriptor.Descriptor(
  name='SlaveDescribeResponse',
  full_name='ci.SlaveDescribeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ci.SlaveDescribeResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='ci.SlaveDescribeResponse.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='executors', full_name='ci.SlaveDescribeResponse.executors', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='node_status', full_name='ci.SlaveDescribeResponse.node_status', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='offline_reason', full_name='ci.SlaveDescribeResponse.offline_reason', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='label', full_name='ci.SlaveDescribeResponse.label', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='provider', full_name='ci.SlaveDescribeResponse.provider', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sku', full_name='ci.SlaveDescribeResponse.sku', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='startup_script', full_name='ci.SlaveDescribeResponse.startup_script', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='ci.SlaveDescribeResponse.created_at', index=9,
      number=10, type=9, cpp_type=9, label=1,
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
  serialized_start=372,
  serialized_end=595,
)


_SLAVEDELETEREQUEST = _descriptor.Descriptor(
  name='SlaveDeleteRequest',
  full_name='ci.SlaveDeleteRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ci.SlaveDeleteRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=597,
  serialized_end=631,
)


_SLAVERESTARTREQUEST = _descriptor.Descriptor(
  name='SlaveRestartRequest',
  full_name='ci.SlaveRestartRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ci.SlaveRestartRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=633,
  serialized_end=668,
)


_SLAVERESTARTRESPONSE = _descriptor.Descriptor(
  name='SlaveRestartResponse',
  full_name='ci.SlaveRestartResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ci.SlaveRestartResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
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
  serialized_start=670,
  serialized_end=724,
)

_SLAVELISTRESPONSE.fields_by_name['status'].message_type = api_dot_dtypes_dot_types__pb2._STATUS
_SLAVELISTRESPONSE.fields_by_name['slaves'].message_type = _SLAVE
_SLAVEDESCRIBERESPONSE.fields_by_name['status'].message_type = api_dot_dtypes_dot_types__pb2._STATUS
_SLAVERESTARTRESPONSE.fields_by_name['status'].message_type = api_dot_dtypes_dot_types__pb2._STATUS
DESCRIPTOR.message_types_by_name['SlaveListResponse'] = _SLAVELISTRESPONSE
DESCRIPTOR.message_types_by_name['Slave'] = _SLAVE
DESCRIPTOR.message_types_by_name['SlaveCreateRequest'] = _SLAVECREATEREQUEST
DESCRIPTOR.message_types_by_name['SlaveDescribeRequest'] = _SLAVEDESCRIBEREQUEST
DESCRIPTOR.message_types_by_name['SlaveDescribeResponse'] = _SLAVEDESCRIBERESPONSE
DESCRIPTOR.message_types_by_name['SlaveDeleteRequest'] = _SLAVEDELETEREQUEST
DESCRIPTOR.message_types_by_name['SlaveRestartRequest'] = _SLAVERESTARTREQUEST
DESCRIPTOR.message_types_by_name['SlaveRestartResponse'] = _SLAVERESTARTRESPONSE

SlaveListResponse = _reflection.GeneratedProtocolMessageType('SlaveListResponse', (_message.Message,), dict(
  DESCRIPTOR = _SLAVELISTRESPONSE,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveListResponse)
  ))
_sym_db.RegisterMessage(SlaveListResponse)

Slave = _reflection.GeneratedProtocolMessageType('Slave', (_message.Message,), dict(
  DESCRIPTOR = _SLAVE,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.Slave)
  ))
_sym_db.RegisterMessage(Slave)

SlaveCreateRequest = _reflection.GeneratedProtocolMessageType('SlaveCreateRequest', (_message.Message,), dict(
  DESCRIPTOR = _SLAVECREATEREQUEST,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveCreateRequest)
  ))
_sym_db.RegisterMessage(SlaveCreateRequest)

SlaveDescribeRequest = _reflection.GeneratedProtocolMessageType('SlaveDescribeRequest', (_message.Message,), dict(
  DESCRIPTOR = _SLAVEDESCRIBEREQUEST,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveDescribeRequest)
  ))
_sym_db.RegisterMessage(SlaveDescribeRequest)

SlaveDescribeResponse = _reflection.GeneratedProtocolMessageType('SlaveDescribeResponse', (_message.Message,), dict(
  DESCRIPTOR = _SLAVEDESCRIBERESPONSE,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveDescribeResponse)
  ))
_sym_db.RegisterMessage(SlaveDescribeResponse)

SlaveDeleteRequest = _reflection.GeneratedProtocolMessageType('SlaveDeleteRequest', (_message.Message,), dict(
  DESCRIPTOR = _SLAVEDELETEREQUEST,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveDeleteRequest)
  ))
_sym_db.RegisterMessage(SlaveDeleteRequest)

SlaveRestartRequest = _reflection.GeneratedProtocolMessageType('SlaveRestartRequest', (_message.Message,), dict(
  DESCRIPTOR = _SLAVERESTARTREQUEST,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveRestartRequest)
  ))
_sym_db.RegisterMessage(SlaveRestartRequest)

SlaveRestartResponse = _reflection.GeneratedProtocolMessageType('SlaveRestartResponse', (_message.Message,), dict(
  DESCRIPTOR = _SLAVERESTARTRESPONSE,
  __module__ = 'slave_pb2'
  # @@protoc_insertion_point(class_scope:ci.SlaveRestartResponse)
  ))
_sym_db.RegisterMessage(SlaveRestartResponse)


# @@protoc_insertion_point(module_scope)
