package postgres

import (
	"encoding/json"
	pb "Golang/genproto"

	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(user *pb.User) (*pb.User, error) {
	query := `INSERT INTO users(
		id,
		assignee,
		title,
		summary,
		deadline,
		status,
		created_at,
		updated_at,
		deleted_at) 
        VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
        RETURNING id,first_name,last_name,email,location,phone`
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Location, res).Scan(
		&
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Location,
		&res,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *TaskRepo) ListUsers(list *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	var byt []byte
	offset := (list.Page - 1) * list.Limit
	query := `SELECT
        id,
        first_name,
        last_name,
        email,
        location,
        phone
        FROM users OFFSET $1 LIMIT $2`
	rows, err := r.db.Query(query, offset, list.Limit)
	if err != nil {
		return nil, err
	}
	allUser := pb.ListUserResponse{}
	err = r.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&allUser.All)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user = pb.User{}
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Location,
			&byt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(byt, &user.Phone)
		if err != nil {
			return nil, err
		}
		allUser.User = append(allUser.User, &user)
	}

	return &allUser, nil
}

func (r *TaskRepo) GetUser(user *pb.User) (*pb.User, error) {
	var byt []byte
	query := `SELECT first_name,last_name,email,location,phone FROM users WHERE id=$1`
	err := r.db.QueryRow(query, user.Id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Location,
		&byt,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byt, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *TaskRepo) DeleteUser(user *pb.User) (*pb.Xabar, error) {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, user.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Xabar{Message: "Ok!"}, nil
}

func (r *TaskRepo) UpdateUser(user *pb.User) (*pb.Xabar, error) {
	query := `UPDATE users SET first_name=$1, last_name=$2, email=$3, location=$4, phone=$5 WHERE id=$6`
	byt, err := json.Marshal(user.Phone)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Location, byt, user.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Xabar{Message: "Ok!"}, nil
}

func (r *TaskRepo) Search(user *pb.SearchUser) (*pb.User, error) {
	userr := pb.User{}
	user.Text += "%"
	if user.Search == "first_name" {
		var byt []byte
		query := `SELECT id,first_name,last_name,email,location,phone FROM users WHERE first_name  LIKE $1`
		err := r.db.QueryRow(query, user.Text).Scan(
			&userr.Id,
			&userr.FirstName,
			&userr.LastName,
			&userr.Email,
			&userr.Location,
			&byt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(byt, &userr.Phone)
		if err != nil {
			return nil, err
		}
	} else if user.Search == "last_name" {
		var byt []byte
		query := `SELECT id,first_name,last_name,email,location,phone FROM users WHERE last_name  LIKE $1`
		err := r.db.QueryRow(query, user.Text).Scan(
			&userr.Id,
			&userr.FirstName,
			&userr.LastName,
			&userr.Email,
			&userr.Location,
			&byt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(byt, &userr.Phone)
		if err != nil {
			return nil, err
		}
	} else if user.Search == "email" {
		var byt []byte
		query := `SELECT id,first_name,last_name,email,location,phone FROM users WHERE email  LIKE $1`
		err := r.db.QueryRow(query, user.Text).Scan(
			&userr.Id,
			&userr.FirstName,
			&userr.LastName,
			&userr.Email,
			&userr.Location,
			&byt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(byt, &userr.Phone)
		if err != nil {
			return nil, err
		}
	} else if user.Search == "location" {
		var byt []byte
		query := `SELECT id,first_name,last_name,email,location,phone FROM users WHERE location  LIKE $1`
		err := r.db.QueryRow(query, user.Text).Scan(
			&userr.Id,
			&userr.FirstName,
			&userr.LastName,
			&userr.Email,
			&userr.Location,
			&byt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(byt, &userr.Phone)
		if err != nil {
			return nil, err
		}
	}
	return &userr, nil
}
