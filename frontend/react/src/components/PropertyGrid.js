import React from 'react';

import PropertyGridItem from './PropertyGridItem';

import '../styles/propertyGrid.css';

const properties = [
    {
        id: "id:0001",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISfozz6yhd1onw0000000000.webp",
        acreage: 200,
        address: "42 Baker St, Whoville, Kentucky 30067",
        currentPrice: 30000,
        recentChange: 1.0,
    },
    {
        id: "id:0002",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjj2m3xc23iwo1000000000.webp",
        acreage: 100,
        address: "1057 Della St SE, Marietta, Georgia 30067",
        currentPrice: 24000,
        recentChange: -8.0,
    },
    {
        id: "id:0003",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjvj7dgj8rsua0000000000.webp",
        acreage: 40,
        address: "420 Baker St, Whoville, Kentucky 30067",
        currentPrice: 140000,
        recentChange: 5.2,
    },
    {
        id: "id:0004",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISfozz6yhd1onw0000000000.webp",
        acreage: 200,
        address: "42 Baker St, Whoville, Kentucky 30067",
        currentPrice: 30000,
        recentChange: 1.0,
    },
    {
        id: "id:0005",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjj2m3xc23iwo1000000000.webp",
        acreage: 100,
        address: "1057 Della St SE, Marietta, Georgia 30067",
        currentPrice: 24000,
        recentChange: -8.0,
    },
    {
        id: "id:0006",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjvj7dgj8rsua0000000000.webp",
        acreage: 40,
        address: "420 Baker St, Whoville, Kentucky 30067",
        currentPrice: 140000,
        recentChange: 5.2,
    },
    {
        id: "id:0007",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISfozz6yhd1onw0000000000.webp",
        acreage: 200,
        address: "42 Baker St, Whoville, Kentucky 30067",
        currentPrice: 30000,
        recentChange: 1.0,
    },
    {
        id: "id:0008",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjj2m3xc23iwo1000000000.webp",
        acreage: 100,
        address: "1057 Della St SE, Marietta, Georgia 30067",
        currentPrice: 24000,
        recentChange: -8.0,
    },
    {
        id: "id:0009",
        url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
        imageURL: "https://photos.zillowstatic.com/o_a/ISb5iqfwjmwe431000000000.webp",
        acreage: 40,
        address: "420 Baker St, Whoville, Kentucky 30067",
        currentPrice: 140000,
        recentChange: 5.2,
    },
];

const PropertiesGrid = () => (
    <div className="property-grid" >
        {properties.map((property) => (
            <PropertyGridItem key={property.id} {...property} />
        ))}
    </div>
);

export default PropertiesGrid;