package main

import (
	"log"

	"github.com/Juakstlomr/Clamtalea/server"
	"github.com/jpillora/opts"
)

var VERSION = "0.8.20-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Clamtalea",
		Port:       3000,
		ConfigPath: "Clamtalea.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
