import React, { useState, useEffect } from 'react';
import axios from 'axios';

const ProductCard = ({ product }) => {
  const [selectedColor, setSelectedColor] = useState(product.colors[0]);
  const [currentImage, setCurrentImage] = useState(product.images[selectedColor]);
  const [showBreakdown, setShowBreakdown] = useState(false);
  const [priceBreakdown, setPriceBreakdown] = useState(null);
  const [loadingBreakdown, setLoadingBreakdown] = useState(false);
  const [imageError, setImageError] = useState(false);

  const formatINR = new Intl.NumberFormat('en-IN', { style: 'currency', currency: 'INR' });

  useEffect(() => {
    setCurrentImage(product.images[selectedColor]);
    setImageError(false); // Reset error state when image changes
  }, [selectedColor, product.images]);

  const fetchPriceBreakdown = async () => {
    if (priceBreakdown) {
      setShowBreakdown(!showBreakdown);
      return;
    }

    setLoadingBreakdown(true);
    try {
      const response = await axios.get(`/api/products/${product.id}/price-breakdown`);
      setPriceBreakdown(response.data);
      setShowBreakdown(true);
    } catch (error) {
      console.error('Error fetching price breakdown:', error);
      alert('Error fetching price breakdown');
    } finally {
      setLoadingBreakdown(false);
    }
  };

  const handleImageError = () => {
    setImageError(true);
  };

  const fallbackImage = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='400' height='400' viewBox='0 0 400 400'%3E%3Crect width='400' height='400' fill='%23f0f0f0'/%3E%3Ctext x='200' y='200' text-anchor='middle' dy='.3em' font-family='Arial, sans-serif' font-size='16' fill='%23999'%3EImage not available%3C/text%3E%3C/svg%3E";

  return (
    <div className="product-card">
      <img 
        src={imageError ? fallbackImage : currentImage} 
        alt={product.name}
        className="product-image"
        onError={handleImageError}
      />
      <div className="product-info">
        <h3 className="product-name">{product.name}</h3>
        <p className="product-description">{product.description}</p>
        <div className="product-price">{formatINR.format(product.current_price)}</div>
        <div className="product-weight">Gold Weight: {product.gold_weight}g</div>
        
        <div className="color-swatches">
          {product.colors.map((color) => (
            <div
              key={color}
              className={`color-swatch ${color} ${selectedColor === color ? 'active' : ''}`}
              onClick={() => setSelectedColor(color)}
              title={color.replace('-', ' ')}
            />
          ))}
        </div>
        
        <button 
          className="btn btn-secondary" 
          onClick={fetchPriceBreakdown}
          disabled={loadingBreakdown}
        >
          {loadingBreakdown ? 'Loading...' : 'Price Breakdown'}
        </button>
        
        <button className="btn">Add to Cart</button>

        {showBreakdown && priceBreakdown && (
          <div className="price-breakdown">
            <h4>Price Breakdown</h4>
            <div className="breakdown-item">
              <span>Gold Cost ({priceBreakdown.gold_weight}g × {formatINR.format(priceBreakdown.gold_price_per_gram)}):</span>
              <span>{formatINR.format(priceBreakdown.gold_cost)}</span>
            </div>
            <div className="breakdown-item">
              <span>Labor Cost:</span>
              <span>{formatINR.format(priceBreakdown.labor_cost)}</span>
            </div>
            <div className="breakdown-item">
              <span>Subtotal:</span>
              <span>{formatINR.format(priceBreakdown.subtotal)}</span>
            </div>
            <div className="breakdown-item">
              <span>Profit ({(priceBreakdown.profit_margin * 100).toFixed(1)}%):</span>
              <span>{formatINR.format(priceBreakdown.profit_amount)}</span>
            </div>
            <div className="breakdown-item total">
              <span>Final Price:</span>
              <span>{formatINR.format(priceBreakdown.final_price)}</span>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default ProductCard;
