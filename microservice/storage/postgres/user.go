package postgres

import (
    "github.com/jmoiron/sqlx"
    pb "github.com/KhurshidbekSobirov/Golang/microservice/genproto"
)

type userRepo struct {
    db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
    return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
    query := `INSERT INTO users(first_name, last_name, username, email)
    VALUES($1,$2,$3,$4) RETURNING id,first_name, last_name,username, email`
  newUser := pb.User{}
  err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Username, user.Email).Scan(&newUser.Id, &newUser.FirstName, &newUser.LastName, &newUser.Username, &newUser.Email)
  if err != nil {
    return &newUser, err
  }
  return &newUser, nil
}
