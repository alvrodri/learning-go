package models

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`

	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
