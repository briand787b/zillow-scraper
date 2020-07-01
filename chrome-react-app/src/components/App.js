import React from 'react';

import Header from './Header';

class App extends React.Component {
    state = {
        serverHost: 'http://localhost:8080',
        properties: null,
    };

    setProperties = (properties) => {
        this.setState({ ...{properties}})
    }

    async componentDidMount() {
        const resp = await fetch(`${this.state.serverHost}/properties?take=100`);
        const body = await resp.json();
        console.log('body', body);
        this.setState({ properties: body.properties });
    }

    render() {
        return (
            <div>
                <Header setProperties={this.setProperties} />
                <p>Body</p>
            </div>
        );
    }
}

export default App;
