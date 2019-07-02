# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/parameter.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from flyteidl.core import types_pb2 as flyteidl_dot_core_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/parameter.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_pb=_b('\n\x1d\x66lyteidl/core/parameter.proto\x12\rflyteidl.core\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x19\x66lyteidl/core/types.proto\"|\n\tParameter\x12$\n\x03var\x18\x01 \x01(\x0b\x32\x17.flyteidl.core.Variable\x12)\n\x07\x64\x65\x66\x61ult\x18\x02 \x01(\x0b\x32\x16.flyteidl.core.LiteralH\x00\x12\x12\n\x08required\x18\x03 \x01(\x08H\x00\x42\n\n\x08\x62\x65havior\"\x9c\x01\n\x0cParameterMap\x12?\n\nparameters\x18\x01 \x03(\x0b\x32+.flyteidl.core.ParameterMap.ParametersEntry\x1aK\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\'\n\x05value\x18\x02 \x01(\x0b\x32\x18.flyteidl.core.Parameter:\x02\x38\x01\x42\x32Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,flyteidl_dot_core_dot_types__pb2.DESCRIPTOR,])




_PARAMETER = _descriptor.Descriptor(
  name='Parameter',
  full_name='flyteidl.core.Parameter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='var', full_name='flyteidl.core.Parameter.var', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default', full_name='flyteidl.core.Parameter.default', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='required', full_name='flyteidl.core.Parameter.required', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
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
      name='behavior', full_name='flyteidl.core.Parameter.behavior',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=105,
  serialized_end=229,
)


_PARAMETERMAP_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='flyteidl.core.ParameterMap.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='flyteidl.core.ParameterMap.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='flyteidl.core.ParameterMap.ParametersEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=313,
  serialized_end=388,
)

_PARAMETERMAP = _descriptor.Descriptor(
  name='ParameterMap',
  full_name='flyteidl.core.ParameterMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='parameters', full_name='flyteidl.core.ParameterMap.parameters', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_PARAMETERMAP_PARAMETERSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=232,
  serialized_end=388,
)

_PARAMETER.fields_by_name['var'].message_type = flyteidl_dot_core_dot_types__pb2._VARIABLE
_PARAMETER.fields_by_name['default'].message_type = flyteidl_dot_core_dot_literals__pb2._LITERAL
_PARAMETER.oneofs_by_name['behavior'].fields.append(
  _PARAMETER.fields_by_name['default'])
_PARAMETER.fields_by_name['default'].containing_oneof = _PARAMETER.oneofs_by_name['behavior']
_PARAMETER.oneofs_by_name['behavior'].fields.append(
  _PARAMETER.fields_by_name['required'])
_PARAMETER.fields_by_name['required'].containing_oneof = _PARAMETER.oneofs_by_name['behavior']
_PARAMETERMAP_PARAMETERSENTRY.fields_by_name['value'].message_type = _PARAMETER
_PARAMETERMAP_PARAMETERSENTRY.containing_type = _PARAMETERMAP
_PARAMETERMAP.fields_by_name['parameters'].message_type = _PARAMETERMAP_PARAMETERSENTRY
DESCRIPTOR.message_types_by_name['Parameter'] = _PARAMETER
DESCRIPTOR.message_types_by_name['ParameterMap'] = _PARAMETERMAP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Parameter = _reflection.GeneratedProtocolMessageType('Parameter', (_message.Message,), dict(
  DESCRIPTOR = _PARAMETER,
  __module__ = 'flyteidl.core.parameter_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.Parameter)
  ))
_sym_db.RegisterMessage(Parameter)

ParameterMap = _reflection.GeneratedProtocolMessageType('ParameterMap', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _PARAMETERMAP_PARAMETERSENTRY,
    __module__ = 'flyteidl.core.parameter_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl.core.ParameterMap.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _PARAMETERMAP,
  __module__ = 'flyteidl.core.parameter_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.ParameterMap)
  ))
_sym_db.RegisterMessage(ParameterMap)
_sym_db.RegisterMessage(ParameterMap.ParametersEntry)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/core'))
_PARAMETERMAP_PARAMETERSENTRY.has_options = True
_PARAMETERMAP_PARAMETERSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)