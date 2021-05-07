import React, { useState } from 'react'
import { useCookies } from 'react-cookie'
import { createNewPartyPOST, guestLoginPOST } from '../../util/APIutils';
import { getPartyIdKey, getTokenKey, KEY_VENDOR_ID, SIX_HOURS, STATE_JOIN_PARTY, STATE_WAIT_ROOM } from '../../util/Constants';

const CreateParty = (props) => {
    const [, setCookie] = useCookies([getTokenKey(props.id), getPartyIdKey(props.id)])

    let [name, setName] = useState("");
    let [, setToken] = useState("");
    let [, setPartyId] = useState("");

    // let vendorId = props.match.params.vendorId
    let vendorId = props.vendorId

    const createNewParty = () => {
        if (vendorId === "") {
            console.log("ERROR: No vendor ID")
            return
        }
        // Login as guest
        handleGuestLogin()
    }

    const handleGuestLogin = () => {
        guestLoginPOST(name, vendorId).then(loginData => {
            setToken(loginData.token)
            handleNewParty(loginData.token)
        }).catch(() => {});
    }

    const handleNewParty = (mToken) => {
        createNewPartyPOST(mToken).then(partyData => {
            let partyTag = partyData.party.tag
            setPartyId(partyTag)
            saveCookies(mToken, partyTag)
            redirectHelper(STATE_WAIT_ROOM, { 
                token: mToken,
                tag: partyTag
            })
        }).catch(() => {})
    }

    const saveCookies = (mToken, partyTag) => {
        setNewCookie(getTokenKey(props.id), mToken)
        setNewCookie(getPartyIdKey(props.id), partyTag)
        setNewCookie(KEY_VENDOR_ID, vendorId)
        localStorage.setItem(KEY_VENDOR_ID, vendorId);
    }

    const redirectHelper = (path, data) => {    
        props.callback(path, data)
    }

    const setNewCookie = (key, value) => {
        let expires = new Date()
        expires.setTime(expires.getTime() + (SIX_HOURS))
        setCookie(key, value, { path: '/',  expires})
    }

    return (
        <div>
            <div className="container">

                <div className="margin-top"></div>

                <div className="right">
                    <a onClick={() => props.callback(STATE_JOIN_PARTY, {})} className="a-tag" >Join a Party</a>
                </div>

                <div className="">
                    <h2 className="left">Create a Party!</h2>
                </div>

                <div className="row justify-content-md-center h-100">
                    <div className="card-body">
                        <div className="form-group">
                            <label htmlFor="name">Name</label>
                            <input autoComplete='off' onChange={(e) => setName(e.target.value)} id="name" type="text" className="form-control" name="name" required autoFocus/>
                        </div>

                        <div className="form-group">
                            <button onClick={(e) => createNewParty()} disabled={name.length===0} type="submit" className="btn btn-danger btn-block">Create Party</button>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    )
}

export default CreateParty;