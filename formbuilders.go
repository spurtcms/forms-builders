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
	fmt.Println("FormResponseList")
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

// change cta status
func (forms *Formbuilders) ChangeFormStatus(id int, isactive int, userid int, tenantid int) (bool, error) {

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

func (forms *Formbuilders) Addctatomycollecton(uid string, tenantid int, userid int) (bool, error) {

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

	err1 := Formsmodel.CreateForm(&NewForms, forms.DB)
	if err1 != nil {

		return false, err1

	}
	return true, nil
}

func (forms *Formbuilders) Removectatomycollecton(uid string, tenantid int, userid int) (bool, error) {

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
