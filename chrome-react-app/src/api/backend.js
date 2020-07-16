import axios from 'axios';

class BackendClient {
    constructor(host) {
        this.host = host;
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

        return resp.data.properties;
    }

    async getGoogleMapsEmbedAPIKey() {
        const resp = await axios.get(this.host + '/secrets/google-maps-embed-api-key')
        console.log('response: ', resp)

        if (resp.status > 299) {
            throw new Error(`(${resp.status}) could not load google maps embed api key`)
        }

        return resp.data.secret
    }
}

export default BackendClient;