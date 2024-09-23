package formbuilders

import (
	"fmt"

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

//FormList 
func (forms *Formbuilders) FormBuildersList(tenantid int) ([]TblForms, error) {

	if AuthErr := AuthandPermission(forms); AuthErr != nil {

		return []TblForms{}, AuthErr
	}

	Formsmodel.DataAccess = forms.DataAccess

	Formsmodel.UserId = forms.UserId

	Form, err := Formsmodel.FormsList(forms.DB, tenantid)

	if err != nil {
		fmt.Println(err)
	}

	return Form, nil
}
