import React from 'react';

import PropertiesList from '../components/PropertiesList';

import '../styles/homePage.css';

const HomePage = ({ properties }) => (
    <>
        <h1 className="home-page-header" >Favorited Properties</h1>
        <PropertiesList {...{properties}} />
    </>
);

export default HomePage;