import React from 'react';

import '../styles/Map.css';

const apiKey = 'AIzaSyBWhb2aqqybdMvZdVDfCXhojQBkpEGIPG4';

const Map = (props) => {
    console.log('map props: ', props);
    return (
        <div class="map">
            <iframe
                    frameborder="0"
                    title="id"
                    src={`https://www.google.com/maps/embed/v1/place?key=${apiKey}&q=${encodeURIComponent(props.address)}`}
                    allowfullscreen
            ></iframe>
        </div>
    );
}

export default Map;