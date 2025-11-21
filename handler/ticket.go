package handler

import (
	"github.com/dandimuzaki/ticket/dto"
	"github.com/dandimuzaki/ticket/service"
)

type TicketHandler struct {
	TicketService service.TicketService
}

// constructor
func NewTicketHandler(ticketService service.TicketService) TicketHandler {
	return TicketHandler{
		TicketService: ticketService,
	}
}

// method get harga
func (ticketHandler *TicketHandler) GetTicket(req dto.NewRequest) (dto.NewResponse, error) {
	ticket, err := ticketHandler.TicketService.GetTicket(req)

	if err != nil {
		return dto.NewResponse{}, err
	}

	return ticket, nil
}