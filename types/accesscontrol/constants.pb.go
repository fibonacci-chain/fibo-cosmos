// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/constants.proto

package accesscontrol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccessType int32

const (
	AccessType_UNKNOWN AccessType = 0
	AccessType_READ    AccessType = 1
	AccessType_WRITE   AccessType = 2
	AccessType_COMMIT  AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "UNKNOWN",
	1: "READ",
	2: "WRITE",
	3: "COMMIT",
}

var AccessType_value = map[string]int32{
	"UNKNOWN": 0,
	"READ":    1,
	"WRITE":   2,
	"COMMIT":  3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{0}
}

type AccessOperationSelectorType int32

const (
	AccessOperationSelectorType_NONE                           AccessOperationSelectorType = 0
	AccessOperationSelectorType_JQ                             AccessOperationSelectorType = 1
	AccessOperationSelectorType_JQ_BECH32_ADDRESS              AccessOperationSelectorType = 2
	AccessOperationSelectorType_JQ_LENGTH_PREFIXED_ADDRESS     AccessOperationSelectorType = 3
	AccessOperationSelectorType_SENDER_BECH32_ADDRESS          AccessOperationSelectorType = 4
	AccessOperationSelectorType_SENDER_LENGTH_PREFIXED_ADDRESS AccessOperationSelectorType = 5
	AccessOperationSelectorType_CONTRACT_ADDRESS               AccessOperationSelectorType = 6
	AccessOperationSelectorType_JQ_MESSAGE_CONDITIONAL         AccessOperationSelectorType = 7
	AccessOperationSelectorType_CONSTANT_STRING_TO_HEX         AccessOperationSelectorType = 8
	AccessOperationSelectorType_CONTRACT_REFERENCE             AccessOperationSelectorType = 9
)

var AccessOperationSelectorType_name = map[int32]string{
	0: "NONE",
	1: "JQ",
	2: "JQ_BECH32_ADDRESS",
	3: "JQ_LENGTH_PREFIXED_ADDRESS",
	4: "SENDER_BECH32_ADDRESS",
	5: "SENDER_LENGTH_PREFIXED_ADDRESS",
	6: "CONTRACT_ADDRESS",
	7: "JQ_MESSAGE_CONDITIONAL",
	8: "CONSTANT_STRING_TO_HEX",
	9: "CONTRACT_REFERENCE",
}

var AccessOperationSelectorType_value = map[string]int32{
	"NONE":                           0,
	"JQ":                             1,
	"JQ_BECH32_ADDRESS":              2,
	"JQ_LENGTH_PREFIXED_ADDRESS":     3,
	"SENDER_BECH32_ADDRESS":          4,
	"SENDER_LENGTH_PREFIXED_ADDRESS": 5,
	"CONTRACT_ADDRESS":               6,
	"JQ_MESSAGE_CONDITIONAL":         7,
	"CONSTANT_STRING_TO_HEX":         8,
	"CONTRACT_REFERENCE":             9,
}

func (x AccessOperationSelectorType) String() string {
	return proto.EnumName(AccessOperationSelectorType_name, int32(x))
}

