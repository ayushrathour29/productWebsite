package main

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Product struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	BasePrice   float64           `json:"base_price"`
	GoldWeight  float64           `json:"gold_weight"` // in grams
	Category    string            `json:"category"`
	Colors      []string          `json:"colors"`
	Images      map[string]string `json:"images"` // color -> image URL
}

type PriceConfig struct {
	GoldPricePerGram float64 `json:"gold_price_per_gram"`
	LaborCost        float64 `json:"labor_cost"`
	ProfitMargin     float64 `json:"profit_margin"`
}

type ProductWithPrice struct {
	Product
	CurrentPrice float64 `json:"current_price"`
}

var products = []Product{
	{1, "Classic Gold Ring", "Elegant 18k gold ring with intricate design", 150.0, 5.0, "Rings", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1605100804763-247f67b3557e?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
	{2, "Diamond Earrings", "Stunning diamond stud earrings", 200.0, 3.0, "Earrings", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1535632066927-ab7c9ab60908?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1506630448388-4e683c67ddb0?w=400",
	}},
	{3, "Pearl Necklace", "Luxurious pearl necklace with gold clasp", 180.0, 8.0, "Necklaces", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1596944924616-7b384c1a7b75?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1596944924616-7b384c1a7b75?w=400",
	}},
	{4, "Emerald Bracelet", "Beautiful emerald bracelet with gold setting", 220.0, 6.0, "Bracelets", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
	{5, "Sapphire Ring", "Exquisite sapphire ring with diamond accents", 300.0, 4.5, "Rings", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1605100804763-247f67b3557e?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
	}},
	{6, "Gold Chain", "Classic gold chain necklace", 120.0, 10.0, "Necklaces", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1573408301185-9146fe634ad0?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1573408301185-9146fe634ad0?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1573408301185-9146fe634ad0?w=400",
	}},
	{7, "Ruby Earrings", "Dazzling ruby drop earrings", 250.0, 3.5, "Earrings", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1535632066927-ab7c9ab60908?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1506630448388-4e683c67ddb0?w=400",
	}},
	{8, "Diamond Bracelet", "Elegant diamond tennis bracelet", 400.0, 7.0, "Bracelets", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
	}},
	{9, "Opal Ring", "Mystical opal ring with gold band", 160.0, 4.0, "Rings", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1605100804763-247f67b3557e?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
	{10, "Pearl Earrings", "Classic pearl stud earrings", 140.0, 2.5, "Earrings", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1535632066927-ab7c9ab60908?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1506630448388-4e683c67ddb0?w=400",
	}},
	{11, "Amethyst Necklace", "Violet amethyst pendant necklace", 190.0, 5.5, "Necklaces", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1617038220319-276d4f445e83?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1617038220319-276d4f445e83?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1617038220319-276d4f445e83?w=400",
	}},
	{12, "Citrine Bracelet", "Sunny citrine charm bracelet", 170.0, 6.5, "Bracelets", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
	}},
	{13, "Topaz Ring", "Brilliant blue topaz ring", 210.0, 3.8, "Rings", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1605100804763-247f67b3557e?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
	{14, "Garnet Earrings", "Deep red garnet earrings", 180.0, 3.2, "Earrings", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1535632066927-ab7c9ab60908?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1506630448388-4e683c67ddb0?w=400",
	}},
	{15, "Aquamarine Necklace", "Ocean blue aquamarine necklace", 230.0, 7.2, "Necklaces", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1583394838336-acd977736f90?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1583394838336-acd977736f90?w=400",
	}},
	{16, "Peridot Bracelet", "Fresh green peridot bracelet", 200.0, 5.8, "Bracelets", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
	{17, "Moonstone Ring", "Mystical moonstone ring", 175.0, 4.2, "Rings", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1605100804763-247f67b3557e?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
	}},
	{18, "Tourmaline Earrings", "Rainbow tourmaline earrings", 220.0, 3.6, "Earrings", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1535632066927-ab7c9ab60908?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1506630448388-4e683c67ddb0?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
	{19, "Zircon Necklace", "Brilliant zircon pendant", 195.0, 6.8, "Necklaces", []string{"yellow-gold", "white-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400",
	}},
	{20, "Spinel Bracelet", "Rare spinel gemstone bracelet", 280.0, 8.5, "Bracelets", []string{"yellow-gold", "white-gold", "rose-gold"}, map[string]string{
		"yellow-gold": "https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=400",
		"white-gold":  "https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=400",
		"rose-gold":   "https://images.unsplash.com/photo-1599643478518-a784e5dc4c8f?w=400",
	}},
}

