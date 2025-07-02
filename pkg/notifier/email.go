package notifier

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To       string
	Proxy    string // 可选，标准库 net/smtp 不支持代理
}

func SendEmail(cfg EmailConfig, subject, body string) error {
	if cfg.Host == "" || cfg.Port == 0 || cfg.Username == "" || cfg.Password == "" || cfg.From == "" || cfg.To == "" {
		return errors.New("邮件配置不完整")
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	header := make(map[string]string)
	header["From"] = cfg.From
	header["To"] = cfg.To
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""

	var msg strings.Builder
	for k, v := range header {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n" + body)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         cfg.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, cfg.Host)
	if err != nil {
		return err
	}
	defer c.Quit()

	if err = c.Auth(smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)); err != nil {
		return err
	}
	if err = c.Mail(cfg.From); err != nil {
		return err
	}
	if err = c.Rcpt(cfg.To); err != nil {
		return err
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(msg.String()))
	if err != nil {
		return err
	}
	return w.Close()
}
