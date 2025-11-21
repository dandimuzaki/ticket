package service

import (
	"errors"
	"strings"

	"github.com/dandimuzaki/ticket/data"
	"github.com/dandimuzaki/ticket/dto"
	"github.com/dandimuzaki/ticket/model"
	"github.com/dandimuzaki/ticket/utils"
)

type TicketService struct{}

// constructor
func NewTicketService() TicketService {
	return TicketService{}
}

// method get harga
func (ticketService *TicketService) GetTicket(req dto.NewRequest) (dto.NewResponse, error) {
	if strings.TrimSpace(req.Penumpang) == "" ||
		strings.TrimSpace(req.Tujuan) == "" {
			return dto.NewResponse{}, errors.New("penumpang dan tujuan harus diisi")
		}
	
	destination, err := utils.StringToMap(data.Destination)
	if err != nil {
		return dto.NewResponse{}, errors.New("tidak bisa mengakses data")
	}

	found := false
	for key, _ := range destination {
		if req.Tujuan == key {
			found = true
		}
	}

	if !found {
		return dto.NewResponse{}, errors.New("tujuan tidak ditemukan")
	}

	ticket := model.Ticket{
		Penumpang: req.Penumpang,
		Tujuan: req.Tujuan,
		Harga: destination[req.Tujuan],
	}

	res := dto.NewResponse{
		Penumpang: ticket.Penumpang,
		Tujuan: ticket.Tujuan,
		Harga: ticket.Harga,
	}

	return res, nil
}