package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AuGhoti/goserver/endpoints/activity"
	"github.com/AuGhoti/goserver/endpoints/auth"
	"github.com/AuGhoti/goserver/endpoints/current"
	"github.com/AuGhoti/goserver/endpoints/historical"
	"github.com/AuGhoti/goserver/endpoints/stats"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initializeAPI() (*chi.Mux, *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.URLFormat,
		render.SetContentType(render.ContentTypeJSON),
	)

	router.Mount("/auth", auth.Routes())
	router.Mount("/v1", v1routes())

	return router, client
}

func v1routes() http.Handler {
	r := chi.NewRouter()
	r.Route("/historical", historical.Routes)
	r.Route("/current", current.Routes)
	r.Route("/activity", activity.Routes)
	r.Route("/stats", stats.Routes)
	return r
}

func main() {
	router, _ := initializeAPI()
	log.Fatal(http.ListenAndServe(":8080", router))
}
