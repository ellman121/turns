package models

import (
	"fmt"
	"math/rand"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// SessionCache - Global in-memory store for active sessions
var sessionCache *gc.Cache

// Session an active game session
type Session struct {
	ID string
}

// NewSession - Return a new session
func NewSession() (*Session, error) {
	id := ""

	for {
		id = newSessionID()
		_, found := sessionCache.Get(id)
		if !found {
			break
		}
	}

	s := &Session{
		ID: id,
	}
	sessionCache.Set(id, s, 0)

	return &Session{
		ID: id,
	}, nil
}

// GetSession - Return an existing session
func GetSession(id string) (*Session, error) {
	if i, found := sessionCache.Get(id); found {
		s := i.(*Session)
		return s, nil
	}
	return nil, fmt.Errorf("Error: Session with id `%s` not found in the store", id)
}

func init() {
	// create the in-memory session store
	sessionCache = gc.New(1*time.Hour, 5*time.Minute)

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

func newSessionID() string {
	// I could make a 6 character string to allow for leading 0s, but it's
	// frankly just easier to  make the minimum 100000 and generate the rest
	return fmt.Sprintf("%d", rand.Intn(900000)+100000)
}
