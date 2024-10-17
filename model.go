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
	CreatedDate      string    `gorm:"-:migration;<-:false"`
	ModifiedDate     string    `gorm:"-:migration;<-:false"`
}

type TblFormRegistrations struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name      string    `gorm:"type:character varying"`
	EmailId   string    `gorm:"type:character varying"`
	MobileNo  string    `gorm:"type:character varying"`
	FormId    int       `gorm:"type:integer;DEFAULT:NULL"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblFormResponse struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormId       int       `gorm:"type:integer;"`
	FormResponse string    `gorm:"type:character varying"`
	UserId       int       `gorm:"type:integer;"`
	IsActive     int       `gorm:"type:integer"`
	IsDeleted    int       `gorm:"type:integer;DEFAULT:0"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId     int       `gorm:"type:integer"`
}

type TblFormResponses struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormId       int       `gorm:"type:integer;"`
	FormResponse string    `gorm:"type:character varying"`
	UserId       int       `gorm:"type:integer;"`
	IsActive     int       `gorm:"type:integer"`
	IsDeleted    int       `gorm:"type:integer;DEFAULT:0"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId     int       `gorm:"type:integer"`
	DateString   string    `gorm:"-"`
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

type FormResponseCount struct {
	ID            int64 `json:"id"`
	ResponseCount int64 `json:"response_count"`
}

var Formsmodel FormModel

// FormList
func (Formsmodel FormModel) FormsList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int, status int) (Forms []TblForms, Count int64, err error) {

	query := DB.Table("tbl_forms").
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

// Response count

func (Formsmodel FormModel) ResponseCount(DB *gorm.DB, tenantid int) ([]FormResponseCount, error) {
	var results []FormResponseCount

	err := DB.Table("tbl_forms").
		Select("tbl_forms.id, COUNT(tbl_form_responses.id) AS response_count").
		Joins("INNER JOIN tbl_form_responses ON tbl_form_responses.form_id = tbl_forms.id").
		Where("tbl_forms.tenant_id = ?", tenantid).
		Group("tbl_forms.id").
		Scan(&results).Error

	return results, err
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

	if err := DB.Table("tbl_forms").Where("id in (?) and tenant_id=?", id, forms.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": forms.IsDeleted, "deleted_on": forms.DeletedOn, "deleted_by": forms.DeletedBy}).Error; err != nil {

		return err
	}

	return nil

}

func (Formsmodel FormModel) MultiSelectStatusChange(forms *TblForm, id []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_forms").Where("id in (?) and tenant_id=?", id, forms.TenantId).UpdateColumns(map[string]interface{}{"status": forms.Status, "modified_on": forms.ModifiedOn, "modified_by": forms.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil

}

// Form Preview
func (Formsmodel FormModel) GetPreview(forms *TblForm, DB *gorm.DB, uuid string) (err error) {

	if err = DB.Table("tbl_forms").Where("uuid = ?", uuid).Find(&forms).Error; err != nil {

		return err
	}

	return nil
}

// Form Response
func (Formsmodel FormModel) CreateResponse(response *TblFormResponse, DB *gorm.DB) (err error) {

	if err = DB.Table("tbl_form_responses").Create(&response).Error; err != nil {

		return err
	}

	return nil
}

func (Formsmodel FormModel) FormResponseList(offset int, limit int, filter Filter, response *TblFormResponses, DB *gorm.DB) (ResponseList []TblFormResponses, Count int64, FormTitle string, err error) {

	query := DB.Table("tbl_form_responses").Where("form_id=? and user_id=? and tenant_id=?", response.FormId, response.UserId, response.TenantId).Order("tbl_form_responses.created_on desc")

	if filter.Keyword != "" {
		query = query.Where("Lower(TRIM(form_title)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if filter.FromDate != "" {
		query = query.Where("tbl_form_responses.created_on >= ? AND tbl_form_responses.created_on <= ?", filter.FromDate, filter.ToDate+" 23:59:59")
	}

	if limit != 0 {
		query.Limit(limit).Offset(offset).Find(&ResponseList)

		return ResponseList, 0, "", err
	}

	err = query.Find(&ResponseList).Count(&Count).Error
	if err != nil {
		return nil, 0, "", err
	}

	var forms TblForm

	if err = DB.Table("tbl_forms").Where("id = ?", response.FormId).First(&forms).Error; err != nil {

		return nil, 0, "", err
	}

	FormTitle = forms.FormTitle

	return ResponseList, Count, FormTitle, err
}
