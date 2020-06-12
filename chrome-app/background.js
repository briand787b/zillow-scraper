class Scraper {
  constructor(favoriteURLs) {
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
  chrome.storage.sync.get(['favorite_urls'], ({ favorite_urls }) => {
    // TODO: remove this temporary override for a value that does not exist
    favorite_urls = [
      'https://www.zillow.com/homedetails/122-Eastwood-Cir-Westminster-SC-29693/70950821_zpid/',
      'https://www.zillow.com/homedetails/328-Holly-Dr-Westminster-SC-29693/218227134_zpid/',
    ];

    new Scraper(favorite_urls).scrape();
  });
});