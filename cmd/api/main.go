package main

import (
	"github.com/zuni-lab/dexon-service/cmd/api/server"
)

func main() {
	s := server.New()
	err := s.Start()
	defer s.Close()
	if err != nil {
		panic(err)
	}
}
