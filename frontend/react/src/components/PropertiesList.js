import React from 'react';

import Property from './Property';

import '../styles/propertyList.css';

const PropertiesList = ({ properties }) => (
    <ul className="property-list">
        {properties.map((zillowProperty) => 
            <li key={zillowProperty.id}><Property {...zillowProperty}/></li>
        )}
    </ul>
);

export default PropertiesList;