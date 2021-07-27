// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/qrcodegenerator.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QRCodeFormat struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The entity fields required to generate the QR code.
	FieldMask            *types.FieldMask `protobuf:"bytes,3,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *QRCodeFormat) Reset()      { *m = QRCodeFormat{} }
func (*QRCodeFormat) ProtoMessage() {}
func (*QRCodeFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_f400aed11530ba72, []int{0}
}
func (m *QRCodeFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QRCodeFormat.Unmarshal(m, b)
}
func (m *QRCodeFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QRCodeFormat.Marshal(b, m, deterministic)
}
func (m *QRCodeFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QRCodeFormat.Merge(m, src)
}
func (m *QRCodeFormat) XXX_Size() int {
	return xxx_messageInfo_QRCodeFormat.Size(m)
}
func (m *QRCodeFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_QRCodeFormat.DiscardUnknown(m)
}

var xxx_messageInfo_QRCodeFormat proto.InternalMessageInfo

func (m *QRCodeFormat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QRCodeFormat) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *QRCodeFormat) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

type QRCodeFormats struct {
	// Available formats. The map key is the format identifier.
	Formats              map[string]*QRCodeFormat `protobuf:"bytes,1,rep,name=formats,proto3" json:"formats,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *QRCodeFormats) Reset()      { *m = QRCodeFormats{} }
func (*QRCodeFormats) ProtoMessage() {}
func (*QRCodeFormats) Descriptor() ([]byte, []int) {
	return fileDescriptor_f400aed11530ba72, []int{1}
}
func (m *QRCodeFormats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QRCodeFormats.Unmarshal(m, b)
}
func (m *QRCodeFormats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QRCodeFormats.Marshal(b, m, deterministic)
}
func (m *QRCodeFormats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QRCodeFormats.Merge(m, src)
}
func (m *QRCodeFormats) XXX_Size() int {
	return xxx_messageInfo_QRCodeFormats.Size(m)
}
func (m *QRCodeFormats) XXX_DiscardUnknown() {
	xxx_messageInfo_QRCodeFormats.DiscardUnknown(m)
}

var xxx_messageInfo_QRCodeFormats proto.InternalMessageInfo

func (m *QRCodeFormats) GetFormats() map[string]*QRCodeFormat {
	if m != nil {
		return m.Formats
	}
	return nil
}

type GetQRCodeFormatRequest struct {
	// QR code format identifier. Enumerate available formats with rpc ListFormats in the EndDeviceQRCodeGenerator service.
	FormatID             string   `protobuf:"bytes,1,opt,name=format_id,json=formatId,proto3" json:"format_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQRCodeFormatRequest) Reset()      { *m = GetQRCodeFormatRequest{} }
func (*GetQRCodeFormatRequest) ProtoMessage() {}
func (*GetQRCodeFormatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f400aed11530ba72, []int{2}
}
func (m *GetQRCodeFormatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQRCodeFormatRequest.Unmarshal(m, b)
}
func (m *GetQRCodeFormatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQRCodeFormatRequest.Marshal(b, m, deterministic)
}
func (m *GetQRCodeFormatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQRCodeFormatRequest.Merge(m, src)
}
func (m *GetQRCodeFormatRequest) XXX_Size() int {
	return xxx_messageInfo_GetQRCodeFormatRequest.Size(m)
}
func (m *GetQRCodeFormatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQRCodeFormatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQRCodeFormatRequest proto.InternalMessageInfo

func (m *GetQRCodeFormatRequest) GetFormatID() string {
	if m != nil {
		return m.FormatID
	}
	return ""
}

type GenerateEndDeviceQRCodeRequest struct {
	// QR code format identifier. Enumerate available formats with rpc ListFormats in the EndDeviceQRCodeGenerator service.
	FormatID string `protobuf:"bytes,1,opt,name=format_id,json=formatId,proto3" json:"format_id,omitempty"`
	// End device to use as input to generate the QR code.
	EndDevice EndDevice `protobuf:"bytes,2,opt,name=end_device,json=endDevice,proto3" json:"end_device"`
	// If set, the server will render the QR code image according to these settings.
	Image                *GenerateEndDeviceQRCodeRequest_Image `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *GenerateEndDeviceQRCodeRequest) Reset()      { *m = GenerateEndDeviceQRCodeRequest{} }
