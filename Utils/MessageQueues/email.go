package MessageQueues

import (
	"TigerPopulation/Domain/Tigers"
	"crypto/tls"

	"github.com/golang/glog"
	gomail "gopkg.in/mail.v2"
)

var fromEmail string = "pankajofcbanshiwal@gmail.com"
var host string = "smtp.gmail.com"

func SendMail(req Tigers.EmailStruct) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", fromEmail)

	// Set E-Mail receivers
	m.SetHeader("To", req.Email)

	// Set E-Mail subject
	m.SetHeader("Subject", req.TigerName+" was spotted recently")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Hey "+req.UserName+"!! \n\n"+req.TigerName+" (tiger which you sighted) was spotted recently in 'Some location'\nThanks & regards")
	//m.SetBody("text/plain", "Hey "+req.UserName+"!! "+req.TigerName+"(tiger which you sighted) was spotted recently in "+req.Location)

	// Settings for SMTP server
	d := gomail.NewDialer(host, 587, fromEmail, "yubc hbfr aidd mktz")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		glog.Errorln("Error sending email = Err", err)
		return err
	}
	glog.V(2).Infoln("Email sent successfully : Info = ", req)
	return nil
}
