// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/static-srv/proto/pubsub/pubsub.proto

/*
Package go_micro_srv_static is a generated protocol buffer package.

It is generated from these files:
	server/static-srv/proto/pubsub/pubsub.proto

It has these top-level messages:
	PublishBulkNotification
	PublishPersonalNotification
	BulkNotification
	PersonalNotification
	Alert
	SubscribeRequest
	SubscribeResponse
*/
package go_micro_srv_static

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PublishBulkNotification struct {
	Notification *BulkNotification `protobuf:"bytes,1,opt,name=notification" json:"notification,omitempty"`
}

func (m *PublishBulkNotification) Reset()                    { *m = PublishBulkNotification{} }
func (m *PublishBulkNotification) String() string            { return proto.CompactTextString(m) }
func (*PublishBulkNotification) ProtoMessage()               {}
func (*PublishBulkNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublishBulkNotification) GetNotification() *BulkNotification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type PublishPersonalNotification struct {
	Notification []*PersonalNotification `protobuf:"bytes,1,rep,name=notification" json:"notification,omitempty"`
}

func (m *PublishPersonalNotification) Reset()                    { *m = PublishPersonalNotification{} }
func (m *PublishPersonalNotification) String() string            { return proto.CompactTextString(m) }
func (*PublishPersonalNotification) ProtoMessage()               {}
func (*PublishPersonalNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PublishPersonalNotification) GetNotification() []*PersonalNotification {
	if m != nil {
		return m.Notification
	}
	return nil
}

// single message to all users
type BulkNotification struct {
	UserIds []string          `protobuf:"bytes,1,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Alert   *Alert            `protobuf:"bytes,3,opt,name=alert" json:"alert,omitempty"`
	Data    map[string]string `protobuf:"bytes,4,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BulkNotification) Reset()                    { *m = BulkNotification{} }
func (m *BulkNotification) String() string            { return proto.CompactTextString(m) }
func (*BulkNotification) ProtoMessage()               {}
func (*BulkNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BulkNotification) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *BulkNotification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BulkNotification) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *BulkNotification) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

// single (personalised) message for every user
type PersonalNotification struct {
	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Alert   *Alert `protobuf:"bytes,3,opt,name=alert" json:"alert,omitempty"`
}

func (m *PersonalNotification) Reset()                    { *m = PersonalNotification{} }
func (m *PersonalNotification) String() string            { return proto.CompactTextString(m) }
func (*PersonalNotification) ProtoMessage()               {}
func (*PersonalNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PersonalNotification) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PersonalNotification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PersonalNotification) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

