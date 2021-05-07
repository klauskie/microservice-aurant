import React, { Component } from 'react'
import { instanceOf } from 'prop-types';
import { withCookies, Cookies } from 'react-cookie';
import { getPartyIdKey, getTokenKey, KEY_VENDOR_ID, STATE_JOIN_PARTY, STATE_MENU } from '../../util/Constants';
import { fetchPartyGET, updatePartyStatusPUT } from '../../util/APIutils';
import './InitialSetup.css';


class WaitRoomComponent extends Component {
    static propTypes = {
        cookies: instanceOf(Cookies).isRequired
    };

    state = {
        pollingCount: 0,
        delay: 3000,
        clients: [],
        tag: this.props.cookies.get(getPartyIdKey(this.props.id)) || "",
        token: this.props.cookies.get(getTokenKey(this.props.id)) || "",
        vendorId: this.props.cookies.get(KEY_VENDOR_ID) || "",
        host: {},
    };

    componentDidMount() {
        this.interval = setInterval(this.tick, this.state.delay)
        this.fetchPartyData()
    }

    componentDidUpdate(prevProps, prevState) {
        if (prevState.delay !== this.state.delay) {
            clearInterval(this.interval)
            this.interval = setInterval(this.tick, this.state.delay)
        } 
    }

    componentWillMount() {
        clearInterval(this.interval)
    }

    tick = () => {
        this.checkForCookie()
        this.setState({
            pollingCount: this.state.pollingCount + 1
        })
        this.fetchPartyData()
    }

    fetchPartyData = () => {
        fetchPartyGET(this.state.tag, this.state.token).then(party => {
            this.setState({
                clients: party.client_list,
                host: party.host
            })
            if (party.is_ok) {
                this.partyReadyNextScreen()
            }
        }).catch(() => {});
    }

    setPartyToReadyState = () => {
        if (!this.isUserHost()) {
            console.log("User not host")
            return
        }

        updatePartyStatusPUT(this.state.tag, this.state.token).then(data => {
            this.partyReadyNextScreen()
        }).catch(() => {})
    }

    isUserHost = () => {
        return this.state.token === this.state.host.Id
    }

    partyReadyNextScreen = () => {
        clearInterval(this.interval)
        this.props.callback(STATE_MENU, {})
    }

    checkForCookie = () => {
        if (this.state.token === "") {
            clearInterval(this.interval)
            this.props.callback(STATE_JOIN_PARTY, {})
            return
        }
    }

    render() {
        return (
            <div>
                <div className="container">
                    <div className="margin-top"></div>
                    <div className="">
                        <h2 className="left">Waitroom</h2>
                    </div>
                    <div className="">
                        <h3 className="left">Share this tag: <span className="color-red">{this.state.tag}</span></h3>
                    </div>

                    <div className="row justify-content-md-center">
                        <div className="card-body">

                            <div className="left">
                                <ul>
                                    {this.state.clients.map((client, index) => {
                                        if (client.Id === this.state.token) {
                                            return <li key={index}><span className="bold color-red">{client.Name}</span></li>
                                        }
                                        return <li key={index}><span className="bold">{client.Name}</span></li>
                                    })}
                                </ul>
                            </div>

                            <div className="form-group">
                                <button onClick={(e) => this.setPartyToReadyState()} disabled={!this.isUserHost()} type="submit" className="btn btn-danger btn-block">Ready!</button>
                            </div>
                        </div>
                    </div>

                </div>
        </div>
        )
    }
}

export default withCookies(WaitRoomComponent);
