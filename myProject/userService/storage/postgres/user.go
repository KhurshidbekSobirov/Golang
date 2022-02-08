package postgres

import (
	"encoding/json"
	pb "myProject/userService/genproto/user_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.UserRes) (*pb.UserReq, error) {
	new_user := pb.UserReq{}
	adress, err := json.Marshal(user.Adress)
	if err != nil {
		return nil, err
	}
	numbers, err := json.Marshal(user.PhoneNumbers)
	
	if err != nil {
		return nil, err
	}
	var a = []byte{}
	var b = []byte{}
	query := `INSERT INTO users(
		id,
		first_name,
		last_name,
		username,
		profile_photo,
		bio,
		email,
		gender,
		adress,
		phone_numbers,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id,
		first_name,
		last_name,
		username,
		profile_photo,
		bio,
		email,
		gender,
		adress,
		phone_numbers,
		created_at`

	uid := uuid.New().String()
	time := time.Now().Format(time.RFC3339)
	err = r.db.QueryRow(query,
		uid,
		user.FirstName,
		user.LastName,
		user.Username,
		user.ProfilePhoto,
		user.Bio,
		user.Email,
		user.Gender,
		adress,
		numbers,
		time,
	).Scan(
		&new_user.Id,
		&new_user.FirstName,
		&new_user.LastName,
		&new_user.Username,
		&new_user.ProfilePhoto,
		&new_user.Bio,
		&new_user.Email,
		&new_user.Gender,
		&a,
		&b,
		&new_user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	

	err = json.Unmarshal(a, &new_user.Adress)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &new_user.PhoneNumbers)
	if err != nil {
		return nil, err
	}
	

	return &new_user, nil

}

func (r *userRepo) GetUser(user *pb.ById) (*pb.UserReq, error) {
	new_user := pb.UserReq{}
	var a = []byte{}
	var b = []byte{}
	query := `SELECT
		id, 
		first_name,
		last_name,
		username,
		profile_photo,
		bio,
		email,
		gender,
		adress,
		phone_numbers,
		created_at
		FROM users
		WHERE id=$1 AND
		deleted_at IS NULL`

	err := r.db.QueryRow(query, user.UserId).Scan(
		&new_user.Id,
		&new_user.FirstName,
		&new_user.LastName,
		&new_user.Username,
		&new_user.ProfilePhoto,
		&new_user.Bio,
		&new_user.Email,
		&new_user.Gender,
		&a,
		&b,
		&new_user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(a, &new_user.Adress)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &new_user.PhoneNumbers)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &new_user, nil
}

func (r *userRepo) Update(user *pb.UserReq) (*pb.Mess, error) {
	time := time.Now().Format(time.RFC3339)
	adress, err := json.Marshal(user.Adress)
	if err != nil {
		return nil, err
	}
	numbers, err := json.Marshal(user.PhoneNumbers)

	if err != nil {
		return nil, err
	}
	query := `UPDATE users SET
		first_name = $1,
		last_name = $2,
		username = $3,
		profile_photo= $4,
		bio = $5,
		email = $6,
		gender = $7,
		adress = $8,
		phone_numbers = $9,
		updated_at = $10
		WHERE id=$11 AND
		deleted_at IS NULL
		`
	_, err = r.db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.ProfilePhoto,
		user.Bio,
		user.Email,
		user.Gender,
		adress,
		numbers,
		time,
	)
	if err != nil {
		return nil, err
	}

	return &pb.Mess{Message: "OK"}, nil
}

func (r *userRepo) Delete(user *pb.ById) (*pb.Mess, error) {
	query := `UPDATE users SET deleted_at=$1 WHERE id=$2`
	time := time.Now().Format(time.RFC3339)
	_, err := r.db.Exec(query, time, user.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Mess{Message: "OK"}, err
}

func (r *userRepo) ListOverdue(req *pb.Mess) (*pb.ListUser, error) {
	var a = []byte{}
	var b = []byte{}
	query := ` SELECT 
		id, 
		first_name,
		last_name,
		username,
		profile_photo,
		bio,
		email,
		gender,
		adress,
		phone_numbers,
		created_at
		FROM users WHERE
		deleted_at IS NULL
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	new_user := pb.UserReq{}
	users := pb.ListUser{}
	for rows.Next() {
		err := rows.Scan(
			&new_user.Id,
			&new_user.FirstName,
			&new_user.LastName,
			&new_user.Username,
			&new_user.ProfilePhoto,
			&new_user.Bio,
			&new_user.Email,
			&new_user.Gender,
			&a,
			&b,
			&new_user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(a, &new_user.Adress)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &new_user.PhoneNumbers)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		users.User = append(users.User, &new_user)
	}
	return &users, err
}
