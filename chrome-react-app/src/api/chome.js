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
    getFavorites(callback) {
        chrome.storage.sync.get(['favorites'], callback)
    }

    setFavorites(favorites, callback=()=>{}) {
        chrome.storage.sync.set({ 'favorites': favorites }, callback);
    }
}

export default Chrome