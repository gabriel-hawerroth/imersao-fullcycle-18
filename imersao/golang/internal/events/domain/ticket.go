package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketKind string

var (
	ErrTicketPriceInvalid = errors.New("ticket price must be greater than 0")
	ErrInvalidTicketKind  = errors.New("invalid ticket kind")
)

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketType TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketType) {
		return nil, ErrInvalidTicketKind
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketType,
		Price:      event.Price,
	}

	ticket.CalculatePrice()

	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceInvalid
	}
	return nil
}
