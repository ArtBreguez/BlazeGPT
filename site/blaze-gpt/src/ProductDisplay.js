import React from 'react';
import './ProductDisplay.css';

const ProductDisplay = (props) => {
  return (
    <div className="product-display">
      <h1 className="product-name">{props.name}</h1>
      <p className="product-description">{props.description}</p>
      <a href={props.link} className="https://go.hotmart.com/T79295798B">Ver Produto</a>
    </div>
  );
}

export default ProductDisplay;
