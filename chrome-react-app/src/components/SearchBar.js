import React from 'react';

class SearchBar extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            placeholder: props.placeholder,
            search: props.search,
        };
    }

    onFormChanged = (event) => {
        this.setState({ search: event.target.value })
        console.log(event.target.value);
    };

    onFormSubmitted = (event) => {
        event.preventDefault();
        // this.state
    };

    render() {
        return (
            <div>
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
