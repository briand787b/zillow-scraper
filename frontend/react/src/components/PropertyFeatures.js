import React from 'react';

import '../styles/propertyFeatures.css';

const PropertyFeatures = ({ acreage }) => (
    <div className="features">
        <h4>Features</h4>
        <ul className="features-list">
            <li>Acreage: {acreage} </li>
        </ul>
    </div>
);

export default PropertyFeatures;