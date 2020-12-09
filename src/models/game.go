package models

import (
	"fmt"
	"math/rand"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// cache - Global in-memory store for active games
var cache *gc.Cache

type playerState struct {
	ID    string
	Score int
}

// Game an active game
type Game struct {
	ID      string        `json:",omitempty"`
	Players []playerState `json:"playerStates"`
}

// NewGame - Return a new game
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

// AddPlayer - Add a player ID to the game and update self in the cache.
// If the game is full, return an error
func (g Game) AddPlayer(playerID string) error {
	if g.Players[0].ID == "" {
		g.Players[0].ID = playerID
		return nil
	} else if g.Players[1].ID == "" {
		g.Players[1].ID = playerID
		return nil
	}
	return fmt.Errorf("[Game.AddPlayer] Cannot add player to full game")
}

// Sanatized - A game returns a copy of itself (does not modify the actual game)
// with all player IDs replaced with '-' except for the ID passed in
func (g Game) Sanatized(keepID string) Game {
	temp := Game{}
	temp.ID = g.ID
	temp.Players = make([]playerState, len(g.Players))
	for i, p := range g.Players {
		temp.Players[i].Score = p.Score
		if p.ID != keepID {
			temp.Players[i].ID = "-"
		} else {
			temp.Players[i].ID = keepID
		}
	}
	return temp
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
