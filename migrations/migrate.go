package migrations

import (
	"log"

	"github.com/m3rashid/cleopatra/app"
	"github.com/m3rashid/cleopatra/pkg/models"
)

func Migrate() {
	log.Println("Initiating migration...")
	err := app.Http.Database.DB.Migrator().AutoMigrate(
		&models.Role{},
		&models.RoleAndPermission{},
		&models.User{},
		&models.UserMeta{},
		&models.UserSetting{},
		&models.File{},
		&models.PaymentMethod{},
		&models.Payment{},
		&models.UserFile{},
		&models.Transaction{},
		&models.UserTransactionLog{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migration Completed...")
}
