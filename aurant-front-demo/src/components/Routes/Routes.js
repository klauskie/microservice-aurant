import React, { Component } from 'react'
import { Switch, Route } from 'react-router-dom'
import ClientGrid from '../../containers/ClientGrid/ClientGrid';
import Menu from '../../containers/Menu/Menu';
import Orders from '../../containers/Orders/Orders';
import CreateParty from '../InitialSetup/CreateParty';
import InitialSelection from '../InitialSetup/InitialSelection';
import JoinParty from '../InitialSetup/JoinParty';
import Waitroom from '../InitialSetup/Waitroom';

class Routes extends Component {
    render() {
        return (
            <Switch>
                <Route exact path="/" component={InitialSelection} />
                <Route exact path="/create-party/:vendorId" component={CreateParty} />
                <Route exact path="/join-party/:vendorId" component={JoinParty} />
                <Route exact path="/waitroom" render={(props) => <Waitroom {...props}/>}/>
                <Route exact path="/menu" component={Menu} />
                <Route exact path="/orders" component={Orders} />
                <Route exact path="/test" component={ClientGrid} />
            </Switch>
        );
    }
}

export default Routes