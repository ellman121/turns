package models

import (
	"fmt"
	"math/rand"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// SessionCache - Global in-memory store for active sessions
var cache *gc.Cache

type state struct {
	ScoreA int `json:"scoreA"`
	ScoreB int `json:"scoreB"`
}

// Game an active game session
type Game struct {
	ID        string
	GameState state `json:"state"`
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
	}
	cache.Set(id, s, 0)

	return &Game{
		ID: id,
		GameState: state{
			ScoreA: 0,
			ScoreB: 0,
		},
	}, nil
}

// GetGame - Return an existing session
func GetGame(id string) (*Game, error) {
	if i, found := cache.Get(id); found {
		s := i.(*Game)
		return s, nil
	}
	return nil, fmt.Errorf("Error: Session with id `%s` not found in the store", id)
}

func init() {
	// Initialize the RNG
	rand.Seed(0)
	// rand.Seed(time.Now().UnixNano())

	// create the in-memory session store
	cache = gc.New(1*time.Hour, 5*time.Minute)

	// Debug Setup

	// sessionCache = gc.New(5*time.Second, 1*time.Second)
	// tick := time.NewTicker(2 * time.Second)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-tick.C:
	// 			log.Println(fmt.Sprintf("Items in cache: %d", sessionCache.ItemCount()))
	// 		}
	// 	}
	// }()
}
