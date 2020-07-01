import axios from 'axios';

const maxTake = 100;

class BackendClient {
    constructor(host) {
        this.host = host;
    }

    #getPictures(skip, take) {
        if (take < 1) {
            throw new Error("take must be greater than 0");
        }

        if (take > maxTake) {
            throw new Error(`take exceeds ${maxTake}`);
        }
    }

    getPictures(skip, take) {
        
    }

    async getAllPictures() {
        const resp = await axios.get(this.host+'/pictures', {
            params: {
                take: maxTake,
            }
        })

        console.log(resp);
    }
}

export default BackendClient;