package main

import (
	"fmt"
	"net/http"
	"sync"
)

var prPool = sync.Pool{
	New: func()interface{} {

		return new(DBConnection)
	},
}


type Config struct{
	URL string
	Name string
}

type DBConnection struct{
	db * DBConnection
	config Config
}

func New() * DBConnection{
	return &DBConnection{
		db:     &DBConnection{},
		config: Config{},
	}
}

func (d DBConnection) Fetch() string{
	return "fetched data"
}

func main() {
	//body := bytes.NewBuffer([]byte{})
	//req, err := http.NewRequest(http.MethodGet, "http://golang.tutorial.org", body)
	//if err != nil{
	//	panic(err)
	//}
	http.HandleFunc("/", handler)
}


// methods to initialise the database connections

func handler(w http.ResponseWriter, r * http.Request){
	// db = DBConnectionP{}
	db, ok := prPool.Get().(* DBConnection)
	prPool.Put(db)
	// initialise the state of the db
	if !ok{
		fmt.Fprintf(w, "wrong data type")
	}
	output := db.Fetch()
	// reset the state of the db object
	fmt.Fprintf(w, output)
}
