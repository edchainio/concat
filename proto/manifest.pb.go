// Code generated by protoc-gen-gogo.
// source: manifest.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Manifest struct {
	Entity    string        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	KeyId     string        `protobuf:"bytes,2,opt,name=keyId,proto3" json:"keyId,omitempty"`
	Body      *ManifestBody `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Timestamp int64         `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature []byte        `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Manifest) Reset()                    { *m = Manifest{} }
func (m *Manifest) String() string            { return proto1.CompactTextString(m) }
func (*Manifest) ProtoMessage()               {}
func (*Manifest) Descriptor() ([]byte, []int) { return fileDescriptorManifest, []int{0} }

func (m *Manifest) GetBody() *ManifestBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type ManifestBody struct {
	// Types that are valid to be assigned to Body:
	//	*ManifestBody_Node
	Body isManifestBody_Body `protobuf_oneof:"body"`
}

func (m *ManifestBody) Reset()                    { *m = ManifestBody{} }
func (m *ManifestBody) String() string            { return proto1.CompactTextString(m) }
func (*ManifestBody) ProtoMessage()               {}
func (*ManifestBody) Descriptor() ([]byte, []int) { return fileDescriptorManifest, []int{1} }

type isManifestBody_Body interface {
	isManifestBody_Body()
}

type ManifestBody_Node struct {
	Node *NodeManifest `protobuf:"bytes,1,opt,name=node,oneof"`
}

func (*ManifestBody_Node) isManifestBody_Body() {}

func (m *ManifestBody) GetBody() isManifestBody_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *ManifestBody) GetNode() *NodeManifest {
	if x, ok := m.GetBody().(*ManifestBody_Node); ok {
		return x.Node
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ManifestBody) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _ManifestBody_OneofMarshaler, _ManifestBody_OneofUnmarshaler, _ManifestBody_OneofSizer, []interface{}{
		(*ManifestBody_Node)(nil),
	}
}

func _ManifestBody_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*ManifestBody)
	// body
	switch x := m.Body.(type) {
	case *ManifestBody_Node:
		_ = b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Node); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ManifestBody.Body has unexpected type %T", x)
	}
	return nil
}

func _ManifestBody_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*ManifestBody)
	switch tag {
	case 1: // body.node
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(NodeManifest)
		err := b.DecodeMessage(msg)
		m.Body = &ManifestBody_Node{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ManifestBody_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*ManifestBody)
	// body
	switch x := m.Body.(type) {
	case *ManifestBody_Node:
		s := proto1.Size(x.Node)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NodeManifest struct {
	Peer      string `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Publisher string `protobuf:"bytes,2,opt,name=publisher,proto3" json:"publisher,omitempty"`
}

func (m *NodeManifest) Reset()                    { *m = NodeManifest{} }
func (m *NodeManifest) String() string            { return proto1.CompactTextString(m) }
func (*NodeManifest) ProtoMessage()               {}
func (*NodeManifest) Descriptor() ([]byte, []int) { return fileDescriptorManifest, []int{2} }

func init() {
	proto1.RegisterType((*Manifest)(nil), "proto.Manifest")
	proto1.RegisterType((*ManifestBody)(nil), "proto.ManifestBody")
	proto1.RegisterType((*NodeManifest)(nil), "proto.NodeManifest")
}

func init() { proto1.RegisterFile("manifest.proto", fileDescriptorManifest) }

var fileDescriptorManifest = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x31, 0x9b, 0x44, 0xec, 0x10, 0x51, 0x0c, 0x08, 0xb9, 0xa0, 0x88, 0xd2, 0x10, 0x9a,
	0x2d, 0xe0, 0x02, 0xb0, 0x15, 0x14, 0x50, 0xf8, 0x06, 0x89, 0x3c, 0x80, 0x05, 0xb1, 0x2d, 0xdb,
	0x5b, 0xf8, 0x30, 0xdc, 0x15, 0xad, 0xe3, 0x95, 0x53, 0xd9, 0xf3, 0xfe, 0xd7, 0x9f, 0x3f, 0x70,
	0x35, 0x8f, 0x5a, 0x7d, 0x92, 0x0f, 0x3b, 0xeb, 0x4c, 0x30, 0x58, 0xa7, 0xa7, 0xff, 0x63, 0x70,
	0xf1, 0x9e, 0x15, 0xbc, 0x85, 0x86, 0x74, 0x50, 0x21, 0x72, 0xd6, 0xb1, 0x61, 0x2b, 0xf2, 0x84,
	0x37, 0x50, 0xff, 0x50, 0x7c, 0x93, 0xfc, 0x3c, 0xe1, 0x65, 0xc0, 0x7b, 0xa8, 0x26, 0x23, 0x23,
	0xdf, 0x74, 0x6c, 0xb8, 0x7c, 0xbc, 0x5e, 0x72, 0x77, 0xa7, 0xb0, 0xbd, 0x91, 0x51, 0x24, 0x03,
	0xde, 0xc1, 0x36, 0xa8, 0x99, 0x7c, 0x18, 0x67, 0xcb, 0xab, 0x8e, 0x0d, 0x1b, 0x51, 0xc0, 0x51,
	0xf5, 0xea, 0x4b, 0x8f, 0xe1, 0xe0, 0x88, 0xd7, 0x1d, 0x1b, 0x5a, 0x51, 0x40, 0xff, 0x02, 0xed,
	0x3a, 0x11, 0x1f, 0xa0, 0xd2, 0x46, 0x52, 0x2a, 0x58, 0x96, 0x7e, 0x18, 0x49, 0x27, 0xdb, 0xeb,
	0x99, 0x48, 0x96, 0x7d, 0xb3, 0xf4, 0xeb, 0x9f, 0xa1, 0x5d, 0xeb, 0x88, 0x50, 0x59, 0x22, 0x97,
	0x6f, 0x4c, 0xff, 0x63, 0x09, 0x7b, 0x98, 0x7e, 0x95, 0xff, 0x26, 0x97, 0xaf, 0x2c, 0x60, 0x6a,
	0xd2, 0x96, 0xa7, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x51, 0xb9, 0xd1, 0x44, 0x01, 0x00,
	0x00,
}
