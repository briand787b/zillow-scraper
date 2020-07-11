import React from 'react';

import '../styles/PropertyBoxHeader.css';

const PropertyBoxHeader = (props) => {
    return (
        <div className="property-box-header">
            {Object.keys(props.detailTypes).map(detail => {
                return (
                    <button onClick={props.handleChangeDetailType(detail)}>
                        {props.detailTypes[detail]}
                    </button>
                );
            })}
        </div>
    )
}

export default PropertyBoxHeader;