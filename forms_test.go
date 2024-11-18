package formbuilders

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "postgres",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "Nov08",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to database:", err)

	}
	if err != nil {

		return nil, err

	}

	return db, nil
}

func TestFormList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    1,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Forms Builder", auth.CRUD, 1)

	Forms := FormSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permission {

		Formlist, TotalFormsCount, responseCount, err := Forms.FormBuildersList(10, 0, Filter{}, 1, 1)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(Formlist, TotalFormsCount, responseCount)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestCreateForms(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    1,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Forms Builder", auth.CRUD, 1)

	Forms := FormSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	tblForm := TblForm{
		FormTitle: "Hello Forms",
		FormData:  "",
		Status:    1,
		IsActive:  1,
		CreatedBy: 2,
		TenantId:  1,
	}

	if permission {

		err := Forms.CreateForms(tblForm)
		if err != nil {
			fmt.Println(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestUpdate(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    1,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Forms Builder", auth.CRUD, 1)

	Forms := FormSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	tblForm := TblForm{
		Id:         9,
		FormTitle:  "Hello",
		FormData:   "",
		Status:     1,
		IsActive:   1,
		ModifiedBy: 2,
	}

	if permission {

		err := Forms.UpdateForms(tblForm, 1)
		if err != nil {
			fmt.Println(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestDeleteForms(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    1,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Forms Builder", auth.CRUD, 1)

	Forms := FormSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})


	if permission {
		err := Forms.Formdelete(10, 2, 1)
		if err != nil {
			fmt.Println(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}
}
