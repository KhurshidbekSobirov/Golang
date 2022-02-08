package postgres

import (
	"log"
	pb "myProject/taskService/genproto/task_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task *pb.TaskRes) (*pb.TaskReq, error) {
	newtask := pb.TaskReq{}
	query := `INSERT INTO tasks(
		id, 
		name,
		deadline,
		summary,
		assigneid,
		status,
		create_at)
        VALUES($1,$2,$3,$4,$5,$6,$7)
        RETURNING id,name,deadline,summary,assigneid,status,create_at`
	uid := uuid.New().String()
	time := time.Now().Format(time.RFC3339)
	err := r.db.QueryRow(query, uid, task.Name, task.Deadline, task.Summary, task.AssigneeId, task.Status, time).Scan(
		&newtask.Id,
		&newtask.Name,
		&newtask.Deadline,
		&newtask.Summary,
		&newtask.AssigneId,
		&newtask.Status,
		&newtask.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &newtask, nil
}

func (r *taskRepo) GetTask(l *pb.ById) (*pb.TaskReq, error) {

	nwtask := pb.TaskReq{}
	query := `SELECT id,name,deadline,summary,assigneid,status FROM tasks WHERE id=$1 AND delete_at IS NULL`
	err := r.db.QueryRow(query, l.TaskId).Scan(
		&nwtask.Id,
		&nwtask.Name,
		&nwtask.Deadline,
		&nwtask.Summary,
		&nwtask.AssigneId,
		&nwtask.Status,
	)
	if err != nil {
		return nil, err
	}

	return &nwtask, nil
}

func (r *taskRepo) UpdateTask(l *pb.TaskReq) (*pb.Mess, error) {

	time := time.Now().Format(time.RFC3339)

	query := `UPDATE tasks SET 
	name = $1,
	deadline = $3,
	summary = $4,
	assigneid = $5,
	status = $6,
	update_at = $7
	WHERE id=$8 AND
	delete_at IS NULL`
	_, err := r.db.Exec(query, l.Name, l.Deadline, l.Summary, l.AssigneId, l.Status, time, l.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Mess{Message: "OK"}, nil
}

func (r *taskRepo) DeleteTask(task *pb.ById) (*pb.Mess, error) {
	query := `UPDATE tasks SET delete_at=$1 WHERE id=$2`
	time := time.Now().Format(time.RFC3339)
	_, err := r.db.Exec(query, time, task.TaskId)
	if err != nil {
		return nil, err
	}
	return &pb.Mess{Message: "Ok!"}, nil
}

func (r *taskRepo) ListOverdue(req *pb.Mess) (*pb.ListTasks, error) {
	query := `SELECT id,name,assigneid,deadline,summary,status,create_at FROM tasks WHERE deadline<$1  AND delete_at IS NULL`
	now_time := time.Now().Format(time.RFC3339)

	rows, err := r.db.Query(query, now_time)
	if err != nil {
		return nil, err
	}
	task := pb.TaskReq{}
	tasks := pb.ListTasks{}
	for rows.Next() {
		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.AssigneId,
			&task.Deadline,
			&task.Summary,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks.Tasks = append(tasks.Tasks, &task)
	}

	return &tasks, err
}

func (r *taskRepo) ListTasks(user *pb.ByUserId) (*pb.ListTasks, error) {

	tasks := pb.ListTasks{}
	
	query := `SELECT 
	 		id,
			name,
			deadline,
			summary,
			assigneid,
			status,
			create_at
			FROM tasks 
			where assigneid = $1`

	rows, err := r.db.Query(query, user.UserId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for rows.Next() {
		task := pb.TaskReq{}
		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Deadline,
			&task.Summary,
			&task.AssigneId,
			&task.Status,
			&task.CreatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		tasks.Tasks = append(tasks.Tasks, &task)
	}

	return &tasks, nil

}
