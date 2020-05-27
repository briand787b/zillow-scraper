import React from 'react';

import Property from './Property';

import '../styles/propertyList.css';

const properties = [
    {
        id: "id:0001",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISv4nxu7zsu6lb0000000000.webp",
        acreage: 200,
        address: "42 Baker St, Whoville, Kentucky 30067"
    },
    {
        id: "id:0002",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISrxmxndgjkdn91000000000.webp",
        acreage: 100,
        address: "1057 Della St SE, Marietta, Georgia 30067"
    },
    {
        id: "id:0003",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/IS76crjh3u08y60000000000.webp",
        acreage: 40,
        address: "420 Baker St, Whoville, Kentucky 30067"
    },
];

const PropertiesList = () => (
    <ul className="property-list">
        {properties.map((zillowProperty) => 
            <Property {...zillowProperty}/>
        )}
    </ul>
);

export default PropertiesList;