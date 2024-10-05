package formbuilders

import (
	"strings"
	"time"

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
func (forms *Formbuilders) FormBuildersList(Limit int, offset int, filter Filter, tenantid int, status int) (formlist []TblForms, count int64, err error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return []TblForms{}, 0, AuthErr
	}

	Formsmodel.DataAccess = forms.DataAccess

	Formsmodel.UserId = forms.UserId

	_, TotalFormsCount, _ := Formsmodel.FormsList(0, 0, filter, forms.DB, tenantid, status)

	Formlist, _, err := Formsmodel.FormsList(offset, Limit, filter, forms.DB, tenantid, status)

	if err != nil {
		return []TblForms{}, 0, err
	}

	return Formlist, TotalFormsCount, nil
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

	Forms.FormTitle = tblform.FormTitle

	Forms.FormSlug = strings.ToLower(strings.ReplaceAll(strings.TrimRight(tblform.FormTitle, " "), " ", "-"))

	Forms.FormData = tblform.FormData

	Forms.Status = tblform.Status

	Forms.IsActive = tblform.IsActive

	Forms.CreatedBy = tblform.CreatedBy

	Forms.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

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
