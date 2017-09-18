# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: msp/msp_config.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='msp/msp_config.proto',
  package='msp',
  syntax='proto3',
  serialized_pb=_b('\n\x14msp/msp_config.proto\x12\x03msp\")\n\tMSPConfig\x12\x0c\n\x04type\x18\x01 \x01(\x05\x12\x0e\n\x06\x63onfig\x18\x02 \x01(\x0c\"\xd6\x02\n\x0f\x46\x61\x62ricMSPConfig\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x12\n\nroot_certs\x18\x02 \x03(\x0c\x12\x1a\n\x12intermediate_certs\x18\x03 \x03(\x0c\x12\x0e\n\x06\x61\x64mins\x18\x04 \x03(\x0c\x12\x17\n\x0frevocation_list\x18\x05 \x03(\x0c\x12\x32\n\x10signing_identity\x18\x06 \x01(\x0b\x32\x18.msp.SigningIdentityInfo\x12@\n\x1forganizational_unit_identifiers\x18\x07 \x03(\x0b\x32\x17.msp.FabricOUIdentifier\x12.\n\rcrypto_config\x18\x08 \x01(\x0b\x32\x17.msp.FabricCryptoConfig\x12\x16\n\x0etls_root_certs\x18\t \x03(\x0c\x12\x1e\n\x16tls_intermediate_certs\x18\n \x03(\x0c\"^\n\x12\x46\x61\x62ricCryptoConfig\x12\x1d\n\x15signature_hash_family\x18\x01 \x01(\t\x12)\n!identity_identifier_hash_function\x18\x02 \x01(\t\"R\n\x13SigningIdentityInfo\x12\x15\n\rpublic_signer\x18\x01 \x01(\x0c\x12$\n\x0eprivate_signer\x18\x02 \x01(\x0b\x32\x0c.msp.KeyInfo\"7\n\x07KeyInfo\x12\x16\n\x0ekey_identifier\x18\x01 \x01(\t\x12\x14\n\x0ckey_material\x18\x02 \x01(\x0c\"Q\n\x12\x46\x61\x62ricOUIdentifier\x12\x13\n\x0b\x63\x65rtificate\x18\x01 \x01(\x0c\x12&\n\x1eorganizational_unit_identifier\x18\x02 \x01(\tB_\n!org.hyperledger.fabric.protos.mspB\x10MspConfigPackageZ(github.com/oxchains/fabric/protos/mspb\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_MSPCONFIG = _descriptor.Descriptor(
  name='MSPConfig',
  full_name='msp.MSPConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='msp.MSPConfig.type', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='config', full_name='msp.MSPConfig.config', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=29,
  serialized_end=70,
)


_FABRICMSPCONFIG = _descriptor.Descriptor(
  name='FabricMSPConfig',
  full_name='msp.FabricMSPConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='msp.FabricMSPConfig.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='root_certs', full_name='msp.FabricMSPConfig.root_certs', index=1,
      number=2, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='intermediate_certs', full_name='msp.FabricMSPConfig.intermediate_certs', index=2,
      number=3, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='admins', full_name='msp.FabricMSPConfig.admins', index=3,
      number=4, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='revocation_list', full_name='msp.FabricMSPConfig.revocation_list', index=4,
      number=5, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='signing_identity', full_name='msp.FabricMSPConfig.signing_identity', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='organizational_unit_identifiers', full_name='msp.FabricMSPConfig.organizational_unit_identifiers', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='crypto_config', full_name='msp.FabricMSPConfig.crypto_config', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tls_root_certs', full_name='msp.FabricMSPConfig.tls_root_certs', index=8,
      number=9, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tls_intermediate_certs', full_name='msp.FabricMSPConfig.tls_intermediate_certs', index=9,
      number=10, type=12, cpp_type=9, label=3,
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
  serialized_end=415,
)


_FABRICCRYPTOCONFIG = _descriptor.Descriptor(
  name='FabricCryptoConfig',
  full_name='msp.FabricCryptoConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='signature_hash_family', full_name='msp.FabricCryptoConfig.signature_hash_family', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='identity_identifier_hash_function', full_name='msp.FabricCryptoConfig.identity_identifier_hash_function', index=1,
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
  serialized_start=417,
  serialized_end=511,
)


_SIGNINGIDENTITYINFO = _descriptor.Descriptor(
  name='SigningIdentityInfo',
  full_name='msp.SigningIdentityInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='public_signer', full_name='msp.SigningIdentityInfo.public_signer', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='private_signer', full_name='msp.SigningIdentityInfo.private_signer', index=1,
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
  serialized_start=513,
  serialized_end=595,
)


