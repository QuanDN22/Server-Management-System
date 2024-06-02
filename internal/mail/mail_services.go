package mail

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/QuanDN22/Server-Management-System/proto/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	gomail "gopkg.in/mail.v2"
)

var (
	email_from     = "nghiemphong12345@gmail.com"
	email_password = "zwbh yweu wuse roog"
	// m              = make(map[string]bool)
	// emails_to      = []string{}
)

func (m *MailService) SendMail(ctx context.Context, in *mail.SendMailRequest) (*emptypb.Empty, error) {
	fmt.Println("Send Mail begin")
	// parse data
	type data_server struct {
		Sum_Server     int64   `json:"sum_server"`;
		Sum_Server_On  int64   `json:"sum_server_on"`;
		Sum_Server_Off int64   `json:"sum_server_off"`;
		Uptime         float32 `json:"uptime"`;
	}

	type data struct {
		Email      []string      `json:"email"`;
		DataServer []data_server `json:"data_send"`;
	}

	var data_recv data

	err := json.Unmarshal(in.GetDataSendMail(), &data_recv)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	fmt.Println(data_recv)

	// // MAIL_SENDER_TEMPLATE
	t, err := template.ParseFiles("./internal/mail/template.html")
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	var body bytes.Buffer
	fmt.Println(data_recv.DataServer[0])
	t.Execute(&body, data_recv.DataServer[0])

	// send mail
	sm := gomail.NewMessage()

	// Set E-Mail sender
	sm.SetHeader("From", email_from)

	// Set E-Mail receivers
	sm.SetHeader("To", data_recv.Email...)

	// Set E-Mail subject
	sm.SetHeader("Subject", "Report for Management System Server")

	// Set E-Mail body. You can set plain text or html with text/html
	sm.SetBody("text/html", body.String())

	// m.Attach("./data_export_example.xlsx")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, email_from, email_password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(sm); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &emptypb.Empty{}, nil
}
