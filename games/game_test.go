package games

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGameMapper_GetGame(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"group_id", "game_id"}).AddRow("aaa","bbbb")

	mock.ExpectQuery("SELECT group_id, game_id FROM kv_group WHERE game_id=1").WithArgs("1").WillReturnRows(rows)

	gm := GameMapper{}
	gm.GetGame("1")
}
