// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: waMultiDevice/WAMultiDevice.proto

package waMultiDevice

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MultiDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *MultiDevice_Payload  `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	Metadata *MultiDevice_Metadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (x *MultiDevice) Reset() {
	*x = MultiDevice{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice) ProtoMessage() {}

func (x *MultiDevice) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice.ProtoReflect.Descriptor instead.
func (*MultiDevice) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0}
}

func (x *MultiDevice) GetPayload() *MultiDevice_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MultiDevice) GetMetadata() *MultiDevice_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MultiDevice_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiDevice_Metadata) Reset() {
	*x = MultiDevice_Metadata{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_Metadata) ProtoMessage() {}

func (x *MultiDevice_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_Metadata.ProtoReflect.Descriptor instead.
func (*MultiDevice_Metadata) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 0}
}

type MultiDevice_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*MultiDevice_Payload_ApplicationData
	//	*MultiDevice_Payload_Signal
	Payload isMultiDevice_Payload_Payload `protobuf_oneof:"payload"`
}

func (x *MultiDevice_Payload) Reset() {
	*x = MultiDevice_Payload{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_Payload) ProtoMessage() {}

func (x *MultiDevice_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_Payload.ProtoReflect.Descriptor instead.
func (*MultiDevice_Payload) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 1}
}

func (m *MultiDevice_Payload) GetPayload() isMultiDevice_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *MultiDevice_Payload) GetApplicationData() *MultiDevice_ApplicationData {
	if x, ok := x.GetPayload().(*MultiDevice_Payload_ApplicationData); ok {
		return x.ApplicationData
	}
	return nil
}

func (x *MultiDevice_Payload) GetSignal() *MultiDevice_Signal {
	if x, ok := x.GetPayload().(*MultiDevice_Payload_Signal); ok {
		return x.Signal
	}
	return nil
}

type isMultiDevice_Payload_Payload interface {
	isMultiDevice_Payload_Payload()
}

type MultiDevice_Payload_ApplicationData struct {
	ApplicationData *MultiDevice_ApplicationData `protobuf:"bytes,1,opt,name=applicationData,oneof"`
}

type MultiDevice_Payload_Signal struct {
	Signal *MultiDevice_Signal `protobuf:"bytes,2,opt,name=signal,oneof"`
}

func (*MultiDevice_Payload_ApplicationData) isMultiDevice_Payload_Payload() {}

func (*MultiDevice_Payload_Signal) isMultiDevice_Payload_Payload() {}

type MultiDevice_ApplicationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ApplicationData:
	//
	//	*MultiDevice_ApplicationData_AppStateSyncKeyShare
	//	*MultiDevice_ApplicationData_AppStateSyncKeyRequest
	ApplicationData isMultiDevice_ApplicationData_ApplicationData `protobuf_oneof:"applicationData"`
}

func (x *MultiDevice_ApplicationData) Reset() {
	*x = MultiDevice_ApplicationData{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData) ProtoMessage() {}

func (x *MultiDevice_ApplicationData) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2}
}

func (m *MultiDevice_ApplicationData) GetApplicationData() isMultiDevice_ApplicationData_ApplicationData {
	if m != nil {
		return m.ApplicationData
	}
	return nil
}

func (x *MultiDevice_ApplicationData) GetAppStateSyncKeyShare() *MultiDevice_ApplicationData_AppStateSyncKeyShareMessage {
	if x, ok := x.GetApplicationData().(*MultiDevice_ApplicationData_AppStateSyncKeyShare); ok {
		return x.AppStateSyncKeyShare
	}
	return nil
}

func (x *MultiDevice_ApplicationData) GetAppStateSyncKeyRequest() *MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage {
	if x, ok := x.GetApplicationData().(*MultiDevice_ApplicationData_AppStateSyncKeyRequest); ok {
		return x.AppStateSyncKeyRequest
	}
	return nil
}

type isMultiDevice_ApplicationData_ApplicationData interface {
	isMultiDevice_ApplicationData_ApplicationData()
}

type MultiDevice_ApplicationData_AppStateSyncKeyShare struct {
	AppStateSyncKeyShare *MultiDevice_ApplicationData_AppStateSyncKeyShareMessage `protobuf:"bytes,1,opt,name=appStateSyncKeyShare,oneof"`
}

type MultiDevice_ApplicationData_AppStateSyncKeyRequest struct {
	AppStateSyncKeyRequest *MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage `protobuf:"bytes,2,opt,name=appStateSyncKeyRequest,oneof"`
}

func (*MultiDevice_ApplicationData_AppStateSyncKeyShare) isMultiDevice_ApplicationData_ApplicationData() {
}

func (*MultiDevice_ApplicationData_AppStateSyncKeyRequest) isMultiDevice_ApplicationData_ApplicationData() {
}

type MultiDevice_Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiDevice_Signal) Reset() {
	*x = MultiDevice_Signal{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_Signal) ProtoMessage() {}

func (x *MultiDevice_Signal) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_Signal.ProtoReflect.Descriptor instead.
func (*MultiDevice_Signal) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 3}
}

type MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyIDs []*MultiDevice_ApplicationData_AppStateSyncKeyId `protobuf:"bytes,1,rep,name=keyIDs" json:"keyIDs,omitempty"`
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage) Reset() {
	*x = MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage) ProtoMessage() {}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage) GetKeyIDs() []*MultiDevice_ApplicationData_AppStateSyncKeyId {
	if x != nil {
		return x.KeyIDs
	}
	return nil
}

type MultiDevice_ApplicationData_AppStateSyncKeyShareMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*MultiDevice_ApplicationData_AppStateSyncKey `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyShareMessage) Reset() {
	*x = MultiDevice_ApplicationData_AppStateSyncKeyShareMessage{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyShareMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData_AppStateSyncKeyShareMessage) ProtoMessage() {}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyShareMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData_AppStateSyncKeyShareMessage.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData_AppStateSyncKeyShareMessage) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyShareMessage) GetKeys() []*MultiDevice_ApplicationData_AppStateSyncKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type MultiDevice_ApplicationData_AppStateSyncKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyID   *MultiDevice_ApplicationData_AppStateSyncKeyId                   `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
	KeyData *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData `protobuf:"bytes,2,opt,name=keyData" json:"keyData,omitempty"`
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey) Reset() {
	*x = MultiDevice_ApplicationData_AppStateSyncKey{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData_AppStateSyncKey) ProtoMessage() {}

func (x *MultiDevice_ApplicationData_AppStateSyncKey) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData_AppStateSyncKey.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData_AppStateSyncKey) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2, 2}
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey) GetKeyID() *MultiDevice_ApplicationData_AppStateSyncKeyId {
	if x != nil {
		return x.KeyID
	}
	return nil
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey) GetKeyData() *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData {
	if x != nil {
		return x.KeyData
	}
	return nil
}

type MultiDevice_ApplicationData_AppStateSyncKeyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyID []byte `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyId) Reset() {
	*x = MultiDevice_ApplicationData_AppStateSyncKeyId{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData_AppStateSyncKeyId) ProtoMessage() {}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyId) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData_AppStateSyncKeyId.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData_AppStateSyncKeyId) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2, 3}
}

