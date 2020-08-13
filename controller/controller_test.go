package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestGetMessages test HTTP Get to "/v1/messages" using ResponseRecorder
func TestGetMessages(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/messages", ResponseHandler(MessageController.Get)).Methods("GET")
	//r.HandleFunc("/v1/messages", getMessages).Methods("GET")
	req, err := http.NewRequest("GET", "/v1/messages", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

// TestGetMessagesWithServer test HTTP Get to "/v1/messages" using Server
func TestGetMessagesWithServer(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/messages", ResponseHandler(MessageController.Get)).Methods("GET")
	//r.HandleFunc("/v1/messages", getMessages).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	messagesURL := fmt.Sprintf("%s/v1/messages", server.URL)
	request, err := http.NewRequest("GET", messagesURL, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}

