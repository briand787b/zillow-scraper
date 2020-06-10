chrome.runtime.onInstalled.addListener(() => {
    chrome.browserAction.onClicked.addListener(() => {
        console.log('Clicked');

        chrome.storage.sync.set({
            urls: [
                "https://www.zillow.com/homedetails/410-Retreat-St-Westminster-SC-29693/111164824_zpid/",
                "https://www.zillow.com/homedetails/246-Little-Choestoea-Rd-Westminster-SC-29693/70955580_zpid/",
                "https://www.zillow.com/homedetails/200-Harbour-West-Dr-Westminster-SC-29693/70960114_zpid/",
            ]
        });

        chrome.storage.sync.get("urls", (result) => {
            console.log(result);
        });
    });
});