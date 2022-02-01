package repo

import (
	pb "app/genproto"
)

//UserStorageI ...
type TaskStorageI interface {
	Create(*pb.Task) (*pb.Task, error)
	GetTask(*pb.Task) (*pb.Task, error)
	Update(*pb.Task) (*pb.Task, error)
	Delete(*pb.Task) (*pb.Mess, error)
	ListOverdue(*pb.Mess) (*pb.ListTask, error)
}
