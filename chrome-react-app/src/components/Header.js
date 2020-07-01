import React from 'react';

import SearchBar from './SearchBar';

const Header = (props) => {
   return (
      <div>
         <h1>Hello</h1>
         <SearchBar host={props.host} setProperties={props.setProperties} />
         <h1>Zillow Scraper</h1>
      </div>
   )
}

export default Header;