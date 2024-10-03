package mysql

import (
	"time"

	"gorm.io/gorm"
)

type TblForms struct {
	Id         int       `gorm:"primaryKey;autoIncrement;type:int"`
	FormTitle  string    `gorm:"type:varchar(255)"`
	FormSlug   string    `gorm:"type:varchar(255)"`
	FormData   string    `gorm:"type:varchar(255)"`
	Status     int       `gorm:"type:int"`
	IsActive   int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:int;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:int;DEFAULT:0"`
	TenantId   int       `gorm:"type:int"`
}

type TblFormRegistrations struct {
	Id        int       `gorm:"primaryKey;autoIncrement;type:int"`
	Name      string    `gorm:"type:varchar(255)"`
	EmailId   string    `gorm:"type:varchar(255)"`
	MobileNo  string    `gorm:"type:varchar(20)"`
	FormId    int       `gorm:"type:int;DEFAULT:NULL"`
	CreatedOn time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
}

func MigrationTables(db *gorm.DB) {

	if err := db.AutoMigrate(

		&TblForms{},
		&TblFormRegistrations{},
	); err != nil {

		panic(err)
	}

}
