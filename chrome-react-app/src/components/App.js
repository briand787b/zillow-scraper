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
