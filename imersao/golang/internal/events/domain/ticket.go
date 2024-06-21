package domain

import "errors"

type TicketKind string

var (
	ErrTicketPriceInvalid = errors.New("ticket price must be greater than 0")
	ErrInvalidTicketKind  = errors.New("invalid ticket kind")
)

const (
	TicketTypeHalf TicketKind = "half"
	TicketTypeFull TicketKind = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func IsValidTicketType(ticketType TicketKind) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceInvalid
	}
	return nil
}
