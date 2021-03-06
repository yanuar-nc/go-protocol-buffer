// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Article_StatusType int32

const (
	Article_DRAFT       Article_StatusType = 0
	Article_PUBLISHED   Article_StatusType = 1
	Article_UNPUBLISHED Article_StatusType = 2
)

var Article_StatusType_name = map[int32]string{
	0: "DRAFT",
	1: "PUBLISHED",
	2: "UNPUBLISHED",
}

var Article_StatusType_value = map[string]int32{
	"DRAFT":       0,
	"PUBLISHED":   1,
	"UNPUBLISHED": 2,
}

func (x Article_StatusType) String() string {
	return proto.EnumName(Article_StatusType_name, int32(x))
}

func (Article_StatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0, 0}
}

type Article struct {
	Id                   int32                                         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string                                        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string                                        `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status               Article_StatusType                            `protobuf:"varint,4,opt,name=status,proto3,enum=pb.Article_StatusType" json:"status,omitempty"`
	Tags                 []*Article_Tags                               `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	SocialMediaStatistic map[string]*Article_SocialMediaStatisticField `protobuf:"bytes,6,rep,name=SocialMediaStatistic,proto3" json:"SocialMediaStatistic,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to ProcessOneof:
	//	*Article_Insert
	//	*Article_Update
	//	*Article_Delete
	//	*Article_Archive
	ProcessOneof         isArticle_ProcessOneof `protobuf_oneof:"process_oneof"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetStatus() Article_StatusType {
	if m != nil {
		return m.Status
	}
	return Article_DRAFT
}

func (m *Article) GetTags() []*Article_Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Article) GetSocialMediaStatistic() map[string]*Article_SocialMediaStatisticField {
	if m != nil {
		return m.SocialMediaStatistic
	}
	return nil
}

type isArticle_ProcessOneof interface {
	isArticle_ProcessOneof()
}

type Article_Insert struct {
	Insert bool `protobuf:"varint,7,opt,name=insert,proto3,oneof"`
}

type Article_Update struct {
	Update bool `protobuf:"varint,8,opt,name=update,proto3,oneof"`
}

type Article_Delete struct {
	Delete bool `protobuf:"varint,9,opt,name=delete,proto3,oneof"`
}

type Article_Archive struct {
	Archive bool `protobuf:"varint,10,opt,name=archive,proto3,oneof"`
}

func (*Article_Insert) isArticle_ProcessOneof() {}

func (*Article_Update) isArticle_ProcessOneof() {}

func (*Article_Delete) isArticle_ProcessOneof() {}

func (*Article_Archive) isArticle_ProcessOneof() {}

func (m *Article) GetProcessOneof() isArticle_ProcessOneof {
	if m != nil {
		return m.ProcessOneof
	}
	return nil
}

func (m *Article) GetInsert() bool {
	if x, ok := m.GetProcessOneof().(*Article_Insert); ok {
		return x.Insert
	}
	return false
}

func (m *Article) GetUpdate() bool {
	if x, ok := m.GetProcessOneof().(*Article_Update); ok {
		return x.Update
	}
	return false
}

func (m *Article) GetDelete() bool {
	if x, ok := m.GetProcessOneof().(*Article_Delete); ok {
		return x.Delete
	}
	return false
}

func (m *Article) GetArchive() bool {
	if x, ok := m.GetProcessOneof().(*Article_Archive); ok {
		return x.Archive
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Article) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Article_OneofMarshaler, _Article_OneofUnmarshaler, _Article_OneofSizer, []interface{}{
		(*Article_Insert)(nil),
		(*Article_Update)(nil),
		(*Article_Delete)(nil),
		(*Article_Archive)(nil),
	}
}

func _Article_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Article)
	// process_oneof
	switch x := m.ProcessOneof.(type) {
	case *Article_Insert:
		t := uint64(0)
		if x.Insert {
			t = 1
		}
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Article_Update:
		t := uint64(0)
		if x.Update {
			t = 1
		}
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Article_Delete:
		t := uint64(0)
		if x.Delete {
			t = 1
		}
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Article_Archive:
		t := uint64(0)
		if x.Archive {
			t = 1
		}
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Article.ProcessOneof has unexpected type %T", x)
	}
	return nil
}