type Alert struct {
	Title        string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Body         string   `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	Subtitle     string   `protobuf:"bytes,3,opt,name=subtitle" json:"subtitle,omitempty"`
	Action       string   `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	ActionLocKey string   `protobuf:"bytes,5,opt,name=actionLocKey" json:"actionLocKey,omitempty"`
	LaunchImage  string   `protobuf:"bytes,6,opt,name=launchImage" json:"launchImage,omitempty"`
	LocKey       string   `protobuf:"bytes,7,opt,name=locKey" json:"locKey,omitempty"`
	TitleLocKey  string   `protobuf:"bytes,8,opt,name=titleLocKey" json:"titleLocKey,omitempty"`
	LocArgs      []string `protobuf:"bytes,9,rep,name=locArgs" json:"locArgs,omitempty"`
	TitleLocArgs []string `protobuf:"bytes,10,rep,name=titleLocArgs" json:"titleLocArgs,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Alert) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Alert) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Alert) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *Alert) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Alert) GetActionLocKey() string {
	if m != nil {
		return m.ActionLocKey
	}
	return ""
}

func (m *Alert) GetLaunchImage() string {
	if m != nil {
		return m.LaunchImage
	}
	return ""
}

func (m *Alert) GetLocKey() string {
	if m != nil {
		return m.LocKey
	}
	return ""
}

func (m *Alert) GetTitleLocKey() string {
	if m != nil {
		return m.TitleLocKey
	}
	return ""
}

func (m *Alert) GetLocArgs() []string {
	if m != nil {
		return m.LocArgs
	}
	return nil
}

func (m *Alert) GetTitleLocArgs() []string {
	if m != nil {
		return m.TitleLocArgs
	}
	return nil
}

type SubscribeRequest struct {
	Channel string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SubscribeRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type SubscribeResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *SubscribeResponse) Reset()                    { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()               {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SubscribeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*PublishBulkNotification)(nil), "go.micro.srv.static.PublishBulkNotification")
	proto.RegisterType((*PublishPersonalNotification)(nil), "go.micro.srv.static.PublishPersonalNotification")
	proto.RegisterType((*BulkNotification)(nil), "go.micro.srv.static.BulkNotification")
	proto.RegisterType((*PersonalNotification)(nil), "go.micro.srv.static.PersonalNotification")
	proto.RegisterType((*Alert)(nil), "go.micro.srv.static.Alert")
	proto.RegisterType((*SubscribeRequest)(nil), "go.micro.srv.static.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "go.micro.srv.static.SubscribeResponse")
}

func init() { proto.RegisterFile("server/static-srv/proto/pubsub/pubsub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x8e, 0x13, 0x31,
	0x10, 0xd4, 0xe4, 0x9d, 0xce, 0x1e, 0x82, 0x59, 0xb1, 0x43, 0xb8, 0x44, 0x23, 0x21, 0x2d, 0x82,
	0x9d, 0xa0, 0xe5, 0x00, 0xe2, 0xb6, 0x3c, 0x0e, 0x11, 0x0f, 0xad, 0x86, 0x0f, 0x40, 0x1e, 0x8f,
	0x49, 0xac, 0x75, 0xc6, 0xc1, 0x6d, 0x47, 0x8a, 0xf8, 0x1c, 0x3e, 0x92, 0x2b, 0xf2, 0x23, 0x51,
	0xb2, 0x99, 0x03, 0x07, 0x4e, 0x71, 0x55, 0x57, 0x57, 0x75, 0xb7, 0x26, 0xf0, 0x1c, 0xb9, 0xde,
	0x70, 0x3d, 0x43, 0x43, 0x8d, 0x60, 0x57, 0xa8, 0x37, 0xb3, 0xb5, 0x56, 0x46, 0xcd, 0xd6, 0xb6,
	0x44, 0x5b, 0xc6, 0x9f, 0xdc, 0x73, 0xe4, 0xe1, 0x42, 0xe5, 0x2b, 0xc1, 0xb4, 0xca, 0x51, 0x6f,
	0xf2, 0xd0, 0x92, 0x55, 0x70, 0x71, 0x6b, 0x4b, 0x29, 0x70, 0xf9, 0xce, 0xca, 0xbb, 0xaf, 0xca,
	0x88, 0x1f, 0x82, 0x51, 0x23, 0x54, 0x4d, 0xe6, 0x70, 0x56, 0x1f, 0xe0, 0x34, 0x99, 0x26, 0x97,
	0xa3, 0xeb, 0xa7, 0x79, 0x83, 0x4d, 0x7e, 0xbf, 0xb9, 0x38, 0x6a, 0xcd, 0x24, 0x3c, 0x89, 0x29,
	0xb7, 0x5c, 0xa3, 0xaa, 0xa9, 0x3c, 0x4a, 0xfa, 0x72, 0x92, 0xd4, 0xbe, 0x1c, 0x5d, 0x3f, 0x6b,
	0x4c, 0x6a, 0x32, 0xb8, 0x97, 0xf6, 0x27, 0x81, 0xf1, 0xc9, 0x36, 0x8f, 0x61, 0x60, 0x91, 0xeb,
	0xef, 0xa2, 0x42, 0xef, 0x3f, 0x2c, 0xfa, 0x0e, 0xcf, 0x2b, 0x24, 0x29, 0xf4, 0x57, 0x1c, 0x91,
	0x2e, 0x78, 0xda, 0x9a, 0x26, 0xae, 0x12, 0x21, 0x79, 0x09, 0x5d, 0x2a, 0xb9, 0x36, 0x69, 0xdb,
	0xef, 0x3e, 0x69, 0x9c, 0xe8, 0xc6, 0x29, 0x8a, 0x20, 0x24, 0xef, 0xa1, 0x53, 0x51, 0x43, 0xd3,
	0x8e, 0x5f, 0x61, 0xf6, 0x4f, 0xc7, 0xca, 0x3f, 0x50, 0x43, 0x3f, 0xd6, 0x46, 0x6f, 0x0b, 0xdf,
	0x3c, 0x79, 0x0d, 0xc3, 0x3d, 0x45, 0xc6, 0xd0, 0xbe, 0xe3, 0x5b, 0x7f, 0xfd, 0x61, 0xe1, 0x9e,
	0xe4, 0x1c, 0xba, 0x1b, 0x2a, 0xed, 0x6e, 0xda, 0x00, 0xde, 0xb6, 0xde, 0x24, 0xd9, 0x2f, 0x38,
	0x6f, 0x3c, 0xf0, 0x05, 0xf4, 0xe3, 0xf2, 0xd1, 0xa7, 0x17, 0x76, 0xff, 0x9f, 0xab, 0x67, 0xbf,
	0x5b, 0xd0, 0xf5, 0x84, 0x1b, 0xd0, 0x08, 0x23, 0x79, 0x0c, 0x0b, 0x80, 0x10, 0xe8, 0x94, 0xaa,
	0xda, 0xc6, 0x20, 0xff, 0x26, 0x13, 0x18, 0xa0, 0x2d, 0x83, 0xb8, 0xed, 0xf9, 0x3d, 0x26, 0x8f,
	0xa0, 0x47, 0x99, 0xff, 0x1e, 0x3a, 0x61, 0xe6, 0x80, 0x48, 0x06, 0x67, 0xe1, 0xf5, 0x59, 0xb1,
	0x4f, 0x7c, 0x9b, 0x76, 0x7d, 0xf5, 0x88, 0x23, 0x53, 0x18, 0x49, 0x6a, 0x6b, 0xb6, 0x9c, 0xaf,
	0xdc, 0x6e, 0x3d, 0x2f, 0x39, 0xa4, 0x9c, 0xbb, 0x0c, 0xfd, 0xfd, 0xe0, 0x2e, 0xf7, 0x9d, 0x3e,
	0x3e, 0x9a, 0x0f, 0x42, 0xe7, 0x01, 0xe5, 0x6e, 0x26, 0x15, 0xbb, 0xd1, 0x0b, 0x4c, 0x87, 0xe1,
	0x43, 0x8a, 0xd0, 0x4d, 0xb6, 0x13, 0xfa, 0x32, 0xf8, 0xf2, 0x11, 0x97, 0xbd, 0x80, 0xf1, 0x37,
	0x5b, 0x22, 0xd3, 0xa2, 0xe4, 0x05, 0xff, 0x69, 0x39, 0x1a, 0xe7, 0xc8, 0x96, 0xb4, 0xae, 0xb9,
	0x8c, 0x17, 0xdb, 0xc1, 0xec, 0x0a, 0x1e, 0x1c, 0xa8, 0x71, 0xad, 0x6a, 0xe4, 0x4e, 0x8e, 0x96,
	0x31, 0x8e, 0xe8, 0xe5, 0x83, 0x62, 0x07, 0xcb, 0x9e, 0xff, 0xa7, 0xbf, 0xfa, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0xbd, 0x86, 0x03, 0x18, 0x04, 0x00, 0x00,
}
