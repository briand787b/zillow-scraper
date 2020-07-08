import React from 'react';
import PropertyList from './PropertyList';
import '../styles/PropertyBox.css';

const PropertyBox = (props) => {
    return (
        <div class="propert-box">
            <PropertyList properties={props.properties} />
        </div>
    );
}

export default PropertyBox;