import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <div id="page-body">
          <Route path="/">
            <HomePage />
          </Route>
          <Route>
            
          </Route>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