func (*GenerateEndDeviceQRCodeRequest) ProtoMessage() {}
func (*GenerateEndDeviceQRCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f400aed11530ba72, []int{3}
}
func (m *GenerateEndDeviceQRCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateEndDeviceQRCodeRequest.Unmarshal(m, b)
}
func (m *GenerateEndDeviceQRCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateEndDeviceQRCodeRequest.Marshal(b, m, deterministic)
}
func (m *GenerateEndDeviceQRCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateEndDeviceQRCodeRequest.Merge(m, src)
}
func (m *GenerateEndDeviceQRCodeRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateEndDeviceQRCodeRequest.Size(m)
}
func (m *GenerateEndDeviceQRCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateEndDeviceQRCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateEndDeviceQRCodeRequest proto.InternalMessageInfo

func (m *GenerateEndDeviceQRCodeRequest) GetFormatID() string {
	if m != nil {
		return m.FormatID
	}
	return ""
}

func (m *GenerateEndDeviceQRCodeRequest) GetEndDevice() EndDevice {
	if m != nil {
		return m.EndDevice
	}
	return EndDevice{}
}

func (m *GenerateEndDeviceQRCodeRequest) GetImage() *GenerateEndDeviceQRCodeRequest_Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type GenerateEndDeviceQRCodeRequest_Image struct {
	// Requested QR code image dimension in pixels.
	ImageSize            uint32   `protobuf:"varint,1,opt,name=image_size,json=imageSize,proto3" json:"image_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateEndDeviceQRCodeRequest_Image) Reset()      { *m = GenerateEndDeviceQRCodeRequest_Image{} }
func (*GenerateEndDeviceQRCodeRequest_Image) ProtoMessage() {}
func (*GenerateEndDeviceQRCodeRequest_Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_f400aed11530ba72, []int{3, 0}
}
func (m *GenerateEndDeviceQRCodeRequest_Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateEndDeviceQRCodeRequest_Image.Unmarshal(m, b)
}
func (m *GenerateEndDeviceQRCodeRequest_Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateEndDeviceQRCodeRequest_Image.Marshal(b, m, deterministic)
}
func (m *GenerateEndDeviceQRCodeRequest_Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateEndDeviceQRCodeRequest_Image.Merge(m, src)
}
func (m *GenerateEndDeviceQRCodeRequest_Image) XXX_Size() int {
	return xxx_messageInfo_GenerateEndDeviceQRCodeRequest_Image.Size(m)
}
func (m *GenerateEndDeviceQRCodeRequest_Image) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateEndDeviceQRCodeRequest_Image.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateEndDeviceQRCodeRequest_Image proto.InternalMessageInfo

func (m *GenerateEndDeviceQRCodeRequest_Image) GetImageSize() uint32 {
	if m != nil {
		return m.ImageSize
	}
	return 0
}

type GenerateQRCodeResponse struct {
	// Text representation of the QR code contents.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// QR code in PNG format, if requested.
	Image                *Picture `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateQRCodeResponse) Reset()      { *m = GenerateQRCodeResponse{} }
func (*GenerateQRCodeResponse) ProtoMessage() {}
func (*GenerateQRCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f400aed11530ba72, []int{4}
}
func (m *GenerateQRCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateQRCodeResponse.Unmarshal(m, b)
}
func (m *GenerateQRCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateQRCodeResponse.Marshal(b, m, deterministic)
}
func (m *GenerateQRCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateQRCodeResponse.Merge(m, src)
}
func (m *GenerateQRCodeResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateQRCodeResponse.Size(m)
}
func (m *GenerateQRCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateQRCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateQRCodeResponse proto.InternalMessageInfo

func (m *GenerateQRCodeResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *GenerateQRCodeResponse) GetImage() *Picture {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*QRCodeFormat)(nil), "ttn.lorawan.v3.QRCodeFormat")
	golang_proto.RegisterType((*QRCodeFormat)(nil), "ttn.lorawan.v3.QRCodeFormat")
	proto.RegisterType((*QRCodeFormats)(nil), "ttn.lorawan.v3.QRCodeFormats")
	golang_proto.RegisterType((*QRCodeFormats)(nil), "ttn.lorawan.v3.QRCodeFormats")
	proto.RegisterMapType((map[string]*QRCodeFormat)(nil), "ttn.lorawan.v3.QRCodeFormats.FormatsEntry")
	golang_proto.RegisterMapType((map[string]*QRCodeFormat)(nil), "ttn.lorawan.v3.QRCodeFormats.FormatsEntry")
	proto.RegisterType((*GetQRCodeFormatRequest)(nil), "ttn.lorawan.v3.GetQRCodeFormatRequest")
	golang_proto.RegisterType((*GetQRCodeFormatRequest)(nil), "ttn.lorawan.v3.GetQRCodeFormatRequest")
	proto.RegisterType((*GenerateEndDeviceQRCodeRequest)(nil), "ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest")
	golang_proto.RegisterType((*GenerateEndDeviceQRCodeRequest)(nil), "ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest")
	proto.RegisterType((*GenerateEndDeviceQRCodeRequest_Image)(nil), "ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest.Image")
	golang_proto.RegisterType((*GenerateEndDeviceQRCodeRequest_Image)(nil), "ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest.Image")
	proto.RegisterType((*GenerateQRCodeResponse)(nil), "ttn.lorawan.v3.GenerateQRCodeResponse")
	golang_proto.RegisterType((*GenerateQRCodeResponse)(nil), "ttn.lorawan.v3.GenerateQRCodeResponse")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/qrcodegenerator.proto", fileDescriptor_f400aed11530ba72)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/qrcodegenerator.proto", fileDescriptor_f400aed11530ba72)
}

