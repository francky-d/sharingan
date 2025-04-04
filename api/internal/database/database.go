package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConnection struct {
	db       *gorm.DB
	migrator gorm.Migrator
}

var dbConnection *databaseConnection

func (dc databaseConnection) Db() *gorm.DB {
	return dc.db
}

func (dc databaseConnection) Migrator() gorm.Migrator {
	return dc.migrator
}

func DbConnection() *databaseConnection {
	return getInstance()
}

func newDatabaseConnection() *databaseConnection {
	//TODO get this value from en file
	dsn := "host=api-db user=sharingan password=sharingan dbname=sharingan port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	return &databaseConnection{
		db:       db,
		migrator: db.Migrator(),
	}
}

func getInstance() *databaseConnection {
	if dbConnection != nil {
		return dbConnection
	}
	return newDatabaseConnection()
}
