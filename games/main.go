package games

import "database/sql"

type Game struct {
	Command interface{
		Get(string) error
		Del(string) error
	}
	db *sql.DB
}

func (g *Game) Get(string string) error {
	return g.Command.Get(string)
}
