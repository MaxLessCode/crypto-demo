package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"crypto-demo/assets"
	"crypto-demo/internal/repository"
	"crypto-demo/ui/pages"

	"github.com/joho/godotenv"
)

func main() {
	InitDotEnv()

	mux := http.NewServeMux()
	repo := repository.NewInMemoryCryptoRepo()

	SetupAssetsRoutes(mux)
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		portfolio := repo.GetPortfolio()
		assets := repo.GetAssets()

		pages.Dashboard(portfolio, assets).Render(r.Context(), w)
	})
	// 1. Route pour traiter l'achat
	// 1. Route GET : Pour OUVRIR la modale (quand on clique sur "Négocier" dans le tableau)
	mux.HandleFunc("GET /trade/{symbol}", func(w http.ResponseWriter, r *http.Request) {
		symbol := r.PathValue("symbol")

		asset, err := repo.GetAssetBySymbol(symbol)
		if err != nil {
			http.Error(w, "Crypto introuvable", http.StatusNotFound)
			return
		}

		// Renvoie le HTML de la modale
		pages.TradeModal(asset).Render(r.Context(), w)
	})

	// 2. Route POST : Pour TRAITER l'achat (quand on clique sur "Confirmer l'ordre" dans la modale)
	mux.HandleFunc("POST /trade/{symbol}", func(w http.ResponseWriter, r *http.Request) {
		symbol := r.PathValue("symbol")

		// Récupération de la quantité
		amountStr := r.FormValue("amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			log.Println("⚠️ Erreur de saisie (utilise un point au lieu d'une virgule) :", amountStr)
		}

		log.Printf("✅ Transaction validée : Achat de %.2f %s\n", amount, symbol)

		// Exécution de la transaction dans notre fausse base de données
		repo.ExecuteTrade(symbol, amount)

		// Déclenche le refresh du dashboard via l'événement HTMX
		w.Header().Set("HX-Trigger", "tradeCompleted")

		// Toast de succès envoyé via OOB swap (hx-swap-oob) — HTMX l'insère dans #toast-container
		pages.TradeSuccessToast(symbol, amount).Render(r.Context(), w)
	})

	// 2. Route pour recharger uniquement les données du Dashboard
	mux.HandleFunc("GET /dashboard/data", func(w http.ResponseWriter, r *http.Request) {
		portfolio := repo.GetPortfolio()
		assets := repo.GetAssets()
		// On ne rend QUE le fragment intérieur, sans le layout global !
		pages.DashboardContent(portfolio, assets).Render(r.Context(), w)
	})

	fmt.Println("Server is running on http://localhost:8090")
	http.ListenAndServe(":8090", mux)
}

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func SetupAssetsRoutes(mux *http.ServeMux) {
	var isDevelopment = os.Getenv("GO_ENV") != "production"

	// On définit le système de fichiers une seule fois
	var root http.FileSystem
	if isDevelopment {
		root = http.Dir("./assets")
	} else {
		root = http.FS(assets.Assets)
	}

	fs := http.FileServer(root)

	// On crée le handler final
	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		// Laisse FileServer gérer le Content-Type tout seul
		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
