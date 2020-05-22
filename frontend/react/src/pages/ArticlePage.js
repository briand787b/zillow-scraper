import React from 'react';
import { useParams } from 'react-router-dom';
import articles from './article-content';
import NotFoundPage from './NotFoundPage'

const ArticlePage = () => {
    const { name } = useParams();

    const article = articles.find(article => article.name === name );

    return article 
        ? (
        <>
        <h1>{article.title}</h1>
        {article.content.map((paragraph, id) => (
            <p key={id}>{paragraph}</p>
        ))}
        </>
    ) : (
        <NotFoundPage/>
    );
}

export default ArticlePage;