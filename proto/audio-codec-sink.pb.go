// Code generated by protoc-gen-go.
// source: github.com/yeyus/dmr-audio-sink/proto/audio-codec-sink.proto
// DO NOT EDIT!

/*
Package com_ea7jmf_dmr_srv_audiocodecsink is a generated protocol buffer package.

It is generated from these files:
	github.com/yeyus/dmr-audio-sink/proto/audio-codec-sink.proto

It has these top-level messages:
	DecodeRequest
	DecodeResponse
	GetVolumeRequest
	SetVolumeRequest
	VolumeResponse
*/
package com_ea7jmf_dmr_srv_audiocodecsink

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type DecodeRequest_Codec int32

const (
	DecodeRequest_CODEC2MODE3200 DecodeRequest_Codec = 0
	DecodeRequest_CODEC2MODE2400 DecodeRequest_Codec = 1
	DecodeRequest_CODEC2MODE1600 DecodeRequest_Codec = 2
	DecodeRequest_CODEC2MODE1400 DecodeRequest_Codec = 3
	DecodeRequest_CODEC2MODE1300 DecodeRequest_Codec = 4
	DecodeRequest_CODEC2MODE1200 DecodeRequest_Codec = 5
	DecodeRequest_CODEC2MODE700  DecodeRequest_Codec = 6
	DecodeRequest_CODEC2MODE700B DecodeRequest_Codec = 7
	DecodeRequest_AMBE3600X2400  DecodeRequest_Codec = 8
	DecodeRequest_AMBE3600X2450  DecodeRequest_Codec = 9
	DecodeRequest_IMBE7200X4400  DecodeRequest_Codec = 10
)

var DecodeRequest_Codec_name = map[int32]string{
	0:  "CODEC2MODE3200",
	1:  "CODEC2MODE2400",
	2:  "CODEC2MODE1600",
	3:  "CODEC2MODE1400",
	4:  "CODEC2MODE1300",
	5:  "CODEC2MODE1200",
	6:  "CODEC2MODE700",
	7:  "CODEC2MODE700B",
	8:  "AMBE3600X2400",
	9:  "AMBE3600X2450",
	10: "IMBE7200X4400",
}
var DecodeRequest_Codec_value = map[string]int32{
	"CODEC2MODE3200": 0,
	"CODEC2MODE2400": 1,
	"CODEC2MODE1600": 2,
	"CODEC2MODE1400": 3,
	"CODEC2MODE1300": 4,
	"CODEC2MODE1200": 5,
	"CODEC2MODE700":  6,
	"CODEC2MODE700B": 7,
	"AMBE3600X2400":  8,
	"AMBE3600X2450":  9,
	"IMBE7200X4400":  10,
}

func (x DecodeRequest_Codec) String() string {
	return proto.EnumName(DecodeRequest_Codec_name, int32(x))
}
func (DecodeRequest_Codec) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type DecodeRequest struct {
	Codec      DecodeRequest_Codec `protobuf:"varint,1,opt,name=codec,enum=com.ea7jmf.dmr.srv.audiocodecsink.DecodeRequest_Codec" json:"codec,omitempty"`
	VoiceBits  []byte              `protobuf:"bytes,2,opt,name=voiceBits,proto3" json:"voiceBits,omitempty"`
	ClientId   uint32              `protobuf:"varint,3,opt,name=clientId" json:"clientId,omitempty"`
	SequenceNo uint64              `protobuf:"varint,4,opt,name=sequenceNo" json:"sequenceNo,omitempty"`
}

func (m *DecodeRequest) Reset()                    { *m = DecodeRequest{} }
func (m *DecodeRequest) String() string            { return proto.CompactTextString(m) }
func (*DecodeRequest) ProtoMessage()               {}
func (*DecodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DecodeRequest) GetCodec() DecodeRequest_Codec {
	if m != nil {
		return m.Codec
	}
	return DecodeRequest_CODEC2MODE3200
}

func (m *DecodeRequest) GetVoiceBits() []byte {
	if m != nil {
		return m.VoiceBits
	}
	return nil
}

func (m *DecodeRequest) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *DecodeRequest) GetSequenceNo() uint64 {
	if m != nil {
		return m.SequenceNo
	}
	return 0
}

type DecodeResponse struct {
	DecodedBits uint32 `protobuf:"varint,1,opt,name=decodedBits" json:"decodedBits,omitempty"`
	Error       string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *DecodeResponse) Reset()                    { *m = DecodeResponse{} }
func (m *DecodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DecodeResponse) ProtoMessage()               {}
func (*DecodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DecodeResponse) GetDecodedBits() uint32 {
	if m != nil {
		return m.DecodedBits
	}
	return 0
}

