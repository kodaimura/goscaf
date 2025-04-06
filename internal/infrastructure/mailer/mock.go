package mailer

import (
	"fmt"
	"strings"
	"mime"
	"goscaf/config"
	"goscaf/internal/core"
)

type MockMailer struct {
	from     string
}

func NewMockMailer() core.MailerI {
	return &MockMailer{
		from: config.MailFrom,
	}
}

func (m *MockMailer) SendText(to []string, subject, body string) error {
	msg := m.composeMessage(to, subject, "text/plain", body)
	fmt.Println("MockMailer (Text):")
	fmt.Println(string(msg))
	return nil
}

func (m *MockMailer) SendHTML(to []string, subject, body string) error {
	msg := m.composeMessage(to, subject, "text/html", body)
	fmt.Println("MockMailer (HTML):")
	fmt.Println(string(msg))
	return nil
}

func (m *MockMailer) composeMessage(to []string, subject, contentType, body string) []byte {
	header := m.defaultHeader(to, subject)
	header += fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", contentType)
	return []byte(header + "\r\n" + body)
}

func (m *MockMailer) defaultHeader(to []string, subject string) string {
	return "From: " + m.from + "\r\n" +
		"To: " + strings.Join(to, ", ") + "\r\n" +
		"Subject: " + mime.QEncoding.Encode("UTF-8", subject) + "\r\n" +
		"MIME-Version: 1.0\r\n"
}
