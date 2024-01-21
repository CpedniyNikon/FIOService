package models

type PersonData struct {
	ID          uint    `gorm:"primarykey" json:"-"`
	FirstName   string  `json:"FirstName"`
	LastName    string  `json:"LastName"`
	Age         float64 `json:"Age"`
	Gender      string  `json:"Gender"`
	Nationality string  `json:"Nationality"`
}
