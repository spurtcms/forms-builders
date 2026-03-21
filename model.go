package formbuilders

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Create Forms
type TblForm struct {
	Id                   int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormTitle            string    `gorm:"type:character varying"`
	FormSlug             string    `gorm:"type:character varying"` //use
	FormData             string    `gorm:"type:character varying"`
	Status               int       `gorm:"type:integer"`
	IsActive             int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy            int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted            int       `gorm:"type:integer;DEFAULT:0"`
	TenantId             string    `gorm:"type:character varying"`
	Uuid                 string    `gorm:"type:character varying"`
	FormImagePath        string    `gorm:"type:character varying"`
	FormDescription      string    `gorm:"type:character varying"`
	ChannelId            string    `gorm:"type:character varying"`
	ChannelName          string    `gorm:"type:character varying"`
	FormPreviewImagepath string    `gorm:"type:character varying"`
	FormPreviewImagename string    `gorm:"type:character varying"`
	ImageName            string    `gorm:"type:character varying"`
	ImagePath            string    `gorm:"type:character varying"`
	// karthi
	MetaTitle       string            `gorm:"type:character varying"`
	MetaDescription string            `gorm:"type:character varying"`
	Keywords        string            `gorm:"type:character varying"`
	Recaptcha       int               `gorm:"type:integer;DEFAULT:NULL"`
	OnScreen        datatypes.JSONMap `gorm:"type:jsonb"`
	EmailContent    datatypes.JSONMap `gorm:"type:jsonb"`
	SmtpProtection  datatypes.JSONMap `gorm:"type:jsonb"`
}

type TblForms struct {
	Id                   int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormTitle            string    `gorm:"type:character varying"`
	FormSlug             string    `gorm:"type:character varying"`
	FormData             string    `gorm:"type:character varying"`
	Status               int       `gorm:"type:integer"`
	IsActive             int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy            int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted            int       `gorm:"type:integer;DEFAULT:0"`
	TenantId             string    `gorm:"type:character varying"`
	Uuid                 string    `gorm:"type:character varying"`
	Username             string    `gorm:"<-:false"`
	ProfileImagePath     string    `gorm:"<-:false"`
	NameString           string    `gorm:"<-:false"`
	FirstName            string    `gorm:"<-:false"`
	LastName             string    `gorm:"<-:false"`
	DateString           string    `gorm:"-"`
	CreatedDate          string    `gorm:"-:migration;<-:false"`
	ModifiedDate         string    `gorm:"-:migration;<-:false"`
	FormImagePath        string    `gorm:"type:character varying"`
	FormDescription      string    `gorm:"type:character varying"`
	ChannelId            string    `gorm:"type:character varying"`
	ChannelName          string    `gorm:"type:character varying"`
	Channelnamearr       []string  `gorm:"-"`
	FormPreviewImagepath string    `gorm:"type:character varying"`
	FormPreviewImagename string    `gorm:"type:character varying"`
	ImageName            string    `gorm:"type:character varying"`
	ImagePath            string    `gorm:"type:character varying"`
	// karthi
	MetaTitle       string            `gorm:"type:character varying"`
	MetaDescription string            `gorm:"type:character varying"`
	Keywords        string            `gorm:"type:character varying"`
	Recaptcha       int               `gorm:"type:integer;DEFAULT:NULL"`
	OnScreen        datatypes.JSONMap `gorm:"type:jsonb"`
	EmailContent    datatypes.JSONMap `gorm:"type:jsonb"`
	SmtpProtection  datatypes.JSONMap `gorm:"type:jsonb"`
}

type Forms struct {
	Id       int    `gorm:"primaryKey;auto_increment;type:serial"`
	Name     string `gorm:"type:character varying"`
	EmailId  string `gorm:"type:character varying"`
	MobileNo string `gorm:"type:character varying"`
	Uuid     string `gorm:"type:character varying"`
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
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId     string    `gorm:"type:character varying"`
	Name         string    `gorm:"type:character varying"`
	Ticket       string    `gorm:"type:character varying"`
	CloseTicket  int       `gorm:"type:integer;DEFAULT:0"`
	EntryId      int       `gorm:"type:integer"`
}

