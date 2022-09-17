// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: task/v1/model.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 任务优先级
type TaskPriority int32

const (
	// 未知
	TaskPriority_TASK_PRIORITY_UNKNOWN TaskPriority = 0
	// 低优先级
	TaskPriority_TASK_PRIORITY_LOW TaskPriority = 1
	// 高优先级
	TaskPriority_TASK_PRIORITY_HiGH TaskPriority = 2
	// 紧急
	TaskPriority_TASK_PRIORITY_URGENT TaskPriority = 3
)

// Enum value maps for TaskPriority.
var (
	TaskPriority_name = map[int32]string{
		0: "TASK_PRIORITY_UNKNOWN",
		1: "TASK_PRIORITY_LOW",
		2: "TASK_PRIORITY_HiGH",
		3: "TASK_PRIORITY_URGENT",
	}
	TaskPriority_value = map[string]int32{
		"TASK_PRIORITY_UNKNOWN": 0,
		"TASK_PRIORITY_LOW":     1,
		"TASK_PRIORITY_HiGH":    2,
		"TASK_PRIORITY_URGENT":  3,
	}
)

func (x TaskPriority) Enum() *TaskPriority {
	p := new(TaskPriority)
	*p = x
	return p
}

func (x TaskPriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskPriority) Descriptor() protoreflect.EnumDescriptor {
	return file_task_v1_model_proto_enumTypes[0].Descriptor()
}

func (TaskPriority) Type() protoreflect.EnumType {
	return &file_task_v1_model_proto_enumTypes[0]
}

func (x TaskPriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskPriority.Descriptor instead.
func (TaskPriority) EnumDescriptor() ([]byte, []int) {
	return file_task_v1_model_proto_rawDescGZIP(), []int{0}
}

// 任务状态
type TaskStatus int32

const (
	// 未知
	TaskStatus_TASK_STATUS_UNKNOWN TaskStatus = 0
	// open 需要员工采取行动
	TaskStatus_TASK_STATUS_OPEN TaskStatus = 1
	// failed 任务无法完成
	TaskStatus_TASK_STATUS_FAILED TaskStatus = 2
	// completed 任务成功
	TaskStatus_TASK_STATUS_COMPLETED TaskStatus = 3
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "TASK_STATUS_UNKNOWN",
		1: "TASK_STATUS_OPEN",
		2: "TASK_STATUS_FAILED",
		3: "TASK_STATUS_COMPLETED",
	}
	TaskStatus_value = map[string]int32{
		"TASK_STATUS_UNKNOWN":   0,
		"TASK_STATUS_OPEN":      1,
		"TASK_STATUS_FAILED":    2,
		"TASK_STATUS_COMPLETED": 3,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_task_v1_model_proto_enumTypes[1].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_task_v1_model_proto_enumTypes[1]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_task_v1_model_proto_rawDescGZIP(), []int{1}
}

// 任务
type TaskDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务ID
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 任务标题
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 任务描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// 任务优先级
	Priority TaskPriority `protobuf:"varint,4,opt,name=priority,proto3,enum=task.v1.TaskPriority" json:"priority,omitempty"`
	// 任务状态
	Status TaskStatus `protobuf:"varint,5,opt,name=status,proto3,enum=task.v1.TaskStatus" json:"status,omitempty"`
	// 所有者员工ID
	OwnerId uint32 `protobuf:"varint,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// 所属医院ID
	HospitalId uint32 `protobuf:"varint,7,opt,name=hospital_id,json=hospitalId,proto3" json:"hospital_id,omitempty"`
	// 任务创建时间
	CreatedAt int64 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 任务更新时间
	UpdatedAt int64 `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TaskDetail) Reset() {
	*x = TaskDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetail) ProtoMessage() {}

func (x *TaskDetail) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetail.ProtoReflect.Descriptor instead.
func (*TaskDetail) Descriptor() ([]byte, []int) {
	return file_task_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *TaskDetail) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskDetail) GetPriority() TaskPriority {
	if x != nil {
		return x.Priority
	}
	return TaskPriority_TASK_PRIORITY_UNKNOWN
}

func (x *TaskDetail) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TASK_STATUS_UNKNOWN
}

func (x *TaskDetail) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *TaskDetail) GetHospitalId() uint32 {
	if x != nil {
		return x.HospitalId
	}
	return 0
}

func (x *TaskDetail) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TaskDetail) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_task_v1_model_proto protoreflect.FileDescriptor

var file_task_v1_model_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0xae,
	0x02, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a,
	0x72, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x19, 0x0a, 0x15, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x48, 0x69, 0x47, 0x48, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x52, 0x47, 0x45, 0x4e,
	0x54, 0x10, 0x03, 0x2a, 0x6e, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x03, 0x42, 0x44, 0x0a, 0x16, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x54,
	0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_task_v1_model_proto_rawDescOnce sync.Once
	file_task_v1_model_proto_rawDescData = file_task_v1_model_proto_rawDesc
)

func file_task_v1_model_proto_rawDescGZIP() []byte {
	file_task_v1_model_proto_rawDescOnce.Do(func() {
		file_task_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_v1_model_proto_rawDescData)
	})
	return file_task_v1_model_proto_rawDescData
}

var file_task_v1_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_task_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_task_v1_model_proto_goTypes = []interface{}{
	(TaskPriority)(0),  // 0: task.v1.TaskPriority
	(TaskStatus)(0),    // 1: task.v1.TaskStatus
	(*TaskDetail)(nil), // 2: task.v1.TaskDetail
}
var file_task_v1_model_proto_depIdxs = []int32{
	0, // 0: task.v1.TaskDetail.priority:type_name -> task.v1.TaskPriority
	1, // 1: task.v1.TaskDetail.status:type_name -> task.v1.TaskStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_task_v1_model_proto_init() }
func file_task_v1_model_proto_init() {
	if File_task_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDetail); i {
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
			RawDescriptor: file_task_v1_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_v1_model_proto_goTypes,
		DependencyIndexes: file_task_v1_model_proto_depIdxs,
		EnumInfos:         file_task_v1_model_proto_enumTypes,
		MessageInfos:      file_task_v1_model_proto_msgTypes,
	}.Build()
	File_task_v1_model_proto = out.File
	file_task_v1_model_proto_rawDesc = nil
	file_task_v1_model_proto_goTypes = nil
	file_task_v1_model_proto_depIdxs = nil
}