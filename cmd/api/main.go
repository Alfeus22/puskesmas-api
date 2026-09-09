package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/Alfeus22/puskesmas-api/graph"
	authMid "github.com/Alfeus22/puskesmas-api/internal/middleware"
	"github.com/Alfeus22/puskesmas-api/internal/service"
	"github.com/Alfeus22/puskesmas-api/internal/storage"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("gagal memuat file .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	// 1. Koneksi Database

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal Koneksi database: %v", err)
	}
	defer db.Close()

	// 2. Dependency Injection (Merakit Layer)
	pasienStorage := storage.NewPasienStorage(db)
	pasienService := service.NewPasienService(pasienStorage)

	// 3. Setup Router menggunakan Chi
	r := chi.NewRouter()

	// Memasang middleware bawaan chi
	r.Use(middleware.Logger)    // Mencatat log request di terminal
	r.Use(middleware.Recoverer) // Mencegah server mati jika panic/bug
	r.Use(authMid.AuthMiddleware())

	// 5. RUTE GRAPHQL (Baru)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			PasienService: pasienService,
		},
	}))

	// Gunakan r.Handle untuk API GraphQL (Tempat data lewat)
	r.Handle("/query", srv)
	// Gunakan r.Get khusus untuk memanggil UI Playground (Halaman Web)
	r.Get("/", playground.Handler("GraphQL playground", "/query"))

	// 6. Jalankan Server
	log.Print("🏥 Server Puskesmas API berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server Mati : %v", err)
	}
}
