import React, { Component } from 'react'
import TopMenu from '../../components/TopMenu/TopMenu';
import { instanceOf } from 'prop-types';
import { withCookies, Cookies } from 'react-cookie';
import OrderTable from '../../components/OrderList/OrderTable';
import ToggleButton from '../../components/ToggleButton/ToggleButton';
import { getPartyIdKey, getTokenKey, KEY_VENDOR_ID } from '../../util/Constants';
import { fetchPartyOrderGET } from '../../util/APIutils';


class Orders extends Component {
    static propTypes = {
        cookies: instanceOf(Cookies).isRequired
    };

    state = {
        list: [],
        tempList: [],
        token: this.props.cookies.get(getTokenKey(this.props.id)) || "",
        tag: this.props.cookies.get(getPartyIdKey(this.props.id)) || "",
        vendorId: this.props.cookies.get(KEY_VENDOR_ID) || "",
    }

    componentDidMount() {
        this.getPartyOrder()
    }

    getPartyOrder = () => {
        // TODO poll server

        fetchPartyOrderGET(this.state.tag, this.state.token).then(data => {
            if (data.Orders != null) {
                this.setState({
                    list: data.Orders,
                    tempList: data.Orders
                })
            }
        }).catch(() => {})
    }

    toggleClientListView = (toggleValue) => {
        if (!toggleValue) {
            this.setState({
                tempList: this.state.list,
            })
        } else {
            this.setState({
                tempList: this.state.list.slice(0,1), // expects first item is client
            })
        }
    }

    render() {
        return (
            <div>
                <TopMenu tagId={this.state.tag} callback={this.props.callback} />
                <div className="container">
                    <div className="align-right"><ToggleButton callback={this.toggleClientListView} /></div>
                    <OrderTable list={this.state.tempList} />
                </div>
            </div>
        )
    }
}

export default withCookies(Orders);