package models;

type Schedule struct {
	Id        int64  `db:"ID" json:"id"`
	Events  string `db:"EVENTS" json:"EVENTS"`
	StartDate  string `db:"STARTDATE" json:"STARTDATE"`
	EndDate string `db:"ENDDATE" json:"ENDDATE"`
	Country  string `db:"COUNTRY" json:"COUNTRY"`
}
