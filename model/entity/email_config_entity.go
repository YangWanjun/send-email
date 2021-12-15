package entity

import (
	"gopkg.in/guregu/null.v3"
)

type EmailConfigEntity struct {
	Id           int64        `json:"id"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	SmtpServer   string       `json:"smtp_server"`
	SmtpPort     int          `json:"smtp_port"`
	Sender       null.String  `json:"sender"`
	DisplayName  null.String  `json:"display_name"`
	IsDefault    bool         `json:"is_default"`
}

func (c *EmailConfigEntity)GetSender() string {
	if c.Sender.Valid {
		return c.Sender.String
	} else {
		return c.Username
	}
}

func (c *EmailConfigEntity)GetDisplayName() string {
	if c.DisplayName.Valid {
		return c.DisplayName.String
	} else {
		return ""
	}
}