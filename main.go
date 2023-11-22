package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type RequestHandler struct {
	db *sql.DB
	fs http.Handler
}

func (c *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := c.db.Exec("INSERT INTO users () values ()")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	c.fs.ServeHTTP(w, r)
}

func main() {
	connectionURL := "root:password@tcp(localhost:3306)/test_db?parseTime=true"
	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", &RequestHandler{
		db: db,
		fs: fs,
	})

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