type TblFormResponses struct {
	Id                   int           `gorm:"primaryKey;auto_increment;type:serial"`
	FormId               int           `gorm:"type:integer;"`
	FormResponse         string        `gorm:"type:character varying"`
	UserId               int           `gorm:"type:integer;"`
	IsActive             int           `gorm:"type:integer"`
	IsDeleted            int           `gorm:"type:integer;DEFAULT:0"`
	CreatedBy            int           `gorm:"type:integer"`
	CreatedOn            time.Time     `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedOn           time.Time     `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId             string        `gorm:"type:character varying"`
	Name                 string        `gorm:"type:character varying"`
	Ticket               string        `gorm:"type:character varying"`
	CloseTicket          int           `gorm:"type:integer;DEFAULT:0"`
	Replycontent         template.HTML `gorm:"-"`
	DateString           string        `gorm:"-"`
	ModifiedOnDateString string        `gorm:"-"`
	FormTitle            string        `gorm:"column:form_title" json:"form_title"`
	EntryId              int           `gorm:"type:integer"`
	Notes                string        `gorm:"type:character varying"`
}

type TblReplyForResponse struct {
	Id           int           `gorm:"primaryKey;auto_increment;type:serial"`
	Ticket       string        `gorm:"type:character varying"`
	Reply        string        `gorm:"type:character varying"`
	CreatedBy    int           `gorm:"type:integer"`
	CreatedOn    time.Time     `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId     string        `gorm:"type:character varying"`
	DateString   string        `gorm:"-"`
	Replycontent template.HTML `gorm:"-"`
}

type Filter struct {
	Keyword     string
	TicketNo    string
	FromDate    string
	ToDate      string
	ChannelSlug string
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
func (Formsmodel FormModel) FormsList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid string, status int, channelslug string, defaultlist int) (Forms []TblForms, Count int64, err error) {

	query := DB.Debug().Table("tbl_forms").
		Select("tbl_forms.*, tbl_users.username,tbl_users.first_name,tbl_users.last_name, tbl_users.profile_image_path").
		Joins("inner join tbl_users on tbl_forms.created_by=tbl_users.id").
		Where("tbl_forms.is_deleted = 0 ").
		Order("tbl_forms.modified_on desc")

	if status == 3 {

		query = query.Where("tbl_forms.status=? and tbl_forms.is_active=?", 1, 1)
	} else {

		query = query.Where("tbl_forms.status=? ", status)
	}
	if channelslug != "" {

		query = query.Where("string_to_array(LOWER(TRIM(tbl_forms.channel_name)), ',') @> ARRAY[?]::TEXT[]", channelslug)
	}

	if defaultlist == 1 {

		query = query.Where("tbl_forms.tenant_id is null")
	} else {

		query = query.Where("tbl_forms.tenant_id=?", tenantid)
	}
	if filter.Keyword != "" {
		query = query.Where("Lower(TRIM(form_title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if filter.FromDate != "" {
		query = query.Where("tbl_forms.modified_on >= ? AND tbl_forms.modified_on <= ?", filter.FromDate, filter.ToDate+" 23:59:59")
	}

	if filter.ChannelSlug != "" {
		query = query.Where("string_to_array(LOWER(TRIM(tbl_forms.channel_name)), ',') @> ARRAY[?]::TEXT[]", strings.ToLower(strings.TrimSpace(filter.ChannelSlug)))
	}

	if limit != 0 {
		query.Limit(limit).Offset(offset).Find(&Forms)

		return Forms, 0, err
	}

	query.Find(&Forms).Count(&Count)

	return Forms, Count, nil
}

// Response count

func (Formsmodel FormModel) ResponseCount(DB *gorm.DB, tenantid string, entryid int) ([]FormResponseCount, error) {
	var results []FormResponseCount

	if entryid == 0 {
		err := DB.Table("tbl_forms").
			Select("tbl_forms.id, COUNT(tbl_form_responses.id) AS response_count").
			Joins("INNER JOIN tbl_form_responses ON tbl_form_responses.form_id = tbl_forms.id").
			Where("tbl_forms.tenant_id = ?", tenantid).
			Group("tbl_forms.id").
			Scan(&results).Error

		return results, err
	} else {
		err := DB.Table("tbl_channel_entries").
			Select("tbl_channel_entries.id, COUNT(tbl_form_responses.entry_id) AS response_count").
			Joins("INNER JOIN tbl_form_responses ON tbl_form_responses.entry_id = tbl_channel_entries.id").
			Where("tbl_channel_entries.tenant_id = ?", tenantid).
			Group("tbl_channel_entries.id").
			Scan(&results).Error

		return results, err
	}
}

//Create Forms

func (Formsmodel FormModel) CreateForm(tblforms *TblForm, DB *gorm.DB) (formdetails TblForm, err error) {
	fmt.Println("makeprint")

	if err := DB.Debug().Table("tbl_forms").Create(&tblforms).Error; err != nil {

		return TblForm{}, err
	}

	return *tblforms, nil
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

func (Formsmodel FormModel) EditForm(id int, tenantid string, DB *gorm.DB) (Forms TblForm, err error) {

	if err := DB.Table("tbl_forms").Where("id=? and tenant_id=?", id, tenantid).First(&Forms).Error; err != nil {

		return TblForm{}, err

	}

	return Forms, nil

}

func (Formsmodel FormModel) UpdateForm(tblforms *TblForm, DB *gorm.DB) error {

	if err := DB.Table("tbl_forms").Where("id=? and tenant_id=?", &tblforms.Id, &tblforms.TenantId).UpdateColumns(map[string]interface{}{"form_title": &tblforms.FormTitle, "form_slug": &tblforms.FormSlug, "form_data": &tblforms.FormData, "status": &tblforms.Status, "modified_by": &tblforms.ModifiedBy, "modified_on": &tblforms.ModifiedOn, "channel_name": &tblforms.ChannelName, "channel_id": &tblforms.ChannelId, "form_preview_imagepath": &tblforms.FormPreviewImagepath, "form_preview_imagename": &tblforms.FormPreviewImagename, "form_description": &tblforms.FormDescription, "image_name": &tblforms.ImageName, "image_path": &tblforms.ImagePath, "meta_title": &tblforms.MetaTitle, "meta_description": &tblforms.MetaDescription, "keywords": &tblforms.Keywords, "on_screen": &tblforms.OnScreen, "email_content": &tblforms.EmailContent, "smtp_protection": &tblforms.SmtpProtection, "recaptcha": &tblforms.Recaptcha}).Error; err != nil {

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

	if err = DB.Debug().Table("tbl_forms").Where("uuid = ? and is_deleted=0", uuid).Find(&forms).Error; err != nil {

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

// Form OverallResponseList
func (Formsmodel FormModel) OverallResponseList(offset int, limit int, filter Filter, tenantid string, DB *gorm.DB) (ResponseList []TblFormResponses, Count int64, err error) {

	query := DB.Debug().
		Table("tbl_form_responses").
		Select("tbl_form_responses.*,tbl_forms.form_title as form_title").
		Joins("INNER JOIN tbl_forms ON tbl_form_responses.form_id = tbl_forms.id").
		Where("tbl_form_responses.tenant_id = ? and tbl_form_responses.is_deleted=0", tenantid).
		Order("tbl_form_responses.created_on desc")

	if filter.Keyword != "" {
		query = query.Where("Lower(TRIM(tbl_forms.form_title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if filter.TicketNo != "" {
		query = query.Where("Lower(TRIM(tbl_form_responses.ticket)) LIKE LOWER(TRIM(?))", "%"+filter.TicketNo+"%")
	}

	if filter.FromDate != "" {
		query = query.Where("tbl_form_responses.created_on >= ? AND tbl_form_responses.created_on <= ?", filter.FromDate, filter.ToDate+" 23:59:59")
	}

	if limit != 0 {
		query.Limit(limit).Offset(offset).Find(&ResponseList)

		return ResponseList, 0, err
	}

	query.Find(&ResponseList).Count(&Count)

	return ResponseList, Count, nil
}

func (Formsmodel FormModel) ResponseDetail(ticket string, tenantid string, DB *gorm.DB) (*TblFormResponses, error) {

	var response TblFormResponses

	result := DB.Table("tbl_form_responses").
		Select("tbl_form_responses.*,tbl_forms.form_title as form_title").
		Joins("INNER JOIN tbl_forms ON tbl_form_responses.form_id = tbl_forms.id").
		Where("tbl_form_responses.ticket=? and tbl_form_responses.tenant_id = ?", ticket, tenantid).
		First(&response)

	if result.Error != nil {
		return nil, result.Error
	}

	return &response, nil

}

func (Formsmodel FormModel) ReplyforResponse(replycontent *TblReplyForResponse, DB *gorm.DB) (bool, error) {

	result := DB.Table("tbl_reply_for_responses").Create(&replycontent)

	if result.Error != nil {
		return false, result.Error
	}

	// true = updated, false = no rows matched
	return result.RowsAffected > 0, nil

}

func (Formsmodel FormModel) ReplyforResponseList(ticket string, tenantid string, DB *gorm.DB) ([]TblReplyForResponse, error) {

	var response []TblReplyForResponse

	result := DB.Table("tbl_reply_for_responses").
		Where("ticket=? and tenant_id = ?", ticket, tenantid).
		Order("tbl_reply_for_responses.created_on desc").
		Find(&response)

	if result.Error != nil {
		return nil, result.Error
	}

	return response, nil

}

func (Formsmodel FormModel) CloseTicket(ticket string, tenantid string, DB *gorm.DB, notes string, modifiedon time.Time) (bool, error) {
	result := DB.Table("tbl_form_responses").
		Where("ticket=? and tenant_id = ?", ticket, tenantid).
		Updates(map[string]interface{}{
			"close_ticket": 1,
			"notes":        notes,
			"modified_on":  modifiedon,
		})

	if result.Error != nil {
		fmt.Println("notesnotesnotesnotes :", result.Error)
		return false, result.Error
	}

	return true, nil

}

func (Formsmodel FormModel) ReopenTicket(ticket string, tenantid string, DB *gorm.DB, ModifiedOn time.Time) (bool, error) {
	result := DB.Table("tbl_form_responses").
		Where("ticket=? and tenant_id = ?", ticket, tenantid).
		Updates(map[string]interface{}{
			"close_ticket": 0,
			"modified_on":  ModifiedOn,
		})

	if result.Error != nil {
		fmt.Println("Error :", result.Error)
		return false, result.Error
	}

	return true, nil

}

func (Formsmodel FormModel) TicketNotes(ticket string, tenantid string, DB *gorm.DB, notes string, modifiedOn time.Time) (bool, error) {
	result := DB.Table("tbl_form_responses").
		Where("ticket=? and tenant_id = ?", ticket, tenantid).
		Updates(map[string]interface{}{
			"notes":       notes,
			"modified_on": modifiedOn,
		})

	if result.Error != nil {
		fmt.Println("Error :", result.Error)
		return false, result.Error
	}

	return true, nil

}

func (Formsmodel FormModel) FormResponseList(offset int, limit int, filter Filter, response *TblFormResponses, DB *gorm.DB) (ResponseList []TblFormResponses, Count int64, FormTitle string, err error) {
	fmt.Println("FormResponseList")

	query := DB.Debug().Table("tbl_form_responses")
	if response.EntryId == 0 {
		query = query.Where("form_id=? and user_id=? and tenant_id=?", response.FormId, response.UserId, response.TenantId).Order("tbl_form_responses.created_on desc")
		fmt.Println("wwww:")
	} else {
		query = query.Debug().Where("entry_id=? and user_id=? and tenant_id=?", response.EntryId, response.UserId, response.TenantId).Order("tbl_form_responses.created_on desc")
		fmt.Println("qqqq:")
	}
	if filter.Keyword != "" {
		query = query.Where("Lower(TRIM(form_title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
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

	fmt.Println("Count:", Count)

	return ResponseList, Count, FormTitle, err
}

/*Isactive cta*/
func (Formsmodel FormModel) FormIsActive(tblform *TblForm, id, val int, DB *gorm.DB, tenantid string) error {

	if err := DB.Debug().Table("tbl_forms").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": val, "modified_on": tblform.ModifiedOn, "modified_by": tblform.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

//Remove cta from mycollection//

func (Formsmodel FormModel) Removecta(form *TblForm, uuid string, tenantid string, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_forms").Where("uuid=? and tenant_id=?", uuid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": form.DeletedOn, "deleted_by": form.DeletedBy}).Error; err != nil {

		return err
	}

	return nil
}

//Get CTA By Id//

func (Formsmodel FormModel) GetCtaById(forms *TblForm, DB *gorm.DB, id int) (err error) {

	if err = DB.Debug().Table("tbl_forms").Where("id = ? and is_deleted=0", id).Find(&forms).Error; err != nil {

		return err
	}

	return nil
}

func (Formsmodel FormModel) GetFormResponses(formId int, forms *[]TblFormResponses, DB *gorm.DB) (err error) {

	// Join tbl_form_responses with tbl_forms to get form_title
	if err = DB.Debug().
		Table("tbl_form_responses as r").
		Select("r.*, f.form_title").
		Joins("left join tbl_forms as f on r.form_id = f.id").
		Where("r.form_id = ? AND r.is_deleted = ?", formId, 0).
		Order("r.id DESC").
		Find(forms).Error; err != nil {

		return err

	}

	return nil
}
