package gigsafe

import (
	"context"
	"log"
	"os"

	"github.com/schollz/jsonstore"
	"intern.mn8.ee/arash/intern"
)

// Server is the implementation of the simulation
type Server struct {
	Port int
}

var didStore *jsonstore.JSONStore

// WipeStore deletes the persistent DID
func wipeStore() error {

	// Remove file
	err := os.Remove("data/did.json")

	// Check for error
	if err != nil {
		log.Fatal(err)
	}

	// Recreate file
	file, err := os.Create("data/did.json")

	// Check for error
	if err != nil {
		log.Fatal(err)
	}

	// Write empty file
	file.WriteString("{}")

	// Close file
	file.Close()

	// Return err
	return err

}

// ForgetMe forgets the DID
func (s *Server) ForgetMe(context.Context, *intern.Empty) (*intern.Success, error) {

	// Log
	log.Print("gRPC : Server.ForgetMe called")

	// Clean DIDs
	err := wipeStore()

	// Handle error
	if err != nil {
		log.Fatal(err)
	}

	// Success message
	response := intern.Success{}

	// Return
	return &response, nil
}

// GetDID gets new DID
// func (s *Server) GetDID(context.Context, *intern.Empty) (*intern.DID, error) {

// }

// // GetEvents returns list of random events
// func (s *Server) GetEvents(context.Context, *intern.Empty) (*intern.Events, error) {}

// // CheckTickets checks there are tickets left
// func (s *Server) CheckTickets(context.Context, *intern.EventIDs) (*intern.Tickets, error) {}

// // RequestTicket gets the QR data
// func (s *Server) RequestTicket(context.Context, *intern.QRCodeRequest) (*intern.QRCodeResponse, error) {
// }
