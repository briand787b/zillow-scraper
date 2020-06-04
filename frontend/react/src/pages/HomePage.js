import React from 'react';

import PropertiesList from '../components/PropertiesList';

import '../styles/homePage.css';

const HomePage = () => (
    <>
        <h1 className="home-page-header" >Favorited Properties</h1>
        <PropertiesList />
    </>
);

export default HomePage;