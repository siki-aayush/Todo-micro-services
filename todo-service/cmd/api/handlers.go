package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"todoService/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JSONPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (app *Config) AddTodo(w http.ResponseWriter, r *http.Request) {

	var requestPayload JSONPayload
	_ = app.readJson(w, r, &requestPayload)

	event := data.TodoEntry{
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
	}

	documentId, err := app.Models.TodoEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	id := documentId.InsertedID.(primitive.ObjectID).Hex()
	err = app.logRequest("Todo Service", fmt.Sprintf("Todo created with ID: %s", id))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Todo created successfully",
		Data: data.TodoEntry{
			ID:          id,
			Name:        event.Name,
			Description: event.Description,
		},
	}

	app.writeJson(w, http.StatusAccepted, resp)
}

func (app *Config) GetAllTodo(w http.ResponseWriter, r *http.Request) {

	todos, err := app.Models.TodoEntry.All()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Data:  todos,
	}
	app.writeJson(w, http.StatusAccepted, resp)

}

type DeleteTodoRequest struct {
	ID string `json:"id"`
}

func (app *Config) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	var requestPayload DeleteTodoRequest
	_ = app.readJson(w, r, &requestPayload)

	event := data.TodoEntry{
		ID: requestPayload.ID,
	}

	err := app.Models.TodoEntry.Delete(event)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.logRequest("Todo Service", fmt.Sprintf("Todo deleted with ID: %s", event.ID))
	if err != nil {
		app.errorJson(w, err)
		return
	}
	resp := jsonResponse{
		Error:   false,
		Message: "Todo deleted successfully",
	}
	app.writeJson(w, http.StatusAccepted, resp)
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service:8082/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
