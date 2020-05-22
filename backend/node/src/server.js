import express from 'express';
import bodyParser from 'body-parser';

const articlesInfo = {
    'learn-react': { upvotes: 0, comments: [] },
    'learn-node': { upvotes: 0, comments: []},
    'my-thoughts-on-resumes': { upvotes: 0, comments: []}
};

const app = express();
app.use(bodyParser.json());

app.get('/hello', (req, res) => res.send('Hello!'));
app.get('hello/:name', (req, res) => {
    const name = req.params.name;
    res.send(`Hello ${name}!`);
});
app.post('/hello', (req, res) => {
    const name = req.body.name;
    res.send(`Hello ${name}!!!!`);
});

app.post('/api/articles/:name/upvote', (req, res) => {
    const articleName = req.params.name;
    articlesInfo[articleName].upvotes += 1;

    res.status(200).send(`Success! ${articleName} now has ${articlesInfo[articleName].upvotes} upvotes!`)
});

app.post('/api/articles/:name/add-comment', (req, res) => {
    const articleName = req.params.name;
    const { comment } = req.body;
    articlesInfo[articleName].comments.push(comment);

    res.status(200).send(articlesInfo[articleName]);
});

app.listen(8000, () => {
    console.log('server is listening on port 8000')
})