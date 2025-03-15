package main

import (
	"github.com/zuni-lab/yexus-api/cmd/api/server"
)

func main() {
	s := server.New()
	err := s.Start()
	defer s.Close()
	if err != nil {
		panic(err)
	}
}
