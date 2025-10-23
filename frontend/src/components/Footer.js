import React from 'react';

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-content">
        <div className="footer-section">
          <h3>LuxeJewels</h3>
          <p>Creating timeless pieces that celebrate life's most precious moments. Our handcrafted jewelry combines traditional techniques with modern design.</p>
        </div>
        <div className="footer-section">
          <h3>Quick Links</h3>
          <p><a href="/">Home</a></p>
          <p><a href="/collection">Collection</a></p>
          <p><a href="/price-calculator">Price Calculator</a></p>
        </div>
        <div className="footer-section">
          <h3>Contact</h3>
          <p>123 Jewelry Street</p>
          <p>New York, NY 10001</p>
          <p>Phone: (555) 123-4567</p>
          <p>Email: info@luxejewels.com</p>
        </div>
        <div className="footer-section">
          <h3>Follow Us</h3>
          <p><a href="#">Instagram</a></p>
          <p><a href="#">Facebook</a></p>
          <p><a href="#">Twitter</a></p>
        </div>
      </div>
      <div className="footer-bottom">
        <p>&copy; 2024 LuxeJewels. All rights reserved.</p>
      </div>
    </footer>
  );
};

export default Footer;
