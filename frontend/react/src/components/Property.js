import React from 'react';

import '../styles/property.css';

const Property = ({ id, url, imageURL, acreage, address }) => (
    <li key={id}>
        <h3 className="address">address: {address}</h3>
        <div className="property">
            <img className="property-image" src={imageURL} />
            <p className="acreage">>acreage: {acreage}</p>
        </div>
        <p className="url" >url: {url}</p>
    </li>
)

export default Property;