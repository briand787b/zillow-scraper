import React from 'react';

import PropertyFeatures from './PropertyFeatures';

import '../styles/property.css';

const Property = ({ id, url, imageURL, acreage, address }) => (
    <li key={id}>
        <div className="property">
            <h3 className="address">{address}</h3>
            <img className="property-image" src={imageURL} alt={address}/>
            <PropertyFeatures {...{acreage}}/>
            <a className="url" href={url}>Visit Property</a>
        </div>
    </li>
)

export default Property;