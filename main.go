package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/migrateMyDatabase", DBMigrationHandler)
	http.ListenAndServe(":8080", nil)
}

func DBMigrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DB is migrated")
}
