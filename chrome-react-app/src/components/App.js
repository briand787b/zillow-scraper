import React from 'react';

import Header from './Header';
import BackendClient from '../api/backend';
import Map from './Map';
import PropertyBox from './PropertyBox';
import SearchBar from './SearchBar';
import '../styles/App.css';

class App extends React.Component {
    state = {
        backendClient: new BackendClient('http://localhost:8080'),
        properties: [],
    };

    async componentDidMount() {
        const properties = await this.state.backendClient.getProperties()
        this.setState({ properties: properties });
    }

    handleSearch = async (term) => {
        const properties = await this.state.backendClient.getProperties(0, term)
        this.setState({ properties: properties });
    }

    render() {
        console.log('state', this.state);
        return (
            <div class="content">
                <Header />
                <div className="search">
                    <SearchBar handleSearch={this.handleSearch} />
                </div>
                <aside>
                    <Map properties={this.state.properties} />
                </aside>
                <main>
                    <PropertyBox properties={this.state.properties} />
                </main>
            </div>
        );
    }
}

export default App;
