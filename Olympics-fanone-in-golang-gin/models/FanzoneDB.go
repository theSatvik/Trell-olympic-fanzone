package models;

type Fanzone struct {
	Country  string `db:"COUNTRY" json:"COUNTRY"`
	Cheers  int64 `db:"CHEERCOUNTER" json:"CHEERCOUNTER"`
}
