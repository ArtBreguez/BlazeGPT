import React from 'react';
import bannerImage from './banner.jpg';

const Banner = () => {
  return (
    <div style={{backgroundColor: '#eb3100', textAlign: 'center'}}>
      <img src={bannerImage} alt='Banner' style={{display: 'block', margin: 'auto'}} />
    </div>
  );
};

export default Banner;
