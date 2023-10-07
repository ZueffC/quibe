package app

import (
	"errors"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	Boards struct {
		gorm.Model
		Name  string
		Index string
	}

	Messages struct {
		gorm.Model
		Author  string
		Text    string
		Content string
		BoardId uint
		TopicId uint //if null or 0 - it must be global message
	}

	Migrator struct {
		Models []interface{}
		db     *gorm.DB
	}
)

func (m Migrator) ConnectToDB() Migrator {
	m.db, _ = gorm.Open(
		sqlite.Open(os.Getenv("DB_NAME")),
		&gorm.Config{},
	)

	return m
}

func (m Migrator) Migrate() Migrator {
	if m.db == nil || len(m.Models) <= 0 {
		panic(errors.New("cannot access to DB"))
	}

	for _, table := range m.Models {
		m.db.AutoMigrate(table)
	}

	return m
}
