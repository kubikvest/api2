package games

import (
	"database/sql"
	"fmt"
)

type GameMapper struct {
	DB *sql.DB
}

type Game struct {
	GroupId string
	GameId  string
}

type GameRepo interface {
	Get(uuid string) (*Game, error)
}

func (gm *GameMapper) GetGame(uuid string) (*Game, error) {
	g := &Game{}
	rows, err := gm.DB.Query("SELECT group_id, game_id FROM kv_group WHERE game_id=? limit 1", uuid)
	if err != nil {
		return g, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&g.GroupId, &g.GameId); err != nil {
			return nil, err
		}
	}

	fmt.Println(g)

	return g, nil
}
