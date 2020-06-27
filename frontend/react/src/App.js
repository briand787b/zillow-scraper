import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import HomePage from './pages/HomePage';
import NotFoundPage from './pages/NotFoundPage';

import NavBar from './components/NavBar';

import './App.css';

const properties = [
  {
      id: "id:0001",
      url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
      imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISfozz6yhd1onw0000000000.webp",
      acreage: 200,
      address: "308 Greentree Ct, Seneca, SC 29672",
      currentPrice: 300000,
      recentChange: 1.0,
  },
  {
      id: "id:0002",
      url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
      imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjj2m3xc23iwo1000000000.webp",
      acreage: 100,
      address: "1057 Della St SE, Marietta, Georgia 30067",
      currentPrice: 240000,
      recentChange: -8.0,
  },
  {
      id: "id:0003",
      url: "https://www.zillow.com/homedetails/1310-W-Oak-Hwy-Westminster-SC-29693/111151308_zpid/",
      imageURL: "https://photos.zillowstatic.com/cc_ft_768/ISjvj7dgj8rsua0000000000.webp",
      acreage: 40,
      address: "717 Twilight Ct, Seneca, SC 29672",
      currentPrice: 140000,
      recentChange: 5.2,
  },
];

class App extends React.Component {
  state = {};

  async componentDidMount() {
    const resp = await fetch('http://localhost:8080/properties?skip=0&take=10');
    const body = await resp.json();
    console.log('body', body);
  }

  render() {
    return (
      <BrowserRouter>
        <div className="container">
          <NavBar />
          <div id="page-body">
            <Switch>
              <Route path="/" exact>
                <HomePage {...{properties}} />
              </Route>
              <Route>
                <NotFoundPage />
              </Route>
            </Switch>
          </div>
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
