package repo

import (
	pb "app/genproto"
)

//UserStorageI ...
type TaskStorageI interface {
	Create(*pb.Task) (*pb.Task, error)
	GetTask(*pb.Task) (*pb.Task, error)
	UpdateTask(*pb.Task) (*pb.Task, error)
	DeleteTask(*pb.Task) (*pb.Mess, error)
	ListOverdue(*pb.Mess) (*pb.ListTask, error)
}
