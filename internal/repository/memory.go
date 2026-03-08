package repository

import (
	"crypto-demo/internal/models"
	"errors"
)

// CryptoRepository définit le contrat de notre source de données.
type CryptoRepository interface {
	GetPortfolio() models.Portfolio
	GetAssets() []models.Asset
	GetTransactions() []models.Transaction
	GetAssetBySymbol(symbol string) (models.Asset, error)
	ExecuteTrade(symbol string, amount float64)
}

// InMemoryCryptoRepo implémente l'interface avec des données mockées.
type InMemoryCryptoRepo struct {
	portfolio    models.Portfolio
	assets       []models.Asset
	transactions []models.Transaction
}

// NewInMemoryCryptoRepo initialise notre mock avec de belles données de démonstration.
func NewInMemoryCryptoRepo() *InMemoryCryptoRepo {
	return &InMemoryCryptoRepo{
		// Un portefeuille attractif pour la démo
		portfolio: models.Portfolio{
			TotalBalance:    124530.50,
			TotalPNL:        14230.20,
			TotalPNLPercent: 12.8,
		},

		// Liste des cryptomonnaies principales
		assets: []models.Asset{
			{ID: "1", Name: "Bitcoin", Symbol: "BTC", Price: 64230.00, Change24h: 2.4, LogoURL: "https://cryptologos.cc/logos/bitcoin-btc-logo.svg", Amount: 0.15},
			{ID: "2", Name: "Ethereum", Symbol: "ETH", Price: 3450.75, Change24h: -1.2, LogoURL: "https://cryptologos.cc/logos/ethereum-eth-logo.svg", Amount: 2.5},
			{ID: "3", Name: "Solana", Symbol: "SOL", Price: 145.20, Change24h: 8.5, LogoURL: "https://cryptologos.cc/logos/solana-sol-logo.svg", Amount: 40.0},
			{ID: "4", Name: "Cardano", Symbol: "ADA", Price: 0.45, Change24h: -0.5, LogoURL: "https://cryptologos.cc/logos/cardano-ada-logo.svg", Amount: 1500.0},
		},

		// Un petit historique pour peupler un futur tableau
		transactions: []models.Transaction{
			{ID: "TX101", Date: "Aujourd'hui, 14:30", Type: "Achat", Symbol: "BTC", Amount: 0.15, Price: 64000.00, Total: 9600.00},
			{ID: "TX102", Date: "Hier, 09:15", Type: "Achat", Symbol: "SOL", Amount: 50.0, Price: 138.00, Total: 6900.00},
			{ID: "TX103", Date: "Il y a 3 jours", Type: "Vente", Symbol: "ETH", Amount: 2.5, Price: 3500.00, Total: 8750.00},
		},
	}
}

func (r *InMemoryCryptoRepo) GetAssetBySymbol(symbol string) (models.Asset, error) {
	for _, a := range r.assets {
		if a.Symbol == symbol {
			return a, nil
		}
	}
	return models.Asset{}, errors.New("actif non trouvé")
}

func (r *InMemoryCryptoRepo) ExecuteTrade(symbol string, amount float64) {
	for i := range r.assets {
		if r.assets[i].Symbol == symbol {
			totalCost := r.assets[i].Price * amount

			r.portfolio.TotalBalance -= totalCost
			r.portfolio.TotalPNL += (totalCost * 0.02)

			r.assets[i].Amount += amount
			break
		}
	}
}

// Implémentation des méthodes de l'interface
func (r *InMemoryCryptoRepo) GetPortfolio() models.Portfolio {
	return r.portfolio
}

func (r *InMemoryCryptoRepo) GetAssets() []models.Asset {
	return r.assets
}

func (r *InMemoryCryptoRepo) GetTransactions() []models.Transaction {
	return r.transactions
}
