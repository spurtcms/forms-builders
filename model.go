package formbuilders

import (
	"time"

	"gorm.io/gorm"
)

// Create Forms
type TblForm struct {
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
	Uuid       string    `gorm:"type:character varying"`
}

type TblForms struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormTitle        string    `gorm:"type:character varying"`
	FormSlug         string    `gorm:"type:character varying"`
	FormData         string    `gorm:"type:character varying"`
	Status           int       `gorm:"type:integer"`
	IsActive         int       `gorm:"type:integer"`
	CreatedBy        int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	TenantId         int       `gorm:"type:integer"`
	Uuid             string    `gorm:"type:character varying"`
	Username         string    `gorm:"<-:false"`
	ProfileImagePath string    `gorm:"<-:false"`
	NameString       string    `gorm:"<-:false"`
	FirstName        string    `gorm:"<-:false"`
	LastName         string    `gorm:"<-:false"`
	DateString       string    `gorm:"-"`
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
	Keyword  string
	FromDate string
	ToDate   string
}

type FormModel struct {
	DataAccess int
	UserId     int
}

var Formsmodel FormModel

// FormList
func (Formsmodel FormModel) FormsList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int, status int) (Forms []TblForms, Count int64, err error) {

	query := DB.Debug().Table("tbl_forms").
		Select("tbl_forms.*, tbl_users.username,tbl_users.first_name,tbl_users.last_name, tbl_users.profile_image_path").
		Joins("inner join tbl_users on tbl_forms.created_by=tbl_users.id").
		Where("tbl_forms.is_deleted = 0 and tbl_forms.tenant_id = ? and tbl_forms.status = ?", tenantid, status).
		Order("tbl_forms.modified_on desc")

	if filter.Keyword != "" {
		query = query.Where("Lower(TRIM(form_title)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if filter.FromDate != "" {
		query = query.Where("tbl_forms.modified_on >= ? AND tbl_forms.modified_on <= ?", filter.FromDate, filter.ToDate+" 23:59:59")
	}

	if limit != 0 {
		query.Limit(limit).Offset(offset).Find(&Forms)

		return Forms, 0, err
	}

	query.Find(&Forms).Count(&Count)

	return Forms, Count, nil
}

//Create Forms

func (Formsmodel FormModel) CreateForm(tblforms *TblForm, DB *gorm.DB) error {

	if err := DB.Table("tbl_forms").Create(&tblforms).Error; err != nil {

		return err
	}

	return nil
}

func (Formsmodel FormModel) ChangeStatus(forms *TblForm, DB *gorm.DB) error {

	if err := DB.Table("tbl_forms").Where("id=? and tenant_id=?", &forms.Id, &forms.TenantId).UpdateColumns(map[string]interface{}{"status": &forms.Status, "modified_by": &forms.ModifiedBy, "modified_on": &forms.ModifiedOn}).Error; err != nil {

		return err

	}

	return nil
}

func (Formsmodel FormModel) FormsDelete(forms *TblForm, DB *gorm.DB) error {

	if err := DB.Table("tbl_forms").Where("id=? and tenant_id=?", &forms.Id, &forms.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": &forms.IsDeleted, "deleted_by": &forms.DeletedBy, "deleted_on": &forms.DeletedOn}).Error; err != nil {

		return err

	}

	return nil

}

func (Formsmodel FormModel) EditForm(id int, tenantid int, DB *gorm.DB) (Forms TblForm, err error) {

	if err := DB.Table("tbl_forms").Where("id=? and tenant_id=?", id, tenantid).First(&Forms).Error; err != nil {

		return TblForm{}, err

	}

	return Forms, nil

}

func (Formsmodel FormModel) UpdateForm(tblforms *TblForm, DB *gorm.DB) error {

	if err := DB.Table("tbl_forms").Where("id=? and tenant_id=?", &tblforms.Id, &tblforms.TenantId).UpdateColumns(map[string]interface{}{"form_title": &tblforms.FormTitle, "form_slug": &tblforms.FormSlug, "form_data": &tblforms.FormData, "status": &tblforms.Status, "modified_by": &tblforms.ModifiedBy, "modified_on": &tblforms.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (Formsmodel FormModel) MultiSelectFormDelete(forms *TblForm, id []int, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_forms").Where("id in (?) and tenant_id=?", id, forms.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": forms.IsDeleted, "deleted_on": forms.DeletedOn, "deleted_by": forms.DeletedBy}).Error; err != nil {

		return err
	}

	return nil

}

func (Formsmodel FormModel) MultiSelectStatusChange(forms *TblForm, id []int, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_forms").Where("id in (?) and tenant_id=?", id, forms.TenantId).UpdateColumns(map[string]interface{}{"status": forms.Status, "modified_on": forms.ModifiedOn, "modified_by": forms.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil

}