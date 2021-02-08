package v1

import (
	"github.com/almendrosl/go-postgres-microblog/internal/data"
	"github.com/go-chi/chi"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/posts", pr.Routes())

	return r
}
