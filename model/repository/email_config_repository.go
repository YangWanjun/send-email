package repository

import (
	"database/sql"
	"fmt"
	"send-email/model/entity"
)

func GetDefaultEmailConfig() (entity.EmailConfigEntity, error) {
	var emailConfig entity.EmailConfigEntity

	row := Db.QueryRow("SELECT * FROM t_email_config WHERE is_default = 1")
	if err := row.Scan(
		&emailConfig.Id,
		&emailConfig.Username,
		&emailConfig.Password,
		&emailConfig.SmtpServer,
		&emailConfig.SmtpPort,
		&emailConfig.Sender,
		&emailConfig.DisplayName,
		&emailConfig.IsDefault); err != nil {
		if err == sql.ErrNoRows {
			return emailConfig, sql.ErrNoRows
		}
		return emailConfig, fmt.Errorf("GetDefaultEmailConfig: %v", err)
	}
	return emailConfig, nil
}

func GetEmailConfigByUsername(username string) (entity.EmailConfigEntity, error) {
	var emailConfig entity.EmailConfigEntity

	row := Db.QueryRow("SELECT * FROM t_email_config WHERE username = ?", username)
	if err := row.Scan(
		&emailConfig.Id,
		&emailConfig.Username,
		&emailConfig.Password,
		&emailConfig.SmtpServer,
		&emailConfig.SmtpPort,
		&emailConfig.Sender,
		&emailConfig.DisplayName,
		&emailConfig.IsDefault); err != nil {
		if err == sql.ErrNoRows {
			return emailConfig, sql.ErrNoRows
		}
		return emailConfig, fmt.Errorf("GetEmailConfigByUsername %s: %v", username, err)
	}
	return emailConfig, nil
}