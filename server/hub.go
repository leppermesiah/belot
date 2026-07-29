package server

import "sync"

type Hub struct {
	mu    sync.Mutex
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{Rooms: map[string]*Room{}}
}

func (h *Hub) Create(targetScore int) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()
	r := NewRoom(targetScore)
	for h.Rooms[r.Code] != nil { // extremely unlikely collision, but be safe
		r = NewRoom(targetScore)
	}
	h.Rooms[r.Code] = r
	return r
}

func (h *Hub) Get(code string) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.Rooms[code]
}
