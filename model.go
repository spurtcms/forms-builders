package formbuilders

import (
	"fmt"
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

// FormList
func (Formsmodel FormModel) FormsList(DB *gorm.DB, tenantid int) ([]TblForms, error) {

	var Forms []TblForms

	if err := DB.Debug().Table("tbl_forms").Select("tbl_forms.*,tbl_users.username,tbl_users.profile_image_path").Joins("INNER JOIN tbl_users ON tbl_forms.created_by=tbl_users.id").Where("is_deleted=0 and (tbl_forms.tenant_id is NULL or tbl_forms.tenant_id = ?)", tenantid).Find(&Forms).Error; err != nil {
		return nil, err
	}

	fmt.Println("Forms:", Forms)

	return Forms, nil
}
