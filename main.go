package main

import (
	"fmt"

	"github.com/dandimuzaki/ticket/dto"
	"github.com/dandimuzaki/ticket/handler"
	"github.com/dandimuzaki/ticket/service"
)

func main() {
	req := dto.NewRequest{"Dandi", "Bandung"}

	service := service.NewTicketService()
	handler := handler.NewTicketHandler(service)

	res, err := handler.GetTicket(req)
	if err != nil {
		fmt.Printf("Ada kesalahan: %v\n", err)
	} else {
		fmt.Printf("=== Harga Tiket ===\nPenumpang : %s\nTujuan    : %s\nHarga     : %.2f\n===================", res.Penumpang, res.Tujuan, res.Harga)
	}
}