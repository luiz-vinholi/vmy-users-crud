package main

import (
	"log"
	"os"
	"testing"

	"github.com/benweissmann/memongo"
)

func TestMain(m *testing.M) {
	mongoServer, err := memongo.Start("4.0.5")
	defer mongoServer.Stop()
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("MONGODB_URI", mongoServer.URI())
	os.Setenv("MONGODB_DATABASE_NAME", memongo.RandomDatabase())
	code := m.Run()
	os.Exit(code)
}
