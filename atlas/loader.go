package main

import (
	"fmt"
	"io"
	"os"

	"gmtc/login/models"

	"ariga.io/atlas-provider-gorm/gormschema"
)

// Using guide from docs: https://atlasgo.io/guides/orms/gorm

func main() {
	stmts, err := gormschema.New("postgres").Load(models.GetAllModels()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
