class Scraper {
  constructor(favoriteURLs) {
    console.log(favoriteURLs);
    this.favoriteURLs = favoriteURLs;
    chrome.runtime.onMessage.addListener(this.receiveMessage);
  }

  scrape() {
    chrome.tabs.create({}, (tab) => {
      this.tab = tab;
      this.travel();
    });
  }

  travel() {
    const url = this.favoriteURLs.shift();
    if (!url) {
      return;
    }

    chrome.tabs.update(this.tab.id, {url});
  }

  receiveMessage = (request, sender, sendResponse) => {
    console.log(request);
    this.travel();
  }
}

// See if there appears to be more than one listener added
// for this event.  If there is, then wrap this in a
// listener for the chrome extension installation
chrome.browserAction.onClicked.addListener(() => {
  chrome.storage.sync.get(['favorites'], ({ favorites }) => {
    new Scraper(favorites.map((f) => f.url)).scrape();
  });
});