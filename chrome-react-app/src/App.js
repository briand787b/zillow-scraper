import React from 'react';
import logo from './logo.svg';
import './App.css';

// IMPORTANT: this app will not work in a chrome extension without this:
// https://stackoverflow.com/a/52498189

class App extends React.Component {
  state = { host: null, favorites: null };

  componentDidMount() {
    // TODO: call backend api for 
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1>Zillow Scraper</h1>
        </header>
      </div>
    );
  }
}

export default App;
