package store

import "cf-backend/models"

type TicketStore interface {
	Add(ticket *models.Ticket) error
	Query(id int) (*models.Ticket, error)
	Update(id int, updatedTicket *models.Ticket) error
	Close() error
}