import React from 'react';

import Header from './Header';
import BackendClient from '../api/backend';
import Chrome from '../api/chome';
import Map from './Map';
import PropertyList from './PropertyList';
import SearchBar from './SearchBar';

import '../styles/App.css';

class App extends React.Component {
    state = {
        backendClient: new BackendClient('http://localhost:8080'),
        chromeClient: new Chrome(),
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
            console.log('handling mapping - property: ', property, ' state: ', this.state);
            const newProperties = this.state.properties.map((p) => {
                if (p.id === property.id) {
                    p.mapped = true;
                    return p;
                }

                p.mapped = false;
                return p;
            });
            this.setState({ properties: newProperties });
        }
    }

    handleFavorite = (property) => {
        return () => {
            console.log('handling favorite - property:', property, ' state', this.state);
            const newProperties = this.state.properties.map((p) => {
                if (p.id === property.id) {
                    p.favorited = true;
                    return p;
                }

                return p;
            });
            this.setState({ properties: newProperties })
        }
    }

    getMapComponent() {
        console.log('getting map component - this.state: ', this.state);
        if (this.state.properties && this.state.properties.length > 0) {
            let mapped = this.state.properties.find((p) => p.mapped)
            if (!mapped) {
                this.handleMapProperty(this.state.properties[0])()
                mapped = this.state.properties[0]
            }
            
            console.log('mapped property: ', mapped)
            return <Map address={mapped.address} />
        }

        return <p>no properties to map</p>
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
                        handleFavorite={this.handleFavorite}
                        handleMapProperty={this.handleMapProperty}
                    />
                </main>
            </div>
        );
    }
}

export default App;
