package models

import (
	"time"
)

// type Countries struct {
// 	gorm.Model
// 	ID           uint       `gorm:"column:id" json:"country_id"`
// 	Name         string     `gorm:"column:name" json:"name"`
// 	Country_code string     `gorm:"column:country_code" json:"country_code"`
// 	Is_highrisk  uint       `gorm:"column:is_highrisk" json:"is_highrisk"`
// 	CreatedAt    time.Time  `gorm:"column:created_at" json:"-"`
// 	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"-"`
// 	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"-"`
// }

type Countries struct {
	ID           uint       ` json:"country_id"`
	Name         string     ` json:"name"`
	Country_code string     ` json:"country_code"`
	Is_highrisk  uint       ` json:"is_highrisk"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time ` json:"-"`
}

// type Provinces struct {
// 	ID         uint         `gorm:"column:id" json:"province_id"`
// 	Country_ID []Country_ID `gorm:"Foreignkey:country_id;association_foreignkey:ID" json:"country_id"`
// 	Name       string       `gorm:"column:name" json:"name"`
// 	CreatedAt  time.Time    `gorm:"column:created_at" json:"-"`
// 	UpdatedAt  time.Time    `gorm:"column:updated_at" json:"-"`
// 	DeletedAt  *time.Time   `gorm:"column:deleted_at" json:"-"`
// }

// type Cities struct {
// 	ID           uint       `gorm:"column:id" json:"cities_id"`
// 	Provinces_ID uint       `gorm:"Foreignkey:province_id;association_foreignkey:ID" json:"province_id"`
// 	Type         string     `gorm:"column:type" json:"type"`
// 	Name         string     `gorm:"column:name" json:"name"`
// 	CreatedAt    time.Time  `gorm:"column:created_at" json:"-"`
// 	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"-"`
// 	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"-"`
// }

// type Districts struct {
// 	ID        uint       `gorm:"column:id" json:"districts_id"`
// 	Cities_ID uint       `gorm:"Foreignkey:cities_id;association_foreignkey:ID" json:"city_id"`
// 	Name      string     `gorm:"column:name" json:"name"`
// 	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
// 	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
// 	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
// }

// type SubDistrict struct {
// 	ID           uint       `gorm:"column:id" json:"subdistrict_id"`
// 	Districts_ID uint       `gorm:"Foreignkey:district_id;association_foreignkey:ID" json:"district_id"`
// 	Name         string     `gorm:"column:name" json:"name"`
// 	CreatedAt    time.Time  `gorm:"column:created_at" json:"-"`
// 	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"-"`
// 	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"-"`
// 	Zipcode      uint       `gorm:"column:zipcode" json:"zipcode"`
// }
