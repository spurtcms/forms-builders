package postgres

import (
	"time"

	"gorm.io/gorm"
)

type TblForms struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormTitle  string    `gorm:"type:character varying"`
	FormSlug   string    `gorm:"type:character varying"`
	FormData   string    `gorm:"type:character varying"`
	Status     int       `gorm:"type:integer"`
	IsActive   int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
	TenantId   int       `gorm:"type:integer"`
}

type TblFormRegistrations struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name      string    `gorm:"type:character varying"`
	EmailId   string    `gorm:"type:character varying"`
	MobileNo  string    `gorm:"type:character varying"`
	FormId    int       `gorm:"type:integer;DEFAULT:NULL"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

func MigrationTables(db *gorm.DB) {

	if err := db.AutoMigrate(

		&TblForms{},
		&TblFormRegistrations{},
	); err != nil {

		panic(err)
	}
	
}