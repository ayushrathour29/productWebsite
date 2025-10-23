import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ProductCard from '../components/ProductCard';

const CollectionPage = () => {
  const [products, setProducts] = useState([]);
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState('All');
  const [loading, setLoading] = useState(true);

  const categories = ['All', 'Rings', 'Earrings', 'Necklaces', 'Bracelets'];

  useEffect(() => {
    fetchProducts();
  }, []);

  // Auto-refresh when price config is updated elsewhere
  useEffect(() => {
    const handler = () => fetchProducts();
    window.addEventListener('price-config-updated', handler);
    return () => window.removeEventListener('price-config-updated', handler);
  }, []);

  useEffect(() => {
    if (selectedCategory === 'All') {
      setFilteredProducts(products);
    } else {
      setFilteredProducts(products.filter(product => product.category === selectedCategory));
    }
  }, [selectedCategory, products]);

  const fetchProducts = async () => {
    try {
      setLoading(true);
      const response = await axios.get('/api/products');
      setProducts(response.data);
      setFilteredProducts(response.data);
    } catch (error) {
      console.error('Error fetching products:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      {/* Collection Header */}
      <section className="collection-header">
        <div className="container">
          <h1>Our Collection</h1>
          <p>Discover our complete range of exquisite jewelry pieces, each crafted with precision and passion.</p>
        </div>
      </section>

      {/* Collection Content */}
      <section className="collection-content">
        <div className="container">
          {/* Filter Section */}
          <div className="filter-section">
            <h3 style={{ marginBottom: '1rem', color: '#333' }}>Filter by Category</h3>
            <div className="filter-buttons">
              {categories.map(category => (
                <button
                  key={category}
                  className={`filter-btn ${selectedCategory === category ? 'active' : ''}`}
                  onClick={() => setSelectedCategory(category)}
                >
                  {category}
                </button>
              ))}
            </div>
          </div>

          {/* Products Grid */}
          {loading ? (
            <div style={{ textAlign: 'center', padding: '4rem' }}>
              <h3>Loading products...</h3>
            </div>
          ) : (
            <>
              <div style={{ marginBottom: '2rem', textAlign: 'center' }}>
                <p style={{ color: '#666', fontSize: '1.1rem' }}>
                  Showing {filteredProducts.length} products
                  {selectedCategory !== 'All' && ` in ${selectedCategory}`}
                </p>
              </div>
              
              <div className="products-grid">
                {filteredProducts.map(product => (
                  <ProductCard key={product.id} product={product} />
                ))}
              </div>

              {filteredProducts.length === 0 && (
                <div style={{ textAlign: 'center', padding: '4rem' }}>
                  <h3>No products found in this category</h3>
                  <p style={{ color: '#666' }}>Try selecting a different category or check back later.</p>
                </div>
              )}
            </>
          )}
        </div>
      </section>
    </div>
  );
};

export default CollectionPage;
