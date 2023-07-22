package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/DoomGuy1818/restapi-server/internal/app/apiserver"
)

var (

	configPath string

)

func init(){
	flag.StringVar(&configPath, "config-path", "configs/APIServer.toml", "path to config file")
}


func main() {
	flag.Parse()


	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
		if err != nil  {
			log.Fatal(err)
		}

	s := apiserver.New(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
}