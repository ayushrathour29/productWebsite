import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import ProductCard from '../components/ProductCard';

const HomePage = () => {
  const [featuredProducts, setFeaturedProducts] = useState([]);

  useEffect(() => {
    fetchFeaturedProducts();
  }, []);

  const fetchFeaturedProducts = async () => {
    try {
      const response = await axios.get('/api/products');
      // Get first 6 products as featured
      setFeaturedProducts(response.data.slice(0, 6));
    } catch (error) {
      console.error('Error fetching featured products:', error);
    }
  };

  return (
    <div>
      {/* Hero Section */}
      <section className="hero">
        <div className="hero-content">
          <h1>Timeless Elegance</h1>
          <p>Discover our exquisite collection of handcrafted jewelry, where tradition meets modern sophistication. Each piece tells a story of luxury and craftsmanship.</p>
          <div style={{ display: 'flex', gap: '1rem', justifyContent: 'center', flexWrap: 'wrap' }}>
            <Link to="/collection" className="btn">Explore Collection</Link>
            <Link to="/price-calculator" className="btn btn-secondary">Price Calculator</Link>
          </div>
        </div>
      </section>

      {/* Featured Products */}
      <section className="featured-section">
        <div className="container">
          <h2 className="section-title">Featured Collection</h2>
          <div className="products-grid">
            {featuredProducts.map(product => (
              <ProductCard key={product.id} product={product} />
            ))}
          </div>
          <div style={{ textAlign: 'center', marginTop: '3rem' }}>
            <Link to="/collection" className="btn">View All Products</Link>
          </div>
        </div>
      </section>

      {/* About Section */}
      <section style={{ padding: '4rem 0', backgroundColor: '#f8f8f8' }}>
        <div className="container">
          <div style={{ textAlign: 'center', maxWidth: '800px', margin: '0 auto' }}>
            <h2 className="section-title">Crafting Excellence Since 1985</h2>
            <p style={{ fontSize: '1.2rem', lineHeight: '1.8', color: '#666' }}>
              For over three decades, LuxeJewels has been creating exceptional jewelry pieces that celebrate life's most precious moments. 
              Our master craftsmen combine traditional techniques with contemporary design, ensuring each piece is a work of art that will 
              be treasured for generations.
            </p>
          </div>
        </div>
      </section>
    </div>
  );
};

export default HomePage;
