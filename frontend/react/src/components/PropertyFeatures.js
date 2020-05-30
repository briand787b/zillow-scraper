import React from 'react';

import '../styles/propertyFeatures.css';

const PropertyFeatures = ({ acreage, currentPrice, recentChange }) => (
    <div className="features">
        <div className="price">
            <span>Hello Class</span>
        </div>
        <div className="cost">
            <span className="cost-currency">$</span>
            <span className="cost-dollars">20</span>
            <span className="cost-cents">.00</span>
        </div>
        <a className="cta-button" href="/pricing">Sign up</a>
    </div >
);

export default PropertyFeatures;