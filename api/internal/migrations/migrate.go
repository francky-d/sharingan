package migrations

import (
	"fmt"
	"reflect"

	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/database"
)

func Migrate() error {

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

		err := dbConnection.Db().AutoMigrate(model)

		if err != nil {
			return fmt.Errorf("error occured during migration %w", err)
		}
		return nil

	}

	fmt.Println("\nMigrations terminated")
	fmt.Print("\n")
}
