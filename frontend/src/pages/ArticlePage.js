import React from 'react';
import { useParams } from 'react-router-dom';
import articles from './article-content';

const ArticlePage = () => {
    const { name } = useParams();

    const article = articles.find(article => article.name == name );

    return (
        <>
        <h1>{article.title}</h1>
        {article.content.map((paragraph, id) => (
            <p key={id}>{paragraph}</p>
        ))}
        </>
    );
}

export default ArticlePage;