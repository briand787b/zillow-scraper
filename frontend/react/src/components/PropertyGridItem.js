import React from 'react';

import '../styles/propertyGridItem.css';

const PropertyPrice = (status, price, pctChange) => (
    status === "For Sale" ? <div></div> : <div></div>
);

const PropertyStatus = (status, price, pctChange) => {
    switch (status) {
        case "For Sale":
            return (<span>${price} {status}</span>)
        default:
            return (<span>{status}</span>)
    }
}

const PropertyGridItem = ({ url, imageURL, acreage, address, currentPrice, priceChangePercent, status }) => (
    <div className="property-grid-item">
        <div className="property-grid-item-header">
            {PropertyStatus(status, currentPrice, priceChangePercent)}
        </div>
        <img src={imageURL} alt={address} />
        <span>Recent Change: {}</span>
    </div>
);

export default PropertyGridItem;