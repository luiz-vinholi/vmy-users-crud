package main

import (
	"log"
	"os"
	"testing"

	"github.com/benweissmann/memongo"
)

func TestMain(m *testing.M) {
	log.Print("----------------------------------------------------------- COMEÇO")
	mongoServer, err := memongo.Start("4.0.5")
	defer mongoServer.Stop()
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("MONGODB_URI", mongoServer.URI())
	os.Setenv("MONGODB_DATABASE_NAME", memongo.RandomDatabase())
	log.Print("----------------------------------------------------------- COMEÇO")
	code := m.Run()
	log.Print("----------------------------------------------------------- FIM")
	os.Exit(code)
}
