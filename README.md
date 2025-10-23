# LuxeJewels - Jewelry Website

A modern jewelry e-commerce website built with React frontend and Go backend, featuring dynamic pricing based on gold prices and color variations for products.

## Features

- **Homepage**: Hero section with featured products
- **Collection Page**: Complete product catalog with 20+ jewelry pieces
- **Dynamic Price Calculator**: Real-time price updates based on gold prices
- **Color Swatches**: Interactive color selection for each product
- **Responsive Design**: Mobile-friendly interface
- **Modern UI**: Elegant jewelry-themed design

## Tech Stack

- **Frontend**: React 18, React Router, Axios
- **Backend**: Go, Gorilla Mux, CORS support
- **Styling**: Custom CSS with responsive design

## Project Structure

```
JewlleryWebsite/
├── backend/
│   ├── main.go          # Go backend server
│   └── go.mod           # Go dependencies
├── frontend/
│   ├── public/
│   │   ├── index.html
│   │   └── manifest.json
│   ├── src/
│   │   ├── components/
│   │   │   ├── Header.js
│   │   │   ├── Footer.js
│   │   │   ├── ProductCard.js
│   │   │   └── PriceCalculator.js
│   │   ├── pages/
│   │   │   ├── HomePage.js
│   │   │   └── CollectionPage.js
│   │   ├── App.js
│   │   ├── App.css
│   │   ├── index.js
│   │   └── index.css
│   └── package.json
└── README.md
```

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- Node.js 16 or higher
- npm or yarn

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install Go dependencies:
```bash
go mod tidy
```

3. Run the backend server:
```bash
go run main.go
```

The backend will start on `http://localhost:8080`

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install npm dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm start
```

The frontend will start on `http://localhost:3000`

## API Endpoints

- `GET /api/products` - Get all products with current prices
- `GET /api/products/{id}` - Get specific product
- `GET /api/price-config` - Get current price configuration
- `POST /api/price-config` - Update price configuration

## Key Features Explained

### Dynamic Price Calculator
The backend calculates product prices based on:
- Current gold price per gram
- Labor cost per piece
- Profit margin percentage

When gold prices change, simply update the configuration through the Price Calculator page, and all product prices will automatically update across the site.

### Color Swatches
Each product supports multiple color variations:
- Yellow Gold
- White Gold  
- Rose Gold

Clicking a color swatch updates the product image to show the selected color variant.

### Product Categories
Products are organized into categories:
- Rings
- Earrings
- Necklaces
- Bracelets

Use the filter buttons on the Collection page to browse by category.

## Usage

1. **Homepage**: View featured products and navigate to different sections
2. **Collection**: Browse all products, filter by category, and interact with color swatches
3. **Price Calculator**: Update gold prices and calculate custom pricing

## Customization

### Adding New Products
Edit the `products` slice in `backend/main.go` to add new jewelry pieces.

### Updating Gold Prices
Use the Price Calculator page in the frontend to update gold prices, or directly modify the `priceConfig` in the backend.

### Styling
Modify `frontend/src/App.css` to customize the appearance and colors.

## Production Deployment

1. Build the React frontend:
```bash
cd frontend
npm run build
```

2. The backend serves the built frontend files from the `build` directory

3. Deploy the Go backend to your preferred hosting platform

## License

This project is created for demonstration purposes.
