package cors

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func WithCors(server *rest.Server) {
	rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
	}, "*")(server)
}
