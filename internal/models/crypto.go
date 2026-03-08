package models

// Asset représente une cryptomonnaie sur le marché.
type Asset struct {
	ID        string
	Name      string
	Symbol    string
	Price     float64
	Change24h float64 // Variation sur 24h en pourcentage
	LogoURL   string  // URL vers le logo (SVG ou WebP)
	Amount    float64
}

// Portfolio représente le solde global de l'utilisateur.
type Portfolio struct {
	TotalBalance    float64
	TotalPNL        float64 // Profit & Loss (Gains/Pertes en valeur)
	TotalPNLPercent float64 // Gains/Pertes en pourcentage
}

// Transaction représente l'historique des opérations.
type Transaction struct {
	ID     string
	Date   string
	Type   string // "Achat" ou "Vente"
	Symbol string // Ex: "BTC"
	Amount float64
	Price  float64
	Total  float64
}
