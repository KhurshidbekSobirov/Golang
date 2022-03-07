package postgres


import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"app/storage/repo"
)

type sendRepo struct {
	db *sqlx.DB
}

// NewSendRepo ...
func NewSendRepo(db *sqlx.DB) repo.SendStorageI {
	return &sendRepo{db: db}
}

func (cm *sendRepo) MakeSent(ID string) error {
	var err error
	makesent := `UPDATE email_send_email SET send_status=true where id = $1`
	cm.db.MustExec(makesent, ID)
	return err
}

func (cm *sendRepo) Send(subject, body string, status bool,val string) error {
	textID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	insert := `
	INSERT INTO
	email_text
	(
		id,
		subject,
		body,
		status

	)
	values($1, $2, $3,$4)
	`
	_, err = cm.db.Exec(insert, textID, subject, body,status)
	if err != nil {
		return err
	}

	
	return nil
}

func (cm *sendRepo) SendS(To, Text string) error {
	textID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	insert := `
	INSERT INTO
	sms_text
	(
		id,
		to_us,
		text_us
	)
	values($1, $2, $3)
	`
	_, err = cm.db.Exec(insert, textID,To, Text)
	if err != nil {
		return err
	}

	
	return nil
}
