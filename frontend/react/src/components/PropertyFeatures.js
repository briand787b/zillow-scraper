import React from 'react';

import '../styles/propertyFeatures.css';

const PropertyFeatures = ({ url, acreage, currentPrice, recentChange }) => (
    <div className="features">
        <div className="cost">
            <span className="cost-currency">$</span>
            <span className="cost-dollars">
                {currentPrice / 1000 }K
            </span>
        </div>
        <div className="cta-button" ></div>
        <div className="specs">
            <span className="acreage">Acreage: {acreage}ac</span>
        </div>
    </div >
);

export default PropertyFeatures;