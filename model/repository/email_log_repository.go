package repository

import (
	"database/sql"
	"fmt"
	"send-email/model/entity"
)

func GetEmailLogsByServer(serverName string) ([]entity.EmailLogEntity, error) {
	emailLogs := []entity.EmailLogEntity{}

	rows, err := Db.Query("SELECT * FROM t_email_log WHERE server_name = ?", serverName)

	if err != nil {
		return nil, fmt.Errorf("GetEmailLogsByServer %q: %v", serverName, err)
	}
	defer rows.Close()

	for rows.Next() {
		var emailLog entity.EmailLogEntity
		if err := rows.Scan(
			&emailLog.Id,
			&emailLog.ActionTime,
			&emailLog.Username,
			&emailLog.Sender,
			&emailLog.Recipient,
			&emailLog.Cc,
			&emailLog.Bcc,
			&emailLog.Title,
			&emailLog.Body,
			&emailLog.PasswordBody,
			&emailLog.Attachments,
			&emailLog.ServerName); err != nil {
			return nil, fmt.Errorf("GetEmailLogsByServer %q: %v", serverName, err)
		}
		emailLogs = append(emailLogs, emailLog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetEmailLogsByServer %q: %v", serverName, err)
	}
	return emailLogs, nil
}

func addEmailLog(emailLog entity.EmailLogEntity) (int64, error) {
	result, err := Db.Exec("INSERT INTO t_email_log (action_time, username, sender, title, recipient, cc, bcc, body, password_body, attachments) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		emailLog.ActionTime,
		emailLog.Username,
		emailLog.Sender,
		emailLog.Title,
		emailLog.Recipient,
		emailLog.Cc,
		emailLog.Bcc,
		emailLog.Body,
		emailLog.PasswordBody,
		emailLog.Attachments)
	if err != nil {
		return 0, fmt.Errorf("addEmailLog: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addEmailLog: %v", err)
	}
	return id, nil
}

func GetEmailLogByID(id int64, serverName string) (entity.EmailLogEntity, error) {
	var emailLog entity.EmailLogEntity

	row := Db.QueryRow("SELECT * FROM t_email_log WHERE id = ? and server_name = ?", id, serverName)
	if err := row.Scan(
		&emailLog.Id,
		&emailLog.ActionTime,
		&emailLog.Username,
		&emailLog.Sender,
		&emailLog.Recipient,
		&emailLog.Cc,
		&emailLog.Bcc,
		&emailLog.Title,
		&emailLog.Body,
		&emailLog.PasswordBody,
		&emailLog.Attachments,
		&emailLog.ServerName); err != nil {
		if err == sql.ErrNoRows {
			return emailLog, sql.ErrNoRows
		}
		return emailLog, fmt.Errorf("GetEmailLogByID %d, %s: %v", id, serverName, err)
	}
	return emailLog, nil
}