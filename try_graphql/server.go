package main

import (
	"database/sql"
	"flag"
	"github.com/mark-by/golang/try_graphql/graph/api/db"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mark-by/golang/try_graphql/graph"
	"github.com/mark-by/golang/try_graphql/graph/generated"
)

const defaultPort = "8080"

var newDB bool

func init()  {
	flag.BoolVar(&newDB, "newdb", false, "Создавать ли новую БД?")
}

func main() {
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbase, err := db.Connect()
	checkErr(err)

	if newDB {
		println("Init db...")
		initDB(dbase)
	} else {
		println("not init")
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Db: dbase}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB(dbase *sql.DB) {
	db.MustExec(dbase,"DROP TABLE IF EXISTS reviews")
	db.MustExec(dbase,"DROP TABLE IF EXISTS screenshots")
	db.MustExec(dbase,"DROP TABLE IF EXISTS videos")
	db.MustExec(dbase,"DROP TABLE IF EXISTS users")
	db.MustExec(dbase,"CREATE TABLE public.users (id SERIAL PRIMARY KEY, name varchar(255), email varchar(255))")
	db.MustExec(dbase,"CREATE TABLE public.videos (id SERIAL PRIMARY KEY, name varchar(255), description varchar(255), url text,created_at TIMESTAMP, user_id int, FOREIGN KEY (user_id) REFERENCES users (id))")
	db.MustExec(dbase,"CREATE TABLE public.screenshots (id SERIAL PRIMARY KEY, video_id int, url text, FOREIGN KEY (video_id) REFERENCES videos (id))")
	db.MustExec(dbase,"CREATE TABLE public.reviews (id SERIAL PRIMARY KEY, video_id int,user_id int, description varchar(255), rating varchar(255), created_at TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users (id), FOREIGN KEY (video_id) REFERENCES videos (id))")
	db.MustExec(dbase,"INSERT INTO users(name, email) VALUES('Ridham', 'contact@ridham.me')")
	db.MustExec(dbase,"INSERT INTO users(name, email) VALUES('Tushar', 'tushar@ridham.me')")
	db.MustExec(dbase,"INSERT INTO users(name, email) VALUES('Dipen', 'dipen@ridham.me')")
	db.MustExec(dbase,"INSERT INTO users(name, email) VALUES('Harsh', 'harsh@ridham.me')")
	db.MustExec(dbase,"INSERT INTO users(name, email) VALUES('Priyank', 'priyank@ridham.me')")
}