import React from 'react';
import './ProductDisplay.css';

const ProductDisplay = (props) => {
return (
  <div className="product-display">
  <h1 className="product-display__title">{props.name}</h1>
  <p className="product-display__description">{props.description}</p>
  <a href={props.link} target="_blank" className="product-display__link btn">Assine Já!</a>
  </div>
);
};

export default ProductDisplay;

