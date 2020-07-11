import React from 'react';

import Header from './Header';
import BackendClient from '../api/backend';
import Chrome from '../api/chrome';
import Map from './Map';
import PropertyList from './PropertyList';
import PropertyBox from './PropertyBox';
import SearchBar from './SearchBar';

import '../styles/App.css';

class App extends React.Component {
    state = {
        backendClient: new BackendClient('http://localhost:8080'),
        properties: [],
        chromeClient: new Chrome(),
        favorites: [],
    };

    async componentDidMount() {
        const properties = await this.state.backendClient.getProperties();
        const favorites = await this.state.chromeClient.getFavorites();
        this.setState({ properties: properties, favorites: favorites });
    }

    handleSearch = async (term) => {
        console.log('handling search');
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
        return async () => {
            console.log('handleFavorite has been triggered - property: ', property, ' state: ', this.state);
            const oldFavorites = this.state.favorites;
            console.log('oldFavorites: ', oldFavorites);
            let newFavorites;
            if (!property.favorited) {
                console.log('adding property to favorites');
                newFavorites = [...oldFavorites, property];
            } else {
                console.log('removing property from favorites');
                newFavorites = oldFavorites.filter(fav => fav.id !== property.id);
            }

            console.log('new favorites: ', newFavorites);
            await this.state.chromeClient.setFavorites(newFavorites);
            console.log('setting local favorites - state: ', this.state);
            this.setState({ favorites: newFavorites });
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
                        favorites={this.state.favorites}
                        handleFavorite={this.handleFavorite}
                        handleMapProperty={this.handleMapProperty}
                    />
                </main>
                <PropertyBox />
            </div>
        );
    }
}

export default App;
