package entity

import (
	"gopkg.in/guregu/null.v3"
)

type ApplicationEntity struct {
	Id           int64        `json:"id"`
	ClientId     string       `json:"client_id"`
	ClientSecret string       `json:"client_secret"`
	Domain       string       `json:"domain"`
	Name         null.String  `json:"name"`
}
