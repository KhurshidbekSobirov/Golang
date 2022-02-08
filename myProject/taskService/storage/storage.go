package storage

import (
    "github.com/jmoiron/sqlx"
    "myProject/taskService/storage/postgres"
    "myProject/taskService/storage/repo"
)

//IStorage ...
type IStorage interface {
    Task() repo.TaskStorageI
}

type storagePg struct {
    db         *sqlx.DB
    taskRepo   repo.TaskStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
    return &storagePg{
        db:         db,
        taskRepo:   postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
    return s.taskRepo
}
