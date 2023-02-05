import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckCircle } from '@fortawesome/free-solid-svg-icons';
import './Advantages.css';


const Advantages = () => {
  return (
    
    <div style={{ marginBottom: '70px' }} className="advantages-container">

      <h2 className="advantages-title">Vantagens da Sala de Sinais Automatizados por IA</h2>
      <ul className="advantages-list">
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Tomada de decisões mais acertadas com a ajuda da inteligência artificial</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Análise automatizada de padrões de jogo com precisão surpreendente</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Aumento da efetividade nas apostas com a análise de dados avançada</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Redução de erros humanos com o poder da inteligência artificial</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Monitoramento constante dos padrões do jogo para garantir vantagem competitiva.</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Receba alertas diretamente em seu dispositivo móvel através do Telegram para auxiliar em suas apostas com o poder da inteligência artificial.</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Previsão de resultados futuros com base em dados históricos e análise estatística.</span>
        </li>
        <li className="advantage">
          <FontAwesomeIcon icon={faCheckCircle} className="advantage-icon"/>
          <span className="advantage-text">Acesso a insights e estatísticas em tempo real para uma tomada de decisão mais informada e precisa.</span>
        </li>
      </ul>
    </div>
  );
};

export default Advantages;
