// Code generated by protoc-gen-go.
// source: proto.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/nvcook42/morgoth/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto1.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type State struct {
	Time             *int64   `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	State            *string  `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Service          *string  `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Host             *string  `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	Description      *string  `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Once             *bool    `protobuf:"varint,6,opt,name=once" json:"once,omitempty"`
	Tags             []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	Ttl              *float32 `protobuf:"fixed32,8,opt,name=ttl" json:"ttl,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *State) Reset()         { *this = State{} }
func (this *State) String() string { return proto1.CompactTextString(this) }
func (*State) ProtoMessage()       {}

func (this *State) GetTime() int64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *State) GetState() string {
	if this != nil && this.State != nil {
		return *this.State
	}
	return ""
}

func (this *State) GetService() string {
	if this != nil && this.Service != nil {
		return *this.Service
	}
	return ""
}

func (this *State) GetHost() string {
	if this != nil && this.Host != nil {
		return *this.Host
	}
	return ""
}

func (this *State) GetDescription() string {
	if this != nil && this.Description != nil {
		return *this.Description
	}
	return ""
}

func (this *State) GetOnce() bool {
	if this != nil && this.Once != nil {
		return *this.Once
	}
	return false
}

func (this *State) GetTags() []string {
	if this != nil {
		return this.Tags
	}
	return nil
}

func (this *State) GetTtl() float32 {
	if this != nil && this.Ttl != nil {
		return *this.Ttl
	}
	return 0
}

type Event struct {
	Time             *int64       `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	State            *string      `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Service          *string      `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Host             *string      `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	Description      *string      `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Tags             []string     `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	Ttl              *float32     `protobuf:"fixed32,8,opt,name=ttl" json:"ttl,omitempty"`
	Attributes       []*Attribute `protobuf:"bytes,9,rep,name=attributes" json:"attributes,omitempty"`
	MetricSint64     *int64       `protobuf:"zigzag64,13,opt,name=metric_sint64" json:"metric_sint64,omitempty"`
	MetricD          *float64     `protobuf:"fixed64,14,opt,name=metric_d" json:"metric_d,omitempty"`
	MetricF          *float32     `protobuf:"fixed32,15,opt,name=metric_f" json:"metric_f,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (this *Event) Reset()         { *this = Event{} }
func (this *Event) String() string { return proto1.CompactTextString(this) }
func (*Event) ProtoMessage()       {}

func (this *Event) GetTime() int64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *Event) GetState() string {
	if this != nil && this.State != nil {
		return *this.State
	}
	return ""
}

func (this *Event) GetService() string {
	if this != nil && this.Service != nil {
		return *this.Service
	}
	return ""
}

func (this *Event) GetHost() string {
	if this != nil && this.Host != nil {
		return *this.Host
	}
	return ""
}

func (this *Event) GetDescription() string {
	if this != nil && this.Description != nil {
		return *this.Description
	}
	return ""
}

func (this *Event) GetTags() []string {
	if this != nil {
		return this.Tags
	}
	return nil
}

func (this *Event) GetTtl() float32 {
	if this != nil && this.Ttl != nil {
		return *this.Ttl
	}
	return 0
}

func (this *Event) GetAttributes() []*Attribute {
	if this != nil {
		return this.Attributes
	}
	return nil
}

func (this *Event) GetMetricSint64() int64 {
	if this != nil && this.MetricSint64 != nil {
		return *this.MetricSint64
	}
	return 0
}

func (this *Event) GetMetricD() float64 {
	if this != nil && this.MetricD != nil {
		return *this.MetricD
	}
	return 0
}

func (this *Event) GetMetricF() float32 {
	if this != nil && this.MetricF != nil {
		return *this.MetricF
	}
	return 0
}

type Query struct {
	String_          *string `protobuf:"bytes,1,opt,name=string" json:"string,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Query) Reset()         { *this = Query{} }
func (this *Query) String() string { return proto1.CompactTextString(this) }
func (*Query) ProtoMessage()       {}

func (this *Query) GetString_() string {
	if this != nil && this.String_ != nil {
		return *this.String_
	}
	return ""
}

type Msg struct {
	Ok               *bool    `protobuf:"varint,2,opt,name=ok" json:"ok,omitempty"`
	Error            *string  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	States           []*State `protobuf:"bytes,4,rep,name=states" json:"states,omitempty"`
	Query            *Query   `protobuf:"bytes,5,opt,name=query" json:"query,omitempty"`
	Events           []*Event `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Msg) Reset()         { *this = Msg{} }
func (this *Msg) String() string { return proto1.CompactTextString(this) }
func (*Msg) ProtoMessage()       {}

func (this *Msg) GetOk() bool {
	if this != nil && this.Ok != nil {
		return *this.Ok
	}
	return false
}

func (this *Msg) GetError() string {
	if this != nil && this.Error != nil {
		return *this.Error
	}
	return ""
}

func (this *Msg) GetStates() []*State {
	if this != nil {
		return this.States
	}
	return nil
}

func (this *Msg) GetQuery() *Query {
	if this != nil {
		return this.Query
	}
	return nil
}

func (this *Msg) GetEvents() []*Event {
	if this != nil {
		return this.Events
	}
	return nil
}

type Attribute struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Attribute) Reset()         { *this = Attribute{} }
func (this *Attribute) String() string { return proto1.CompactTextString(this) }
func (*Attribute) ProtoMessage()       {}

func (this *Attribute) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *Attribute) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

func init() {
}
