package migrations

import (
	"fmt"
	"reflect"

	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/database"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/models"
)

func Migrate() {

	dbConnection := database.DbConnection()

	fmt.Println("Migrating...")
	fmt.Print("\n")
	for table, model := range Migrations() {

		modelName := reflect.TypeOf(model).Elem().Name()

		if dbConnection.Migrator().HasTable(table) {
			fmt.Printf("Table '%s' associated to model '%v' already exitst\n", table, modelName)
			continue
		}

		fmt.Printf("Migrating table '%s' associated to model '%v' \n", table, modelName)
		dbConnection.Db().AutoMigrate(model)
	}

	//TODO: REMOVE this: it shall be dynamic using keycloack
	dbConnection.Db().Create(&models.User{
		Name:     "Franck",
		Email:    "djacotofranck@gmail.com",
		Password: "password",
	})

	fmt.Println("\nMigrations terminated")
	fmt.Print("\n")
}
