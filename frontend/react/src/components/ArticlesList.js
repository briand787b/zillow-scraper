import React from 'react';
import { Link } from 'react-router-dom';

const ArticlesList = ({ articles }) => (
    <>
    {articles.map(article => (
        <Link 
            to={`/article/${article.name}`}
            className="article-list-item" 
            key={article.name} 
        >
            <h3>{article.name}</h3>
            <p>{article.content[0].substring(0, 100)}...</p>
        </Link>
    ))}
    </>
);

export default ArticlesList;