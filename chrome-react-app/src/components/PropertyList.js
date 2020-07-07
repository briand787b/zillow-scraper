import React from 'react';

import '../styles/PropertyList.css';

const PropertyList = (props) => {
    return (
        <div className="property-list">
            <div>
                {props.properties.map((property) => {
                    return (<div className="property">
                        <a href={property.url}>{property.address}</a>
                        <button onClick={props.handleMapProperty(property)}>Map</button>
                        <p>Fav</p>
                    </div>);
                })}
            </div>
        </div>
    );
};

export default PropertyList;