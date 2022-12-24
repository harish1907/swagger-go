package initializers

import "github.com/harish1907/zgolang/models"


func MigrateDatabaseSystem() {
	DB.AutoMigrate(&models.UserModel{})
}