func _Article_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Article)
	switch tag {
	case 7: // process_oneof.insert
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProcessOneof = &Article_Insert{x != 0}
		return true, err
	case 8: // process_oneof.update
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProcessOneof = &Article_Update{x != 0}
		return true, err
	case 9: // process_oneof.delete
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProcessOneof = &Article_Delete{x != 0}
		return true, err
	case 10: // process_oneof.archive
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProcessOneof = &Article_Archive{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Article_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Article)
	// process_oneof
	switch x := m.ProcessOneof.(type) {
	case *Article_Insert:
		n += 1 // tag and wire
		n += 1
	case *Article_Update:
		n += 1 // tag and wire
		n += 1
	case *Article_Delete:
		n += 1 // tag and wire
		n += 1
	case *Article_Archive:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Article_Tags struct {
	TagId                int32    `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article_Tags) Reset()         { *m = Article_Tags{} }
func (m *Article_Tags) String() string { return proto.CompactTextString(m) }
func (*Article_Tags) ProtoMessage()    {}
func (*Article_Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0, 0}
}

func (m *Article_Tags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article_Tags.Unmarshal(m, b)
}
func (m *Article_Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article_Tags.Marshal(b, m, deterministic)
}
func (m *Article_Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article_Tags.Merge(m, src)
}
func (m *Article_Tags) XXX_Size() int {
	return xxx_messageInfo_Article_Tags.Size(m)
}
func (m *Article_Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_Article_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_Article_Tags proto.InternalMessageInfo

func (m *Article_Tags) GetTagId() int32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *Article_Tags) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Article_SocialMediaStatisticField struct {
	Like                 int32    `protobuf:"varint,1,opt,name=Like,proto3" json:"Like,omitempty"`
	Share                int32    `protobuf:"varint,2,opt,name=Share,proto3" json:"Share,omitempty"`
	Comments             int32    `protobuf:"varint,3,opt,name=Comments,proto3" json:"Comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article_SocialMediaStatisticField) Reset()         { *m = Article_SocialMediaStatisticField{} }
func (m *Article_SocialMediaStatisticField) String() string { return proto.CompactTextString(m) }
func (*Article_SocialMediaStatisticField) ProtoMessage()    {}
func (*Article_SocialMediaStatisticField) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0, 1}
}

func (m *Article_SocialMediaStatisticField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article_SocialMediaStatisticField.Unmarshal(m, b)
}
func (m *Article_SocialMediaStatisticField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article_SocialMediaStatisticField.Marshal(b, m, deterministic)
}
func (m *Article_SocialMediaStatisticField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article_SocialMediaStatisticField.Merge(m, src)
}
func (m *Article_SocialMediaStatisticField) XXX_Size() int {
	return xxx_messageInfo_Article_SocialMediaStatisticField.Size(m)
}
func (m *Article_SocialMediaStatisticField) XXX_DiscardUnknown() {
	xxx_messageInfo_Article_SocialMediaStatisticField.DiscardUnknown(m)
}

var xxx_messageInfo_Article_SocialMediaStatisticField proto.InternalMessageInfo

func (m *Article_SocialMediaStatisticField) GetLike() int32 {
	if m != nil {
		return m.Like
	}
	return 0
}

func (m *Article_SocialMediaStatisticField) GetShare() int32 {
	if m != nil {
		return m.Share
	}
	return 0
}

func (m *Article_SocialMediaStatisticField) GetComments() int32 {
	if m != nil {
		return m.Comments
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.Article_StatusType", Article_StatusType_name, Article_StatusType_value)
	proto.RegisterType((*Article)(nil), "pb.Article")
	proto.RegisterMapType((map[string]*Article_SocialMediaStatisticField)(nil), "pb.Article.SocialMediaStatisticEntry")
	proto.RegisterType((*Article_Tags)(nil), "pb.Article.Tags")
	proto.RegisterType((*Article_SocialMediaStatisticField)(nil), "pb.Article.SocialMediaStatisticField")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x8d, 0x14, 0xaf, 0x6c, 0x4d, 0x70, 0x62, 0x86, 0xb4, 0x6c, 0x75, 0x12, 0xa1, 0x05, 0x9d,
	0x04, 0x4d, 0x0f, 0x2d, 0xed, 0x29, 0x69, 0x12, 0x12, 0x48, 0x4b, 0x59, 0x3b, 0x87, 0x9e, 0xc2,
	0x5a, 0x9a, 0x3a, 0x4b, 0x64, 0x49, 0x68, 0xc7, 0x01, 0xff, 0x57, 0x3f, 0xb0, 0x68, 0x25, 0x25,
	0x2d, 0xf8, 0xd0, 0xdb, 0xbc, 0x79, 0x33, 0xef, 0xed, 0xbe, 0x5d, 0x98, 0xea, 0x86, 0x4d, 0x56,
	0x50, 0x5a, 0x37, 0x15, 0x57, 0xe8, 0xd7, 0xcb, 0x93, 0xdf, 0x02, 0xc6, 0x67, 0x5d, 0x17, 0x0f,
	0xc1, 0x37, 0xb9, 0xf4, 0x62, 0x2f, 0x11, 0xca, 0x37, 0x39, 0x1e, 0x83, 0x60, 0xc3, 0x05, 0x49,
	0x3f, 0xf6, 0x92, 0x50, 0x75, 0x00, 0x25, 0x8c, 0xb3, 0xaa, 0x64, 0x2a, 0x59, 0xee, 0xbb, 0xfe,
	0x00, 0x31, 0x85, 0xc0, 0xb2, 0xe6, 0x8d, 0x95, 0xa3, 0xd8, 0x4b, 0x0e, 0x4f, 0x5f, 0xa7, 0xf5,
	0x32, 0xed, 0xc5, 0xd3, 0xb9, 0x63, 0x16, 0xdb, 0x9a, 0x54, 0x3f, 0x85, 0x6f, 0x61, 0xc4, 0x7a,
	0x65, 0xa5, 0x88, 0xf7, 0x93, 0x83, 0xd3, 0xd9, 0xdf, 0xd3, 0x0b, 0xbd, 0xb2, 0xca, 0xb1, 0xf8,
	0x13, 0x8e, 0xe7, 0x55, 0x66, 0x74, 0xf1, 0x8d, 0x72, 0xa3, 0x5b, 0x19, 0x63, 0xd9, 0x64, 0x32,
	0x70, 0x5b, 0xef, 0xfe, 0xf1, 0xd8, 0x31, 0x77, 0x59, 0x72, 0xb3, 0x55, 0x3b, 0x25, 0x50, 0x42,
	0x60, 0x4a, 0x4b, 0x0d, 0xcb, 0x71, 0xec, 0x25, 0x93, 0xeb, 0x3d, 0xd5, 0xe3, 0x96, 0xd9, 0xd4,
	0xb9, 0x66, 0x92, 0x93, 0x81, 0xe9, 0x70, 0xcb, 0xe4, 0x54, 0x10, 0x93, 0x0c, 0x07, 0xa6, 0xc3,
	0x18, 0xc1, 0x58, 0x37, 0xd9, 0x83, 0x79, 0x22, 0x09, 0x3d, 0x35, 0x34, 0xa2, 0xf7, 0x30, 0x6a,
	0xaf, 0x84, 0xaf, 0x20, 0x60, 0xbd, 0xba, 0x7f, 0x8e, 0x59, 0xb0, 0x5e, 0xdd, 0xe4, 0x88, 0x30,
	0x2a, 0xf5, 0x7a, 0x08, 0xda, 0xd5, 0x91, 0x86, 0x37, 0xbb, 0x0e, 0x7d, 0x65, 0xa8, 0x70, 0x0b,
	0xb7, 0xe6, 0x91, 0x7a, 0x15, 0x57, 0xb7, 0xcf, 0x35, 0x7f, 0xd0, 0x4d, 0xa7, 0x22, 0x54, 0x07,
	0x30, 0x82, 0xc9, 0xd7, 0x6a, 0xbd, 0xa6, 0x92, 0xad, 0x7b, 0x2f, 0xa1, 0x9e, 0x71, 0x54, 0xee,
	0xb6, 0x70, 0x91, 0xe1, 0x0c, 0xf6, 0x1f, 0x69, 0xeb, 0x1c, 0x42, 0xd5, 0x96, 0xf8, 0x05, 0xc4,
	0x93, 0x2e, 0x36, 0x9d, 0xc1, 0x7f, 0x44, 0xef, 0x8e, 0xaa, 0xba, 0x9d, 0xcf, 0xfe, 0x27, 0xef,
	0xe4, 0x23, 0xc0, 0xcb, 0x37, 0xc0, 0x10, 0xc4, 0x85, 0x3a, 0xbb, 0x5a, 0xcc, 0xf6, 0x70, 0x0a,
	0xe1, 0x8f, 0xbb, 0xf3, 0xdb, 0x9b, 0xf9, 0xf5, 0xe5, 0xc5, 0xcc, 0xc3, 0x23, 0x38, 0xb8, 0xfb,
	0xfe, 0xd2, 0xf0, 0xcf, 0x8f, 0x60, 0x5a, 0x37, 0x55, 0x46, 0xd6, 0xde, 0x57, 0x25, 0x55, 0xbf,
	0x96, 0x81, 0xfb, 0xc1, 0x1f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x18, 0x31, 0x98, 0x8f, 0xd2,
	0x02, 0x00, 0x00,
}
