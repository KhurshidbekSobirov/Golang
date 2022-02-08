package storage

import (
    "github.com/jmoiron/sqlx"
    "myProject/userService/storage/postgres"
    "myProject/userService/storage/repo"
)

//IStorage ...
type IStorage interface {
    User() repo.UserStorageI
}

type storagePg struct {
    db         *sqlx.DB
    taskRepo   repo.UserStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
    return &storagePg{
        db:         db,
        taskRepo:   postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
    return s.taskRepo
}
