import React from 'react';

import Header from './Header';
import BackendClient from '../api/backend';

class App extends React.Component {
    state = {
        backendClient: new BackendClient('http://localhost:8080'),
        properties: null,
    };

    setProperties = (properties) => {
        this.setState({ ...{properties}})
    }

    async componentDidMount() {
        const properties = await this.state.backendClient.getProperties(10)
        this.setState({ properties: properties});
    }

    render() {
        console.log('state', this.state);
        return (
            <div>
                <Header setProperties={this.setProperties} />
                <p>Body</p>
            </div>
        );
    }
}

export default App;
