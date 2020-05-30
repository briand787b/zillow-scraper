import React from 'react';

import '../styles/propertyGridItem.css';

const PropertyGridItem = ({ url, imageURL, acreage, address, currentPrice, recentChange }) => (
    <div className="property-grid-item">
        <span>I'm a grid item</span>
        <img src={imageURL} alt={address} />
    </div>
);

export default PropertyGridItem;