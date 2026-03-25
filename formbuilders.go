package formbuilders

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spurtcms/auth/migration"
)

func FormSetup(config Config) *Formbuilders {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Formbuilders{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		Permissions:      config.Permissions,
	}

}

// FormList
func (forms *Formbuilders) FormBuildersList(Limit int, offset int, filter Filter, tenantid string, status int, entryid int, channelslug string, defaultlist int) (formlist []TblForms, count int64, ResponseCount []FormResponseCount, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return []TblForms{}, 0, []FormResponseCount{}, AuthErr
	}

	Formsmodel.DataAccess = forms.DataAccess

	Formsmodel.UserId = forms.UserId

	_, TotalFormsCount, _ := Formsmodel.FormsList(0, 0, filter, forms.DB, tenantid, status, channelslug, defaultlist)

	Formlist, _, err := Formsmodel.FormsList(offset, Limit, filter, forms.DB, tenantid, status, channelslug, defaultlist)

	ResponseCount, _ = Formsmodel.ResponseCount(forms.DB, tenantid, entryid)

	if err != nil {
		return []TblForms{}, 0, []FormResponseCount{}, err
	}

	return Formlist, TotalFormsCount, ResponseCount, nil
}

//Create functionality

func (forms *Formbuilders) CreateForms(tblform TblForm) (TblForm, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return TblForm{}, AuthErr

	}

	if tblform.FormTitle == "" {
		return TblForm{}, ErrorFormName
	}
	var Forms TblForm

	uuid := (uuid.New()).String()

	arr := strings.Split(uuid, "-")

	Forms.Uuid = arr[len(arr)-1]

	Forms.FormTitle = tblform.FormTitle

	Forms.FormSlug = strings.ToLower(strings.ReplaceAll(strings.TrimRight(tblform.FormTitle, " "), " ", "-"))

	Forms.FormData = tblform.FormData

	Forms.Status = tblform.Status

	Forms.IsActive = tblform.IsActive

	Forms.CreatedBy = tblform.CreatedBy

	Forms.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.TenantId = tblform.TenantId

	Forms.ChannelId = tblform.ChannelId

	Forms.ChannelName = tblform.ChannelName

	Forms.FormDescription = tblform.FormDescription

	Forms.FormImagePath = tblform.FormImagePath

	Forms.FormPreviewImagepath = tblform.FormPreviewImagepath

	Forms.ImageName = tblform.ImageName

	Forms.ImagePath = tblform.ImagePath

	// added by karthikeyan

	Forms.FormSlug = tblform.FormSlug
	Forms.MetaTitle = tblform.MetaTitle
	Forms.MetaDescription = tblform.MetaDescription
	Forms.Keywords = tblform.Keywords
	Forms.Recaptcha = tblform.Recaptcha
	Forms.OnScreen = tblform.OnScreen
	Forms.EmailContent = tblform.EmailContent
	Forms.SmtpProtection = tblform.SmtpProtection

	formdetails, err := Formsmodel.CreateForm(&Forms, forms.DB)
	if err != nil {

		return TblForm{}, err

	}
	return formdetails, nil
}

func (forms *Formbuilders) StatusChange(id int, status int, modifiedby int, tenantid string) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}
	var Forms TblForm
	Forms.Id = id

	Forms.Status = status

	Forms.ModifiedBy = modifiedby

	Forms.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.TenantId = tenantid

	err := Formsmodel.ChangeStatus(&Forms, forms.DB)
	if err != nil {

		return err

	}
	return nil
}

func (forms *Formbuilders) Formdelete(id, deletedby int, tenantid string) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}

	var Forms TblForm

	Forms.Id = id

	Forms.DeletedBy = deletedby

	Forms.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.IsDeleted = 1

	Forms.TenantId = tenantid

	err := Formsmodel.FormsDelete(&Forms, forms.DB)
	if err != nil {

		return err

	}
	return nil
}

func (forms *Formbuilders) FormsEdit(id int, TenantId string) (FormList TblForm, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return TblForm{}, AuthErr

	}

	Forms, err := Formsmodel.EditForm(id, TenantId, forms.DB)
	if err != nil {

		return TblForm{}, err

	}
	return Forms, nil

}