func (m *DecodeResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetVolumeRequest struct {
}

func (m *GetVolumeRequest) Reset()                    { *m = GetVolumeRequest{} }
func (m *GetVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeRequest) ProtoMessage()               {}
func (*GetVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SetVolumeRequest struct {
	Volume float32 `protobuf:"fixed32,1,opt,name=volume" json:"volume,omitempty"`
}

func (m *SetVolumeRequest) Reset()                    { *m = SetVolumeRequest{} }
func (m *SetVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*SetVolumeRequest) ProtoMessage()               {}
func (*SetVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetVolumeRequest) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type VolumeResponse struct {
	Volume float32 `protobuf:"fixed32,1,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeResponse) Reset()                    { *m = VolumeResponse{} }
func (m *VolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeResponse) ProtoMessage()               {}
func (*VolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VolumeResponse) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func init() {
	proto.RegisterType((*DecodeRequest)(nil), "com.ea7jmf.dmr.srv.audiocodecsink.DecodeRequest")
	proto.RegisterType((*DecodeResponse)(nil), "com.ea7jmf.dmr.srv.audiocodecsink.DecodeResponse")
	proto.RegisterType((*GetVolumeRequest)(nil), "com.ea7jmf.dmr.srv.audiocodecsink.GetVolumeRequest")
	proto.RegisterType((*SetVolumeRequest)(nil), "com.ea7jmf.dmr.srv.audiocodecsink.SetVolumeRequest")
	proto.RegisterType((*VolumeResponse)(nil), "com.ea7jmf.dmr.srv.audiocodecsink.VolumeResponse")
	proto.RegisterEnum("com.ea7jmf.dmr.srv.audiocodecsink.DecodeRequest_Codec", DecodeRequest_Codec_name, DecodeRequest_Codec_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Publisher API

type Publisher interface {
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

type publisher struct {
	c     client.Client
	topic string
}

func (p *publisher) Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error {
	return p.c.Publish(ctx, p.c.NewPublication(p.topic, msg), opts...)
}

func NewPublisher(topic string, c client.Client) Publisher {
	if c == nil {
		c = client.NewClient()
	}
	return &publisher{c, topic}
}

// Subscriber API

func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return s.Subscribe(s.NewSubscriber(topic, h, opts...))
}

// Client API for AudioCodecSink service

type AudioCodecSinkClient interface {
	Decode(ctx context.Context, in *DecodeRequest, opts ...client.CallOption) (*DecodeResponse, error)
	SetVolume(ctx context.Context, in *SetVolumeRequest, opts ...client.CallOption) (*VolumeResponse, error)
	GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...client.CallOption) (*VolumeResponse, error)
}

type audioCodecSinkClient struct {
	c           client.Client
	serviceName string
}

func NewAudioCodecSinkClient(serviceName string, c client.Client) AudioCodecSinkClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.ea7jmf.dmr.srv.audiocodecsink"
	}
	return &audioCodecSinkClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *audioCodecSinkClient) Decode(ctx context.Context, in *DecodeRequest, opts ...client.CallOption) (*DecodeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AudioCodecSink.Decode", in)
	out := new(DecodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioCodecSinkClient) SetVolume(ctx context.Context, in *SetVolumeRequest, opts ...client.CallOption) (*VolumeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AudioCodecSink.SetVolume", in)
	out := new(VolumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioCodecSinkClient) GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...client.CallOption) (*VolumeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AudioCodecSink.GetVolume", in)
	out := new(VolumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AudioCodecSink service

type AudioCodecSinkHandler interface {
	Decode(context.Context, *DecodeRequest, *DecodeResponse) error
	SetVolume(context.Context, *SetVolumeRequest, *VolumeResponse) error
	GetVolume(context.Context, *GetVolumeRequest, *VolumeResponse) error
}

func RegisterAudioCodecSinkHandler(s server.Server, hdlr AudioCodecSinkHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AudioCodecSink{hdlr}, opts...))
}

type AudioCodecSink struct {
	AudioCodecSinkHandler
}

func (h *AudioCodecSink) Decode(ctx context.Context, in *DecodeRequest, out *DecodeResponse) error {
	return h.AudioCodecSinkHandler.Decode(ctx, in, out)
}

func (h *AudioCodecSink) SetVolume(ctx context.Context, in *SetVolumeRequest, out *VolumeResponse) error {
	return h.AudioCodecSinkHandler.SetVolume(ctx, in, out)
}

func (h *AudioCodecSink) GetVolume(ctx context.Context, in *GetVolumeRequest, out *VolumeResponse) error {
	return h.AudioCodecSinkHandler.GetVolume(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/yeyus/dmr-audio-sink/proto/audio-codec-sink.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x9b, 0xfd, 0x67, 0xf7, 0xe8, 0x86, 0xf1, 0x20, 0xb2, 0x14, 0x91, 0x98, 0xab, 0x20,
	0x74, 0xf6, 0x34, 0xa9, 0xbb, 0x37, 0xde, 0x74, 0xff, 0xb0, 0x16, 0x5c, 0x0b, 0x29, 0x48, 0x6f,
	0xdb, 0x64, 0xd4, 0xd8, 0x26, 0x53, 0x33, 0xc9, 0x42, 0x1f, 0xc5, 0xa7, 0xf3, 0x11, 0x7c, 0x05,
	0x99, 0x49, 0xb7, 0x6d, 0x02, 0x62, 0x17, 0xbc, 0x3c, 0xbf, 0xfd, 0xf6, 0xfb, 0xce, 0x7c, 0x33,
	0x81, 0xf7, 0x5f, 0x93, 0xe2, 0x5b, 0x79, 0xc1, 0x23, 0x99, 0x8e, 0x6e, 0xc4, 0x4d, 0xa9, 0x46,
	0x71, 0x9a, 0xef, 0x9f, 0x97, 0x71, 0x22, 0xf7, 0x55, 0x92, 0x5d, 0x8e, 0xae, 0x73, 0x59, 0xc8,
	0x51, 0x05, 0x22, 0x19, 0x8b, 0xc8, 0x60, 0x6e, 0x30, 0xbe, 0x89, 0x64, 0xca, 0xc5, 0xf9, 0xe4,
	0x7b, 0xfa, 0x85, 0xc7, 0x69, 0xce, 0x55, 0xbe, 0xe6, 0x46, 0x6a, 0x94, 0x5a, 0xe8, 0xfe, 0x6c,
	0xc3, 0x60, 0x2e, 0xf4, 0x1c, 0x8a, 0x1f, 0xa5, 0x50, 0x05, 0x7e, 0x84, 0xae, 0xf9, 0x79, 0x68,
	0x39, 0x96, 0x67, 0xfb, 0x63, 0xfe, 0x4f, 0x13, 0x5e, 0x33, 0xe0, 0x33, 0xcd, 0xc3, 0xca, 0x04,
	0x5f, 0x41, 0x7f, 0x2d, 0x93, 0x48, 0x4c, 0x93, 0x42, 0x0d, 0x5b, 0x8e, 0xe5, 0x3d, 0x0b, 0xef,
	0x01, 0xee, 0xc1, 0x6e, 0x74, 0x95, 0x88, 0xac, 0x38, 0x8e, 0x87, 0x6d, 0xc7, 0xf2, 0x06, 0xe1,
	0xdd, 0x8c, 0xaf, 0x01, 0x94, 0x76, 0xcc, 0x22, 0xf1, 0x49, 0x0e, 0x3b, 0x8e, 0xe5, 0x75, 0xc2,
	0x07, 0xc4, 0xfd, 0x65, 0x41, 0xd7, 0x44, 0x21, 0x82, 0x3d, 0x3b, 0x99, 0x2f, 0x66, 0xfe, 0xea,
	0x64, 0xbe, 0x08, 0x7c, 0x22, 0xb6, 0x53, 0x67, 0xfe, 0x21, 0x11, 0xb3, 0xea, 0xec, 0x60, 0x4c,
	0xc4, 0x5a, 0x0d, 0xa6, 0x75, 0xed, 0x06, 0x0b, 0x88, 0x58, 0xa7, 0xc1, 0x74, 0x46, 0x17, 0x9f,
	0xc3, 0xe0, 0x9e, 0x4d, 0x88, 0x58, 0xaf, 0x2e, 0x9b, 0x10, 0x4d, 0xd9, 0x13, 0x2d, 0x3b, 0x5a,
	0x4d, 0x17, 0xc1, 0x98, 0xe8, 0xcc, 0x6c, 0xb2, 0xdb, 0x40, 0xef, 0x88, 0xf5, 0x35, 0x3a, 0x5e,
	0x4d, 0x17, 0x13, 0x9f, 0xe8, 0xec, 0x50, 0xab, 0xc0, 0xfd, 0x00, 0xf6, 0xa6, 0x59, 0x75, 0x2d,
	0x33, 0x25, 0xd0, 0x81, 0xa7, 0xb1, 0x21, 0xb1, 0xe9, 0xd3, 0x32, 0x95, 0x3d, 0x44, 0xf8, 0x02,
	0xba, 0x22, 0xcf, 0x65, 0x6e, 0xba, 0xee, 0x87, 0xd5, 0xe0, 0x22, 0xb0, 0xa5, 0x28, 0x3e, 0xcb,
	0xab, 0x32, 0xdd, 0x5c, 0x93, 0xfb, 0x16, 0xd8, 0x69, 0x83, 0xe1, 0x4b, 0xe8, 0xad, 0x0d, 0x30,
	0xd6, 0xad, 0xf0, 0x76, 0x72, 0x3d, 0xb0, 0x37, 0xc2, 0xdb, 0x4d, 0xfe, 0xa2, 0xf4, 0x7f, 0xb7,
	0xc0, 0x3e, 0xd2, 0xaf, 0xc3, 0x5c, 0xcd, 0x69, 0x92, 0x5d, 0xa2, 0x84, 0x5e, 0x75, 0x0c, 0xa4,
	0x6d, 0xdf, 0xd2, 0xde, 0xc1, 0x16, 0xff, 0xa8, 0x36, 0x73, 0x77, 0xb0, 0x84, 0xfe, 0xdd, 0xc9,
	0x30, 0x78, 0x84, 0x43, 0xb3, 0x87, 0x47, 0xc5, 0xd6, 0x0b, 0xa9, 0x62, 0x97, 0x5b, 0xc5, 0x2e,
	0xff, 0x47, 0xec, 0x45, 0xcf, 0x7c, 0xeb, 0xc1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf5,
	0x55, 0xdf, 0x2b, 0x04, 0x00, 0x00,
}
