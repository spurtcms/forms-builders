package formbuilders

import (
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
func (forms *Formbuilders) FormBuildersList(Limit int, offset int, filter Filter, tenantid int, status int) (formlist []TblForms, count int64, ResponseCount []FormResponseCount, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return []TblForms{}, 0, []FormResponseCount{}, AuthErr
	}

	Formsmodel.DataAccess = forms.DataAccess

	Formsmodel.UserId = forms.UserId

	_, TotalFormsCount, _ := Formsmodel.FormsList(0, 0, filter, forms.DB, tenantid, status)

	Formlist, _, err := Formsmodel.FormsList(offset, Limit, filter, forms.DB, tenantid, status)

	ResponseCount, _ = Formsmodel.ResponseCount(forms.DB, tenantid)

	if err != nil {
		return []TblForms{}, 0, []FormResponseCount{}, err
	}

	return Formlist, TotalFormsCount, ResponseCount, nil
}

//Create functionality

func (forms *Formbuilders) CreateForms(tblform TblForm) error {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return AuthErr

	}

	if tblform.FormTitle == "" {
		return ErrorFormName
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

	err := Formsmodel.CreateForm(&Forms, forms.DB)
	if err != nil {

		return err

	}
	return nil
}

func (forms *Formbuilders) StatusChange(id int, status int, modifiedby int, tenantid int) error {

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

func (forms *Formbuilders) Formdelete(id, deletedby, tenantid int) error {

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

func (forms *Formbuilders) FormsEdit(id int, TenantId int) (FormList TblForm, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return TblForm{}, AuthErr

	}

	Forms, err := Formsmodel.EditForm(id, TenantId, forms.DB)
	if err != nil {

		return TblForm{}, err

	}
	return Forms, nil

}

func (forms *Formbuilders) UpdateForms(tblforms TblForm, tenantid int) error {

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

	err := Formsmodel.UpdateForm(&Forms, forms.DB)
	if err != nil {

		return err

	}
	return nil

}

func (forms *Formbuilders) MultiSelectDeleteForm(formids []int, modifiedby int, tenantid int) error {

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

func (forms *Formbuilders) MultiSelectStatus(formids []int, status int, modifiedby int, tenantid int) error {

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

	Response.TenantId = response.TenantId

	err := Formsmodel.CreateResponse(&Response, forms.DB)
	if err != nil {

		return err

	}

	return nil
}

func (forms *Formbuilders) FormDetailLists(Limit int, offset int, filter Filter, formid, userid, tenantid int) (response []TblFormResponses, count int64, FormTitle string, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return []TblFormResponses{}, 0, "", AuthErr

	}

	var Response TblFormResponses

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