func (AccessOperationSelectorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

type ResourceType int32

const (
	ResourceType_ANY                                      ResourceType = 0
	ResourceType_KV                                       ResourceType = 1
	ResourceType_Mem                                      ResourceType = 2
	ResourceType_DexMem                                   ResourceType = 3
	ResourceType_KV_BANK                                  ResourceType = 4
	ResourceType_KV_STAKING                               ResourceType = 5
	ResourceType_KV_WASM                                  ResourceType = 6
	ResourceType_KV_ORACLE                                ResourceType = 7
	ResourceType_KV_DEX                                   ResourceType = 8
	ResourceType_KV_EPOCH                                 ResourceType = 9
	ResourceType_KV_TOKENFACTORY                          ResourceType = 10
	ResourceType_KV_ORACLE_VOTE_TARGETS                   ResourceType = 11
	ResourceType_KV_ORACLE_AGGREGATE_VOTES                ResourceType = 12
	ResourceType_KV_ORACLE_FEEDERS                        ResourceType = 13
	ResourceType_KV_STAKING_DELEGATION                    ResourceType = 14
	ResourceType_KV_STAKING_VALIDATOR                     ResourceType = 15
	ResourceType_KV_AUTH                                  ResourceType = 16
	ResourceType_KV_AUTH_ADDRESS_STORE                    ResourceType = 17
	ResourceType_KV_BANK_SUPPLY                           ResourceType = 18
	ResourceType_KV_BANK_DENOM                            ResourceType = 19
	ResourceType_KV_BANK_BALANCES                         ResourceType = 20
	ResourceType_KV_TOKENFACTORY_DENOM                    ResourceType = 21
	ResourceType_KV_TOKENFACTORY_METADATA                 ResourceType = 22
	ResourceType_KV_TOKENFACTORY_ADMIN                    ResourceType = 23
	ResourceType_KV_TOKENFACTORY_CREATOR                  ResourceType = 24
	ResourceType_KV_ORACLE_EXCHANGE_RATE                  ResourceType = 25
	ResourceType_KV_ORACLE_VOTE_PENALTY_COUNTER           ResourceType = 26
	ResourceType_KV_ORACLE_PRICE_SNAPSHOT                 ResourceType = 27
	ResourceType_KV_STAKING_VALIDATION_POWER              ResourceType = 28
	ResourceType_KV_STAKING_TOTAL_POWER                   ResourceType = 29
	ResourceType_KV_STAKING_VALIDATORS_CON_ADDR           ResourceType = 30
	ResourceType_KV_STAKING_UNBONDING_DELEGATION          ResourceType = 31
	ResourceType_KV_STAKING_UNBONDING_DELEGATION_VAL      ResourceType = 32
	ResourceType_KV_STAKING_REDELEGATION                  ResourceType = 33
	ResourceType_KV_STAKING_REDELEGATION_VAL_SRC          ResourceType = 34
	ResourceType_KV_STAKING_REDELEGATION_VAL_DST          ResourceType = 35
	ResourceType_KV_STAKING_REDELEGATION_QUEUE            ResourceType = 36
	ResourceType_KV_STAKING_VALIDATOR_QUEUE               ResourceType = 37
	ResourceType_KV_STAKING_HISTORICAL_INFO               ResourceType = 38
	ResourceType_KV_STAKING_UNBONDING                     ResourceType = 39
	ResourceType_KV_STAKING_VALIDATORS_BY_POWER           ResourceType = 41
	ResourceType_KV_DISTRIBUTION                          ResourceType = 40
	ResourceType_KV_DISTRIBUTION_FEE_POOL                 ResourceType = 42
	ResourceType_KV_DISTRIBUTION_PROPOSER_KEY             ResourceType = 43
	ResourceType_KV_DISTRIBUTION_OUTSTANDING_REWARDS      ResourceType = 44
	ResourceType_KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR  ResourceType = 45
	ResourceType_KV_DISTRIBUTION_DELEGATOR_STARTING_INFO  ResourceType = 46
	ResourceType_KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS   ResourceType = 47
	ResourceType_KV_DISTRIBUTION_VAL_CURRENT_REWARDS      ResourceType = 48
	ResourceType_KV_DISTRIBUTION_VAL_ACCUM_COMMISSION     ResourceType = 49
	ResourceType_KV_DISTRIBUTION_SLASH_EVENT              ResourceType = 50
	ResourceType_KV_DEX_CONTRACT_LONGBOOK                 ResourceType = 51
	ResourceType_KV_DEX_CONTRACT_SHORTBOOK                ResourceType = 52
	ResourceType_KV_DEX_SETTLEMENT                        ResourceType = 53
	ResourceType_KV_DEX_PAIR_PREFIX                       ResourceType = 54
	ResourceType_KV_DEX_TWAP                              ResourceType = 55
	ResourceType_KV_DEX_PRICE                             ResourceType = 56
	ResourceType_KV_DEX_SETTLEMENT_ENTRY                  ResourceType = 57
	ResourceType_KV_DEX_REGISTERED_PAIR                   ResourceType = 58
	ResourceType_KV_DEX_ORDER                             ResourceType = 60
	ResourceType_KV_DEX_CANCEL                            ResourceType = 61
	ResourceType_KV_DEX_ACCOUNT_ACTIVE_ORDERS             ResourceType = 62
	ResourceType_KV_DEX_ASSET_LIST                        ResourceType = 64
	ResourceType_KV_DEX_NEXT_ORDER_ID                     ResourceType = 65
	ResourceType_KV_DEX_NEXT_SETTLEMENT_ID                ResourceType = 66
	ResourceType_KV_DEX_MATCH_RESULT                      ResourceType = 67
	ResourceType_KV_DEX_SETTLEMENT_ORDER_ID               ResourceType = 68
	ResourceType_KV_DEX_ORDER_BOOK                        ResourceType = 69
	ResourceType_KV_ACCESSCONTROL                         ResourceType = 71
	ResourceType_KV_ACCESSCONTROL_WASM_DEPENDENCY_MAPPING ResourceType = 72
	ResourceType_KV_WASM_CODE                             ResourceType = 73
	ResourceType_KV_WASM_CONTRACT_ADDRESS                 ResourceType = 74
	ResourceType_KV_WASM_CONTRACT_STORE                   ResourceType = 75
	ResourceType_KV_WASM_SEQUENCE_KEY                     ResourceType = 76
	ResourceType_KV_WASM_CONTRACT_CODE_HISTORY            ResourceType = 77
	ResourceType_KV_WASM_CONTRACT_BY_CODE_ID              ResourceType = 78
	ResourceType_KV_WASM_PINNED_CODE_INDEX                ResourceType = 79
	ResourceType_KV_AUTH_GLOBAL_ACCOUNT_NUMBER            ResourceType = 80
	ResourceType_KV_AUTHZ                                 ResourceType = 81
	ResourceType_KV_FEEGRANT                              ResourceType = 82
	ResourceType_KV_FEEGRANT_ALLOWANCE                    ResourceType = 83
	ResourceType_KV_SLASHING                              ResourceType = 84
	ResourceType_KV_SLASHING_VAL_SIGNING_INFO             ResourceType = 85
	ResourceType_KV_SLASHING_ADDR_PUBKEY_RELATION_KEY     ResourceType = 86
	ResourceType_KV_DEX_MEM_ORDER                         ResourceType = 87
	ResourceType_KV_DEX_MEM_CANCEL                        ResourceType = 88
	ResourceType_KV_DEX_MEM_DEPOSIT                       ResourceType = 89
	ResourceType_KV_DEX_CONTRACT                          ResourceType = 90
)

var ResourceType_name = map[int32]string{
	0:  "ANY",
	1:  "KV",
	2:  "Mem",
	3:  "DexMem",
	4:  "KV_BANK",
	5:  "KV_STAKING",
	6:  "KV_WASM",
	7:  "KV_ORACLE",
	8:  "KV_DEX",
	9:  "KV_EPOCH",
	10: "KV_TOKENFACTORY",
	11: "KV_ORACLE_VOTE_TARGETS",
	12: "KV_ORACLE_AGGREGATE_VOTES",
	13: "KV_ORACLE_FEEDERS",
	14: "KV_STAKING_DELEGATION",
	15: "KV_STAKING_VALIDATOR",
	16: "KV_AUTH",
	17: "KV_AUTH_ADDRESS_STORE",
	18: "KV_BANK_SUPPLY",
	19: "KV_BANK_DENOM",
	20: "KV_BANK_BALANCES",
	21: "KV_TOKENFACTORY_DENOM",
	22: "KV_TOKENFACTORY_METADATA",
	23: "KV_TOKENFACTORY_ADMIN",
	24: "KV_TOKENFACTORY_CREATOR",
	25: "KV_ORACLE_EXCHANGE_RATE",
	26: "KV_ORACLE_VOTE_PENALTY_COUNTER",
	27: "KV_ORACLE_PRICE_SNAPSHOT",
	28: "KV_STAKING_VALIDATION_POWER",
	29: "KV_STAKING_TOTAL_POWER",
	30: "KV_STAKING_VALIDATORS_CON_ADDR",
	31: "KV_STAKING_UNBONDING_DELEGATION",
	32: "KV_STAKING_UNBONDING_DELEGATION_VAL",
	33: "KV_STAKING_REDELEGATION",
	34: "KV_STAKING_REDELEGATION_VAL_SRC",
	35: "KV_STAKING_REDELEGATION_VAL_DST",
	36: "KV_STAKING_REDELEGATION_QUEUE",
	37: "KV_STAKING_VALIDATOR_QUEUE",
	38: "KV_STAKING_HISTORICAL_INFO",
	39: "KV_STAKING_UNBONDING",
	41: "KV_STAKING_VALIDATORS_BY_POWER",
	40: "KV_DISTRIBUTION",
	42: "KV_DISTRIBUTION_FEE_POOL",
	43: "KV_DISTRIBUTION_PROPOSER_KEY",
	44: "KV_DISTRIBUTION_OUTSTANDING_REWARDS",
	45: "KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR",
	46: "KV_DISTRIBUTION_DELEGATOR_STARTING_INFO",
	47: "KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS",
	48: "KV_DISTRIBUTION_VAL_CURRENT_REWARDS",
	49: "KV_DISTRIBUTION_VAL_ACCUM_COMMISSION",
	50: "KV_DISTRIBUTION_SLASH_EVENT",
	51: "KV_DEX_CONTRACT_LONGBOOK",
	52: "KV_DEX_CONTRACT_SHORTBOOK",
	53: "KV_DEX_SETTLEMENT",
	54: "KV_DEX_PAIR_PREFIX",
	55: "KV_DEX_TWAP",
	56: "KV_DEX_PRICE",
	57: "KV_DEX_SETTLEMENT_ENTRY",
	58: "KV_DEX_REGISTERED_PAIR",
	60: "KV_DEX_ORDER",
	61: "KV_DEX_CANCEL",
	62: "KV_DEX_ACCOUNT_ACTIVE_ORDERS",
	64: "KV_DEX_ASSET_LIST",
	65: "KV_DEX_NEXT_ORDER_ID",
	66: "KV_DEX_NEXT_SETTLEMENT_ID",
	67: "KV_DEX_MATCH_RESULT",
	68: "KV_DEX_SETTLEMENT_ORDER_ID",
	69: "KV_DEX_ORDER_BOOK",
	71: "KV_ACCESSCONTROL",
	72: "KV_ACCESSCONTROL_WASM_DEPENDENCY_MAPPING",
	73: "KV_WASM_CODE",
	74: "KV_WASM_CONTRACT_ADDRESS",
	75: "KV_WASM_CONTRACT_STORE",
	76: "KV_WASM_SEQUENCE_KEY",
	77: "KV_WASM_CONTRACT_CODE_HISTORY",
	78: "KV_WASM_CONTRACT_BY_CODE_ID",
	79: "KV_WASM_PINNED_CODE_INDEX",
	80: "KV_AUTH_GLOBAL_ACCOUNT_NUMBER",
	81: "KV_AUTHZ",
	82: "KV_FEEGRANT",
	83: "KV_FEEGRANT_ALLOWANCE",
	84: "KV_SLASHING",
	85: "KV_SLASHING_VAL_SIGNING_INFO",
	86: "KV_SLASHING_ADDR_PUBKEY_RELATION_KEY",
	87: "KV_DEX_MEM_ORDER",
	88: "KV_DEX_MEM_CANCEL",
	89: "KV_DEX_MEM_DEPOSIT",
	90: "KV_DEX_CONTRACT",
}

var ResourceType_value = map[string]int32{
	"ANY":                                      0,
	"KV":                                       1,
	"Mem":                                      2,
	"DexMem":                                   3,
	"KV_BANK":                                  4,
	"KV_STAKING":                               5,
	"KV_WASM":                                  6,
	"KV_ORACLE":                                7,
	"KV_DEX":                                   8,
	"KV_EPOCH":                                 9,
	"KV_TOKENFACTORY":                          10,
	"KV_ORACLE_VOTE_TARGETS":                   11,
	"KV_ORACLE_AGGREGATE_VOTES":                12,
	"KV_ORACLE_FEEDERS":                        13,
	"KV_STAKING_DELEGATION":                    14,
	"KV_STAKING_VALIDATOR":                     15,
	"KV_AUTH":                                  16,
	"KV_AUTH_ADDRESS_STORE":                    17,
	"KV_BANK_SUPPLY":                           18,
	"KV_BANK_DENOM":                            19,
	"KV_BANK_BALANCES":                         20,
	"KV_TOKENFACTORY_DENOM":                    21,
	"KV_TOKENFACTORY_METADATA":                 22,
	"KV_TOKENFACTORY_ADMIN":                    23,
	"KV_TOKENFACTORY_CREATOR":                  24,
	"KV_ORACLE_EXCHANGE_RATE":                  25,
	"KV_ORACLE_VOTE_PENALTY_COUNTER":           26,
	"KV_ORACLE_PRICE_SNAPSHOT":                 27,
	"KV_STAKING_VALIDATION_POWER":              28,
	"KV_STAKING_TOTAL_POWER":                   29,
	"KV_STAKING_VALIDATORS_CON_ADDR":           30,
	"KV_STAKING_UNBONDING_DELEGATION":          31,
	"KV_STAKING_UNBONDING_DELEGATION_VAL":      32,
	"KV_STAKING_REDELEGATION":                  33,
	"KV_STAKING_REDELEGATION_VAL_SRC":          34,
	"KV_STAKING_REDELEGATION_VAL_DST":          35,
	"KV_STAKING_REDELEGATION_QUEUE":            36,
	"KV_STAKING_VALIDATOR_QUEUE":               37,
	"KV_STAKING_HISTORICAL_INFO":               38,
	"KV_STAKING_UNBONDING":                     39,
	"KV_STAKING_VALIDATORS_BY_POWER":           41,
	"KV_DISTRIBUTION":                          40,
	"KV_DISTRIBUTION_FEE_POOL":                 42,
	"KV_DISTRIBUTION_PROPOSER_KEY":             43,
	"KV_DISTRIBUTION_OUTSTANDING_REWARDS":      44,
	"KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR":  45,
	"KV_DISTRIBUTION_DELEGATOR_STARTING_INFO":  46,
	"KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS":   47,
	"KV_DISTRIBUTION_VAL_CURRENT_REWARDS":      48,
	"KV_DISTRIBUTION_VAL_ACCUM_COMMISSION":     49,
	"KV_DISTRIBUTION_SLASH_EVENT":              50,
	"KV_DEX_CONTRACT_LONGBOOK":                 51,
	"KV_DEX_CONTRACT_SHORTBOOK":                52,
	"KV_DEX_SETTLEMENT":                        53,
	"KV_DEX_PAIR_PREFIX":                       54,
	"KV_DEX_TWAP":                              55,
	"KV_DEX_PRICE":                             56,
	"KV_DEX_SETTLEMENT_ENTRY":                  57,
	"KV_DEX_REGISTERED_PAIR":                   58,
	"KV_DEX_ORDER":                             60,
	"KV_DEX_CANCEL":                            61,
	"KV_DEX_ACCOUNT_ACTIVE_ORDERS":             62,
	"KV_DEX_ASSET_LIST":                        64,
	"KV_DEX_NEXT_ORDER_ID":                     65,
	"KV_DEX_NEXT_SETTLEMENT_ID":                66,
	"KV_DEX_MATCH_RESULT":                      67,
	"KV_DEX_SETTLEMENT_ORDER_ID":               68,
	"KV_DEX_ORDER_BOOK":                        69,
	"KV_ACCESSCONTROL":                         71,
	"KV_ACCESSCONTROL_WASM_DEPENDENCY_MAPPING": 72,
	"KV_WASM_CODE":                             73,
	"KV_WASM_CONTRACT_ADDRESS":                 74,
	"KV_WASM_CONTRACT_STORE":                   75,
	"KV_WASM_SEQUENCE_KEY":                     76,
	"KV_WASM_CONTRACT_CODE_HISTORY":            77,
	"KV_WASM_CONTRACT_BY_CODE_ID":              78,
	"KV_WASM_PINNED_CODE_INDEX":                79,
	"KV_AUTH_GLOBAL_ACCOUNT_NUMBER":            80,
	"KV_AUTHZ":                                 81,
	"KV_FEEGRANT":                              82,
	"KV_FEEGRANT_ALLOWANCE":                    83,
	"KV_SLASHING":                              84,
	"KV_SLASHING_VAL_SIGNING_INFO":             85,
	"KV_SLASHING_ADDR_PUBKEY_RELATION_KEY":     86,
	"KV_DEX_MEM_ORDER":                         87,
	"KV_DEX_MEM_CANCEL":                        88,
	"KV_DEX_MEM_DEPOSIT":                       89,
	"KV_DEX_CONTRACT":                          90,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{2}
}

type WasmMessageSubtype int32

const (
	WasmMessageSubtype_QUERY   WasmMessageSubtype = 0
	WasmMessageSubtype_EXECUTE WasmMessageSubtype = 1
)

var WasmMessageSubtype_name = map[int32]string{
	0: "QUERY",
	1: "EXECUTE",
}

var WasmMessageSubtype_value = map[string]int32{
	"QUERY":   0,
	"EXECUTE": 1,
}

func (x WasmMessageSubtype) String() string {
	return proto.EnumName(WasmMessageSubtype_name, int32(x))
}

func (WasmMessageSubtype) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{3}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessOperationSelectorType", AccessOperationSelectorType_name, AccessOperationSelectorType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.WasmMessageSubtype", WasmMessageSubtype_name, WasmMessageSubtype_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 1404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x5b, 0x77, 0x13, 0xb7,
	0x16, 0xce, 0xfd, 0xa2, 0x04, 0xd8, 0x88, 0x3b, 0x09, 0x06, 0x02, 0x07, 0x38, 0x01, 0x12, 0x2e,
	0xe7, 0x0a, 0xe7, 0x1c, 0x8e, 0x3c, 0xb3, 0x3d, 0x9e, 0xcc, 0x8c, 0x34, 0x96, 0x34, 0xbe, 0xf0,
	0xa2, 0x95, 0xb8, 0x5e, 0x94, 0x55, 0x12, 0xb3, 0x62, 0xd3, 0xd5, 0xfe, 0x86, 0xbe, 0xf4, 0x67,
	0xf5, 0x91, 0xc7, 0x3e, 0x76, 0xc1, 0x53, 0xff, 0x45, 0x97, 0x66, 0x64, 0x33, 0x36, 0xa1, 0x3c,
	0x25, 0xde, 0xdf, 0xa7, 0x3d, 0xda, 0xdf, 0xbe, 0x89, 0xdc, 0xee, 0xf6, 0x07, 0x87, 0xfd, 0xc1,
	0xee, 0x7e, 0xb7, 0xdb, 0x1b, 0x0c, 0xba, 0xfd, 0xa3, 0xe1, 0x71, 0xff, 0xcd, 0x6e, 0xb7, 0x7f,
	0x34, 0x18, 0xee, 0x1f, 0x0d, 0x07, 0x3b, 0x6f, 0x8f, 0xfb, 0xc3, 0x3e, 0xdd, 0x2c, 0x58, 0x3b,
	0x13, 0xac, 0x9d, 0xef, 0x1f, 0x1f, 0xf4, 0x86, 0xfb, 0x8f, 0xb7, 0x9f, 0x11, 0xc2, 0x72, 0x40,
	0xff, 0xf8, 0xb6, 0x47, 0xd7, 0xc8, 0x72, 0xc6, 0x23, 0x2e, 0x5a, 0x1c, 0x66, 0xe8, 0x0a, 0x59,
	0x90, 0xc8, 0x7c, 0x98, 0xa5, 0xab, 0x64, 0xb1, 0x25, 0x43, 0x8d, 0x30, 0x47, 0x09, 0x59, 0xf2,
	0x44, 0x92, 0x84, 0x1a, 0xe6, 0xb7, 0x7f, 0x9a, 0x23, 0x1b, 0xc5, 0x61, 0xf1, 0xb6, 0x77, 0xbc,
	0x3f, 0x7c, 0xdd, 0x3f, 0x52, 0xbd, 0x37, 0xbd, 0xee, 0xb0, 0x7f, 0x9c, 0x7b, 0x5b, 0x21, 0x0b,
	0x5c, 0x70, 0x84, 0x19, 0xba, 0x44, 0xe6, 0xf6, 0x1a, 0x30, 0x4b, 0x2f, 0x90, 0xb3, 0x7b, 0x0d,
	0x53, 0x45, 0xaf, 0xfe, 0xf4, 0x89, 0x61, 0xbe, 0x2f, 0x51, 0x29, 0x98, 0xa3, 0x15, 0x72, 0x75,
	0xaf, 0x61, 0x62, 0xe4, 0x81, 0xae, 0x9b, 0x54, 0x62, 0x2d, 0x6c, 0xa3, 0x3f, 0xc6, 0xe7, 0xe9,
	0x15, 0x72, 0x41, 0x21, 0xf7, 0x51, 0x4e, 0x1f, 0x5d, 0xa0, 0x5b, 0xa4, 0xe2, 0xa0, 0x2f, 0x1d,
	0x5f, 0xa4, 0xe7, 0x09, 0x78, 0x82, 0x6b, 0xc9, 0x3c, 0x3d, 0xb6, 0x2e, 0xd1, 0xab, 0xe4, 0xe2,
	0x5e, 0xc3, 0x24, 0xa8, 0x14, 0x0b, 0xd0, 0x78, 0x82, 0xfb, 0xa1, 0x0e, 0x05, 0x67, 0x31, 0x2c,
	0x5b, 0xcc, 0x13, 0x5c, 0x69, 0xc6, 0xb5, 0x51, 0x5a, 0x86, 0x3c, 0x30, 0x5a, 0x98, 0x3a, 0xb6,
	0x61, 0x85, 0x5e, 0x24, 0x74, 0xec, 0x4d, 0x62, 0x0d, 0x25, 0x72, 0x0f, 0x61, 0x75, 0xfb, 0x77,
	0x4a, 0xd6, 0x65, 0x6f, 0xd0, 0x7f, 0x77, 0xdc, 0xed, 0xe5, 0xe1, 0x2f, 0x93, 0x79, 0xc6, 0x3b,
	0x45, 0xf4, 0x51, 0x13, 0x66, 0xad, 0x21, 0xe9, 0x1d, 0x16, 0x22, 0xfa, 0xbd, 0x1f, 0xec, 0xff,
	0xf3, 0x56, 0xf2, 0xa8, 0x69, 0xaa, 0x8c, 0x47, 0xb0, 0x40, 0x4f, 0x13, 0x12, 0x35, 0x8d, 0xd2,
	0x2c, 0x0a, 0x79, 0x00, 0x8b, 0x0e, 0x6c, 0x31, 0x95, 0xc0, 0x12, 0x3d, 0x45, 0x56, 0xa3, 0xa6,
	0x11, 0x92, 0x79, 0x31, 0xc2, 0xb2, 0x75, 0x12, 0x35, 0x8d, 0x9f, 0xdf, 0x69, 0x9d, 0xac, 0x44,
	0x4d, 0x83, 0xa9, 0xf0, 0xea, 0xb0, 0x4a, 0xcf, 0x91, 0x33, 0x51, 0xd3, 0x68, 0x11, 0x21, 0xaf,
	0x31, 0x4f, 0x0b, 0xd9, 0x01, 0x62, 0x43, 0x1a, 0x9f, 0x36, 0x4d, 0xa1, 0xd1, 0x68, 0x26, 0x03,
	0xd4, 0x0a, 0xd6, 0xe8, 0x35, 0x72, 0xe5, 0x13, 0xc6, 0x82, 0x40, 0x62, 0xc0, 0x74, 0xc1, 0x52,
	0xb0, 0x6e, 0xb3, 0xf6, 0x09, 0xae, 0x21, 0xfa, 0x28, 0x15, 0x9c, 0xb2, 0x59, 0xf9, 0x74, 0x59,
	0xe3, 0x63, 0x6c, 0x4f, 0x85, 0x82, 0xc3, 0x69, 0x7a, 0x99, 0x9c, 0x2f, 0x41, 0x4d, 0x16, 0x87,
	0x3e, 0xd3, 0x42, 0xc2, 0x19, 0x17, 0x11, 0xcb, 0x74, 0x1d, 0xc0, 0x79, 0xb0, 0x3f, 0x46, 0x79,
	0x31, 0x4a, 0x0b, 0x89, 0x70, 0x96, 0x52, 0x72, 0xda, 0xc9, 0x62, 0x54, 0x96, 0xa6, 0x71, 0x07,
	0x28, 0x3d, 0x4b, 0x4e, 0x8d, 0x6c, 0x3e, 0x72, 0x91, 0xc0, 0x39, 0x9b, 0xda, 0x91, 0xa9, 0xca,
	0x62, 0xc6, 0x3d, 0x54, 0x70, 0xde, 0xf9, 0x2d, 0x0b, 0xe0, 0x0e, 0x5c, 0xa0, 0x9b, 0xe4, 0xf2,
	0x34, 0x94, 0xa0, 0x66, 0x3e, 0xd3, 0x0c, 0x2e, 0x9e, 0x74, 0x90, 0xf9, 0x49, 0xc8, 0xe1, 0x12,
	0xdd, 0x20, 0x97, 0xa6, 0x21, 0x4f, 0x62, 0x1e, 0xd5, 0x65, 0x07, 0x3a, 0x85, 0xb0, 0xed, 0xd5,
	0x19, 0x0f, 0xd0, 0x48, 0xa6, 0x11, 0xae, 0xd8, 0x12, 0x9d, 0x52, 0x3e, 0x45, 0xce, 0x62, 0xdd,
	0x31, 0x9e, 0xc8, 0xb8, 0x46, 0x09, 0x57, 0xdd, 0xb5, 0x1c, 0x27, 0x95, 0xa1, 0x87, 0x46, 0x71,
	0x96, 0xaa, 0xba, 0xd0, 0xb0, 0x41, 0xaf, 0x93, 0x8d, 0xcf, 0xe5, 0x0c, 0x05, 0x37, 0xa9, 0x68,
	0xa1, 0x84, 0x4d, 0x97, 0xdc, 0x11, 0x41, 0x0b, 0xcd, 0x62, 0x87, 0x5d, 0x73, 0x9f, 0xff, 0x2c,
	0x17, 0xca, 0x96, 0x7c, 0x2e, 0x3b, 0x54, 0xe8, 0x2d, 0x72, 0xbd, 0xc4, 0xc9, 0x78, 0xd5, 0x76,
	0xc3, 0x64, 0x52, 0xaf, 0xd3, 0xbb, 0xe4, 0xd6, 0x57, 0x48, 0xd6, 0x3b, 0xdc, 0x70, 0x6a, 0x8c,
	0x88, 0x12, 0x4b, 0x5e, 0x6e, 0x4e, 0x7d, 0xaa, 0x0c, 0xda, 0xd3, 0x46, 0x49, 0x0f, 0xb6, 0xbe,
	0x46, 0xf2, 0x95, 0x86, 0x5b, 0xf4, 0x26, 0xb9, 0xf6, 0x25, 0x52, 0x23, 0xc3, 0x0c, 0xe1, 0xb6,
	0x1d, 0x2c, 0x27, 0xc5, 0xee, 0xf0, 0xbf, 0x4c, 0xe1, 0xf5, 0xd0, 0x56, 0x5f, 0xe8, 0xb1, 0xd8,
	0x84, 0xbc, 0x26, 0xe0, 0xce, 0x54, 0x1d, 0x8f, 0x43, 0x86, 0xbb, 0x5f, 0x56, 0xb5, 0xda, 0x71,
	0xca, 0xff, 0xd5, 0xf5, 0xa1, 0x1f, 0xda, 0x09, 0x52, 0xcd, 0xf2, 0xf8, 0xef, 0xb9, 0x4c, 0x97,
	0x8d, 0xb6, 0xa5, 0x4c, 0x2a, 0x44, 0x0c, 0xdb, 0xf4, 0x06, 0xd9, 0x9c, 0x46, 0x53, 0x29, 0x52,
	0xa1, 0x50, 0x9a, 0x08, 0x3b, 0x70, 0xdf, 0x65, 0x61, 0x82, 0x21, 0x32, 0x6d, 0x47, 0x95, 0x5f,
	0xc8, 0xd0, 0x62, 0xd2, 0x57, 0xf0, 0x80, 0xde, 0x27, 0x77, 0xa7, 0x89, 0x4e, 0x21, 0x21, 0x4d,
	0x2b, 0xd4, 0x75, 0x5f, 0xb2, 0x56, 0x51, 0x00, 0x0f, 0xff, 0x9c, 0xac, 0x34, 0x93, 0xda, 0x3a,
	0xcf, 0x55, 0xd9, 0xa1, 0xdb, 0xe4, 0xce, 0x34, 0xd9, 0x66, 0xa5, 0x24, 0xdf, 0xe8, 0x16, 0xbb,
	0x27, 0x5d, 0xd7, 0x72, 0xbd, 0x4c, 0x4a, 0xe4, 0x7a, 0x4c, 0x7c, 0x44, 0xef, 0x91, 0xdb, 0x27,
	0x11, 0x99, 0xe7, 0x65, 0x89, 0xc9, 0x57, 0x8e, 0x52, 0x56, 0xc1, 0xc7, 0xae, 0x1b, 0x26, 0x98,
	0x2a, 0x66, 0xaa, 0x6e, 0xb0, 0x89, 0x5c, 0xc3, 0x93, 0x91, 0xc4, 0xd8, 0x36, 0xe3, 0x41, 0x1d,
	0x0b, 0x1e, 0x54, 0x85, 0x88, 0xe0, 0xa9, 0x1b, 0x76, 0x13, 0xa8, 0xaa, 0x0b, 0xa9, 0x73, 0xf8,
	0x6f, 0x6e, 0xd8, 0x59, 0x58, 0xa1, 0xd6, 0x31, 0x26, 0xd6, 0xe7, 0xdf, 0xed, 0xd4, 0x77, 0xe6,
	0x94, 0x85, 0xd2, 0x6d, 0x19, 0xf8, 0x07, 0x3d, 0x43, 0xd6, 0x9c, 0x5d, 0xb7, 0x58, 0x0a, 0xff,
	0xa4, 0x40, 0xd6, 0x47, 0x44, 0xdb, 0xc6, 0xf0, 0x2f, 0xd7, 0x0e, 0x93, 0x1e, 0x0d, 0x72, 0x2d,
	0x3b, 0xf0, 0x6f, 0xd7, 0xb9, 0x16, 0x94, 0x18, 0x84, 0x4a, 0xa3, 0x44, 0x3f, 0xff, 0x04, 0x3c,
	0x2b, 0xb9, 0x12, 0xd2, 0x47, 0x09, 0xff, 0x71, 0x13, 0x30, 0xbf, 0xbb, 0x9d, 0x75, 0x31, 0xfc,
	0x77, 0x54, 0x31, 0xd8, 0xb6, 0x52, 0xd9, 0x79, 0x62, 0x98, 0xa7, 0xc3, 0x26, 0x16, 0x67, 0x14,
	0xfc, 0xaf, 0x14, 0x11, 0x53, 0x0a, 0xb5, 0x89, 0x43, 0xa5, 0xe1, 0xff, 0xae, 0xb6, 0xad, 0x99,
	0x63, 0x5b, 0x17, 0x74, 0x13, 0xfa, 0xc0, 0x4a, 0x0a, 0xe5, 0x48, 0xe9, 0xd6, 0xa1, 0x0f, 0x55,
	0x7a, 0x89, 0x9c, 0x73, 0x70, 0xc2, 0xb4, 0x57, 0x37, 0x12, 0x55, 0x16, 0x6b, 0xf0, 0x5c, 0x37,
	0x4d, 0x05, 0x3a, 0xf6, 0xeb, 0x97, 0x2e, 0x52, 0x18, 0x73, 0xc5, 0xd1, 0xcd, 0x70, 0xe6, 0x79,
	0xa8, 0x54, 0x9e, 0x12, 0x11, 0x43, 0x40, 0x1f, 0x90, 0x7b, 0xd3, 0xd6, 0x7c, 0x11, 0x1a, 0x1f,
	0x53, 0xbb, 0xf0, 0xb9, 0xd7, 0x31, 0x09, 0x4b, 0x53, 0xdb, 0x8e, 0x75, 0x27, 0x55, 0x8e, 0x7b,
	0xc2, 0x47, 0x08, 0x5d, 0x11, 0x38, 0xcb, 0xd4, 0xf2, 0xdf, 0x73, 0xb2, 0x4f, 0xa2, 0xc5, 0xea,
	0x89, 0x9c, 0x30, 0x39, 0xa6, 0xb0, 0x91, 0xd9, 0xf5, 0x9e, 0xf7, 0x5e, 0xec, 0x26, 0xce, 0xe4,
	0x29, 0xfb, 0x39, 0x57, 0xfa, 0x1d, 0x48, 0x5c, 0x71, 0x4e, 0x52, 0xaa, 0x9d, 0x82, 0x15, 0xfa,
	0xc0, 0x9d, 0xb8, 0x39, 0x21, 0x0d, 0x39, 0x47, 0xdf, 0x61, 0xdc, 0x6e, 0x72, 0xe1, 0x3e, 0x91,
	0xaf, 0xc4, 0x20, 0x16, 0xd5, 0xa2, 0x03, 0xf2, 0xb4, 0xf2, 0x2c, 0xa9, 0xa2, 0x84, 0xd4, 0x2d,
	0x7b, 0x4b, 0x79, 0x09, 0x0d, 0x57, 0x80, 0x35, 0xc4, 0x40, 0x32, 0xae, 0x41, 0xba, 0x1d, 0x36,
	0x32, 0x18, 0x16, 0xc7, 0xa2, 0x65, 0x8b, 0x05, 0x94, 0xe3, 0xe6, 0xcd, 0x62, 0x65, 0xd3, 0xae,
	0x78, 0x46, 0x86, 0x62, 0x00, 0x87, 0x01, 0x1f, 0xf7, 0x7a, 0xe6, 0xda, 0x72, 0xcc, 0xb0, 0x0a,
	0x9a, 0x34, 0xab, 0x46, 0xd8, 0x31, 0x12, 0xe3, 0x62, 0xda, 0x5a, 0x71, 0x9a, 0x2e, 0x8d, 0x79,
	0x59, 0x60, 0xe2, 0x2a, 0xb6, 0x55, 0xca, 0xb9, 0xb5, 0xba, 0xaa, 0x6d, 0x97, 0xda, 0xc9, 0x9a,
	0x7d, 0x4c, 0x85, 0x0a, 0x35, 0x74, 0x46, 0x23, 0xb3, 0xd4, 0x9c, 0xf0, 0x72, 0x6b, 0x61, 0xe5,
	0x39, 0x3c, 0xdf, 0x5a, 0x58, 0x79, 0x01, 0x2f, 0xb6, 0x16, 0x56, 0x6a, 0x50, 0xdb, 0x7e, 0x40,
	0x68, 0x6b, 0x7f, 0x70, 0x98, 0xf4, 0x06, 0x83, 0xfd, 0x57, 0x3d, 0xf5, 0xee, 0x60, 0x68, 0x1f,
	0x5c, 0xab, 0x64, 0xb1, 0x91, 0xa1, 0xb4, 0x4f, 0xae, 0x35, 0xb2, 0x8c, 0x6d, 0xf4, 0x32, 0x8d,
	0x30, 0x5b, 0xdd, 0xfb, 0xe5, 0x43, 0x65, 0xf6, 0xfd, 0x87, 0xca, 0xec, 0x6f, 0x1f, 0x2a, 0xb3,
	0x3f, 0x7f, 0xac, 0xcc, 0xbc, 0xff, 0x58, 0x99, 0xf9, 0xf5, 0x63, 0x65, 0xe6, 0xe5, 0xa3, 0x57,
	0xaf, 0x87, 0xdf, 0xbe, 0x3b, 0xd8, 0xe9, 0xf6, 0x0f, 0x77, 0xdd, 0x63, 0xba, 0xf8, 0xf3, 0x70,
	0xf0, 0xcd, 0x77, 0xbb, 0xd6, 0xe9, 0xd4, 0xeb, 0xfa, 0x60, 0x29, 0x7f, 0x54, 0x3f, 0xfd, 0x23,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0x2d, 0x6e, 0x38, 0x7c, 0x0b, 0x00, 0x00,
}
