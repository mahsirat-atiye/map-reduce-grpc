// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: worker_driver.proto

package worker_driver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskSubmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DonTaskType string `protobuf:"bytes,1,opt,name=donTaskType,proto3" json:"donTaskType,omitempty"`
	DoneTaskIdx int32  `protobuf:"varint,2,opt,name=doneTaskIdx,proto3" json:"doneTaskIdx,omitempty"`
}

func (x *TaskSubmissionRequest) Reset() {
	*x = TaskSubmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSubmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSubmissionRequest) ProtoMessage() {}

func (x *TaskSubmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSubmissionRequest.ProtoReflect.Descriptor instead.
func (*TaskSubmissionRequest) Descriptor() ([]byte, []int) {
	return file_worker_driver_proto_rawDescGZIP(), []int{0}
}

func (x *TaskSubmissionRequest) GetDonTaskType() string {
	if x != nil {
		return x.DonTaskType
	}
	return ""
}

func (x *TaskSubmissionRequest) GetDoneTaskIdx() int32 {
	if x != nil {
		return x.DoneTaskIdx
	}
	return 0
}

type TaskAdmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TaskAdmissionResponse) Reset() {
	*x = TaskAdmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAdmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAdmissionResponse) ProtoMessage() {}

func (x *TaskAdmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAdmissionResponse.ProtoReflect.Descriptor instead.
func (*TaskAdmissionResponse) Descriptor() ([]byte, []int) {
	return file_worker_driver_proto_rawDescGZIP(), []int{1}
}

func (x *TaskAdmissionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TaskAssigningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *TaskAssigningRequest) Reset() {
	*x = TaskAssigningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAssigningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAssigningRequest) ProtoMessage() {}

func (x *TaskAssigningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAssigningRequest.ProtoReflect.Descriptor instead.
func (*TaskAssigningRequest) Descriptor() ([]byte, []int) {
	return file_worker_driver_proto_rawDescGZIP(), []int{2}
}

func (x *TaskAssigningRequest) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

type TaskAssigningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                  string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	TaskType                string   `protobuf:"bytes,2,opt,name=taskType,proto3" json:"taskType,omitempty"`
	TaskIdx                 int32    `protobuf:"varint,3,opt,name=taskIdx,proto3" json:"taskIdx,omitempty"`
	NumOfComplementaryTasks int32    `protobuf:"varint,4,opt,name=numOfComplementaryTasks,proto3" json:"numOfComplementaryTasks,omitempty"`
	TargetInputFiles        []string `protobuf:"bytes,5,rep,name=targetInputFiles,proto3" json:"targetInputFiles,omitempty"`
}

func (x *TaskAssigningResponse) Reset() {
	*x = TaskAssigningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAssigningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAssigningResponse) ProtoMessage() {}

func (x *TaskAssigningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAssigningResponse.ProtoReflect.Descriptor instead.
func (*TaskAssigningResponse) Descriptor() ([]byte, []int) {
	return file_worker_driver_proto_rawDescGZIP(), []int{3}
}

func (x *TaskAssigningResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskAssigningResponse) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *TaskAssigningResponse) GetTaskIdx() int32 {
	if x != nil {
		return x.TaskIdx
	}
	return 0
}

func (x *TaskAssigningResponse) GetNumOfComplementaryTasks() int32 {
	if x != nil {
		return x.NumOfComplementaryTasks
	}
	return 0
}

func (x *TaskAssigningResponse) GetTargetInputFiles() []string {
	if x != nil {
		return x.TargetInputFiles
	}
	return nil
}

var File_worker_driver_proto protoreflect.FileDescriptor

var file_worker_driver_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x15, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x6e, 0x54, 0x61,
	0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6f, 0x6e, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x6f, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x78, 0x22, 0x2f, 0x0a, 0x15, 0x54, 0x61, 0x73, 0x6b,
	0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x54, 0x61, 0x73,
	0x6b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x22, 0xcb, 0x01, 0x0a, 0x15, 0x54, 0x61, 0x73, 0x6b,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x78, 0x12,
	0x38, 0x0a, 0x17, 0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x17, 0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x32, 0xbe, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x53, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x15, 0x41, 0x64, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x79,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x62, 0x3b, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_worker_driver_proto_rawDescOnce sync.Once
	file_worker_driver_proto_rawDescData = file_worker_driver_proto_rawDesc
)

func file_worker_driver_proto_rawDescGZIP() []byte {
	file_worker_driver_proto_rawDescOnce.Do(func() {
		file_worker_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_driver_proto_rawDescData)
	})
	return file_worker_driver_proto_rawDescData
}

var file_worker_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_worker_driver_proto_goTypes = []interface{}{
	(*TaskSubmissionRequest)(nil), // 0: driver.TaskSubmissionRequest
	(*TaskAdmissionResponse)(nil), // 1: driver.TaskAdmissionResponse
	(*TaskAssigningRequest)(nil),  // 2: driver.TaskAssigningRequest
	(*TaskAssigningResponse)(nil), // 3: driver.TaskAssigningResponse
}
var file_worker_driver_proto_depIdxs = []int32{
	2, // 0: driver.ManageTaskPool.AssignTaskToWorker:input_type -> driver.TaskAssigningRequest
	0, // 1: driver.ManageTaskPool.AdmitTaskDoneByWorker:input_type -> driver.TaskSubmissionRequest
	3, // 2: driver.ManageTaskPool.AssignTaskToWorker:output_type -> driver.TaskAssigningResponse
	1, // 3: driver.ManageTaskPool.AdmitTaskDoneByWorker:output_type -> driver.TaskAdmissionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_worker_driver_proto_init() }
func file_worker_driver_proto_init() {
	if File_worker_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSubmissionRequest); i {
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
		file_worker_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAdmissionResponse); i {
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
		file_worker_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAssigningRequest); i {
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
		file_worker_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAssigningResponse); i {
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
			RawDescriptor: file_worker_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_driver_proto_goTypes,
		DependencyIndexes: file_worker_driver_proto_depIdxs,
		MessageInfos:      file_worker_driver_proto_msgTypes,
	}.Build()
	File_worker_driver_proto = out.File
	file_worker_driver_proto_rawDesc = nil
	file_worker_driver_proto_goTypes = nil
	file_worker_driver_proto_depIdxs = nil
}
