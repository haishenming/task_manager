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
	hospitalModel := biz.Hospital{
		Name:    req.Name,
		Address: req.Address,
	}

	hospital, err := s.uc.CreateHospital(ctx, &hospitalModel)
	if err != nil {
		return nil, err
	}

	return &pb.CreateHospitalReply{
		Id: uint32(hospital.ID),
	}, nil
}
func (s *TaskService) RegisterEmployee(ctx context.Context, req *pb.RegisterEmployeeRequest) (*pb.RegisterEmployeeReply, error) {
	employeeModel := biz.Employee{
		Name:       req.Name,
		HospitalID: int(req.HospitalId),
	}

	employee, err := s.uc.CreateEmployee(ctx, &employeeModel)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterEmployeeReply{
		Id: uint32(employee.ID),
	}, nil
}
func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskReply, error) {
	taskModel := biz.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		HospitalID:  int(req.HospitalId),
	}

	task, err := s.uc.CreateTask(ctx, &taskModel)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskReply{
		Id: uint32(task.ID),
	}, nil
}
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskReply, error) {
	taskModel := biz.Task{
		ID:          int(req.Id),
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status,
	}

	task, err := s.uc.UpdateTask(ctx, &taskModel)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskReply{
		Task: &pb.TaskDetail{
			Id:          uint32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			Status:      task.Status,
			EmployeeId:  uint32(task.EmployeeID),
			HospitalId:  uint32(task.HospitalID),
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		},
	}, nil
}
func (s *TaskService) AssignTask(ctx context.Context, req *pb.AssignTaskRequest) (*pb.AssignTaskReply, error) {
	err := s.uc.AssignTask(ctx, int(req.TaskId), int(req.EmployeeId))
	if err != nil {
		return nil, err
	}
	return &pb.AssignTaskReply{}, nil
}
func (s *TaskService) GetEmployeeTasks(ctx context.Context, req *pb.GetEmployeeTasksRequest) (*pb.GetEmployeeTasksReply, error) {
	tasks, count, err := s.uc.ListByEmployeeID(ctx, int(req.EmployeeId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var taskDetails []*pb.TaskDetail
	for _, task := range tasks {
		taskDetails = append(taskDetails, &pb.TaskDetail{
			Id:          uint32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			Status:      task.Status,
			EmployeeId:  uint32(task.EmployeeID),
			HospitalId:  uint32(task.HospitalID),
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		})
	}

	return &pb.GetEmployeeTasksReply{
		Tasks: taskDetails,
		Count: uint32(count),
	}, nil
}
func (s *TaskService) GetHospitalTasks(ctx context.Context, req *pb.GetHospitalTasksRequest) (*pb.GetHospitalTasksReply, error) {
	tasks, count, err := s.uc.ListByEmployeeID(ctx, int(req.HospitalId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var taskDetails []*pb.TaskDetail
	for _, task := range tasks {
		taskDetails = append(taskDetails, &pb.TaskDetail{
			Id:          uint32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			Status:      task.Status,
			EmployeeId:  uint32(task.EmployeeID),
			HospitalId:  uint32(task.HospitalID),
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		})
	}

	return &pb.GetHospitalTasksReply{
		Tasks: taskDetails,
		Count: uint32(count),
	}, nil
}
