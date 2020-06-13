class Scraper {
    constructor() {
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

    receiveMessage = (request, sender, sendResponse) => {
        console.log("request", request);
        if (request.price) {
            console.log('POSTing to server')
            // TODO: POST to server
        }

        this.travel();
    }
}

chrome.runtime.onInstalled.addListener(() => {
    const scraper = new Scraper();
    chrome.browserAction.onClicked.addListener(() => {
        chrome.storage.sync.get(['host', 'favorites'], ({ host, favorites }) => {
            scraper.scrape(host, favorites.map((f) => f.url));
        });
    });
});