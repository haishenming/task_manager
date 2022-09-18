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

// CreateHospital creates a hospital
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

// RegisterEmployee registers an employee
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

// CreateTask creates a task
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

// UpdateTask updates a task
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

// AssignTask assigns a task to an employee
func (s *TaskService) AssignTask(ctx context.Context, req *pb.AssignTaskRequest) (*pb.AssignTaskReply, error) {
	err := s.uc.AssignTask(ctx, int(req.TaskId), int(req.EmployeeId))
	if err != nil {
		return nil, err
	}
	return &pb.AssignTaskReply{}, nil
}

// GetEmployeeTasks returns all tasks of an employee
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

// GetHospitalTasks returns all tasks of a hospital
func (s *TaskService) GetHospitalTasks(ctx context.Context, req *pb.GetHospitalTasksRequest) (*pb.GetHospitalTasksReply, error) {
	tasks, count, err := s.uc.ListByHospitalID(ctx, int(req.HospitalId), int(req.Limit), int(req.Offset))
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

// GetHospitals returns all hospitals
func (s *TaskService) GetHospitals(ctx context.Context, request *pb.GetHospitalsRequest) (*pb.GetHospitalsReply, error) {
	hospitals, count, err := s.uc.GetHospitals(ctx, int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	var hospitalDetails []*pb.HospitalDetail
	for _, hospital := range hospitals {
		hospitalDetails = append(hospitalDetails, &pb.HospitalDetail{
			Id:        uint32(hospital.ID),
			Name:      hospital.Name,
			Address:   hospital.Address,
			CreatedAt: hospital.CreatedAt.Unix(),
			UpdatedAt: hospital.UpdatedAt.Unix(),
		})
	}

	return &pb.GetHospitalsReply{
		Hospitals: hospitalDetails,
		Count:     uint32(count),
	}, nil
}

// GetEmployees returns all employees
func (s *TaskService) GetEmployees(ctx context.Context, req *pb.GetEmployeesRequest) (*pb.GetEmployeesReply, error) {
	employees, count, err := s.uc.GetEmployeesByHospital(ctx, int(req.HospitalId), int(req.Limit), int(req.Offset))

	if err != nil {
		return nil, err
	}

	var employeeDetails []*pb.EmployeeDetail
	for _, employee := range employees {
		employeeDetails = append(employeeDetails, &pb.EmployeeDetail{
			Id:         uint32(employee.ID),
			Name:       employee.Name,
			HospitalId: uint32(employee.HospitalID),
			CreatedAt:  employee.CreatedAt.Unix(),
			UpdatedAt:  employee.UpdatedAt.Unix(),
		})
	}

	return &pb.GetEmployeesReply{
		Employees: employeeDetails,
		Count:     uint32(count),
	}, nil
}
