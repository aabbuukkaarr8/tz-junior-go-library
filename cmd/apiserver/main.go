package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"tz-junior-go-library/internal/apiserver"
	booksHandler "tz-junior-go-library/internal/handler/books"
	booksRepo "tz-junior-go-library/internal/repository/books"
	booksSrv "tz-junior-go-library/internal/service/books"
	"tz-junior-go-library/internal/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	db := store.New()
	err = db.Open(config.Store.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	//repo
	booksRepo := booksRepo.NewRepository(db)
	//service
	booksSrv := booksSrv.NewService(booksRepo)
	//handler
	booksHandler := booksHandler.NewHandler(booksSrv)
	s := apiserver.New(config)
	s.ConfigureRouter(booksHandler)
	if err := s.Run(); err != nil {
		panic(err)
	}
}
