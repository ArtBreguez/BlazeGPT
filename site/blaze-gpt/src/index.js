import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import ProductDisplay from './ProductDisplay';


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
    <ProductDisplay
        name="Double Blaze AI Signals Room"
        description="Transforme sua estratégia de jogo com o Blaze AI Signals Room, uma sala exclusiva no Telegram que oferece sinais precisos e inteligentes para o jogo Double da Blaze, aumentando suas chances de vitória."
        link="https://go.hotmart.com/T79295798B"
      />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
