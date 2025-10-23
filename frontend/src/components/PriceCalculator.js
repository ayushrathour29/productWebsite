import React, { useState, useEffect } from 'react';
import axios from 'axios';

const PriceCalculator = () => {
  const [priceConfig, setPriceConfig] = useState({
    gold_price_per_gram: 100.0,
    labor_cost: 50.0,
    profit_margin: 0.3
  });
  const [products, setProducts] = useState([]);
  const [selectedProduct, setSelectedProduct] = useState(null);
  const [customWeight, setCustomWeight] = useState(5.0);
  const [calculatedPrice, setCalculatedPrice] = useState(0);
  const formatINR = new Intl.NumberFormat('en-IN', { style: 'currency', currency: 'INR' });

  useEffect(() => {
    fetchPriceConfig();
    fetchProducts();
  }, []);

  const fetchPriceConfig = async () => {
    try {
      const response = await axios.get('/api/price-config');
      setPriceConfig(response.data);
    } catch (error) {
      console.error('Error fetching price config:', error);
    }
  };

  const fetchProducts = async () => {
    try {
      const response = await axios.get('/api/products');
      setProducts(response.data);
    } catch (error) {
      console.error('Error fetching products:', error);
    }
  };

  const updatePriceConfig = async () => {
    try {
      const response = await axios.post('/api/price-config', priceConfig);
      alert('Price configuration updated successfully!');
      // Update products with the new prices returned from the server
      setProducts(response.data.products);
      // Notify other components/pages to refresh (e.g., CollectionPage)
      window.dispatchEvent(new CustomEvent('price-config-updated'));
    } catch (error) {
      console.error('Error updating price config:', error);
      if (error.response && error.response.data) {
        alert(`Error: ${error.response.data}`);
      } else {
        alert('Error updating price configuration');
      }
    }
  };

  const calculatePrice = (weight) => {
    const goldCost = weight * priceConfig.gold_price_per_gram;
    const totalCost = goldCost + priceConfig.labor_cost;
    const profit = totalCost * priceConfig.profit_margin;
    return totalCost + profit;
  };

  const handleWeightChange = (weight) => {
    setCustomWeight(weight);
    setCalculatedPrice(calculatePrice(weight));
  };

  const handleProductSelect = (product) => {
    setSelectedProduct(product);
    setCustomWeight(product.gold_weight);
    setCalculatedPrice(product.current_price);
  };

  return (
    <div className="price-calculator">
      <div className="container">
        <h2 className="section-title">Price Calculator & Configuration</h2>
        
        <div className="calculator-container">
          <h3>Update Gold Price Configuration</h3>
          <div className="calculator-form">
            <div className="form-group">
              <label>Gold Price per Gram (₹)</label>
              <input
                type="number"
                step="0.01"
                value={priceConfig.gold_price_per_gram}
                onChange={(e) => setPriceConfig({
                  ...priceConfig,
                  gold_price_per_gram: parseFloat(e.target.value)
                })}
              />
            </div>
            <div className="form-group">
              <label>Labor Cost per Piece (₹)</label>
              <input
                type="number"
                step="0.01"
                value={priceConfig.labor_cost}
                onChange={(e) => setPriceConfig({
                  ...priceConfig,
                  labor_cost: parseFloat(e.target.value)
                })}
              />
            </div>
            <div className="form-group">
              <label>Profit Margin (%)</label>
              <input
                type="number"
                step="0.01"
                value={priceConfig.profit_margin * 100}
                onChange={(e) => setPriceConfig({
                  ...priceConfig,
                  profit_margin: parseFloat(e.target.value) / 100
                })}
              />
            </div>
            <button className="btn" onClick={updatePriceConfig}>
              Update Configuration
            </button>
          </div>
        </div>

        <div className="calculator-container" style={{ marginTop: '2rem' }}>
          <h3>Calculate Custom Price</h3>
          <div className="calculator-form">
            <div className="form-group">
              <label>Select Product (for reference)</label>
              <select 
                value={selectedProduct?.id || ''} 
                onChange={(e) => {
                  const product = products.find(p => p.id === parseInt(e.target.value));
                  if (product) handleProductSelect(product);
                }}
              >
                <option value="">Choose a product...</option>
                {products.map(product => (
                  <option key={product.id} value={product.id}>
                    {product.name} ({product.gold_weight}g)
                  </option>
                ))}
              </select>
            </div>
            <div className="form-group">
              <label>Gold Weight (grams)</label>
              <input
                type="number"
                step="0.1"
                value={customWeight}
                onChange={(e) => handleWeightChange(parseFloat(e.target.value))}
              />
            </div>
            
            <div className="calculator-result">
              <h3>Calculated Price</h3>
              <div className="price-display">{formatINR.format(calculatedPrice)}</div>
              <p>Based on current gold price: {formatINR.format(priceConfig.gold_price_per_gram)} / gram</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default PriceCalculator;
