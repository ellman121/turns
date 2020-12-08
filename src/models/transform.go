package models

import "fmt"

// A Transform is a received message from a client.  These are processed using the Process
// Transform method
type Transform struct {
	Instruction string `json:"instruction"`
	GameID      string `json:"gameID"`
	PlayerID    string `json:"playerID"`
	Value       int    `json:"value"`
}

// ProcessTransform - apply a Transform to a game state.  Returns the updated game state
func ProcessTransform(t Transform) (*Game, error) {
	g, err := GetGame(t.GameID)
	if err != nil {
		return nil, err
	}

	// Check the player IDs
	playerIndex := 0
	if t.PlayerID == g.Players[1].ID {
		playerIndex = 1
	} else if t.PlayerID != g.Players[0].ID {
		return nil, fmt.Errorf("[ProcessTransform] Invalid player ID sent with transform")
	}

	// Check the instrucitons
	switch t.Instruction {
	case "Reset":
		g.Players[0].Score = t.Value
		g.Players[1].Score = t.Value
	case "Apply":
		g.Players[playerIndex].Score += t.Value
	}

	cache.Set(t.GameID, g, 0)
	return g, nil
}
