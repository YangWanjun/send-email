package repository

import (
	"fmt"
	"send-email/model/entity"
)

func GetApplications() ([]entity.ApplicationEntity, error) {
	var apps []entity.ApplicationEntity

	rows, err := Db.Query("SELECT * FROM t_application")

	if err != nil {
		return nil, fmt.Errorf("GetApplications: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var app entity.ApplicationEntity
		if err := rows.Scan(
			&app.Id,
			&app.ClientId,
			&app.ClientSecret,
			&app.Domain,
			&app.Name); err != nil {
			return nil, fmt.Errorf("GetApplications: %v", err)
		}
		apps = append(apps, app)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetApplications: %v", err)
	}
	return apps, nil
}

func CreateApplication(app entity.ApplicationEntity) (int64, error) {
	result, err := Db.Exec("INSERT INTO t_application (client_id, client_secret, domain, name) VALUES (?, ?, ?, ?)",
		app.ClientId,
		app.ClientSecret,
		app.Domain,
		app.Name)
	if err != nil {
		return 0, fmt.Errorf("addApplication: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addApplication: %v", err)
	}
	return id, nil
}