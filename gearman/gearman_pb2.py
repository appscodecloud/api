# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gearman.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from api.ci.v0.1 import slave_pb2 as api_dot_ci_dot_v0_dot_1_dot_slave__pb2
from api.kubernetes.v0.1 import clusters_pb2 as api_dot_kubernetes_dot_v0_dot_1_dot_clusters__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='gearman.proto',
  package='gearman',
  syntax='proto3',
  serialized_pb=_b('\n\rgearman.proto\x12\x07gearman\x1a\x17\x61pi/ci/v0.1/slave.proto\x1a\"api/kubernetes/v0.1/clusters.proto\"\xa6\x04\n\tOperation\x12\x0c\n\x04phid\x18\x01 \x01(\t\x12$\n\x04type\x18\x02 \x01(\x0e\x32\x16.gearman.OperationType\x12\x42\n\x16\x63luster_create_request\x18\x03 \x01(\x0b\x32 .kubernetes.ClusterCreateRequestH\x00\x12@\n\x15\x63luster_scale_request\x18\x04 \x01(\x0b\x32\x1f.kubernetes.ClusterScaleRequestH\x00\x12\x42\n\x16\x63luster_delete_request\x18\x05 \x01(\x0b\x32 .kubernetes.ClusterDeleteRequestH\x00\x12\x39\n\x17\x63i_slave_create_request\x18\x06 \x01(\x0b\x32\x16.ci.SlaveCreateRequestH\x00\x12\x39\n\x17\x63i_slave_delete_request\x18\x07 \x01(\x0b\x32\x16.ci.SlaveDeleteRequestH\x00\x12\x42\n\x16\x63luster_update_request\x18\x0b \x01(\x0b\x32 .kubernetes.ClusterUpdateRequestH\x00\x12\x1b\n\x04\x61uth\x18\x08 \x01(\x0b\x32\r.gearman.Auth\x12\x14\n\x0crequest_time\x18\t \x01(\x03\x12#\n\x08metadata\x18\n \x01(\x0b\x32\x11.gearman.MetadataB\t\n\x07request\":\n\x04\x41uth\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x10\n\x08username\x18\x02 \x01(\t\x12\r\n\x05token\x18\x03 \x01(\t\"%\n\x08Metadata\x12\x19\n\x11subscription_phid\x18\x01 \x01(\t*\xa5\x01\n\rOperationType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x12\n\x0e\x43LUSTER_CREATE\x10\x01\x12\x11\n\rCLUSTER_SCALE\x10\x02\x12\x12\n\x0e\x43LUSTER_DELETE\x10\x03\x12\x14\n\x10NAMESPACE_CREATE\x10\x04\x12\x10\n\x0cSLAVE_CREATE\x10\x05\x12\x10\n\x0cSLAVE_DELETE\x10\x06\x12\x12\n\x0e\x43LUSTER_UPDATE\x10\x07\x62\x06proto3')
  ,
  dependencies=[api_dot_ci_dot_v0_dot_1_dot_slave__pb2.DESCRIPTOR,api_dot_kubernetes_dot_v0_dot_1_dot_clusters__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_OPERATIONTYPE = _descriptor.EnumDescriptor(
  name='OperationType',
  full_name='gearman.OperationType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CLUSTER_CREATE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CLUSTER_SCALE', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CLUSTER_DELETE', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NAMESPACE_CREATE', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SLAVE_CREATE', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SLAVE_DELETE', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CLUSTER_UPDATE', index=7, number=7,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=740,
  serialized_end=905,
)
_sym_db.RegisterEnumDescriptor(_OPERATIONTYPE)

OperationType = enum_type_wrapper.EnumTypeWrapper(_OPERATIONTYPE)
UNKNOWN = 0
CLUSTER_CREATE = 1
CLUSTER_SCALE = 2
CLUSTER_DELETE = 3
NAMESPACE_CREATE = 4
SLAVE_CREATE = 5
SLAVE_DELETE = 6
CLUSTER_UPDATE = 7



_OPERATION = _descriptor.Descriptor(
  name='Operation',
  full_name='gearman.Operation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phid', full_name='gearman.Operation.phid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='gearman.Operation.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cluster_create_request', full_name='gearman.Operation.cluster_create_request', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cluster_scale_request', full_name='gearman.Operation.cluster_scale_request', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cluster_delete_request', full_name='gearman.Operation.cluster_delete_request', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ci_slave_create_request', full_name='gearman.Operation.ci_slave_create_request', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ci_slave_delete_request', full_name='gearman.Operation.ci_slave_delete_request', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cluster_update_request', full_name='gearman.Operation.cluster_update_request', index=7,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='auth', full_name='gearman.Operation.auth', index=8,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='request_time', full_name='gearman.Operation.request_time', index=9,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='gearman.Operation.metadata', index=10,
      number=10, type=11, cpp_type=10, label=1,
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
    _descriptor.OneofDescriptor(
      name='request', full_name='gearman.Operation.request',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=88,
  serialized_end=638,
)


_AUTH = _descriptor.Descriptor(
  name='Auth',
  full_name='gearman.Auth',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='gearman.Auth.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='username', full_name='gearman.Auth.username', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='token', full_name='gearman.Auth.token', index=2,
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
  serialized_start=640,
  serialized_end=698,
)


_METADATA = _descriptor.Descriptor(
  name='Metadata',
  full_name='gearman.Metadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='subscription_phid', full_name='gearman.Metadata.subscription_phid', index=0,
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
  serialized_start=700,
  serialized_end=737,
)

_OPERATION.fields_by_name['type'].enum_type = _OPERATIONTYPE
_OPERATION.fields_by_name['cluster_create_request'].message_type = api_dot_kubernetes_dot_v0_dot_1_dot_clusters__pb2._CLUSTERCREATEREQUEST
_OPERATION.fields_by_name['cluster_scale_request'].message_type = api_dot_kubernetes_dot_v0_dot_1_dot_clusters__pb2._CLUSTERSCALEREQUEST
_OPERATION.fields_by_name['cluster_delete_request'].message_type = api_dot_kubernetes_dot_v0_dot_1_dot_clusters__pb2._CLUSTERDELETEREQUEST
_OPERATION.fields_by_name['ci_slave_create_request'].message_type = api_dot_ci_dot_v0_dot_1_dot_slave__pb2._SLAVECREATEREQUEST
_OPERATION.fields_by_name['ci_slave_delete_request'].message_type = api_dot_ci_dot_v0_dot_1_dot_slave__pb2._SLAVEDELETEREQUEST
_OPERATION.fields_by_name['cluster_update_request'].message_type = api_dot_kubernetes_dot_v0_dot_1_dot_clusters__pb2._CLUSTERUPDATEREQUEST
_OPERATION.fields_by_name['auth'].message_type = _AUTH
_OPERATION.fields_by_name['metadata'].message_type = _METADATA
_OPERATION.oneofs_by_name['request'].fields.append(
  _OPERATION.fields_by_name['cluster_create_request'])
_OPERATION.fields_by_name['cluster_create_request'].containing_oneof = _OPERATION.oneofs_by_name['request']
_OPERATION.oneofs_by_name['request'].fields.append(
  _OPERATION.fields_by_name['cluster_scale_request'])
_OPERATION.fields_by_name['cluster_scale_request'].containing_oneof = _OPERATION.oneofs_by_name['request']
_OPERATION.oneofs_by_name['request'].fields.append(
  _OPERATION.fields_by_name['cluster_delete_request'])
_OPERATION.fields_by_name['cluster_delete_request'].containing_oneof = _OPERATION.oneofs_by_name['request']
_OPERATION.oneofs_by_name['request'].fields.append(
  _OPERATION.fields_by_name['ci_slave_create_request'])
_OPERATION.fields_by_name['ci_slave_create_request'].containing_oneof = _OPERATION.oneofs_by_name['request']
_OPERATION.oneofs_by_name['request'].fields.append(
  _OPERATION.fields_by_name['ci_slave_delete_request'])
_OPERATION.fields_by_name['ci_slave_delete_request'].containing_oneof = _OPERATION.oneofs_by_name['request']
_OPERATION.oneofs_by_name['request'].fields.append(
  _OPERATION.fields_by_name['cluster_update_request'])
_OPERATION.fields_by_name['cluster_update_request'].containing_oneof = _OPERATION.oneofs_by_name['request']
DESCRIPTOR.message_types_by_name['Operation'] = _OPERATION
DESCRIPTOR.message_types_by_name['Auth'] = _AUTH
DESCRIPTOR.message_types_by_name['Metadata'] = _METADATA
DESCRIPTOR.enum_types_by_name['OperationType'] = _OPERATIONTYPE

Operation = _reflection.GeneratedProtocolMessageType('Operation', (_message.Message,), dict(
  DESCRIPTOR = _OPERATION,
  __module__ = 'gearman_pb2'
  # @@protoc_insertion_point(class_scope:gearman.Operation)
  ))
_sym_db.RegisterMessage(Operation)

Auth = _reflection.GeneratedProtocolMessageType('Auth', (_message.Message,), dict(
  DESCRIPTOR = _AUTH,
  __module__ = 'gearman_pb2'
  # @@protoc_insertion_point(class_scope:gearman.Auth)
  ))
_sym_db.RegisterMessage(Auth)

Metadata = _reflection.GeneratedProtocolMessageType('Metadata', (_message.Message,), dict(
  DESCRIPTOR = _METADATA,
  __module__ = 'gearman_pb2'
  # @@protoc_insertion_point(class_scope:gearman.Metadata)
  ))
_sym_db.RegisterMessage(Metadata)


# @@protoc_insertion_point(module_scope)
