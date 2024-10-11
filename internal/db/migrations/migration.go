package main

import (
	"io"
	"log/slog"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"

	"github.com/Anuolu-2020/hackernews-api-clone/internal/db"
)

func main() {
	// Define all models to be included in the migration
	modelsToMigrate := []interface{}{
		&db.Users{},
		&db.Links{},
	}

	stmts, err := gormschema.New("postgres").Load(modelsToMigrate...)
	if err != nil {
		slog.Error("Failed to load GORM schema", "err", err.Error())
		os.Exit(1)
	}

	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		slog.Error("Failed to write GORM schema", "err", err.Error())
		os.Exit(1)
	}
}
