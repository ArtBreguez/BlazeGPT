import React from 'react';
import Banner from './Banner/Banner';
import ProductDisplay from './ProductDisplay/ProductDisplay';
import Advantages from './Advantages/Advantages';
import Footer from './Footer/Footer';
import FAQ from './FAQ/FAQ';

function App() {
  return (
    <div>
       <Banner />
       <ProductDisplay
        name="Double Blaze AI Signals Room"
        description="Aumente sua taxa de sucesso no jogo Double da Blaze com o Blaze AI Signals Room, a solução definitiva para apostadores sérios. Com sinais precisos e inteligentes gerados pela nossa inteligência artificial avançada, você terá uma vantagem competitiva e tomará decisões informadas de forma rápida e segura. Junte-se a nossa sala exclusiva no Telegram e transforme seu jogo hoje mesmo!"
        link="https://go.hotmart.com/T79295798B"
        />
        <Advantages />
        <FAQ  />
        <Footer  />
    </div>
  );
}

export default App;
