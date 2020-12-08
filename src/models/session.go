package models

import (
	"fmt"
	"math/rand"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// SessionCache - Global in-memory store for active sessions
var cache *gc.Cache

type playerState struct {
	ID    string
	Score int
}

// Game an active game session
type Game struct {
	ID      string        `json:",omitempty"`
	Players []playerState `json:"playerStates"`
}

// NewGame - Return a new session
func NewGame() (*Game, error) {
	id := ""

	for {
		id = fmt.Sprintf("%d", rand.Intn(900000)+100000)
		_, found := cache.Get(id)
		if !found {
			break
		}
	}

	s := &Game{
		ID: id,
		Players: []playerState{
			playerState{
				ID:    "",
				Score: 0,
			},
			playerState{
				ID:    "",
				Score: 0,
			},
		},
	}
	cache.Set(id, s, 0)

	return s, nil
}

// GetGame - Return an existing game
func GetGame(id string) (*Game, error) {
	if i, found := cache.Get(id); found {
		s := i.(*Game)
		return s, nil
	}
	return nil, fmt.Errorf("Game with id `%s` not found in the store", id)
}

func init() {
	// Initialize the RNG
	rand.Seed(0)
	// rand.Seed(time.Now().UnixNano())

	// create the in-memory game store
	cache = gc.New(1*time.Hour, 5*time.Minute)

	// Debug Setup

	// cache = gc.New(5*time.Second, 1*time.Second)
	// tick := time.NewTicker(2 * time.Second)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-tick.C:
	// 			log.Println(fmt.Sprintf("Items in cache: %d", cache.ItemCount()))
	// 		}
	// 	}
	// }()
}
