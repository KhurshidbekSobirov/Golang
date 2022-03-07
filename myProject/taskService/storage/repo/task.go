package repo

import (
	pb "myProject/taskService/genproto/task_service"
	pbe "myProject/taskService/genproto/email_service"
)

//UserStorageI ...
type TaskStorageI interface {
	Create(*pb.TaskRes) (*pb.TaskReq, error)
	GetTask(*pb.ById) (*pb.TaskReq, error)
	UpdateTask(*pb.TaskReq) (*pb.Mess, error)
	DeleteTask(*pb.ById) (*pb.Mess, error)
	ListOverdue(*pb.Mess) (*pb.ListTasks, error)
	ListTasks(*pb.ByUserId) (*pb.ListTasks, error)

}

type EmailStorageI interface{
	Send(*pbe.Email)(*pbe.Empty,error)
}

