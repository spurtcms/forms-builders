# forms-builders

# Installation

``` bash
go get github.com/spurtcms/forms-builders 
```


# Usage Example

``` bash
import (
	"github.com/spurtcms/auth"
	"github.com/spurtcms/forms-builders"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		DB: &gorm.DB{},
		RoleId: 1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Forms Builder", auth.CRUD, 1)

	Forms := FormSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		//list forms-builders
		Formlist, TotalFormsCount, responseCount, err := Forms.FormBuildersList(10, 0, Filter{}, 1, 1)

		if err != nil {
			fmt.Println(err)
		}

		//create forms-builders
		err := Forms.CreateForms(tblForm)
		if err != nil {
			fmt.Println(err)
		}

		//update forms-builders
		err := Forms.UpdateForms(tblForm, 1)
		if err != nil {
			fmt.Println(err)
		}

		// delete forms-builders
		err := Forms.Formdelete(10, 2, 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

```
# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/forms-builders/issues]. 
