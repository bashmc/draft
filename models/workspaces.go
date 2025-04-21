package models

import "time"

type Workspace struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	OwnedBy string `json:"ownedBy"`
	CreatedAt time.Time `json:"createdAt"`
}