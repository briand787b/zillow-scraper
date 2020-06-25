import React from 'react';

import PropertyFeatures from './PropertyFeatures';

import '../styles/property.css';

const Property = ({ id, url, imageURL, acreage, address, currentPrice, recentChange }) => (
        <div className="property">
            <img className="property-image" src={imageURL} alt={address}/>
            <PropertyFeatures {...{acreage, address, currentPrice, id, recentChange, url}}/>
        </div>
)

export default Property;