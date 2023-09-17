package models

import "time"

type Customer struct {
	CstId         int       `json:"cst_id"`
	NationalityId int       `json:"nationality_id"`
	CstName       string    `json:"cst_name"`
	CstDob        time.Time `gorm:"default:current_timestamp" json:"created"`
	CstPhoneNum   string    `json:"cst_phone_num"`
	CstEmail      string    `json:"cst_email"`
}

type Familylist struct {
	CstId      int    `json:"cst_id"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dobe"`
}

type Nationality struct {
	NationalityId   int    `json:"nationality_id"`
	NationalityCode string `json:"nationality_code"`
	NationalityName string `json:"nationality_name"`
}
