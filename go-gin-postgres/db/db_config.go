package db

import (
	"fmt"
	"os"
)

func Config() string {
	return fmt.Sprintf("host=%s port=%s user=%s sslmode=disable dbname=%s password=%s",
		os.Getenv("POSTGRES_HOST"), "5432", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
}
