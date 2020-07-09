import axios from 'axios';

class BackendClient {
    constructor(host, chromeClient) {
        this.host = host;
        this.chromeClient = chromeClient;
    }

    async getProperties(take = 0, address = '') {
        const resp = await axios.get(this.host + '/properties', {
            params: {
                take: take,
                address: address,
            }
        })

        console.log('backend response: ', resp);
        if (resp.status > 299) {
            throw new Error(`(${resp.status}) could not load properties: ` + resp.data)
        }

        const favorites = await this.chromeClient.getFavorites();
        console.log('favorites: ', favorites);

        return resp.data.properties;
    }
}

export default BackendClient;