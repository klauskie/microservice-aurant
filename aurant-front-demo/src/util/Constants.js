export const CATALOG_API_URL = 'http://localhost:8080'
export const PARTY_API_URL = 'http://localhost:8081'
export const SESSION_API_URL = 'http://localhost:8083'

export const KEY_ACCESS_TOKEN = 'access_token'
export const KEY_TAG = 'tag'
export const KEY_VENDOR_ID = 'vendor_id'
export const SIX_HOURS = 6*60*60*1000;

// DEMO UTIL CONSTANTS
export const STATE_CREATE_PARTY = 0
export const STATE_JOIN_PARTY = 1
export const STATE_WAIT_ROOM = 2
export const STATE_MENU = 3
export const STATE_ORDER = 4

export const VENDOR_ID = "74d760b9-83cc-4baa-bdbb-9e07debb58e1"

// Helper Functions
const getTokenKey = (clientNumber) => {
    return KEY_ACCESS_TOKEN + "_" + clientNumber
}

const getPartyIdKey = (clientNumber) => {
    return KEY_TAG + "_" + clientNumber
}

export {
    getTokenKey,
    getPartyIdKey,
}