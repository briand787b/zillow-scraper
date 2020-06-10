document.getElementById('server-host-url-form').addEventListener('submit', (event) => {
    const host = document.getElementById('server-host-url').value;
    chrome.storage.sync.set({ host: host }, () => { });
    event.preventDefault();
});

// chrome.storage.sync.set({ favorite_ids: ["id:0001"] }, () => { });

chrome.storage.sync.get(['host', 'favorite_ids'], async ({ host, favorite_ids }) => {
    console.log(host);
    console.log(favorite_ids);

    if (typeof (host) !== "string") {
        throw new Error(`Host is not a string value, is ${typeof host}`)
    }

    document.getElementById('server-host-url').setAttribute('value', host);

    const resp = await fetch(host);
    const respJSON = await resp.json();

    console.log(respJSON);

    const propertySelectionForm = document.getElementById('property-selection');
    const propertiesList = respJSON.properties
    for (let i = 0; i < propertiesList.length; i++) {
        const favoriteDiv = document.createElement('div');
        favoriteDiv.setAttribute('class', 'favorite-input-group')

        const input = document.createElement('input');
        input.setAttribute('type', 'checkbox');
        input.setAttribute('id', propertiesList[i].id);
        input.setAttribute('name', propertiesList[i].address);
        input.setAttribute('value', propertiesList[i].address);
        if (favorite_ids && favorite_ids.length > 1) {
            const found = favorite_ids.find(id => id === propertiesList[i].id)
            console.log(found);
            if (found !== undefined) {
                input.checked = true;
            }
        }
        favoriteDiv.appendChild(input);

        const label = document.createElement('label');
        label.setAttribute('for', propertiesList[i].id);
        label.innerHTML = propertiesList[i].address;
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
            newFavorites.push(inputElem.getAttribute('id'))
        }
    }

    console.log(newFavorites);
    chrome.storage.sync.set({'favorite_ids': newFavorites}, () => {
        console.log('saved new favorites');
    });

    event.preventDefault();
});