import React, { Component } from 'react'
import CreateParty from '../../components/InitialSetup/CreateParty';
import JoinParty from '../../components/InitialSetup/JoinParty';
import Waitroom from '../../components/InitialSetup/Waitroom';
import { STATE_CREATE_PARTY, STATE_JOIN_PARTY, STATE_WAIT_ROOM, STATE_MENU, STATE_ORDER, VENDOR_ID } from '../../util/Constants';
import Menu from '../Menu/Menu';
import Orders from '../Orders/Orders';
import './Client.css'

class Client extends Component {

    constructor(props) {
        super(props);
        this.state = {
            id: props.id,
            showing: 0,
            data: {}
        };
    }

    updateShowingState(newState, data) {
        this.setState({
            showing: newState,
            data: data
        })
    }

    wrapComponent(comp) {
        return <div className="client-container">{comp}</div>
    }

    render() {
        switch(this.state.showing) {
            case STATE_CREATE_PARTY:
                return this.wrapComponent(<CreateParty id={this.state.id} vendorId={VENDOR_ID} callback={(s) => this.updateShowingState(s)} />)
            case STATE_JOIN_PARTY:
                return this.wrapComponent(<JoinParty id={this.state.id} vendorId={VENDOR_ID} callback={(s) => this.updateShowingState(s)} />)
            case STATE_WAIT_ROOM:
                return this.wrapComponent(<Waitroom id={this.state.id} vendorId={VENDOR_ID} callback={(s) => this.updateShowingState(s)} {...this.state.data} />)
            case STATE_MENU:
                return this.wrapComponent(<Menu id={this.state.id} vendorId={VENDOR_ID} callback={(s) => this.updateShowingState(s)} />)
            case STATE_ORDER:
                return this.wrapComponent(<Orders id={this.state.id} vendorId={VENDOR_ID} callback={(s) => this.updateShowingState(s)} />)
            default:
                return this.wrapComponent(<CreateParty id={this.state.id} vendorId={VENDOR_ID} callback={(s) => this.updateShowingState(s)} />)
        }
    }
}

export default Client;