var fileDescriptor_f400aed11530ba72 = []byte{
	// 822 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xf6, 0xd8, 0x31, 0x89, 0xc7, 0x29, 0xaa, 0x46, 0xa2, 0x2c, 0xdb, 0x76, 0x63, 0xad, 0x42,
	0x70, 0x42, 0x77, 0x17, 0xd6, 0x1c, 0x68, 0x2e, 0x15, 0x4b, 0x13, 0x2b, 0x08, 0xa4, 0xb2, 0x08,
	0x09, 0x11, 0x15, 0x6b, 0xec, 0x1d, 0x6f, 0x16, 0xdb, 0x33, 0x9b, 0xd9, 0xb1, 0x5b, 0xa7, 0xaa,
	0x84, 0x10, 0x27, 0x4e, 0x08, 0xc4, 0x85, 0x13, 0x47, 0xfe, 0x06, 0x4e, 0x1c, 0x7b, 0xe7, 0xc2,
	0x85, 0x0a, 0x1c, 0x0e, 0x3d, 0x72, 0xf6, 0x09, 0xed, 0xcc, 0x6e, 0xea, 0x1f, 0xc5, 0x82, 0x43,
	0x4f, 0x7e, 0xe3, 0xf7, 0xed, 0x7b, 0xdf, 0x7b, 0xef, 0x7b, 0x0f, 0xbe, 0xd6, 0x67, 0x1c, 0xdf,
	0xc3, 0xd4, 0x4a, 0x04, 0xee, 0xf4, 0x1c, 0x1c, 0x47, 0xce, 0x29, 0xef, 0xb0, 0x80, 0x84, 0x84,
	0x12, 0x8e, 0x05, 0xe3, 0x76, 0xcc, 0x99, 0x60, 0xe8, 0x45, 0x21, 0xa8, 0x9d, 0x81, 0xed, 0x51,
	0x43, 0x7f, 0x27, 0x8c, 0xc4, 0xc9, 0xb0, 0x6d, 0x77, 0xd8, 0xc0, 0x21, 0x74, 0xc4, 0xc6, 0x31,
	0x67, 0xf7, 0xc7, 0x8e, 0x04, 0x77, 0xac, 0x90, 0x50, 0x6b, 0x84, 0xfb, 0x51, 0x80, 0x05, 0x71,
	0x96, 0x0c, 0x15, 0x52, 0xb7, 0x66, 0x42, 0x84, 0x2c, 0x64, 0xea, 0xe3, 0xf6, 0xb0, 0x2b, 0x5f,
	0xf2, 0x21, 0xad, 0x0c, 0x7e, 0x2d, 0x64, 0x2c, 0xec, 0x13, 0xc9, 0x11, 0x53, 0xca, 0x04, 0x16,
	0x11, 0xa3, 0x49, 0xe6, 0xbd, 0x9a, 0x79, 0x2f, 0x62, 0x90, 0x41, 0x2c, 0xc6, 0x99, 0xb3, 0xb6,
	0xe8, 0xec, 0x46, 0xa4, 0x1f, 0xb4, 0x06, 0x38, 0xe9, 0x65, 0x08, 0x73, 0xb9, 0x0f, 0x84, 0x06,
	0xad, 0x80, 0x8c, 0xa2, 0x4e, 0xce, 0x77, 0x6b, 0x19, 0x13, 0x47, 0x1d, 0x31, 0xe4, 0x19, 0xc0,
	0xfc, 0x16, 0xc0, 0xcd, 0x0f, 0xfd, 0x77, 0x59, 0x40, 0x0e, 0x19, 0x1f, 0x60, 0x81, 0xae, 0xc2,
	0x35, 0x8a, 0x07, 0x44, 0x03, 0x35, 0x50, 0xaf, 0x78, 0xeb, 0x53, 0x6f, 0x8d, 0x17, 0xb5, 0xc0,
	0x97, 0x7f, 0xa2, 0x3d, 0x58, 0x0d, 0x48, 0xd2, 0xe1, 0x51, 0x9c, 0xd6, 0xa1, 0x15, 0x25, 0x66,
	0x63, 0xea, 0x95, 0x79, 0x49, 0x7b, 0x04, 0xfc, 0x59, 0x27, 0xba, 0x09, 0xe1, 0x53, 0xca, 0x5a,
	0xa9, 0x06, 0xea, 0x55, 0x57, 0xb7, 0x55, 0x55, 0x76, 0x5e, 0x95, 0x7d, 0x98, 0x42, 0x3e, 0xc0,
	0x49, 0xcf, 0xaf, 0x74, 0x73, 0xd3, 0xfc, 0x1d, 0xc0, 0x4b, 0xb3, 0xa4, 0x12, 0xc4, 0xe1, 0x7a,
	0x57, 0x99, 0x1a, 0xa8, 0x95, 0xea, 0x55, 0x77, 0xcf, 0x9e, 0x1f, 0xae, 0x3d, 0x87, 0xb7, 0xb3,
	0xdf, 0x03, 0x2a, 0xf8, 0xd8, 0xbb, 0x31, 0xf5, 0x76, 0x7f, 0x00, 0x3b, 0xe6, 0x36, 0x37, 0xb5,
	0x6d, 0xd7, 0xf8, 0xec, 0x18, 0x5b, 0x67, 0x6f, 0x58, 0x37, 0xef, 0xd6, 0x6f, 0xed, 0x1f, 0x5b,
	0x77, 0x6f, 0xe5, 0xcf, 0xdd, 0x07, 0xee, 0x8d, 0x87, 0xdb, 0x7e, 0x9e, 0x48, 0xff, 0x04, 0x6e,
	0xce, 0x86, 0x41, 0x97, 0x61, 0xa9, 0x47, 0xc6, 0xaa, 0x31, 0x7e, 0x6a, 0x22, 0x17, 0x96, 0x47,
	0xb8, 0x3f, 0x24, 0xb2, 0x11, 0x55, 0xf7, 0xda, 0x2a, 0x4e, 0xbe, 0x82, 0xee, 0x17, 0xdf, 0x06,
	0xe6, 0xe7, 0xf0, 0x4a, 0x93, 0x88, 0x39, 0x2f, 0x39, 0x1d, 0x92, 0x44, 0xa0, 0x3b, 0xb0, 0xa2,
	0xd2, 0xb7, 0xa2, 0x20, 0x1b, 0x41, 0x63, 0xea, 0xfd, 0x27, 0xe2, 0x93, 0xc7, 0x5b, 0x1b, 0x2a,
	0xd8, 0xd1, 0x6d, 0x7f, 0x43, 0x45, 0x39, 0x0a, 0xcc, 0x9f, 0x8b, 0xd0, 0x68, 0xaa, 0xc5, 0x20,
	0x07, 0x34, 0xb8, 0x2d, 0xd5, 0xa1, 0x52, 0x3f, 0xb7, 0xa4, 0xe8, 0x10, 0xc2, 0xa7, 0x52, 0xcc,
	0xba, 0xf3, 0xca, 0x62, 0x77, 0x2e, 0xd8, 0x78, 0x9b, 0x53, 0xaf, 0xfc, 0x35, 0x28, 0x5e, 0x06,
	0x8f, 0x1e, 0x6f, 0x15, 0xfc, 0x0a, 0xc9, 0x1d, 0xe8, 0x3d, 0x58, 0x8e, 0x06, 0x38, 0x24, 0x99,
	0x7c, 0xde, 0x5a, 0x0c, 0xb1, 0xba, 0x30, 0xfb, 0x28, 0xfd, 0xd6, 0x57, 0x21, 0x74, 0x17, 0x96,
	0xe5, 0x1b, 0xed, 0x42, 0x28, 0xff, 0x69, 0x25, 0xd1, 0x99, 0xd2, 0xf9, 0x25, 0x0f, 0x4e, 0xbd,
	0xf5, 0xbd, 0xb2, 0xf6, 0x64, 0xbd, 0x0e, 0xfd, 0x8a, 0xf4, 0x7e, 0x14, 0x9d, 0x11, 0xf3, 0x38,
	0x1d, 0x94, 0x4a, 0x91, 0x47, 0x4e, 0x62, 0x46, 0x13, 0x82, 0x10, 0x5c, 0x13, 0xe4, 0xbe, 0xc8,
	0xd4, 0x20, 0x6d, 0x64, 0xe5, 0x6c, 0x55, 0xc1, 0x2f, 0x2f, 0xb2, 0xbd, 0xa3, 0x36, 0x2f, 0x23,
	0xe4, 0x7e, 0x5f, 0x82, 0xda, 0x02, 0xf1, 0x66, 0x7e, 0xc1, 0xd0, 0x57, 0x00, 0x56, 0x9a, 0x44,
	0x64, 0x4b, 0xb9, 0xb3, 0x5c, 0xf8, 0xb3, 0xe4, 0xa3, 0xaf, 0x54, 0xa0, 0xf9, 0xe6, 0x97, 0xbf,
	0xfe, 0xf5, 0x5d, 0xf1, 0x75, 0xb4, 0xeb, 0x9c, 0x72, 0x2b, 0x3d, 0x98, 0x49, 0x7a, 0x30, 0x2c,
	0x35, 0xa5, 0xc4, 0xc9, 0x74, 0xef, 0x3c, 0xb8, 0x10, 0xc3, 0x43, 0xd4, 0x83, 0xd5, 0xf7, 0xa3,
	0x44, 0xe4, 0x6b, 0x78, 0x65, 0x69, 0x7f, 0x0f, 0xd2, 0x93, 0xa5, 0x5f, 0x5f, 0xb9, 0x8d, 0xe6,
	0xab, 0x32, 0xf1, 0x16, 0xba, 0xbe, 0x32, 0x71, 0x5a, 0xf3, 0x46, 0xde, 0x6e, 0x64, 0xff, 0xbf,
	0x59, 0xeb, 0x3b, 0xff, 0x86, 0x9f, 0x1f, 0x9c, 0x59, 0x93, 0x5c, 0x74, 0xf3, 0xa5, 0x67, 0x72,
	0xd9, 0x07, 0x7b, 0xde, 0xc7, 0xbf, 0xfd, 0x69, 0x14, 0xbe, 0x98, 0x18, 0xe0, 0xa7, 0x89, 0x01,
	0xfe, 0x98, 0x18, 0xe0, 0xc9, 0xc4, 0x28, 0xfc, 0x3d, 0x31, 0xc0, 0x37, 0xe7, 0x46, 0xe1, 0xc7,
	0x73, 0xa3, 0xf0, 0xcb, 0xb9, 0x01, 0x3e, 0x75, 0x42, 0x66, 0x8b, 0x13, 0x22, 0x4e, 0x22, 0x1a,
	0x26, 0x36, 0x25, 0xe2, 0x1e, 0xe3, 0x3d, 0x67, 0xfe, 0xdc, 0x8e, 0x1a, 0x4e, 0xdc, 0x0b, 0x1d,
	0x21, 0x68, 0xdc, 0x6e, 0xbf, 0x20, 0x7b, 0xd6, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x4a,
	0x70, 0x9d, 0xbf, 0x06, 0x00, 0x00,
}

