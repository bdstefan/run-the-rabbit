package model

import "time"

type Consumer struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"type:varchar(100);" valid:"length(3|50)" json:"name"`
	Host      Host
	Vhost     string     `gorm:"type:varchar(100);" valid:"length(3|50)" json:"vhost"`
	Queue     string     `gorm:"type:varchar(100);" valid:"length(3|50)" json:"queue"`
	Status    bool       `gorm:"type:tinyint(1);" json:"status"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
}