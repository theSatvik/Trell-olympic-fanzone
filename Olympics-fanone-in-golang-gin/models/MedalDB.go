package models;


type Medal struct {
	Id        int64  `db:"ID" json:"id"`
	Country  string `db:"COUNTRY" json:"COUNTRY"`
	GoldCount  int64 `db:"GOLD" json:"GOLD"`
	SilverCount int64 `db:"SILVER" json:"SILVER"`
	BronzeCount  int64 `db:"BRONZE" json:"BRONZE"`
}