func (forms *Formbuilders) UpdateForms(tblforms TblForm, tenantid string) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}
	var Forms TblForm

	Forms.Id = tblforms.Id

	Forms.FormTitle = tblforms.FormTitle

	Forms.FormSlug = strings.ToLower(strings.ReplaceAll(strings.TrimRight(tblforms.FormTitle, " "), " ", "-"))

	Forms.FormData = tblforms.FormData

	Forms.Status = tblforms.Status

	Forms.ModifiedBy = tblforms.ModifiedBy

	Forms.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.TenantId = tenantid

	Forms.ChannelId = tblforms.ChannelId

	Forms.ChannelName = tblforms.ChannelName

	Forms.FormPreviewImagepath = tblforms.FormPreviewImagepath

	Forms.FormPreviewImagename = tblforms.FormPreviewImagename

	Forms.FormDescription = tblforms.FormDescription

	Forms.ImageName = tblforms.ImageName

	Forms.ImagePath = tblforms.ImagePath

	Forms.FormSlug = tblforms.FormSlug
	Forms.MetaTitle = tblforms.MetaTitle
	Forms.MetaDescription = tblforms.MetaDescription
	Forms.Keywords = tblforms.Keywords
	Forms.Recaptcha = tblforms.Recaptcha
	Forms.OnScreen = tblforms.OnScreen
	Forms.EmailContent = tblforms.EmailContent
	Forms.SmtpProtection = tblforms.SmtpProtection

	err := Formsmodel.UpdateForm(&Forms, forms.DB)
	if err != nil {

		return err

	}
	return nil

}

func (forms *Formbuilders) MultiSelectDeleteForm(formids []int, modifiedby int, tenantid string) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}

	var Forms TblForm

	Forms.DeletedBy = modifiedby

	Forms.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.IsDeleted = 1

	Forms.TenantId = tenantid

	err := Formsmodel.MultiSelectFormDelete(&Forms, formids, forms.DB)
	if err != nil {

		return err

	}
	return nil
}

func (forms *Formbuilders) MultiSelectStatus(formids []int, status int, modifiedby int, tenantid string) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}

	var Forms TblForm

	Forms.ModifiedBy = modifiedby

	Forms.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Forms.Status = status

	Forms.TenantId = tenantid

	err := Formsmodel.MultiSelectStatusChange(&Forms, formids, forms.DB)
	if err != nil {

		return err

	}
	return nil

}

// Froms Preview
func (forms *Formbuilders) FormPreview(uuid string) (Form TblForm, Err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return TblForm{}, AuthErr

	}
	var Forms TblForm

	err := Formsmodel.GetPreview(&Forms, forms.DB, uuid)
	if err != nil {

		return Forms, nil

	}

	return Forms, nil

}

func (forms *Formbuilders) CreateFormResponse(response TblFormResponse) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}

	var Response TblFormResponse

	Response.FormId = response.FormId

	Response.FormResponse = response.FormResponse

	Response.UserId = response.UserId

	Response.IsActive = 1

	Response.IsDeleted = 0

	Response.CreatedBy = response.UserId

	Response.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Response.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Response.TenantId = response.TenantId

	Response.EntryId = response.EntryId

	Response.Name = response.Name

	Response.Ticket = response.Ticket

	err := Formsmodel.CreateResponse(&Response, forms.DB)
	if err != nil {

		return err

	}

	return nil
}

func (forms *Formbuilders) FormDetailLists(Limit int, offset int, filter Filter, formid, userid, entryid int, tenantid string) (response []TblFormResponses, count int64, FormTitle string, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return []TblFormResponses{}, 0, "", AuthErr

	}
	fmt.Println("FormResponseList")
	var Response TblFormResponses

	Response.EntryId = entryid

	Response.FormId = formid

	Response.UserId = userid

	Response.TenantId = tenantid

	_, TotalResponseCount, _, _ := Formsmodel.FormResponseList(0, 0, filter, &Response, forms.DB)

	responselist, _, _, err := Formsmodel.FormResponseList(offset, Limit, filter, &Response, forms.DB)

	_, _, formtitle, _ := Formsmodel.FormResponseList(0, 0, filter, &Response, forms.DB)

	if err != nil {

		return []TblFormResponses{}, 0, "", err

	}

	return responselist, TotalResponseCount, formtitle, nil

}

// change cta status
func (forms *Formbuilders) ChangeFormStatus(id int, isactive int, userid int, tenantid string) (bool, error) {

	autherr := AuthandPermission(forms)

	if autherr != nil {

		return false, autherr
	}

	var form TblForm

	form.ModifiedBy = userid

	form.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Formsmodel.FormIsActive(&form, id, isactive, forms.DB, tenantid)

	return true, nil
}

//Add to mycollection//

func (forms *Formbuilders) Addctatomycollecton(uid string, tenantid string, userid int, channelid string) (bool, error) {

	fmt.Println("dfdfdfdf")

	autherr := AuthandPermission(forms)

	if autherr != nil {

		return false, autherr
	}

	var Forms TblForm

	err := Formsmodel.GetPreview(&Forms, forms.DB, uid)
	if err != nil {

		return false, nil

	}

	var NewForms TblForm

	uuid := (uuid.New()).String()

	arr := strings.Split(uuid, "-")

	NewForms.Uuid = arr[len(arr)-1]

	NewForms.FormTitle = Forms.FormTitle

	NewForms.FormSlug = strings.ToLower(strings.ReplaceAll(strings.TrimRight(Forms.FormTitle, " "), " ", "-"))

	NewForms.FormData = Forms.FormData

	NewForms.Status = Forms.Status

	NewForms.IsActive = Forms.IsActive

	NewForms.CreatedBy = userid

	NewForms.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	NewForms.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	NewForms.TenantId = tenantid

	NewForms.FormImagePath = Forms.FormImagePath

	NewForms.FormDescription = Forms.FormDescription

	NewForms.ChannelId = channelid

	NewForms.ChannelName = Forms.ChannelName

	NewForms.FormPreviewImagename = Forms.FormPreviewImagename

	NewForms.FormPreviewImagepath = Forms.FormPreviewImagepath

	_, err1 := Formsmodel.CreateForm(&NewForms, forms.DB)
	if err1 != nil {

		return false, err1

	}
	return true, nil
}

