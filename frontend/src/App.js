import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';

import AboutPage from './pages/AboutPage';
import ArticlePage from './pages/ArticlePage';
import ArticleListPage from './pages/ArticlesListPage';
import HomePage from './pages/HomePage';

import NavBar from './components/NavBar';

import './App.css';

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <NavBar />
        <div id="page-body">
          <Route path="/" exact>
            <HomePage />
          </Route>
          <Route path="/about" >
            <AboutPage />
          </Route>
          <Route path="/articles-list">
            <ArticleListPage />
          </Route>
          <Route path="/article/:name">
            <ArticlePage />
          </Route>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
