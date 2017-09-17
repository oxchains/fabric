# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: peer/transaction.proto

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


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from peer import proposal_response_pb2 as peer_dot_proposal__response__pb2
from common import common_pb2 as common_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='peer/transaction.proto',
  package='protos',
  syntax='proto3',
  serialized_pb=_b('\n\x16peer/transaction.proto\x12\x06protos\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cpeer/proposal_response.proto\x1a\x13\x63ommon/common.proto\"A\n\x11SignedTransaction\x12\x19\n\x11transaction_bytes\x18\x01 \x01(\x0c\x12\x11\n\tsignature\x18\x02 \x01(\x0c\"]\n\x14ProcessedTransaction\x12-\n\x13transactionEnvelope\x18\x01 \x01(\x0b\x32\x10.common.Envelope\x12\x16\n\x0evalidationCode\x18\x02 \x01(\x05\"9\n\x0bTransaction\x12*\n\x07\x61\x63tions\x18\x01 \x03(\x0b\x32\x19.protos.TransactionAction\"4\n\x11TransactionAction\x12\x0e\n\x06header\x18\x01 \x01(\x0c\x12\x0f\n\x07payload\x18\x02 \x01(\x0c\"m\n\x16\x43haincodeActionPayload\x12\"\n\x1a\x63haincode_proposal_payload\x18\x01 \x01(\x0c\x12/\n\x06\x61\x63tion\x18\x02 \x01(\x0b\x32\x1f.protos.ChaincodeEndorsedAction\"g\n\x17\x43haincodeEndorsedAction\x12!\n\x19proposal_response_payload\x18\x01 \x01(\x0c\x12)\n\x0c\x65ndorsements\x18\x02 \x03(\x0b\x32\x13.protos.Endorsement*\xc2\x03\n\x10TxValidationCode\x12\t\n\x05VALID\x10\x00\x12\x10\n\x0cNIL_ENVELOPE\x10\x01\x12\x0f\n\x0b\x42\x41\x44_PAYLOAD\x10\x02\x12\x15\n\x11\x42\x41\x44_COMMON_HEADER\x10\x03\x12\x19\n\x15\x42\x41\x44_CREATOR_SIGNATURE\x10\x04\x12 \n\x1cINVALID_ENDORSER_TRANSACTION\x10\x05\x12\x1e\n\x1aINVALID_CONFIG_TRANSACTION\x10\x06\x12\x1a\n\x16UNSUPPORTED_TX_PAYLOAD\x10\x07\x12\x15\n\x11\x42\x41\x44_PROPOSAL_TXID\x10\x08\x12\x12\n\x0e\x44UPLICATE_TXID\x10\t\x12\x1e\n\x1a\x45NDORSEMENT_POLICY_FAILURE\x10\n\x12\x16\n\x12MVCC_READ_CONFLICT\x10\x0b\x12\x19\n\x15PHANTOM_READ_CONFLICT\x10\x0c\x12\x13\n\x0fUNKNOWN_TX_TYPE\x10\r\x12\x1a\n\x16TARGET_CHAIN_NOT_FOUND\x10\x0e\x12\x14\n\x10MARSHAL_TX_ERROR\x10\x0f\x12\x10\n\x0cNIL_TXACTION\x10\x10\x12\x19\n\x14INVALID_OTHER_REASON\x10\xff\x01\x42\x63\n\"org.hyperledger.fabric.protos.peerB\x12TransactionPackageZ)github.com/oxchains/fabric/protos/peerb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,peer_dot_proposal__response__pb2.DESCRIPTOR,common_dot_common__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_TXVALIDATIONCODE = _descriptor.EnumDescriptor(
  name='TxValidationCode',
  full_name='protos.TxValidationCode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='VALID', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NIL_ENVELOPE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BAD_PAYLOAD', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BAD_COMMON_HEADER', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BAD_CREATOR_SIGNATURE', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INVALID_ENDORSER_TRANSACTION', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INVALID_CONFIG_TRANSACTION', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNSUPPORTED_TX_PAYLOAD', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BAD_PROPOSAL_TXID', index=8, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DUPLICATE_TXID', index=9, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ENDORSEMENT_POLICY_FAILURE', index=10, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MVCC_READ_CONFLICT', index=11, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PHANTOM_READ_CONFLICT', index=12, number=12,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN_TX_TYPE', index=13, number=13,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TARGET_CHAIN_NOT_FOUND', index=14, number=14,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MARSHAL_TX_ERROR', index=15, number=15,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NIL_TXACTION', index=16, number=16,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INVALID_OTHER_REASON', index=17, number=255,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=610,
  serialized_end=1060,
)
_sym_db.RegisterEnumDescriptor(_TXVALIDATIONCODE)

TxValidationCode = enum_type_wrapper.EnumTypeWrapper(_TXVALIDATIONCODE)
VALID = 0
NIL_ENVELOPE = 1
BAD_PAYLOAD = 2
BAD_COMMON_HEADER = 3
BAD_CREATOR_SIGNATURE = 4
INVALID_ENDORSER_TRANSACTION = 5
INVALID_CONFIG_TRANSACTION = 6
UNSUPPORTED_TX_PAYLOAD = 7
BAD_PROPOSAL_TXID = 8
DUPLICATE_TXID = 9
ENDORSEMENT_POLICY_FAILURE = 10
MVCC_READ_CONFLICT = 11
PHANTOM_READ_CONFLICT = 12
UNKNOWN_TX_TYPE = 13
TARGET_CHAIN_NOT_FOUND = 14
MARSHAL_TX_ERROR = 15
NIL_TXACTION = 16
INVALID_OTHER_REASON = 255



