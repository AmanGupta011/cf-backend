package controllers

import (
	"sync"

	"cf-backend/pkg/store"
)

type Application struct {
	Counter     int
	TicketStore store.TicketStore
	Channel     chan bool
	sync.Mutex
}