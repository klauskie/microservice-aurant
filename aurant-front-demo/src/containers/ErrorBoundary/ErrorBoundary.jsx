import React from 'react'
import { instanceOf } from 'prop-types';
import { withCookies, Cookies } from 'react-cookie';
import './ErrorPage.css';
import errorImg from'./error.png';
import { KEY_ACCESS_TOKEN, KEY_TAG, KEY_VENDOR_ID } from '../../util/Constants';

class ErrorBoundary extends React.Component {
    static propTypes = {
        cookies: instanceOf(Cookies).isRequired
    };

    constructor(props) {
        super(props);
        this.state = { 
            hasError: false,
            token: this.props.cookies.get(KEY_ACCESS_TOKEN) || "",
            tag: this.props.cookies.get(KEY_TAG) || ""
        };
    }
  
    static getDerivedStateFromError(error) {
        // Update state so the next render will show the fallback UI.
        return { hasError: true };
    }
  
    componentDidCatch(error, errorInfo) {
        // You can also log the error to an error reporting service
        //logErrorToMyService(error, errorInfo);
    }

    redirectToSafety() {
        const vendorId = localStorage.getItem(KEY_VENDOR_ID);

        let data = {}
        if ( this.state.token !== "") {
            this.props.history.replace("/menu", data)
        } else if (vendorId != null) {
            this.props.history.replace(`/join-party/${vendorId}`, data)
        }
    }
  
    render() {
        if (!this.state.hasError) {
            return this.props.children; 
        }
  
        return (
            <div id="error-container">
                <div className="container">
                    <div>
                        <img src={errorImg} alt="error parachute icon" />
                    </div>
                </div>
                <div className="content"></div>
                <div className="text-box">
                    <div className="content-inner">
                        <h3>Uh oh!</h3>
                        <p>Something wierd happended.</p>
                        <button onClick={() => this.redirectToSafety()} className="btn btn-danger">Try Again</button>
                    </div>
                </div>
            </div>
        )   
    }
}


export default withCookies(ErrorBoundary);