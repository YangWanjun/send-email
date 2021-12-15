package entity

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

type EmailLogEntity struct {
	Id           int64        `json:"id"`
	ActionTime   time.Time    `json:"action_time"`
	Username     null.String  `json:"username"`
	Sender       string       `json:"sender"`
	Recipient    string       `json:"recipient"`
	Cc           null.String  `json:"cc"`
	Bcc          null.String  `json:"bcc"`
	Title        string       `json:"title"`
	Body         string       `json:"body"`
	PasswordBody null.String  `json:"password_body"`
	Attachments  null.String  `json:"attachments"`
	ServerName   null.String  `json:"server_name"`
}
