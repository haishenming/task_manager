package service

import (
	"context"

	pb "task_manager/api/task/v1"
	"task_manager/internal/biz"
)

type TaskService struct {
	pb.UnimplementedTaskServer

	uc *biz.TaskUsecase
}

func NewTaskService(uc *biz.TaskUsecase) *TaskService {
	return &TaskService{
		uc: uc,
	}
}

func (s *TaskService) CreateHospital(ctx context.Context, req *pb.CreateHospitalRequest) (*pb.CreateHospitalReply, error) {
	return &pb.CreateHospitalReply{}, nil
}
func (s *TaskService) RegisterEmployee(ctx context.Context, req *pb.RegisterEmployeeRequest) (*pb.RegisterEmployeeReply, error) {
	return &pb.RegisterEmployeeReply{}, nil
}
func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskReply, error) {
	return &pb.CreateTaskReply{}, nil
}
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskReply, error) {
	return &pb.UpdateTaskReply{}, nil
}
func (s *TaskService) AssignTask(ctx context.Context, req *pb.AssignTaskRequest) (*pb.AssignTaskReply, error) {
	return &pb.AssignTaskReply{}, nil
}
func (s *TaskService) GetEmployeeTasks(ctx context.Context, req *pb.GetEmployeeTasksRequest) (*pb.GetEmployeeTasksReply, error) {
	return &pb.GetEmployeeTasksReply{}, nil
}
func (s *TaskService) GetHospitalTasks(ctx context.Context, req *pb.GetHospitalTasksRequest) (*pb.GetHospitalTasksReply, error) {
	return &pb.GetHospitalTasksReply{}, nil
}
