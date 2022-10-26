package eventos

import (
	"fmt"
	"net/http"
	"sync"
)

type EventMessage struct {
	EventName string
	Data      any
}
type HandlerEvent struct {
	m        sync.Mutex
	clientes map[string]*cliente
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clientes: make(map[string]*cliente),
	}
}

func (h *HandlerEvent) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Println("Error ID")
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := newCliente(id)
	h.register(c)
	fmt.Println("Connected:", id)
	c.OnLine(r.Context(), w, flusher)
	fmt.Println("Desconected:", id)
	h.removeCliente(id)
}
func (h *HandlerEvent) register(c *cliente) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clientes[c.ID] = c
}
func (h *HandlerEvent) removeCliente(id string) {
	h.m.Lock()
	delete(h.clientes, id)
	h.m.Unlock()
}

func (h *HandlerEvent) Broadcast(m EventMessage) {
	h.m.Lock()
	defer h.m.Unlock()
	for _, c := range h.clientes {
		c.sendMessage <- m
	}
}
