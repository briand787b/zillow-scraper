import React from 'react';

import '../styles/Map.css';

const Map = (props) => {
    if (!props.properties || props.properties.length < 1) {
        return <p>No properties to map</p>
    }

    const address = props.properties[0].address
    return (
        <div class="map">
            <iframe
                    frameborder="0"
                    title="id"
                    src={`https://www.google.com/maps/embed/v1/place?key=AIzaSyB76MwBO4v3QDDYLy6o4r7DpOBESqFcZ7A&q=${encodeURIComponent(address)}`}
                    allowfullscreen
            ></iframe>
        </div>
    );
}

export default Map;