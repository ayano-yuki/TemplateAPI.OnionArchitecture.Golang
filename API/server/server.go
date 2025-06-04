package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"

	"templateapi/onionarchitecture/golang/config"
	"templateapi/onionarchitecture/golang/infra/command"
	"templateapi/onionarchitecture/golang/infra/query"
	"templateapi/onionarchitecture/golang/ui/handler"
	"templateapi/onionarchitecture/golang/usecase"
)

type Server struct {
	router http.Handler
	config *config.Config
}

func NewServer() *Server {
	cfg := config.Load()

	db, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		panic(err)
	}

	healthQuery := query.NewHealthQuery(db)
	healthCommand := command.NewHealthCommand(db)

	healthRepo := struct {
		*query.HealthQuery
		*command.HealthCommand
	}{
		healthQuery,
		healthCommand,
	}

	usecase := usecase.NewHealthUsecase(&healthRepo)
	healthHandler := handler.NewHealthHandler(usecase)

	r := chi.NewRouter()
	r.Route("/health", func(r chi.Router) {
		r.Get("/api", healthHandler.HealthAPI)
		r.Get("/db1", healthHandler.HealthDB1)
		r.Get("/db2", healthHandler.HealthDB2)
	})

	return &Server{router: r, config: cfg}
}

func (s *Server) Start() {
	http.ListenAndServe(s.config.Addr(), s.router)
}
