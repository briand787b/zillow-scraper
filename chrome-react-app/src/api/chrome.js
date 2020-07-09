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
        return new Promise((resolve, rejct) => {
            chrome.storage.sync.get(['favorites'], result => {
                resolve(result);
            }); 
        });
    }

    setFavorites(favorites) {
        return new Promise((resolve, reject) => {
            chrome.storage.sync.set({ 'favorites': favorites }, () => {
                resolve();
            });
        })
    }
}

export default Chrome;