func (this *QRCodeFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QRCodeFormat)
	if !ok {
		that2, ok := that.(QRCodeFormat)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.FieldMask.Equal(that1.FieldMask) {
		return false
	}
	return true
}
func (this *QRCodeFormats) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QRCodeFormats)
	if !ok {
		that2, ok := that.(QRCodeFormats)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Formats) != len(that1.Formats) {
		return false
	}
	for i := range this.Formats {
		if !this.Formats[i].Equal(that1.Formats[i]) {
			return false
		}
	}
	return true
}
func (this *GetQRCodeFormatRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetQRCodeFormatRequest)
	if !ok {
		that2, ok := that.(GetQRCodeFormatRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FormatID != that1.FormatID {
		return false
	}
	return true
}
func (this *GenerateEndDeviceQRCodeRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenerateEndDeviceQRCodeRequest)
	if !ok {
		that2, ok := that.(GenerateEndDeviceQRCodeRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FormatID != that1.FormatID {
		return false
	}
	if !this.EndDevice.Equal(&that1.EndDevice) {
		return false
	}
	if !this.Image.Equal(that1.Image) {
		return false
	}
	return true
}
func (this *GenerateEndDeviceQRCodeRequest_Image) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenerateEndDeviceQRCodeRequest_Image)
	if !ok {
		that2, ok := that.(GenerateEndDeviceQRCodeRequest_Image)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ImageSize != that1.ImageSize {
		return false
	}
	return true
}
func (this *GenerateQRCodeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenerateQRCodeResponse)
	if !ok {
		that2, ok := that.(GenerateQRCodeResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Text != that1.Text {
		return false
	}
	if !this.Image.Equal(that1.Image) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndDeviceQRCodeGeneratorClient is the client API for EndDeviceQRCodeGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndDeviceQRCodeGeneratorClient interface {
	// Return the QR code format.
	GetFormat(ctx context.Context, in *GetQRCodeFormatRequest, opts ...grpc.CallOption) (*QRCodeFormat, error)
	// Returns the supported formats.
	ListFormats(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*QRCodeFormats, error)
	// Generates a QR code.
	Generate(ctx context.Context, in *GenerateEndDeviceQRCodeRequest, opts ...grpc.CallOption) (*GenerateQRCodeResponse, error)
}

type endDeviceQRCodeGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewEndDeviceQRCodeGeneratorClient(cc *grpc.ClientConn) EndDeviceQRCodeGeneratorClient {
	return &endDeviceQRCodeGeneratorClient{cc}
}

func (c *endDeviceQRCodeGeneratorClient) GetFormat(ctx context.Context, in *GetQRCodeFormatRequest, opts ...grpc.CallOption) (*QRCodeFormat, error) {
	out := new(QRCodeFormat)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/GetFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceQRCodeGeneratorClient) ListFormats(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*QRCodeFormats, error) {
	out := new(QRCodeFormats)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/ListFormats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceQRCodeGeneratorClient) Generate(ctx context.Context, in *GenerateEndDeviceQRCodeRequest, opts ...grpc.CallOption) (*GenerateQRCodeResponse, error) {
	out := new(GenerateQRCodeResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceQRCodeGeneratorServer is the server API for EndDeviceQRCodeGenerator service.
type EndDeviceQRCodeGeneratorServer interface {
	// Return the QR code format.
	GetFormat(context.Context, *GetQRCodeFormatRequest) (*QRCodeFormat, error)
	// Returns the supported formats.
	ListFormats(context.Context, *types.Empty) (*QRCodeFormats, error)
	// Generates a QR code.
	Generate(context.Context, *GenerateEndDeviceQRCodeRequest) (*GenerateQRCodeResponse, error)
}

// UnimplementedEndDeviceQRCodeGeneratorServer can be embedded to have forward compatible implementations.
type UnimplementedEndDeviceQRCodeGeneratorServer struct {
}

func (*UnimplementedEndDeviceQRCodeGeneratorServer) GetFormat(ctx context.Context, req *GetQRCodeFormatRequest) (*QRCodeFormat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFormat not implemented")
}
func (*UnimplementedEndDeviceQRCodeGeneratorServer) ListFormats(ctx context.Context, req *types.Empty) (*QRCodeFormats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFormats not implemented")
}
func (*UnimplementedEndDeviceQRCodeGeneratorServer) Generate(ctx context.Context, req *GenerateEndDeviceQRCodeRequest) (*GenerateQRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterEndDeviceQRCodeGeneratorServer(s *grpc.Server, srv EndDeviceQRCodeGeneratorServer) {
	s.RegisterService(&_EndDeviceQRCodeGenerator_serviceDesc, srv)
}

func _EndDeviceQRCodeGenerator_GetFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQRCodeFormatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).GetFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/GetFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).GetFormat(ctx, req.(*GetQRCodeFormatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceQRCodeGenerator_ListFormats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).ListFormats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/ListFormats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).ListFormats(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceQRCodeGenerator_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateEndDeviceQRCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).Generate(ctx, req.(*GenerateEndDeviceQRCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndDeviceQRCodeGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceQRCodeGenerator",
	HandlerType: (*EndDeviceQRCodeGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFormat",
			Handler:    _EndDeviceQRCodeGenerator_GetFormat_Handler,
		},
		{
			MethodName: "ListFormats",
			Handler:    _EndDeviceQRCodeGenerator_ListFormats_Handler,
		},
		{
			MethodName: "Generate",
			Handler:    _EndDeviceQRCodeGenerator_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/qrcodegenerator.proto",
}

func (m *QRCodeFormat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	if m.FieldMask != nil {
		l = m.FieldMask.Size()
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	return n
}

func (m *QRCodeFormats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Formats) > 0 {
		for k, v := range m.Formats {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovQrcodegenerator(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovQrcodegenerator(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovQrcodegenerator(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *GetQRCodeFormatRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FormatID)
	if l > 0 {
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	return n
}

func (m *GenerateEndDeviceQRCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FormatID)
	if l > 0 {
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	l = m.EndDevice.Size()
	n += 1 + l + sovQrcodegenerator(uint64(l))
	if m.Image != nil {
		l = m.Image.Size()
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	return n
}

func (m *GenerateEndDeviceQRCodeRequest_Image) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ImageSize != 0 {
		n += 1 + sovQrcodegenerator(uint64(m.ImageSize))
	}
	return n
}

func (m *GenerateQRCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	if m.Image != nil {
		l = m.Image.Size()
		n += 1 + l + sovQrcodegenerator(uint64(l))
	}
	return n
}

func sovQrcodegenerator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQrcodegenerator(x uint64) (n int) {
	return sovQrcodegenerator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QRCodeFormat) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QRCodeFormat{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`FieldMask:` + strings.Replace(fmt.Sprintf("%v", this.FieldMask), "FieldMask", "types.FieldMask", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QRCodeFormats) String() string {
	if this == nil {
		return "nil"
	}
	keysForFormats := make([]string, 0, len(this.Formats))
	for k := range this.Formats {
		keysForFormats = append(keysForFormats, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForFormats)
	mapStringForFormats := "map[string]*QRCodeFormat{"
	for _, k := range keysForFormats {
		mapStringForFormats += fmt.Sprintf("%v: %v,", k, this.Formats[k])
	}
	mapStringForFormats += "}"
	s := strings.Join([]string{`&QRCodeFormats{`,
		`Formats:` + mapStringForFormats + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetQRCodeFormatRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetQRCodeFormatRequest{`,
		`FormatID:` + fmt.Sprintf("%v", this.FormatID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GenerateEndDeviceQRCodeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GenerateEndDeviceQRCodeRequest{`,
		`FormatID:` + fmt.Sprintf("%v", this.FormatID) + `,`,
		`EndDevice:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EndDevice), "EndDevice", "EndDevice", 1), `&`, ``, 1) + `,`,
		`Image:` + strings.Replace(fmt.Sprintf("%v", this.Image), "GenerateEndDeviceQRCodeRequest_Image", "GenerateEndDeviceQRCodeRequest_Image", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GenerateEndDeviceQRCodeRequest_Image) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GenerateEndDeviceQRCodeRequest_Image{`,
		`ImageSize:` + fmt.Sprintf("%v", this.ImageSize) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GenerateQRCodeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GenerateQRCodeResponse{`,
		`Text:` + fmt.Sprintf("%v", this.Text) + `,`,
		`Image:` + strings.Replace(fmt.Sprintf("%v", this.Image), "Picture", "Picture", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQrcodegenerator(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
