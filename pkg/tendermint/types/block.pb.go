// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.12
// source: tendermint/types/block.proto

package types

import (
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *Header       `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data       *Data         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Evidence   *EvidenceList `protobuf:"bytes,3,opt,name=evidence,proto3" json:"evidence,omitempty"`
	LastCommit *Commit       `protobuf:"bytes,4,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_tendermint_types_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Block) GetEvidence() *EvidenceList {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *Block) GetLastCommit() *Commit {
	if x != nil {
		return x.LastCommit
	}
	return nil
}

var File_tendermint_types_block_proto protoreflect.FileDescriptor

var file_tendermint_types_block_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x08, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_block_proto_rawDescOnce sync.Once
	file_tendermint_types_block_proto_rawDescData = file_tendermint_types_block_proto_rawDesc
)

func file_tendermint_types_block_proto_rawDescGZIP() []byte {
	file_tendermint_types_block_proto_rawDescOnce.Do(func() {
		file_tendermint_types_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_block_proto_rawDescData)
	})
	return file_tendermint_types_block_proto_rawDescData
}

var file_tendermint_types_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tendermint_types_block_proto_goTypes = []interface{}{
	(*Block)(nil),        // 0: tendermint.types.Block
	(*Header)(nil),       // 1: tendermint.types.Header
	(*Data)(nil),         // 2: tendermint.types.Data
	(*EvidenceList)(nil), // 3: tendermint.types.EvidenceList
	(*Commit)(nil),       // 4: tendermint.types.Commit
}
var file_tendermint_types_block_proto_depIdxs = []int32{
	1, // 0: tendermint.types.Block.header:type_name -> tendermint.types.Header
	2, // 1: tendermint.types.Block.data:type_name -> tendermint.types.Data
	3, // 2: tendermint.types.Block.evidence:type_name -> tendermint.types.EvidenceList
	4, // 3: tendermint.types.Block.last_commit:type_name -> tendermint.types.Commit
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tendermint_types_block_proto_init() }
func file_tendermint_types_block_proto_init() {
	if File_tendermint_types_block_proto != nil {
		return
	}
	file_tendermint_types_types_proto_init()
	file_tendermint_types_evidence_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_types_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_block_proto_goTypes,
		DependencyIndexes: file_tendermint_types_block_proto_depIdxs,
		MessageInfos:      file_tendermint_types_block_proto_msgTypes,
	}.Build()
	File_tendermint_types_block_proto = out.File
	file_tendermint_types_block_proto_rawDesc = nil
	file_tendermint_types_block_proto_goTypes = nil
	file_tendermint_types_block_proto_depIdxs = nil
}