package formbuilders

import (
	"time"

	"gorm.io/gorm"
)

type TblForms struct {
	Id               int                    `gorm:"primaryKey;auto_increment;type:serial"`
	FormTitle        string                 `gorm:"type:character varying"`
	FormData         map[string]interface{} `gorm:"type:json"`
	Status           int                    `gorm:"type:integer"`
	IsActive         int                    `gorm:"type:integer"`
	CreatedBy        int                    `gorm:"type:integer"`
	CreatedOn        time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int                    `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn       time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int                    `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int                    `gorm:"type:integer;DEFAULT:0"`
	TenantId         int                    `gorm:"type:integer"`
	Username         string                 `gorm:"<-:false"`
	ProfileImagePath string                 `gorm:"<-:false"`
	DateString       string                 `gorm:"-"`
}

type TblFormRegistrations struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name      string    `gorm:"type:character varying"`
	EmailId   string    `gorm:"type:character varying"`
	MobileNo  string    `gorm:"type:character varying"`
	FormId    int       `gorm:"type:integer;DEFAULT:NULL"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type Filter struct {
	Keyword string
}

type FormModel struct {
	DataAccess int
	UserId     int
}

var Formsmodel FormModel

// FormList
func (Formsmodel FormModel) FormsList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int, status int) (Forms []TblForms, Count int64, err error) {

	query := DB.Debug().Table("tbl_forms").
		Select("tbl_forms.*, tbl_users.username, tbl_users.profile_image_path").
		Joins("inner join tbl_users on tbl_forms.created_by=tbl_users.id").
		Where("tbl_forms.is_deleted = 0 and tbl_forms.tenant_id = ? and tbl_forms.status = ?", tenantid, status).
		Order("tbl_forms.id desc")

	if filter.Keyword != "" {
		query = query.Where("Lower(TRIM(form_title)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if limit != 0 {
		query.Limit(limit).Offset(offset).Find(&Forms)

		return Forms, 0, err
	}

	query.Find(&Forms).Count(&Count)

	return Forms, Count, nil
}