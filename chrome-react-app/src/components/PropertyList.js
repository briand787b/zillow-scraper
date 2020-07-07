import React from 'react';

import '../styles/PropertyList.css';

const PropertyList = (props) => {
    return (
        <div className="property-list">
            <div>
                {props.properties.map((property) => {
                    return (<div className="property">
                        <a href={property.url}>{property.address}</a>
                    </div>);
                })}
            </div>
        </div>
    );
};

export default PropertyList;