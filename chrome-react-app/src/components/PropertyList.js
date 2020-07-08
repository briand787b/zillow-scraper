import React from 'react';

import '../styles/PropertyList.css';

// TODO: make this actually work
const capitalizeAddress = (address) => {
    return address;
}

const PropertyList = (props) => {
    console.log('rendering PropertyList - props: ', props);
    return (
        <div className="property-list">
            <div>
                {props.properties.map((property) => {
                    // console.log('property: ', property);
                    const propertyClass = property.mapped ? 'property mapped' : 'property'
                    return (<div className={propertyClass}>
                        <button onClick={props.handleMapProperty(property)}>Map</button>
                        <button 
                            onClick={props.handleFavorite(property)}
                            className={property.favorited ? 'favorite' : 'non-favorite'}
                        >
                            Fav
                        </button>
                        <a href={property.url}>{capitalizeAddress(property.address)}</a>
                    </div>);
                })}
            </div>
        </div>
    );
};

export default PropertyList;