document.getElementById('server-host-url-form').addEventListener('submit', (event) => {
    const host = document.getElementById('server-host-url').value;
    chrome.storage.sync.set({ host: host }, () => { });
    event.preventDefault();
});

// chrome.storage.sync.set({ favorites: [{"id:0001": "address}"] }, () => { });

chrome.storage.sync.get(['host', 'favorites'], async ({ host, favorites }) => {
    console.log(host);
    console.log(favorites);

    if (typeof (host) !== "string") {
        throw new Error(`Host is not a string value, is ${typeof host}`)
    }

    document.getElementById('server-host-url').setAttribute('value', host);

    const resp = await fetch(host + "/properties?take=10");
    const respJSON = await resp.json();

    console.log(respJSON);

    const propertySelectionForm = document.getElementById('property-selection');
    const propertiesList = respJSON.properties
    for (const property of propertiesList) {
        const favoriteDiv = document.createElement('div');
        favoriteDiv.setAttribute('class', 'favorite-input-group')

        const input = document.createElement('input');
        input.setAttribute('type', 'checkbox');
        input.setAttribute('id', property.id);
        input.setAttribute('name', property.address);
        input.setAttribute('value', property.address);
        input.setAttribute('url', property.url);
        if (favorites && favorites.length > 1) {
            const found = favorites.find(({id}) => id === property.id);
            console.log(found);
            if (found !== undefined) {
                input.checked = true;
            }
        }
        favoriteDiv.appendChild(input);

        const label = document.createElement('label');
        label.setAttribute('for', property.id);
        label.innerHTML = property.address;
        favoriteDiv.appendChild(label);

        propertySelectionForm.appendChild(favoriteDiv);
    }

    const propertyListSubmitDiv = document.createElement('div')
    propertyListSubmitDiv.setAttribute("class", "submit");
    propertySelectionForm.appendChild(propertyListSubmitDiv);

    const propertyListButton = document.createElement('button');
    propertyListButton.setAttribute('type', 'submit');
    propertyListButton.setAttribute('class', 'property-list-submit');
    propertyListButton.innerHTML = 'Save Favorites';
    propertyListSubmitDiv.appendChild(propertyListButton);

    console.log('done populating list');
});

document.getElementById('property-selection').addEventListener('submit', (event) => {
    // document.getElementById().getAttribute();
    console.log(event.target);
    console.log(event.target.tagName);
    console.log(event.target.children);
    console.log(event.target.children.item(0));

    const newFavorites = [];
    let inputElem;
    for (const child of event.target.children) {
        console.log(child);
        if (child.className !== 'favorite-input-group') {
            continue;
        }

        inputElem = child.children.item(0);
        console.log(`checked? ${inputElem.checked}`);
        if (inputElem.checked) {
            newFavorites.push({
                id: inputElem.getAttribute('id'),
                address: inputElem.getAttribute('value'),
                url: inputElem.getAttribute('url'),
            });
        }
    }

    console.log(newFavorites);
    chrome.storage.sync.set({ 'favorites': newFavorites }, () => {
        console.log('saved new favorites');
    });

    event.preventDefault();
});