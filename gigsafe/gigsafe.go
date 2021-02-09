package gigsafe

import (
	"context"

	"intern.mn8.ee/arash/intern"
)

// Server is the implementation of the simulation
type Server struct {
	Port int
}

// ForgetMe forgets the DID
func (s *Server) ForgetMe(context.Context, *intern.Empty) (*intern.Success, error) {

}

// GetDID gets new DID
func (s *Server) GetDID(context.Context, *intern.Empty) (*intern.DID, error) {}

// GetEvents returns list of random events
func (s *Server) GetEvents(context.Context, *intern.Empty) (*intern.Events, error) {}

// CheckTickets checks there are tickets left
func (s *Server) CheckTickets(context.Context, *intern.EventIDs) (*intern.Tickets, error) {}

// RequestTicket gets the QR data
func (s *Server) RequestTicket(context.Context, *intern.QRCodeRequest) (*intern.QRCodeResponse, error) {
}
