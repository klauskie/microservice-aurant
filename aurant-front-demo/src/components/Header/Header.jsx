import React from 'react'
import './Header.css'

const Header = () => {
    return (
        <div className="header-container">
            <h2 className="">Aurant Demo!</h2>
            <div className="description-container">
                <h3 id="description-label" >Description: </h3>
                <p id="description"> Food ordering system for you and your friends inside the restaurant. </p>
            </div>
        </div>
    )
}

export default Header;