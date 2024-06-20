package domain

type TicketType string

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}
