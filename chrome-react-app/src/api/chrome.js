/*global chrome*/

// storage schema
// {
//    "favorites": [
//      {
//        "id": "string",
//        "address": "string",
//        "url": "string"
//      } 
//    ]  
// }
// 

class Chrome {
    getFavorites() {
        return new Promise((resolve, reject) => {
            chrome.storage.sync.get(['favorites'], result => {
                console.log('favorites from chrome: ', result);
                resolve(result.favorites);
            }); 
        });
    }

    setFavorites(favorites) {
        return new Promise((resolve, reject) => {
            chrome.storage.sync.set({ 'favorites': favorites }, () => {
                console.log('done setting favorites in chrome');
                resolve();
            });
        })
    }
}

export default Chrome;