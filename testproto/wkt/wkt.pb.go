// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: wkt/wkt.proto

package wkt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageWithWKT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Any         *anypb.Any              `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	Duration    *durationpb.Duration    `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Empty       *emptypb.Empty          `protobuf:"bytes,3,opt,name=empty,proto3" json:"empty,omitempty"`
	FieldMask   *fieldmaskpb.FieldMask  `protobuf:"bytes,4,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	Timestamp   *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DoubleValue *wrapperspb.DoubleValue `protobuf:"bytes,6,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	FloatValue  *wrapperspb.FloatValue  `protobuf:"bytes,7,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	Int64Value  *wrapperspb.Int64Value  `protobuf:"bytes,8,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint64Value *wrapperspb.UInt64Value `protobuf:"bytes,9,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	Int32Value  *wrapperspb.Int32Value  `protobuf:"bytes,10,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Uint32Value *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	BoolValue   *wrapperspb.BoolValue   `protobuf:"bytes,12,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BytesValue  *wrapperspb.BytesValue  `protobuf:"bytes,14,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
}

func (x *MessageWithWKT) Reset() {
	*x = MessageWithWKT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wkt_wkt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithWKT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithWKT) ProtoMessage() {}

func (x *MessageWithWKT) ProtoReflect() protoreflect.Message {
	mi := &file_wkt_wkt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithWKT.ProtoReflect.Descriptor instead.
func (*MessageWithWKT) Descriptor() ([]byte, []int) {
	return file_wkt_wkt_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithWKT) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *MessageWithWKT) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *MessageWithWKT) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *MessageWithWKT) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *MessageWithWKT) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MessageWithWKT) GetDoubleValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *MessageWithWKT) GetFloatValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *MessageWithWKT) GetInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *MessageWithWKT) GetUint64Value() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *MessageWithWKT) GetInt32Value() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *MessageWithWKT) GetUint32Value() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *MessageWithWKT) GetBoolValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *MessageWithWKT) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *MessageWithWKT) GetBytesValue() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

var File_wkt_wkt_proto protoreflect.FileDescriptor

var file_wkt_wkt_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x6b, 0x74, 0x2f, 0x77, 0x6b, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x06, 0x0a, 0x0e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x57, 0x4b, 0x54, 0x12, 0x26, 0x0a,
	0x03, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6b, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wkt_wkt_proto_rawDescOnce sync.Once
	file_wkt_wkt_proto_rawDescData = file_wkt_wkt_proto_rawDesc
)

func file_wkt_wkt_proto_rawDescGZIP() []byte {
	file_wkt_wkt_proto_rawDescOnce.Do(func() {
		file_wkt_wkt_proto_rawDescData = protoimpl.X.CompressGZIP(file_wkt_wkt_proto_rawDescData)
	})
	return file_wkt_wkt_proto_rawDescData
}

var file_wkt_wkt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wkt_wkt_proto_goTypes = []interface{}{
	(*MessageWithWKT)(nil),         // 0: MessageWithWKT
	(*anypb.Any)(nil),              // 1: google.protobuf.Any
	(*durationpb.Duration)(nil),    // 2: google.protobuf.Duration
	(*emptypb.Empty)(nil),          // 3: google.protobuf.Empty
	(*fieldmaskpb.FieldMask)(nil),  // 4: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*wrapperspb.DoubleValue)(nil), // 6: google.protobuf.DoubleValue
	(*wrapperspb.FloatValue)(nil),  // 7: google.protobuf.FloatValue
	(*wrapperspb.Int64Value)(nil),  // 8: google.protobuf.Int64Value
	(*wrapperspb.UInt64Value)(nil), // 9: google.protobuf.UInt64Value
	(*wrapperspb.Int32Value)(nil),  // 10: google.protobuf.Int32Value
	(*wrapperspb.UInt32Value)(nil), // 11: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),   // 12: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil), // 13: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),  // 14: google.protobuf.BytesValue
}
var file_wkt_wkt_proto_depIdxs = []int32{
	1,  // 0: MessageWithWKT.any:type_name -> google.protobuf.Any
	2,  // 1: MessageWithWKT.duration:type_name -> google.protobuf.Duration
	3,  // 2: MessageWithWKT.empty:type_name -> google.protobuf.Empty
	4,  // 3: MessageWithWKT.field_mask:type_name -> google.protobuf.FieldMask
	5,  // 4: MessageWithWKT.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 5: MessageWithWKT.double_value:type_name -> google.protobuf.DoubleValue
	7,  // 6: MessageWithWKT.float_value:type_name -> google.protobuf.FloatValue
	8,  // 7: MessageWithWKT.int64_value:type_name -> google.protobuf.Int64Value
	9,  // 8: MessageWithWKT.uint64_value:type_name -> google.protobuf.UInt64Value
	10, // 9: MessageWithWKT.int32_value:type_name -> google.protobuf.Int32Value
	11, // 10: MessageWithWKT.uint32_value:type_name -> google.protobuf.UInt32Value
	12, // 11: MessageWithWKT.bool_value:type_name -> google.protobuf.BoolValue
	13, // 12: MessageWithWKT.string_value:type_name -> google.protobuf.StringValue
	14, // 13: MessageWithWKT.bytes_value:type_name -> google.protobuf.BytesValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_wkt_wkt_proto_init() }
func file_wkt_wkt_proto_init() {
	if File_wkt_wkt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wkt_wkt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithWKT); i {
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
			RawDescriptor: file_wkt_wkt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wkt_wkt_proto_goTypes,
		DependencyIndexes: file_wkt_wkt_proto_depIdxs,
		MessageInfos:      file_wkt_wkt_proto_msgTypes,
	}.Build()
	File_wkt_wkt_proto = out.File
	file_wkt_wkt_proto_rawDesc = nil
	file_wkt_wkt_proto_goTypes = nil
	file_wkt_wkt_proto_depIdxs = nil
}
