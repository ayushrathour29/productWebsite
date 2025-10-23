import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <header className="header">
      <div className="header-container">
        <Link to="/" className="logo">
          LuxeJewels
        </Link>
        <nav>
          <ul className="nav">
            <li><Link to="/">Home</Link></li>
            <li><Link to="/collection">Collection</Link></li>
            <li><Link to="/price-calculator">Price Calculator</Link></li>
          </ul>
        </nav>
      </div>
    </header>
  );
};

export default Header;
