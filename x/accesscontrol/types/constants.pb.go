// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/constants.proto

package types

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

type ResourceType int32

const (
	ResourceType_ANY    ResourceType = 0
	ResourceType_KV     ResourceType = 1
	ResourceType_Mem    ResourceType = 2
	ResourceType_DexMem ResourceType = 3
)

var ResourceType_name = map[int32]string{
	0: "ANY",
	1: "KV",
	2: "Mem",
	3: "DexMem",
}

var ResourceType_value = map[string]int32{
	"ANY":    0,
	"KV":     1,
	"Mem":    2,
	"DexMem": 3,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0x4e, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf,
	0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x2e, 0x49, 0xcc, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x81, 0xa8, 0xd2, 0x43, 0x51, 0xa5, 0x57, 0x66, 0x98, 0x94, 0x5a, 0x92, 0x68, 0xa8,
	0x65, 0xc5, 0xc5, 0xe5, 0x08, 0x96, 0x08, 0xa9, 0x2c, 0x48, 0x15, 0xe2, 0xe6, 0x62, 0x0f, 0xf5,
	0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0xe2, 0xe0, 0x62, 0x09, 0x72, 0x75, 0x74, 0x11,
	0x60, 0x14, 0xe2, 0xe4, 0x62, 0x0d, 0x0f, 0xf2, 0x0c, 0x71, 0x15, 0x60, 0x12, 0xe2, 0xe2, 0x62,
	0x73, 0xf6, 0xf7, 0xf5, 0xf5, 0x0c, 0x11, 0x60, 0xd6, 0x32, 0xe1, 0xe2, 0x09, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x05, 0xeb, 0x66, 0xe7, 0x62, 0x76, 0xf4, 0x8b, 0x14, 0x60, 0x10, 0x62,
	0xe3, 0x62, 0xf2, 0x0e, 0x13, 0x60, 0x04, 0x09, 0xf8, 0xa6, 0xe6, 0x42, 0x74, 0xb9, 0xa4, 0x56,
	0x80, 0xd8, 0xcc, 0x4e, 0x3e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65,
	0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xf5, 0x1a, 0x84, 0xd2,
	0x2d, 0x4e, 0xc9, 0xd6, 0xaf, 0x40, 0xf3, 0x67, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8,
	0x93, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x1e, 0xf7, 0xf6, 0x0c, 0x01, 0x00, 0x00,
}
