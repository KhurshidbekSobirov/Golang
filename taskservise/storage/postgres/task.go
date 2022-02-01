package postgres

import (
	pb "app/genproto"
	"time"

	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(user *pb.Task) (*pb.Task, error) {
	task := pb.Task{}
	query := `INSERT INTO users(
		id,
		assignee,
		title,
		deadline,
		status,
		created_at,
		updated_at,
		deleted_at) 
        VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
        RETURNING id,assignee,title,deadline,status,created_at,updated_at,deleted_at`
	err := r.db.QueryRow(query, task.Id, task.Assignee, task.Title, task.Deadline, task.Status, task.CreatedAt, task.UpdatedAt, task.DeletedAt).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *taskRepo) GetTask(l *pb.Task) (*pb.Task, error) {

	task := pb.Task{}
	query := `SELECT id,assignee,title,deadline,status,created_at,updated_at,deleted_at FROM users WHERE id=$1`
	err := r.db.QueryRow(query, task.Id).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) Update(l *pb.Task) (*pb.Task,error){
	task := pb.Task{}

	query := `UPDATE tasks SET  assignee = $1,title = $2,deadline = $3,status = $4,created_at = $5,updated_at = $6,deleted_at = $7`

	_, err := r.db.Exec(query, task.Id, task.Assignee, task.Title, task.Deadline, task.Status, task.CreatedAt, task.UpdatedAt, task.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &pb.Task{}, nil
}

func (r *taskRepo) Delete(task *pb.Task) (*pb.Mess, error) {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, task.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Mess{Message: "Ok!"}, nil
}

func (r *taskRepo) ListOverdue(req *pb.Mess) (*pb.ListTask, error) {
	query := `SELECT id,assignee,title,deadline,status,created_at,updated_at,deleted_at WHERE deadline>$1`
	now_time := time.Now().Format(time.RFC3339)

	rows, err := r.db.Query(query,now_time)
	if err != nil {
		return nil,err
	}
	task := pb.Task{}
	tasks := pb.ListTask{}
	for rows.Next(){
		err := rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil{
			return nil,err
		}

		tasks.Task = append(tasks.Task,&task)
	}

	return &tasks,err
}

