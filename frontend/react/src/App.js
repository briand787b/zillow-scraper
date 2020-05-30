import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import HomePage from './pages/HomePage';
import GridPage from './pages/GridPage';
import NotFoundPage from './pages/NotFoundPage';

import NavBar from './components/NavBar';

import './App.css';

function App() {
  return (
    <BrowserRouter>
      <div className="container">
        <NavBar />
        <div id="page-body">
          <Switch>
          <Route path="/" exact>
            <HomePage />
          </Route>
          <Route  path="/grid">
            <GridPage />
          </Route>
          <Route>  
            <NotFoundPage/>
          </Route>
          </Switch>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
