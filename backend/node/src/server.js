import express from 'express';
import bodyParser from 'body-parser';
import { MongoClient } from 'mongodb';

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

app.get('/api/articles/:name', async (req, res) => {
    const articleName = req.params.name;
    
    try {
        const client = await MongoClient.connect(
            'mongodb://zcrapr:zcrapr@mongo:27017',
            {useNewUrlParser: true, useUnifiedTopology: true}
        );
    
        const db = client.db('zcrapr');
        
        const articleInfo = await db.collection('articles')
            .findOne({ name: articleName });
        
        client.close();
        
        if (!articleInfo) {
            return res.status(404).send('article not found');
        }
    
        res.status(200).json(articleInfo);
    } catch (e) {
        console.log(e);
        res.status(500).send('something went wrong');
    }
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