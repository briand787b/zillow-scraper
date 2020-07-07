import React from 'react';

import '../styles/SearchBar.css';

class SearchBar extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            placeholder: props.placeholder,
            search: props.search,
            handleSearch: props.handleSearch,
        };
    }

    onFormChanged = (event) => {
        this.setState({ search: event.target.value }, () => {
            this.state.handleSearch(this.state.search);
        });
    };

    // SearchBar continuously searches, so no need to handle submission
    onFormSubmitted = (event) => {
        event.preventDefault();
    };

    render() {
        return (
            <div className="search-bar">
                <form onSubmit={this.onFormSubmitted}>
                    <input
                        type="text"
                        placeholder={this.state.placeholder}
                        value={this.state.search}
                        onChange={this.onFormChanged}>
                    </input>
                </form>
            </div>
        )
    }
}

SearchBar.defaultProps = {
    placeholder: 'search'
};

export default SearchBar;