var priceConfig = PriceConfig{
	GoldPricePerGram: 100.0, // Current gold price per gram
	LaborCost:        50.0,  // Labor cost per piece
	ProfitMargin:     0.3,   // 30% profit margin
}

func calculatePrice(product Product) float64 {
	goldCost := product.GoldWeight * priceConfig.GoldPricePerGram
	totalCost := goldCost + priceConfig.LaborCost
	profit := totalCost * priceConfig.ProfitMargin
	return totalCost + profit
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var productsWithPrice []ProductWithPrice
	for _, product := range products {
		productsWithPrice = append(productsWithPrice, ProductWithPrice{
			Product:      product,
			CurrentPrice: calculatePrice(product),
		})
	}

	json.NewEncoder(w).Encode(productsWithPrice)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			productWithPrice := ProductWithPrice{
				Product:      product,
				CurrentPrice: calculatePrice(product),
			}
			json.NewEncoder(w).Encode(productWithPrice)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateGoldPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newConfig PriceConfig
	err := json.NewDecoder(r.Body).Decode(&newConfig)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate the configuration
	if newConfig.GoldPricePerGram <= 0 {
		http.Error(w, "Gold price per gram must be greater than 0", http.StatusBadRequest)
		return
	}
	if newConfig.LaborCost < 0 {
		http.Error(w, "Labor cost cannot be negative", http.StatusBadRequest)
		return
	}
	if newConfig.ProfitMargin < 0 || newConfig.ProfitMargin > 1 {
		http.Error(w, "Profit margin must be between 0 and 1", http.StatusBadRequest)
		return
	}

	priceConfig = newConfig

	// Calculate updated prices for all products
	var updatedProducts []ProductWithPrice
	for _, product := range products {
		updatedProducts = append(updatedProducts, ProductWithPrice{
			Product:      product,
			CurrentPrice: calculatePrice(product),
		})
	}

	response := map[string]interface{}{
		"message":  "Gold price updated successfully",
		"config":   priceConfig,
		"products": updatedProducts,
	}

	json.NewEncoder(w).Encode(response)
}

func getPriceConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(priceConfig)
}

func getPriceBreakdown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			goldCost := product.GoldWeight * priceConfig.GoldPricePerGram
			totalCost := goldCost + priceConfig.LaborCost
			profit := totalCost * priceConfig.ProfitMargin
			finalPrice := totalCost + profit

			breakdown := map[string]interface{}{
				"product_id":          product.ID,
				"product_name":        product.Name,
				"gold_weight":         product.GoldWeight,
				"gold_price_per_gram": priceConfig.GoldPricePerGram,
				"gold_cost":           goldCost,
				"labor_cost":          priceConfig.LaborCost,
				"subtotal":            totalCost,
				"profit_margin":       priceConfig.ProfitMargin,
				"profit_amount":       profit,
				"final_price":         finalPrice,
			}

			json.NewEncoder(w).Encode(breakdown)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products/{id}/price-breakdown", getPriceBreakdown).Methods("GET")
	r.HandleFunc("/api/price-config", getPriceConfig).Methods("GET")
	r.HandleFunc("/api/price-config", updateGoldPrice).Methods("POST", "OPTIONS")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/build/")))

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default for local development
		}
	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, corsHandler)

}