_SIGNEDTRANSACTION = _descriptor.Descriptor(
  name='SignedTransaction',
  full_name='protos.SignedTransaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='transaction_bytes', full_name='protos.SignedTransaction.transaction_bytes', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='signature', full_name='protos.SignedTransaction.signature', index=1,
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
  serialized_start=118,
  serialized_end=183,
)


_PROCESSEDTRANSACTION = _descriptor.Descriptor(
  name='ProcessedTransaction',
  full_name='protos.ProcessedTransaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='transactionEnvelope', full_name='protos.ProcessedTransaction.transactionEnvelope', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='validationCode', full_name='protos.ProcessedTransaction.validationCode', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=185,
  serialized_end=278,
)


_TRANSACTION = _descriptor.Descriptor(
  name='Transaction',
  full_name='protos.Transaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='actions', full_name='protos.Transaction.actions', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=280,
  serialized_end=337,
)


_TRANSACTIONACTION = _descriptor.Descriptor(
  name='TransactionAction',
  full_name='protos.TransactionAction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='protos.TransactionAction.header', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='payload', full_name='protos.TransactionAction.payload', index=1,
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
  serialized_start=339,
  serialized_end=391,
)


_CHAINCODEACTIONPAYLOAD = _descriptor.Descriptor(
  name='ChaincodeActionPayload',
  full_name='protos.ChaincodeActionPayload',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='chaincode_proposal_payload', full_name='protos.ChaincodeActionPayload.chaincode_proposal_payload', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='action', full_name='protos.ChaincodeActionPayload.action', index=1,
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
  serialized_start=393,
  serialized_end=502,
)


_CHAINCODEENDORSEDACTION = _descriptor.Descriptor(
  name='ChaincodeEndorsedAction',
  full_name='protos.ChaincodeEndorsedAction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='proposal_response_payload', full_name='protos.ChaincodeEndorsedAction.proposal_response_payload', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='endorsements', full_name='protos.ChaincodeEndorsedAction.endorsements', index=1,
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
  serialized_start=504,
  serialized_end=607,
)

_PROCESSEDTRANSACTION.fields_by_name['transactionEnvelope'].message_type = common_dot_common__pb2._ENVELOPE
_TRANSACTION.fields_by_name['actions'].message_type = _TRANSACTIONACTION
_CHAINCODEACTIONPAYLOAD.fields_by_name['action'].message_type = _CHAINCODEENDORSEDACTION
_CHAINCODEENDORSEDACTION.fields_by_name['endorsements'].message_type = peer_dot_proposal__response__pb2._ENDORSEMENT
DESCRIPTOR.message_types_by_name['SignedTransaction'] = _SIGNEDTRANSACTION
DESCRIPTOR.message_types_by_name['ProcessedTransaction'] = _PROCESSEDTRANSACTION
DESCRIPTOR.message_types_by_name['Transaction'] = _TRANSACTION
DESCRIPTOR.message_types_by_name['TransactionAction'] = _TRANSACTIONACTION
DESCRIPTOR.message_types_by_name['ChaincodeActionPayload'] = _CHAINCODEACTIONPAYLOAD
DESCRIPTOR.message_types_by_name['ChaincodeEndorsedAction'] = _CHAINCODEENDORSEDACTION
DESCRIPTOR.enum_types_by_name['TxValidationCode'] = _TXVALIDATIONCODE

SignedTransaction = _reflection.GeneratedProtocolMessageType('SignedTransaction', (_message.Message,), dict(
  DESCRIPTOR = _SIGNEDTRANSACTION,
  __module__ = 'peer.transaction_pb2'
  # @@protoc_insertion_point(class_scope:protos.SignedTransaction)
  ))
_sym_db.RegisterMessage(SignedTransaction)

ProcessedTransaction = _reflection.GeneratedProtocolMessageType('ProcessedTransaction', (_message.Message,), dict(
  DESCRIPTOR = _PROCESSEDTRANSACTION,
  __module__ = 'peer.transaction_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProcessedTransaction)
  ))
_sym_db.RegisterMessage(ProcessedTransaction)

Transaction = _reflection.GeneratedProtocolMessageType('Transaction', (_message.Message,), dict(
  DESCRIPTOR = _TRANSACTION,
  __module__ = 'peer.transaction_pb2'
  # @@protoc_insertion_point(class_scope:protos.Transaction)
  ))
_sym_db.RegisterMessage(Transaction)

TransactionAction = _reflection.GeneratedProtocolMessageType('TransactionAction', (_message.Message,), dict(
  DESCRIPTOR = _TRANSACTIONACTION,
  __module__ = 'peer.transaction_pb2'
  # @@protoc_insertion_point(class_scope:protos.TransactionAction)
  ))
_sym_db.RegisterMessage(TransactionAction)

ChaincodeActionPayload = _reflection.GeneratedProtocolMessageType('ChaincodeActionPayload', (_message.Message,), dict(
  DESCRIPTOR = _CHAINCODEACTIONPAYLOAD,
  __module__ = 'peer.transaction_pb2'
  # @@protoc_insertion_point(class_scope:protos.ChaincodeActionPayload)
  ))
_sym_db.RegisterMessage(ChaincodeActionPayload)

ChaincodeEndorsedAction = _reflection.GeneratedProtocolMessageType('ChaincodeEndorsedAction', (_message.Message,), dict(
  DESCRIPTOR = _CHAINCODEENDORSEDACTION,
  __module__ = 'peer.transaction_pb2'
  # @@protoc_insertion_point(class_scope:protos.ChaincodeEndorsedAction)
  ))
_sym_db.RegisterMessage(ChaincodeEndorsedAction)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\"org.hyperledger.fabric.protos.peerB\022TransactionPackageZ)github.com/oxchains/fabric/protos/peer'))
try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