func (forms *Formbuilders) Removectatomycollecton(uid string, tenantid string, userid int) (bool, error) {

	autherr := AuthandPermission(forms)

	if autherr != nil {

		return false, autherr
	}
	var form TblForm

	form.DeletedBy = userid

	form.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Formsmodel.Removecta(&form, uid, tenantid, forms.DB)
	if err1 != nil {

		return false, err1

	}
	return true, nil
}

func (forms *Formbuilders) GetCtaById(ctaid int) (form TblForm, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return TblForm{}, AuthErr

	}
	var Forms TblForm

	err = Formsmodel.GetCtaById(&Forms, forms.DB, ctaid)
	if err != nil {

		return Forms, nil

	}

	return Forms, nil

}

func (forms *Formbuilders) OverAllFormResponses(Limit int, offset int, filter Filter, tenantid string) (ResponseList []TblFormResponses, count int64, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return nil, 0, AuthErr

	}

	response, _, err := Formsmodel.OverallResponseList(offset, Limit, filter, tenantid, forms.DB)

	_, count, err = Formsmodel.OverallResponseList(0, 0, filter, tenantid, forms.DB)

	if err != nil {

		return nil, 0, err

	}

	return response, count, nil

}

func (forms *Formbuilders) ResponseDetail(ticket string, TenantId string) (*TblFormResponses, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return nil, AuthErr

	}

	response, err := Formsmodel.ResponseDetail(ticket, TenantId, forms.DB)
	if err != nil {

		return nil, err

	}
	return response, nil

}

func (forms *Formbuilders) ReplyForResponses(replycontent TblReplyForResponse) (bool, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return false, AuthErr

	}

	status, err := Formsmodel.ReplyforResponse(&replycontent, forms.DB)
	if err != nil {

		return false, err

	}
	return status, nil

}

func (forms *Formbuilders) ReplyForResponseList(ticket string, TenantId string) ([]TblReplyForResponse, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return nil, AuthErr

	}

	response, err := Formsmodel.ReplyforResponseList(ticket, TenantId, forms.DB)
	if err != nil {

		return nil, err

	}
	return response, nil

}

func (forms *Formbuilders) Closeticket(ticket string, TenantId string, notes string, modifiedon time.Time) (bool, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return false, AuthErr

	}
	status, err := Formsmodel.CloseTicket(ticket, TenantId, forms.DB, notes, modifiedon)
	if err != nil {

		return false, err

	}
	return status, nil

}

// ticket reopen
func (forms *Formbuilders) Reopenticket(ticketstatus string, TenantId string, ModifiedOn time.Time) (bool, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return false, AuthErr

	}
	status, err := Formsmodel.ReopenTicket(ticketstatus, TenantId, forms.DB, ModifiedOn)
	if err != nil {

		return false, err

	}
	return status, nil

}

func (forms *Formbuilders) TicketNotes(ticket string, TenantId string, notes string, modifiedOn time.Time) (bool, error) {
	fmt.Println("notesnotesnotesnotes :", notes)
	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return false, AuthErr

	}
	status, err := Formsmodel.TicketNotes(ticket, TenantId, forms.DB, notes, modifiedOn)
	if err != nil {

		return false, err

	}
	return status, nil

}

func (forms *Formbuilders) GetFormResponses(formId int, tenantId string) ([]TblFormResponses, error) {

	// Check permissions
	if authErr := AuthandPermission(forms); authErr != nil {
		return nil, authErr
	}

	// Prepare a slice to hold responses
	var responses []TblFormResponses

	// Call the model method to fetch responses
	if err := Formsmodel.GetFormResponses(formId, &responses, forms.DB); err != nil {
		return nil, err
	}

	// Return the fetched responses
	return responses, nil
}
func (forms *Formbuilders) GetFormById(formId int, tenantId string) (TblForms, error) {

	// Check permissions
	if authErr := AuthandPermission(forms); authErr != nil {
		return TblForms{}, authErr
	}

	// Prepare a slice to hold responses
	var form TblForms

	// Call the model method to fetch responses
	if err := Formsmodel.GetFormDetailById(formId, &form, forms.DB); err != nil {
		return TblForms{}, err
	}

	// Return the fetched responses
	return form, nil
}
