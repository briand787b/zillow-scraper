import axios from 'axios';

const maxTake = 1;

class BackendClient {
    constructor(host) {
        this.host = host;
    }

    async getProperties(take = 0) {
        const resp = await axios.get(this.host + '/properties', {
            params: {
                take: take,
            }
        })

        console.log(resp);
        if (resp.status > 299) {
            throw new Error(`(${resp.status}) could not load properties: ` + resp.data)
        }

        return resp.data.properties;
    }
}

export default BackendClient;