func (x *MultiDevice_ApplicationData_AppStateSyncKeyId) GetKeyID() []byte {
	if x != nil {
		return x.KeyID
	}
	return nil
}

type MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyData     []byte                                                                                      `protobuf:"bytes,1,opt,name=keyData" json:"keyData,omitempty"`
	Fingerprint *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
	Timestamp   *int64                                                                                      `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) Reset() {
	*x = MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) ProtoMessage() {}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2, 2, 0}
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) GetKeyData() []byte {
	if x != nil {
		return x.KeyData
	}
	return nil
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) GetFingerprint() *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

type MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawID         *uint32  `protobuf:"varint,1,opt,name=rawID" json:"rawID,omitempty"`
	CurrentIndex  *uint32  `protobuf:"varint,2,opt,name=currentIndex" json:"currentIndex,omitempty"`
	DeviceIndexes []uint32 `protobuf:"varint,3,rep,packed,name=deviceIndexes" json:"deviceIndexes,omitempty"`
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) Reset() {
	*x = MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint{}
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) ProtoMessage() {
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) ProtoReflect() protoreflect.Message {
	mi := &file_waMultiDevice_WAMultiDevice_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint.ProtoReflect.Descriptor instead.
func (*MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) Descriptor() ([]byte, []int) {
	return file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP(), []int{0, 2, 2, 0, 0}
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) GetRawID() uint32 {
	if x != nil && x.RawID != nil {
		return *x.RawID
	}
	return 0
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) GetCurrentIndex() uint32 {
	if x != nil && x.CurrentIndex != nil {
		return *x.CurrentIndex
	}
	return 0
}

func (x *MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint) GetDeviceIndexes() []uint32 {
	if x != nil {
		return x.DeviceIndexes
	}
	return nil
}

var File_waMultiDevice_WAMultiDevice_proto protoreflect.FileDescriptor

//go:embed WAMultiDevice.pb.raw
var file_waMultiDevice_WAMultiDevice_proto_rawDesc []byte

var (
	file_waMultiDevice_WAMultiDevice_proto_rawDescOnce sync.Once
	file_waMultiDevice_WAMultiDevice_proto_rawDescData = file_waMultiDevice_WAMultiDevice_proto_rawDesc
)

func file_waMultiDevice_WAMultiDevice_proto_rawDescGZIP() []byte {
	file_waMultiDevice_WAMultiDevice_proto_rawDescOnce.Do(func() {
		file_waMultiDevice_WAMultiDevice_proto_rawDescData = protoimpl.X.CompressGZIP(file_waMultiDevice_WAMultiDevice_proto_rawDescData)
	})
	return file_waMultiDevice_WAMultiDevice_proto_rawDescData
}

var file_waMultiDevice_WAMultiDevice_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_waMultiDevice_WAMultiDevice_proto_goTypes = []any{
	(*MultiDevice)(nil),                                                                                // 0: WAMultiDevice.MultiDevice
	(*MultiDevice_Metadata)(nil),                                                                       // 1: WAMultiDevice.MultiDevice.Metadata
	(*MultiDevice_Payload)(nil),                                                                        // 2: WAMultiDevice.MultiDevice.Payload
	(*MultiDevice_ApplicationData)(nil),                                                                // 3: WAMultiDevice.MultiDevice.ApplicationData
	(*MultiDevice_Signal)(nil),                                                                         // 4: WAMultiDevice.MultiDevice.Signal
	(*MultiDevice_ApplicationData_AppStateSyncKeyRequestMessage)(nil),                                  // 5: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyRequestMessage
	(*MultiDevice_ApplicationData_AppStateSyncKeyShareMessage)(nil),                                    // 6: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyShareMessage
	(*MultiDevice_ApplicationData_AppStateSyncKey)(nil),                                                // 7: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey
	(*MultiDevice_ApplicationData_AppStateSyncKeyId)(nil),                                              // 8: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyId
	(*MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData)(nil),                            // 9: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.AppStateSyncKeyData
	(*MultiDevice_ApplicationData_AppStateSyncKey_AppStateSyncKeyData_AppStateSyncKeyFingerprint)(nil), // 10: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.AppStateSyncKeyData.AppStateSyncKeyFingerprint
}
var file_waMultiDevice_WAMultiDevice_proto_depIdxs = []int32{
	2,  // 0: WAMultiDevice.MultiDevice.payload:type_name -> WAMultiDevice.MultiDevice.Payload
	1,  // 1: WAMultiDevice.MultiDevice.metadata:type_name -> WAMultiDevice.MultiDevice.Metadata
	3,  // 2: WAMultiDevice.MultiDevice.Payload.applicationData:type_name -> WAMultiDevice.MultiDevice.ApplicationData
	4,  // 3: WAMultiDevice.MultiDevice.Payload.signal:type_name -> WAMultiDevice.MultiDevice.Signal
	6,  // 4: WAMultiDevice.MultiDevice.ApplicationData.appStateSyncKeyShare:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyShareMessage
	5,  // 5: WAMultiDevice.MultiDevice.ApplicationData.appStateSyncKeyRequest:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyRequestMessage
	8,  // 6: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyRequestMessage.keyIDs:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyId
	7,  // 7: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyShareMessage.keys:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey
	8,  // 8: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.keyID:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKeyId
	9,  // 9: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.keyData:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.AppStateSyncKeyData
	10, // 10: WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.AppStateSyncKeyData.fingerprint:type_name -> WAMultiDevice.MultiDevice.ApplicationData.AppStateSyncKey.AppStateSyncKeyData.AppStateSyncKeyFingerprint
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_waMultiDevice_WAMultiDevice_proto_init() }
func file_waMultiDevice_WAMultiDevice_proto_init() {
	if File_waMultiDevice_WAMultiDevice_proto != nil {
		return
	}
	file_waMultiDevice_WAMultiDevice_proto_msgTypes[2].OneofWrappers = []any{
		(*MultiDevice_Payload_ApplicationData)(nil),
		(*MultiDevice_Payload_Signal)(nil),
	}
	file_waMultiDevice_WAMultiDevice_proto_msgTypes[3].OneofWrappers = []any{
		(*MultiDevice_ApplicationData_AppStateSyncKeyShare)(nil),
		(*MultiDevice_ApplicationData_AppStateSyncKeyRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waMultiDevice_WAMultiDevice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waMultiDevice_WAMultiDevice_proto_goTypes,
		DependencyIndexes: file_waMultiDevice_WAMultiDevice_proto_depIdxs,
		MessageInfos:      file_waMultiDevice_WAMultiDevice_proto_msgTypes,
	}.Build()
	File_waMultiDevice_WAMultiDevice_proto = out.File
	file_waMultiDevice_WAMultiDevice_proto_rawDesc = nil
	file_waMultiDevice_WAMultiDevice_proto_goTypes = nil
	file_waMultiDevice_WAMultiDevice_proto_depIdxs = nil
}