_KEYINFO = _descriptor.Descriptor(
  name='KeyInfo',
  full_name='msp.KeyInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key_identifier', full_name='msp.KeyInfo.key_identifier', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='key_material', full_name='msp.KeyInfo.key_material', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_end=652,
)


_FABRICOUIDENTIFIER = _descriptor.Descriptor(
  name='FabricOUIdentifier',
  full_name='msp.FabricOUIdentifier',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='certificate', full_name='msp.FabricOUIdentifier.certificate', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='organizational_unit_identifier', full_name='msp.FabricOUIdentifier.organizational_unit_identifier', index=1,
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
  serialized_start=654,
  serialized_end=735,
)

_FABRICMSPCONFIG.fields_by_name['signing_identity'].message_type = _SIGNINGIDENTITYINFO
_FABRICMSPCONFIG.fields_by_name['organizational_unit_identifiers'].message_type = _FABRICOUIDENTIFIER
_FABRICMSPCONFIG.fields_by_name['crypto_config'].message_type = _FABRICCRYPTOCONFIG
_SIGNINGIDENTITYINFO.fields_by_name['private_signer'].message_type = _KEYINFO
DESCRIPTOR.message_types_by_name['MSPConfig'] = _MSPCONFIG
DESCRIPTOR.message_types_by_name['FabricMSPConfig'] = _FABRICMSPCONFIG
DESCRIPTOR.message_types_by_name['FabricCryptoConfig'] = _FABRICCRYPTOCONFIG
DESCRIPTOR.message_types_by_name['SigningIdentityInfo'] = _SIGNINGIDENTITYINFO
DESCRIPTOR.message_types_by_name['KeyInfo'] = _KEYINFO
DESCRIPTOR.message_types_by_name['FabricOUIdentifier'] = _FABRICOUIDENTIFIER

MSPConfig = _reflection.GeneratedProtocolMessageType('MSPConfig', (_message.Message,), dict(
  DESCRIPTOR = _MSPCONFIG,
  __module__ = 'msp.msp_config_pb2'
  # @@protoc_insertion_point(class_scope:msp.MSPConfig)
  ))
_sym_db.RegisterMessage(MSPConfig)

FabricMSPConfig = _reflection.GeneratedProtocolMessageType('FabricMSPConfig', (_message.Message,), dict(
  DESCRIPTOR = _FABRICMSPCONFIG,
  __module__ = 'msp.msp_config_pb2'
  # @@protoc_insertion_point(class_scope:msp.FabricMSPConfig)
  ))
_sym_db.RegisterMessage(FabricMSPConfig)

FabricCryptoConfig = _reflection.GeneratedProtocolMessageType('FabricCryptoConfig', (_message.Message,), dict(
  DESCRIPTOR = _FABRICCRYPTOCONFIG,
  __module__ = 'msp.msp_config_pb2'
  # @@protoc_insertion_point(class_scope:msp.FabricCryptoConfig)
  ))
_sym_db.RegisterMessage(FabricCryptoConfig)

SigningIdentityInfo = _reflection.GeneratedProtocolMessageType('SigningIdentityInfo', (_message.Message,), dict(
  DESCRIPTOR = _SIGNINGIDENTITYINFO,
  __module__ = 'msp.msp_config_pb2'
  # @@protoc_insertion_point(class_scope:msp.SigningIdentityInfo)
  ))
_sym_db.RegisterMessage(SigningIdentityInfo)

KeyInfo = _reflection.GeneratedProtocolMessageType('KeyInfo', (_message.Message,), dict(
  DESCRIPTOR = _KEYINFO,
  __module__ = 'msp.msp_config_pb2'
  # @@protoc_insertion_point(class_scope:msp.KeyInfo)
  ))
_sym_db.RegisterMessage(KeyInfo)

FabricOUIdentifier = _reflection.GeneratedProtocolMessageType('FabricOUIdentifier', (_message.Message,), dict(
  DESCRIPTOR = _FABRICOUIDENTIFIER,
  __module__ = 'msp.msp_config_pb2'
  # @@protoc_insertion_point(class_scope:msp.FabricOUIdentifier)
  ))
_sym_db.RegisterMessage(FabricOUIdentifier)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n!org.hyperledger.fabric.protos.mspB\020MspConfigPackageZ(github.com/oxchains/fabric/protos/msp'))
try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
