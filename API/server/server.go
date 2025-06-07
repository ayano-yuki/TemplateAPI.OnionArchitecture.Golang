package server

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	_ "API/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"API/config"
	"API/infra"
	"API/ui/api/health"
	"API/usecase"
)

type Server struct {
	router *chi.Mux
	db     *sql.DB
}

func NewServer() *Server {
	// DB設定の読み込み
	dbCfg := config.LoadDBConfig()

	// DSN構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)

	// DB接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// リポジトリ統合実装
	repo := infra.NewHealthRepositoryImpl(db)

	// ユースケース作成
	healthUC := usecase.NewHealthUsecase(repo)

	// ハンドラ生成
	healthHandler := health.NewHandler(healthUC)

	// ルーター設定
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  log.New(os.Stdout, "", log.LstdFlags),
		NoColor: false,
	}))

	log.SetOutput(io.Discard)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Route("/health", func(r chi.Router) {
		r.Get("/api", healthHandler.CheckAPIConnection)
		r.Get("/db-query", healthHandler.GetDBTexts)
		r.Post("/db-command", healthHandler.InsertDBText)
	})

	return &Server{
		router: r,
		db:     db,
	}
}

func (s *Server) Start() {
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", s.router)
}
