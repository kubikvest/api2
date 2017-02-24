package games

import (
	"net/http"
	"github.com/kubikvest/api2/app"
)

func SetHandlers(mux *http.ServeMux, ctx *app.Context){
	mux.Handle("/game/create", app.Handler{ctx, Create})
}

func Create(a *app.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}
