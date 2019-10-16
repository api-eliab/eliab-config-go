package eliab

import (
	"github.com/jgolang/log"

	"github.com/BurntSushi/toml"
)

func loadConfiguration() {
	path := "./config.toml"

	if _, err := toml.DecodeFile(path, &Get); err != nil {
		log.Printf("Couldn't read config file at [%s]\n", path)
		log.Fatal(err)
	}
}

// Init doc ...
func init() {
	loadConfiguration()
	if !dbConnect() {
		log.Panic("¡Error al conectar a la base de datos! \n")
	}
}
