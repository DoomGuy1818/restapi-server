package store

import (
	"os"
	"testing"
)


var(
	databaseURL string
)

func TestMain(m *testing.M){
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "root:musled_03@tcp(localhost:3306)/rest_api_test?multiStatements=true" 
	}

	os.Exit(m.Run())
}