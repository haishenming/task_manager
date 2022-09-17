// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.6
// source: task/v1/task.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTaskAssignTask = "/task.v1.Task/AssignTask"
const OperationTaskCreateHospital = "/task.v1.Task/CreateHospital"
const OperationTaskCreateTask = "/task.v1.Task/CreateTask"
const OperationTaskGetEmployeeTasks = "/task.v1.Task/GetEmployeeTasks"
const OperationTaskGetHospitalTasks = "/task.v1.Task/GetHospitalTasks"
const OperationTaskRegisterEmployee = "/task.v1.Task/RegisterEmployee"
const OperationTaskUpdateTask = "/task.v1.Task/UpdateTask"

type TaskHTTPServer interface {
	AssignTask(context.Context, *AssignTaskRequest) (*AssignTaskReply, error)
	CreateHospital(context.Context, *CreateHospitalRequest) (*CreateHospitalReply, error)
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskReply, error)
	GetEmployeeTasks(context.Context, *GetEmployeeTasksRequest) (*GetEmployeeTasksReply, error)
	GetHospitalTasks(context.Context, *GetHospitalTasksRequest) (*GetHospitalTasksReply, error)
	RegisterEmployee(context.Context, *RegisterEmployeeRequest) (*RegisterEmployeeReply, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskReply, error)
}

func RegisterTaskHTTPServer(s *http.Server, srv TaskHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/task/hospital", _Task_CreateHospital0_HTTP_Handler(srv))
	r.POST("/v1/task/staff", _Task_RegisterEmployee0_HTTP_Handler(srv))
	r.POST("/v1/task", _Task_CreateTask0_HTTP_Handler(srv))
	r.PUT("/v1/task", _Task_UpdateTask0_HTTP_Handler(srv))
	r.PUT("/v1/task/assign", _Task_AssignTask0_HTTP_Handler(srv))
	r.GET("/v1/task/staff/{employee_id}", _Task_GetEmployeeTasks0_HTTP_Handler(srv))
	r.GET("/v1/task/hospital/{hospital_id}", _Task_GetHospitalTasks0_HTTP_Handler(srv))
}

func _Task_CreateHospital0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateHospitalRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateHospital)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateHospital(ctx, req.(*CreateHospitalRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateHospitalReply)
		return ctx.Result(200, reply)
	}
}

func _Task_RegisterEmployee0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterEmployeeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskRegisterEmployee)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterEmployee(ctx, req.(*RegisterEmployeeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterEmployeeReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CreateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTask(ctx, req.(*CreateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_AssignTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssignTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskAssignTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AssignTask(ctx, req.(*AssignTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssignTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetEmployeeTasks0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEmployeeTasksRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetEmployeeTasks)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEmployeeTasks(ctx, req.(*GetEmployeeTasksRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEmployeeTasksReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetHospitalTasks0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHospitalTasksRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetHospitalTasks)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHospitalTasks(ctx, req.(*GetHospitalTasksRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetHospitalTasksReply)
		return ctx.Result(200, reply)
	}
}

type TaskHTTPClient interface {
	AssignTask(ctx context.Context, req *AssignTaskRequest, opts ...http.CallOption) (rsp *AssignTaskReply, err error)
	CreateHospital(ctx context.Context, req *CreateHospitalRequest, opts ...http.CallOption) (rsp *CreateHospitalReply, err error)
	CreateTask(ctx context.Context, req *CreateTaskRequest, opts ...http.CallOption) (rsp *CreateTaskReply, err error)
	GetEmployeeTasks(ctx context.Context, req *GetEmployeeTasksRequest, opts ...http.CallOption) (rsp *GetEmployeeTasksReply, err error)
	GetHospitalTasks(ctx context.Context, req *GetHospitalTasksRequest, opts ...http.CallOption) (rsp *GetHospitalTasksReply, err error)
	RegisterEmployee(ctx context.Context, req *RegisterEmployeeRequest, opts ...http.CallOption) (rsp *RegisterEmployeeReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *UpdateTaskReply, err error)
}

type TaskHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskHTTPClient(client *http.Client) TaskHTTPClient {
	return &TaskHTTPClientImpl{client}
}

func (c *TaskHTTPClientImpl) AssignTask(ctx context.Context, in *AssignTaskRequest, opts ...http.CallOption) (*AssignTaskReply, error) {
	var out AssignTaskReply
	pattern := "/v1/task/assign"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskAssignTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) CreateHospital(ctx context.Context, in *CreateHospitalRequest, opts ...http.CallOption) (*CreateHospitalReply, error) {
	var out CreateHospitalReply
	pattern := "/v1/task/hospital"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateHospital))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...http.CallOption) (*CreateTaskReply, error) {
	var out CreateTaskReply
	pattern := "/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetEmployeeTasks(ctx context.Context, in *GetEmployeeTasksRequest, opts ...http.CallOption) (*GetEmployeeTasksReply, error) {
	var out GetEmployeeTasksReply
	pattern := "/v1/task/staff/{employee_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetEmployeeTasks))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetHospitalTasks(ctx context.Context, in *GetHospitalTasksRequest, opts ...http.CallOption) (*GetHospitalTasksReply, error) {
	var out GetHospitalTasksReply
	pattern := "/v1/task/hospital/{hospital_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetHospitalTasks))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) RegisterEmployee(ctx context.Context, in *RegisterEmployeeRequest, opts ...http.CallOption) (*RegisterEmployeeReply, error) {
	var out RegisterEmployeeReply
	pattern := "/v1/task/staff"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskRegisterEmployee))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*UpdateTaskReply, error) {
	var out UpdateTaskReply
	pattern := "/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
