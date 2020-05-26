import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import HomePage from './pages/HomePage';
import NotFoundPage from './pages/NotFoundPage';

import NavBar from './components/NavBar';

import './App.css';

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <NavBar />
        <div id="page-body">
          <Switch>
          <Route path="/" exact>
            <HomePage />
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
