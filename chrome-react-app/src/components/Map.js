import React from 'react';

import '../styles/Map.css';

const Map = (props) => {
    console.log('map props: ', props);
    return (
        <div class="map">
            <iframe
                    frameborder="0"
                    title="id"
                    src={`https://www.google.com/maps/embed/v1/place?key=AIzaSyB76MwBO4v3QDDYLy6o4r7DpOBESqFcZ7A&q=${encodeURIComponent(props.address)}`}
                    allowfullscreen
            ></iframe>
        </div>
    );
}

export default Map;