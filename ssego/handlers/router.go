package handlers

import (
	"encoding/json"
	"net/http"
	"ssego/handlers/eventos"
)

func InitRoutes(r *http.ServeMux) {
	handlerEvents := eventos.NewHandlerEvent()

	r.HandleFunc("/notify", handlerEvents.Handler)
	r.HandleFunc("/test1", HandlerTest1(handlerEvents))
	r.HandleFunc("/test2", HandlerTest2(handlerEvents))
	r.Handle("/", http.FileServer(http.Dir("./public")))
}

func HandlerTest1(notifier *eventos.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]any{}
		json.NewDecoder(r.Body).Decode(&data)
		notifier.Broadcast(eventos.EventMessage{
			EventName: "saludar",
			Data:      data,
		})
	}
}

func HandlerTest2(notifier *eventos.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]any{}
		json.NewDecoder(r.Body).Decode(&data)
		notifier.Broadcast(eventos.EventMessage{
			EventName: "saltar",
			Data:      data,
		})
	}
}
