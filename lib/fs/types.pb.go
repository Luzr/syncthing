// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/fs/types.proto

package fs

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

type FilesystemType int32

const (
	FilesystemTypeBasic FilesystemType = 0
	FilesystemTypeFake  FilesystemType = 1
)

var FilesystemType_name = map[int32]string{
	0: "FILESYSTEM_TYPE_BASIC",
	1: "FILESYSTEM_TYPE_FAKE",
}

var FilesystemType_value = map[string]int32{
	"FILESYSTEM_TYPE_BASIC": 0,
	"FILESYSTEM_TYPE_FAKE":  1,
}

func (FilesystemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b556f45c4309ad5d, []int{0}
}

func init() {
	proto.RegisterEnum("fs.FilesystemType", FilesystemType_name, FilesystemType_value)
}

func init() { proto.RegisterFile("lib/fs/types.proto", fileDescriptor_b556f45c4309ad5d) }

var fileDescriptor_b556f45c4309ad5d = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xc9, 0x4c, 0xd2,
	0x4f, 0x2b, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4a, 0x2b, 0xd6, 0x2a, 0xe3, 0xe2, 0x73, 0xcb, 0xcc, 0x49, 0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd,
	0x0d, 0xa9, 0x2c, 0x48, 0x15, 0x32, 0xe2, 0x12, 0x75, 0xf3, 0xf4, 0x71, 0x0d, 0x8e, 0x0c, 0x0e,
	0x71, 0xf5, 0x8d, 0x0f, 0x89, 0x0c, 0x70, 0x8d, 0x77, 0x72, 0x0c, 0xf6, 0x74, 0x16, 0x60, 0x90,
	0x12, 0xef, 0x9a, 0xab, 0x20, 0x8c, 0xaa, 0xdc, 0x29, 0xb1, 0x38, 0x33, 0x59, 0xc8, 0x80, 0x4b,
	0x04, 0x5d, 0x8f, 0x9b, 0xa3, 0xb7, 0xab, 0x00, 0xa3, 0x94, 0x58, 0xd7, 0x5c, 0x05, 0x21, 0x54,
	0x2d, 0x6e, 0x89, 0xd9, 0xa9, 0x4e, 0xce, 0x27, 0x1e, 0xca, 0x31, 0x5c, 0x78, 0x28, 0xc7, 0xf0,
	0xe2, 0x91, 0x1c, 0xc3, 0x84, 0xc7, 0x72, 0x0c, 0x0b, 0x1e, 0xcb, 0x31, 0x5e, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x7e, 0x71, 0x65, 0x5e, 0x72, 0x49, 0x46, 0x66, 0x5e, 0x3a, 0x12, 0x0b, 0xe2, 0x99, 0x24,
	0x36, 0xb0, 0x3f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe8, 0x6b, 0xac, 0xdd, 0x00,
	0x00, 0x00,
}
