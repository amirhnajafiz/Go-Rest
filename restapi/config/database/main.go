package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"restapi/restapi/internal/database/migrations"
	"restapi/restapi/internal/database/seeder"
)

var DB *gorm.DB

// Connect : creates a connection to our database
func Connect(migrate bool) {
	dsn := "root:Amir2222@tcp(127.0.0.1:3306)/books"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	log.Println("Database connected on 127.0.0.1:3306 ...")

	// Migrating the database and seeding the data into database
	if migrate {
		migrations.Migrate(db)
		seeder.Seed(db)
	}

	DB = db
}
