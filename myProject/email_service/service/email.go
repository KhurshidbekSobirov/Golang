package service

import (
	"app/config"
	pb "app/genproto/email"
	"app/storage"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gomail "gopkg.in/gomail.v2"
)

// SendService ...
type SendService struct {
	storage storage.I
	conf    config.Config
}

// NewSendService ...
func NewSendService(db *sqlx.DB, cfg config.Config) *SendService {
	return &SendService{storage: storage.NewStoragePg(db), conf: cfg}
}

//Send ...
func (s *SendService) Send(ctx context.Context, req *pb.Email) (*pb.Empty, error) {

	statuss := true
	fmt.Println(req)
	err := s.sendEmail(req.Subject, req.Body, req.Recipient)
	log.Print(err, "gkhlj;khjgvh")
	if err != nil {
		statuss = false
		err := s.storage.SendEmail().Send(req.Subject, req.Body, statuss, req.Recipient)
		if err != nil {
			return &pb.Empty{}, status.Error(codes.Internal, "Internal server error")
		}
	} else {
		statuss = true
		err := s.storage.SendEmail().Send(req.Subject, req.Body, statuss, req.Recipient)
		if err != nil {
			return &pb.Empty{}, status.Error(codes.Internal, "Internal server error")
		}

	}

	return &pb.Empty{}, nil
}

func (s *SendService) sendEmail(subject string, body string, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.conf.EmailFromHeader)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Send the email to
	d := gomail.NewPlainDialer(s.conf.SMTPHost, s.conf.SMTPPort, s.conf.SMTPUser, s.conf.SMTPUserPass)

	if err := d.DialAndSend(m); err != nil {
		log.Print(err)
		panic(err)
	}
	log.Print("Sent")
	return nil
}

func (s *SendService) SendSms(ctx context.Context, req *pb.Sms) (*pb.Empty, error) {
// 	err := s.storage.SendS().SendS(req.To, req.Text)
// 	if err != nil {
// 		return nil, err
// 	}

// 	smsBody, err := json.Marshal(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := http.Post(s.conf.Smshost, "application/json", bytes.NewBuffer(smsBody))
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	respBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	 fmt.Println(string(respBody))
 	return &pb.Empty{}, nil
 }
