import React from 'react';
import './ProductDisplay.css';
import productImage from './phone.jpeg';

const ProductDisplay = (props) => {
  return (
    <div className="product-display">
      <h1 className="product-display__title">{props.name}</h1>
      <p className="product-display__description">{props.description}</p>
      <img src={productImage} alt={props.name} className="product-display__image" />
    <a href={props.link} target="_blank" className="product-display__link btn">Assine Já!</a>
  </div>
);
};

export default ProductDisplay;
