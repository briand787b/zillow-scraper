class Scraper {
    constructor(backendClient) {
        this.backendClient = backendClient;
        chrome.runtime.onMessage.addListener(this.receiveMessage);
    }

    scrape(host, favoriteURLs) {
        console.log(favoriteURLs);
        this.favoriteURLs = favoriteURLs;
        this.host = host;
        chrome.tabs.create({ active: false }, (tab) => {
            this.tab = tab;
            this.travel();
        });
    }

    travel() {
        const url = this.favoriteURLs.shift();
        if (!url) {
            return;
        }

        chrome.tabs.update(this.tab.id, { url });
    }

    receiveMessage = async (msg, sender, sendResponse) => {
        console.log("POSTing to server - msg: ", msg);
        await this.backendClient.postCapture(msg);
        this.travel();
    }
}

class BackendClient {
    constructor(host) {
        this.host = host;
    }

    async postCapture(contentMsg) {
        const status = contentMsg.price ? 'For Sale' : 'Off Market';
        const request = {
            property: {
                url: contentMsg.url,
            },
            status: status,
            price: contentMsg.price,
        };

        console.log('request: ', request);
        const response = await fetch(`${this.host}/captures`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(request),
        });

        console.log('response: ', response);
    }
}

chrome.runtime.onInstalled.addListener(() => {
    const scraper = new Scraper(new BackendClient('http://localhost:8080'));
    chrome.browserAction.onClicked.addListener(() => {
        chrome.storage.sync.get(['host', 'favorites'], ({ host, favorites }) => {
            scraper.scrape(host, favorites.map((f) => f.url));
        });
    });
});