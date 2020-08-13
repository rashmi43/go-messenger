package router

import (
	"github.com/gorilla/mux"
	c "github.com/rashmi43/go-messenger/controller"
)

// SetMessageRoutes registers routes for messages
func SetMessageRoutes(messageController *c.MessageController) *mux.Router {
	messageRouter := mux.NewRouter()

	messageRouter.Handle("/v1/messages", c.ResponseHandler(messageController.Post)).Methods("POST")
	messageRouter.Handle("/v1/messages", c.ResponseHandler(messageController.Get)).Methods("GET")
	messageRouter.Handle("/v1/messages/{id}", c.ResponseHandler(messageController.Put)).Methods("PUT")
	messageRouter.Handle("/v1/messages/{id}", c.ResponseHandler(messageController.Delete)).Methods("DELETE")
	messageRouter.Handle("/v1/messages/{id}", c.ResponseHandler(messageController.GetById)).Methods("GET")
	// Applying authorization middleware
	//router.PathPrefix("/v1/messages").Handler(auth.AuthorizeRequest(messageRouter))
	return messageRouter
}
