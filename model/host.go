package model

type Host struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"type:varchar(100);" valid:"length(3|50)" json:"name"`
}
