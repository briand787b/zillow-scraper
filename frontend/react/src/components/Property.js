import React from 'react';

const Property = ({id, url, acreage, address}) => (
    <div>
        <p>id: {id}</p>
        <p>url: {url}</p>
        <p>acreage: {acreage}</p>
        <p>address: {address}</p>
    </div>
)

export default Property;