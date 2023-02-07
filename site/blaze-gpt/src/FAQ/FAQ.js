import React from 'react';
import './FAQ.css'

const FAQ = () => {
  return (
    <div className="faq">
      <h2 className="faq__title">Perguntas Frequentes</h2>
      <p className="faq__description">Aqui estão as respostas para algumas das perguntas mais comuns sobre nosso produto:</p>
      <ul className="faq__list">
        <li className="faq__item">
          <h3 className="faq__item-title">Como funciona o produto?</h3>
          <p className="faq__item-description">Nosso produto funciona da seguinte maneira:</p>
        </li>
        <li className="faq__item">
          <h3 className="faq__item-title">Quais são as vantagens do produto?</h3>
          <p className="faq__item-description">Algumas das vantagens do nosso produto são:</p>
        </li>
      </ul>
    </div>
  );
};

export default FAQ;
