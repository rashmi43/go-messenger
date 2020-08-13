package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rashmi43/go-messenger/domain"
	util "github.com/rashmi43/go-messenger/util"
	"net/http"
	//"log"
	"encoding/json"
	apputil "github.com/shijuvar/gokit/examples/http-app/pkg/apputil"
)

// response used to send HTTP responses
type response struct {
	Data interface{} `json:"data"`
}

// Generic handler for writing response header and body for all handler functions
func ResponseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r) // execute application handler
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if data != nil {
			// Send JSON response back to the client application
			err = json.NewEncoder(w).Encode(response{Data: data})
			if err != nil {
				apputil.Error.Printf("Error from Handler: %s\n", err.Error())
			}
		}

	})
}

type MessageController struct {
	// Explicit dependency and declarative programming that hides dependent logic
	Store domain.MessageStore
	// It can be any Store including MapStore
}

// HTTP Post - /v1/messages
func (ctl MessageController) Post(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)
	var message domain.Message
	err := decoder.Decode(&message)
	if err != nil {
                // not a good practise, panic is if we cannot continue at all
		panic(err)
	}
	err = ctl.Store.Create(message)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, http.StatusBadRequest, errors.Wrap(err, "Error on inserting Message")
	}
	fmt.Println("A new message has been created")
	return message, http.StatusCreated, nil

}

func (ctl MessageController) Get(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	mList, err := ctl.Store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, http.StatusBadRequest, errors.Wrap(err, "Error on getting Message")
	}
	fmt.Println("Message retrieved", mList)
	return mList, http.StatusOK, nil

}

func (ctl MessageController) GetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	mId := vars["id"]
	var message domain.Message
	fmt.Println("Fetch message id:", mId)
	message, err := ctl.Store.GetById(mId)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, http.StatusBadRequest, errors.Wrap(err, "Error on getting Message")
	}
	fmt.Println("Message retrieved", message.Text)
	isPalindrome := util.IsPalindrome(message.Text)
	if isPalindrome == true {
		fmt.Println("Message is a palindrome!!")
	}
	return message, http.StatusOK, nil

}

func (ctl MessageController) Put(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	mId := vars["id"]
	fmt.Println("Updating message id:", mId)
	decoder := json.NewDecoder(r.Body)
	var c domain.Message
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
        // Allow id update
        if (c.ID != mId) {
          fmt.Println("Id has changed, cannot update")
          return nil, http.StatusBadRequest, errors.Wrap(err, "ID has changed, cannot update")
        } else {
	  err = ctl.Store.Update(mId, c)
        }
	if err != nil {
		fmt.Println("Error:", err)
		return nil, http.StatusBadRequest, errors.Wrap(err, "Error updating message"+mId)
	}
	fmt.Println("Message has been updated", c)
	return c, http.StatusOK, nil

}

func (ctl MessageController) Delete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	mId := vars["id"]
	fmt.Println("Deleting message id:", mId)
	err := ctl.Store.Delete(mId)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, http.StatusBadRequest, errors.Wrap(err, "Error deleting message"+mId)
	}
	fmt.Println("Message has been deleted")
	return nil, http.StatusOK, nil

}
