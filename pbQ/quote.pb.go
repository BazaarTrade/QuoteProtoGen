// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/quote.proto

package pbQ

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair string `protobuf:"bytes,1,opt,name=Pair,proto3" json:"Pair,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{0}
}

func (x *Pair) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

type PairParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair            string  `protobuf:"bytes,1,opt,name=Pair,proto3" json:"Pair,omitempty"`
	PricePrecisions []int32 `protobuf:"varint,2,rep,packed,name=PricePrecisions,proto3" json:"PricePrecisions,omitempty"`
	QtyPrecision    int32   `protobuf:"varint,3,opt,name=QtyPrecision,proto3" json:"QtyPrecision,omitempty"`
}

func (x *PairParams) Reset() {
	*x = PairParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairParams) ProtoMessage() {}

func (x *PairParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairParams.ProtoReflect.Descriptor instead.
func (*PairParams) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{1}
}

func (x *PairParams) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *PairParams) GetPricePrecisions() []int32 {
	if x != nil {
		return x.PricePrecisions
	}
	return nil
}

func (x *PairParams) GetQtyPrecision() int32 {
	if x != nil {
		return x.QtyPrecision
	}
	return 0
}

type Trades struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair   string   `protobuf:"bytes,1,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Trades []*Trade `protobuf:"bytes,2,rep,name=Trades,proto3" json:"Trades,omitempty"`
}

func (x *Trades) Reset() {
	*x = Trades{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trades) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trades) ProtoMessage() {}

func (x *Trades) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trades.ProtoReflect.Descriptor instead.
func (*Trades) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{2}
}

func (x *Trades) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Trades) GetTrades() []*Trade {
	if x != nil {
		return x.Trades
	}
	return nil
}

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsBid bool                   `protobuf:"varint,1,opt,name=IsBid,proto3" json:"IsBid,omitempty"`
	Price string                 `protobuf:"bytes,2,opt,name=Price,proto3" json:"Price,omitempty"`
	Qty   string                 `protobuf:"bytes,3,opt,name=Qty,proto3" json:"Qty,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{3}
}

func (x *Trade) GetIsBid() bool {
	if x != nil {
		return x.IsBid
	}
	return false
}

func (x *Trade) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Trade) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *Trade) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type PrecisedOrderBookSnapshots struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrecisedOrderBookSnapshot map[int32]*OrderBookSnapshot `protobuf:"bytes,1,rep,name=PrecisedOrderBookSnapshot,proto3" json:"PrecisedOrderBookSnapshot,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PrecisedOrderBookSnapshots) Reset() {
	*x = PrecisedOrderBookSnapshots{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrecisedOrderBookSnapshots) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrecisedOrderBookSnapshots) ProtoMessage() {}

func (x *PrecisedOrderBookSnapshots) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrecisedOrderBookSnapshots.ProtoReflect.Descriptor instead.
func (*PrecisedOrderBookSnapshots) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{4}
}

func (x *PrecisedOrderBookSnapshots) GetPrecisedOrderBookSnapshot() map[int32]*OrderBookSnapshot {
	if x != nil {
		return x.PrecisedOrderBookSnapshot
	}
	return nil
}

type OrderBookSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair    string   `protobuf:"bytes,1,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Bids    []*Limit `protobuf:"bytes,2,rep,name=Bids,proto3" json:"Bids,omitempty"`
	Asks    []*Limit `protobuf:"bytes,3,rep,name=Asks,proto3" json:"Asks,omitempty"`
	BidsQty string   `protobuf:"bytes,4,opt,name=BidsQty,proto3" json:"BidsQty,omitempty"`
	AsksQty string   `protobuf:"bytes,5,opt,name=AsksQty,proto3" json:"AsksQty,omitempty"`
}

func (x *OrderBookSnapshot) Reset() {
	*x = OrderBookSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBookSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBookSnapshot) ProtoMessage() {}

func (x *OrderBookSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBookSnapshot.ProtoReflect.Descriptor instead.
func (*OrderBookSnapshot) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{5}
}

func (x *OrderBookSnapshot) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *OrderBookSnapshot) GetBids() []*Limit {
	if x != nil {
		return x.Bids
	}
	return nil
}

func (x *OrderBookSnapshot) GetAsks() []*Limit {
	if x != nil {
		return x.Asks
	}
	return nil
}

func (x *OrderBookSnapshot) GetBidsQty() string {
	if x != nil {
		return x.BidsQty
	}
	return ""
}

func (x *OrderBookSnapshot) GetAsksQty() string {
	if x != nil {
		return x.AsksQty
	}
	return ""
}

type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price string `protobuf:"bytes,1,opt,name=Price,proto3" json:"Price,omitempty"`
	Qty   string `protobuf:"bytes,2,opt,name=Qty,proto3" json:"Qty,omitempty"`
}

func (x *Limit) Reset() {
	*x = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit) ProtoMessage() {}

func (x *Limit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{6}
}

func (x *Limit) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Limit) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

var File_proto_quote_proto protoreflect.FileDescriptor

var file_proto_quote_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x62, 0x51, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x69, 0x72, 0x22, 0x6e, 0x0a, 0x0a, 0x50, 0x61, 0x69, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x51, 0x74, 0x79, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x51, 0x74, 0x79, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x22, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x51, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x06, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x22, 0x75, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x73, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73,
	0x42, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x51, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x51, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x1a,
	0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x7c, 0x0a, 0x19, 0x50, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x70, 0x62, 0x51, 0x2e, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x2e, 0x50,
	0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x19, 0x50,
	0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x1a, 0x64, 0x0a, 0x1e, 0x50, 0x72, 0x65, 0x63,
	0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x51, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b,
	0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x42, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x51, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x04, 0x42, 0x69, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x41, 0x73, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x51, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x04, 0x41, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x69, 0x64, 0x73,
	0x51, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x69, 0x64, 0x73, 0x51,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x73, 0x6b, 0x73, 0x51, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x73, 0x6b, 0x73, 0x51, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x51,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x51, 0x74, 0x79, 0x32, 0x84, 0x02,
	0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12,
	0x09, 0x2e, 0x70, 0x62, 0x51, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x51,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x1f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x09,
	0x2e, 0x70, 0x62, 0x51, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x51, 0x2e,
	0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x51, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x09, 0x2e, 0x70, 0x62, 0x51, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x62, 0x51, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quote_proto_rawDescOnce sync.Once
	file_proto_quote_proto_rawDescData = file_proto_quote_proto_rawDesc
)

func file_proto_quote_proto_rawDescGZIP() []byte {
	file_proto_quote_proto_rawDescOnce.Do(func() {
		file_proto_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quote_proto_rawDescData)
	})
	return file_proto_quote_proto_rawDescData
}

var file_proto_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_quote_proto_goTypes = []any{
	(*Pair)(nil),                       // 0: pbQ.Pair
	(*PairParams)(nil),                 // 1: pbQ.PairParams
	(*Trades)(nil),                     // 2: pbQ.Trades
	(*Trade)(nil),                      // 3: pbQ.Trade
	(*PrecisedOrderBookSnapshots)(nil), // 4: pbQ.PrecisedOrderBookSnapshots
	(*OrderBookSnapshot)(nil),          // 5: pbQ.OrderBookSnapshot
	(*Limit)(nil),                      // 6: pbQ.Limit
	nil,                                // 7: pbQ.PrecisedOrderBookSnapshots.PrecisedOrderBookSnapshotEntry
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_proto_quote_proto_depIdxs = []int32{
	3,  // 0: pbQ.Trades.Trades:type_name -> pbQ.Trade
	8,  // 1: pbQ.Trade.Time:type_name -> google.protobuf.Timestamp
	7,  // 2: pbQ.PrecisedOrderBookSnapshots.PrecisedOrderBookSnapshot:type_name -> pbQ.PrecisedOrderBookSnapshots.PrecisedOrderBookSnapshotEntry
	6,  // 3: pbQ.OrderBookSnapshot.Bids:type_name -> pbQ.Limit
	6,  // 4: pbQ.OrderBookSnapshot.Asks:type_name -> pbQ.Limit
	5,  // 5: pbQ.PrecisedOrderBookSnapshots.PrecisedOrderBookSnapshotEntry.value:type_name -> pbQ.OrderBookSnapshot
	0,  // 6: pbQ.Quote.StreamPrecisedTrades:input_type -> pbQ.Pair
	0,  // 7: pbQ.Quote.StreamPrecisedOrderBookSnapshot:input_type -> pbQ.Pair
	1,  // 8: pbQ.Quote.CreateOrderBook:input_type -> pbQ.PairParams
	0,  // 9: pbQ.Quote.DeleteOrderBook:input_type -> pbQ.Pair
	2,  // 10: pbQ.Quote.StreamPrecisedTrades:output_type -> pbQ.Trades
	4,  // 11: pbQ.Quote.StreamPrecisedOrderBookSnapshot:output_type -> pbQ.PrecisedOrderBookSnapshots
	9,  // 12: pbQ.Quote.CreateOrderBook:output_type -> google.protobuf.Empty
	9,  // 13: pbQ.Quote.DeleteOrderBook:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_quote_proto_init() }
func file_proto_quote_proto_init() {
	if File_proto_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quote_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Pair); i {
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
		file_proto_quote_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PairParams); i {
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
		file_proto_quote_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Trades); i {
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
		file_proto_quote_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Trade); i {
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
		file_proto_quote_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PrecisedOrderBookSnapshots); i {
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
		file_proto_quote_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OrderBookSnapshot); i {
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
		file_proto_quote_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Limit); i {
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
			RawDescriptor: file_proto_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quote_proto_goTypes,
		DependencyIndexes: file_proto_quote_proto_depIdxs,
		MessageInfos:      file_proto_quote_proto_msgTypes,
	}.Build()
	File_proto_quote_proto = out.File
	file_proto_quote_proto_rawDesc = nil
	file_proto_quote_proto_goTypes = nil
	file_proto_quote_proto_depIdxs = nil
}
