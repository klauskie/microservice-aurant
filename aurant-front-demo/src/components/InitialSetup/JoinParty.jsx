import React, { useState } from 'react'
import { useCookies } from 'react-cookie'
import { SIX_HOURS, KEY_VENDOR_ID, getPartyIdKey, STATE_WAIT_ROOM, STATE_CREATE_PARTY, getTokenKey } from '../../util/Constants';
import { guestLoginPOST, joinPartyPUT } from '../../util/APIutils';
import './InitialSetup.css';


const JoinParty = (props) => {
    const [, setCookie] = useCookies([getTokenKey(props.id), getPartyIdKey(props.id)])
    
    let [name, setName] = useState("");
    let [tag, setTag] = useState("");
    let [, setToken] = useState("");

    let vendorId = props.vendorId

    const joinParty = () => {
        if (vendorId === "") {
            return
        }
        // Login as guest
        handleGuestLogin()
    }

    const handleGuestLogin = () => {
        guestLoginPOST(name, vendorId).then(loginData => {
            setToken(loginData.token)
            handleJoinParty(loginData.token)
        }).catch(() => {});
    }

    const handleJoinParty = (mToken) => {
        joinPartyPUT(tag, mToken).then(partyData => {
            let partyTag = partyData.party.tag
            saveCookies(mToken, partyTag)

            redirectHelper(STATE_WAIT_ROOM, { 
                token: mToken,
                tag: partyTag
            })
        }).catch(() => {});
    }

    const saveCookies = (mToken, partyId) => {
        setNewCookie(getTokenKey(props.id), mToken)
        setNewCookie(getPartyIdKey(props.id), partyId)
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
                    <a onClick={() => props.callback(STATE_CREATE_PARTY, {})} className="a-tag" >Create Party</a>
                </div>

                <div className="">
                    <h2 className="left">Join in!</h2>
                </div>

                <div className="row justify-content-md-center h-100">
                    <div className="card-body">
                        <div className="form-group">
                            <label htmlFor="name">Name</label>
                            <input autoComplete='off' onChange={(e) => setName(e.target.value)} id="name" type="text" className="form-control" name="name" required autofocus/>
                        </div>

                        <div className="form-group">
                            <label htmlFor="party-tag">Party TAG</label>
                            <input autoComplete='off' onInput={(e) => e.target.value = e.target.value.toUpperCase()} onChange={(e) => setTag(e.target.value.toUpperCase())} id="party-tag" type="text" className="form-control" name="party-tag" required data-eye/>
                        </div>

                        <div className="form-group">
                            <button onClick={(e) => joinParty()} disabled={name.length===0 || tag.length===0} type="submit" className="btn btn-danger btn-block">Join Party</button>
                        </div>

                    </div>
                </div>
            
            </div>
        </div>
    )
}

export default JoinParty;