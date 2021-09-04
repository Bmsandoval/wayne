package db

import (
	"database/sql"
	"fmt"
	"github.com/bmsandoval/wayne/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

var DB *sql.DB

type Connection struct {
	Tx *sql.Tx
	*sql.DB
}

func Start(config configs.Configuration) (*Connection, error){
	// Open and ping database
	if err := open(config); err != nil {
		return nil, err }
	// Run UP migrations
	if err := migrate(config); err != nil {
		return nil, err }
	// Bind models for Squalor

	return &Connection{
		DB: DB,
	}, nil
}

func Stop() error {
	return DB.Close()
}

func open(config configs.Configuration) error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DbUser,
		config.DbPass,
		config.DbHost,
		config.DbPort,
		config.DbSchema)

	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err }

	if err := ping(); err != nil {
		return err }

	return nil
}

func ping() error {
	return DB.Ping()
}

func migrate(config configs.Configuration) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err }
	if err := goose.Up(DB, config.DbMigrationLocation); err != nil {
		return err }
	return nil
}
