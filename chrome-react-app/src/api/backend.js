import axios from 'axios';

const maxTake = 1;

class BackendClient {
    constructor(host) {
        this.host = host;
    }

    async getProperties(take, skip=0) {
        let currTake;
        let properties = [];
        while (take > 0) {
            currTake = take > maxTake ? maxTake : take;

            const resp = await axios.get(this.host+'/properties', {
                params: {
                    take: currTake,
                    skip: skip,
                }
            })
    
            // console.log(resp);
            if (resp.status > 299) {
                throw new Error(`(${resp.status}) could not load properties: ` + resp.data)
            }

            let props = resp.data.properties;
            if (props.length < 1) {
                break;
            }

            properties = [...properties, ...props];
            
            console.log('currTake: ', currTake, ' take: ', take, ' skip: ', skip, 'properties: ', properties);
            take -= currTake;
            skip += currTake;
        }

        return properties;
    }

    async getAllProperties() {
        let properties = [];
        let take = Infinity;
        let skip = 0;
        while (take > 0) {
            const resp = await axios.get(this.host+'/properties', {
                params: {
                    take: take,
                    skip: skip,
                }
            })

            // console.log('response', resp);
            if (resp.status > 299) {
                throw new Error(`(${resp.status}) could not load properties: ` + resp.data)
            }

            properties = [...properties, ...resp.data.properties];
            skip = resp.data.skip;
            take = resp.data.take;
        }
    }
}

export default BackendClient;