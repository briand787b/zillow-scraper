import React from 'react';

import PropertyFeatures from './PropertyFeatures';

import '../styles/property.css';

const Property = ({ url, imageURL, acreage, address, currentPrice, recentChange }) => (
        <div className="property">
            <h3 className="address">{address}</h3>
            <img className="property-image" src={imageURL} alt={address}/>
            <PropertyFeatures {...{acreage, currentPrice, recentChange}}/>
            <a className="url" href={url}>Visit Property</a>
        </div>
)

export default Property;