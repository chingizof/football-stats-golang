package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chingizof/football-stats-golang/internal/models"
	"github.com/chingizof/football-stats-golang/internal/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store

	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
		store:       store,

		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	// REST
	// сущность/идентификатор

	// /electronics/laptops
	// /electronics/phones
	// /clothes/t-shirts
	r.Post("/market/players", func(w http.ResponseWriter, r *http.Request) {
		player := new(models.Player)
		if err := json.NewDecoder(r.Body).Decode(player); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Players().Create(r.Context(), player)
	})
	r.Get("/players", func(w http.ResponseWriter, r *http.Request) {
		players, err := s.store.Players().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, players)
	})
	r.Get("/players/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		player, err := s.store.Players().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, player)
	})
	r.Put("/players", func(w http.ResponseWriter, r *http.Request) {
		player := new(models.Player)
		if err := json.NewDecoder(r.Body).Decode(player); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Players().Update(r.Context(), player)
	})
	r.Delete("/players/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Players().Delete(r.Context(), id)
	})

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  5,
		WriteTimeout: 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}
