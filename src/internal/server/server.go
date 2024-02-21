package server

import (
	"rs/zerolog/log"
)

type Server interface {
	Run()
}

type server struct {
	log log.Logger
}
