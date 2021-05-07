import axios from 'axios';
import { PARTY_API_URL, SESSION_API_URL, CATALOG_API_URL } from './Constants';

const guestLoginPOST = async (name, vendorId) => {
    const url = `${SESSION_API_URL}/api/guest-login`
    // console.log("API CALL: " + url)

    let requestData = {
        name: name,
        vendor_id: vendorId
    }
    try {
        const resp = await axios.post(url, requestData);
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't post data. " + err);
        return null
    }
};

const createNewPartyPOST = async (token) => {
    const url = `${PARTY_API_URL}/api/party`
    // console.log("API CALL: " + url)

    let requestData = {}
    try {
        const resp = await axios.post(url, requestData, {headers: getHeaders(token)} );
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't post data. " + err);
        return null
    }
}

const joinPartyPUT = async (partyId, token) => {
    const url = `${PARTY_API_URL}/api/party/${partyId}`
    // console.log("API CALL: " + url)

    let requestData = {}
    try {
        const resp = await axios.put(url, requestData, {headers: getHeaders(token)} );
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't put data. " + err);
        return null
    }
}

const fetchMenuGET = async (vendorId) => {
    const url = CATALOG_API_URL + `/api/category/restaurant/${vendorId}`
    console.log("API CALL: " + url)

    try {
        const resp = await axios.get(url);
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't fetch data. " + err);
        return null
    }
}

const sendClientOrderPOST = async (requestData, partyId, token) => {
    const url = `${PARTY_API_URL}/api/order/${partyId}`
    // console.log("API CALL: " + url)

    try {
        const resp = await axios.post(url, requestData, {headers: getHeaders(token)} );
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't post data. " + err);
        return null
    }
}

const fetchPartyOrderGET = async (partyId, token) => {
    const url = `${PARTY_API_URL}/api/party-order/${partyId}`
    // console.log("API CALL: " + url)

    try {
        const resp = await axios.get(url, {headers: getHeaders(token)});
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't fetch data. " + err);
        return null
    }
}

const fetchPartyGET = async (partyId, token) => {
    const url = `${PARTY_API_URL}/api/party/${partyId}`
    console.log("API CALL: " + url)

    try {
        const resp = await axios.get(url, {headers: getHeaders(token)});
        console.log(resp.data);
        return resp.data.party
    } catch (err) {
        console.error("Couldn't fetch data. " + err);
        return null
    }
}

const updatePartyStatusPUT = async (partyId, token) => {
    const url = `${PARTY_API_URL}/api/party-status/${partyId}`
    // console.log("API CALL: " + url)

    let requestData = {ready: true}
    try {
        const resp = await axios.put(url, requestData, {headers: getHeaders(token)} );
        // console.log(resp.data);
        return resp.data
    } catch (err) {
        // console.error("Couldn't put data. " + err);
        return null
    }
}

const getHeaders = (token) => {
    return {
        'Content-Type': 'application/json',
        'token': token
    }
}

export { 
    guestLoginPOST,
    createNewPartyPOST,
    joinPartyPUT,
    sendClientOrderPOST,
    fetchMenuGET,
    fetchPartyOrderGET,
    fetchPartyGET,
    updatePartyStatusPUT,
}