// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"log"
	"net"
	"net/smtp"
)

// SendMail 发送邮件
func SendMail(params *SendMailParams) (err error) {
	if len(params.ReceiverEmails) <= 0 {
		return errors.New("接收邮箱不能为空")
	}
	host, err := Instance().GetSmtpHost()
	if err != nil {
		return err
	}
	port, err := Instance().GetSmtpPort()
	if err != nil {
		return err
	}
	email, err := Instance().GetSmtpEmail()
	if err != nil {
		return err
	}
	password, err := Instance().GetSmtpEmailPassword()
	if err != nil {
		return err
	}
	glog.Line().Debug(host)
	glog.Line().Debug(port)
	glog.Line().Debug(email)
	glog.Line().Debug(password)
	nickname := ""
	emailArr := gstr.Explode("@", email)
	if len(emailArr) > 0 {
		nickname = emailArr[0]
	}
	header := make(map[string]string)
	header["From"] = nickname + "<" + email + ">"
	header["To"] = gstr.Join(params.ReceiverEmails, ",")
	header["Subject"] = params.MsgSubject
	if params.MsgContentType != "" {
		header["Content-Type"] = "text/html; charset=UTF-8"
	}

	body := params.MsgBody

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	glog.Line().Debug(message)

	smtpAuth := smtp.PlainAuth("", email, password, host)

	return SendMailUsingTLS(fmt.Sprintf("%s:%d", host, port), smtpAuth, email, params.ReceiverEmails, []byte(message))
}

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
