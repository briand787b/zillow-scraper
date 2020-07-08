import React from 'react';

import '../styles/PropertyList.css';

const capitalizeAddresses = (address) => {

}

const PropertyList = (props) => {
    return (
        <div className="property-list">
            <div>
                {props.properties.map((property) => {
                    return (<div className="property">
                        <button onClick={props.handleMapProperty(property)}>Map</button>
                        <button 
                            className={property.favorited ? 'favorite' : 'non-favorite'}
                        />
                        <a href={property.url}>{property.address}</a>
                    </div>);
                })}
            </div>
        </div>
    );
};

export default PropertyList;