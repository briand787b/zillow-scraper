import React from 'react';

import '../styles/propertyFeatures.css';

const PropertyFeatures = ({ acreage, address, currentPrice, id, recentChange, url }) => (
    <div className="features">
        <div className="cost">
            <span className="cost-currency">$</span>
            <span className="cost-dollars">
                {currentPrice / 1000}K
            </span>
        </div>
        <h3 className="address">{address}</h3>
        <div className="cta-button" ></div>
        <div className="specs">
            <span className="acreage">{acreage} Acres | ${currentPrice / acreage}/Acre</span>
            <iframe
                    width="300"
                    height="300"
                    frameborder="0"
                    title="id"
                    src={`https://www.google.com/maps/embed/v1/place?key=AIzaSyB76MwBO4v3QDDYLy6o4r7DpOBESqFcZ7A&q=${encodeURIComponent(address)}`}
                    allowfullscreen
            ></iframe>
            <a className="url" href={url}>Visit Property</a>
        </div>
    </div >
);

export default PropertyFeatures;