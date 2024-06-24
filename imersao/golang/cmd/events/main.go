package main

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-hawerroth/imersao-fullcycle-18/golang/internal/events/infra/repository"
	"github.com/gabriel-hawerroth/imersao-fullcycle-18/golang/internal/events/infra/service"
	"github.com/gabriel-hawerroth/imersao-fullcycle-18/golang/internal/events/usecase"

	httpHandler "github.com/gabriel-hawerroth/imersao-fullcycle-18/golang/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=healthcare_desenv")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewPostgresEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnersBaseURL := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnersBaseURL)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("GET /events", eventsHandler.ListEvents)
	r.HandleFunc("GET /events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("GET /events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	http.ListenAndServe(":8080", r)
}
