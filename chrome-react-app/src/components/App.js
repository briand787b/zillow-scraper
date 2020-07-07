import React from 'react';

import Header from './Header';
import BackendClient from '../api/backend';
import Map from './Map';
import PropertyList from './PropertyList';
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

    handleMapProperty = (property) => {
        return () => {
            this.setState({ mappedProperty: property });
        }
    }

    getMapComponent() {
        console.log('getting map component');
        if (this.state.mappedProperty) {
            return <Map address={this.state.mappedProperty.address}/>
        }

        if (this.state.properties && this.state.properties.length > 0) {
            return <Map address={this.state.properties[0].address} />
        }

        console.log('no properties to map')
        return <p>Loading...</p>
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
                    {this.getMapComponent()}
                </aside>
                <main>
                    <PropertyList 
                        properties={this.state.properties} 
                        handleMapProperty={this.handleMapProperty}
                    />
                </main>
            </div>
        );
    }
}

export default App;
