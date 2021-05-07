import React from 'react'
import { FiShoppingCart } from 'react-icons/fi';
import { AiOutlineHome } from 'react-icons/ai';
import './TopMenu.css'; 
import { STATE_MENU, STATE_ORDER } from '../../util/Constants';


const TopMenu = ( {tagId, callback} ) => {
    return (
        <div>
            <nav className="navbar navbar-expand-sm navbar-light bg-light">
            <ul className="nav mr-auto">
                    <li className="nav-item">
                        <h3>{tagId}</h3>
                    </li>
                </ul>
                <ul className="nav ml-auto">
                    <li className="nav-item" onClick={() => callback(STATE_MENU, {}) }>
                        <span className="nav-link">
                            <AiOutlineHome size={30} color="gray" />
                        </span>
                    </li>
                    <li className="nav-item" onClick={() => callback(STATE_ORDER, {}) }>
                        <span className="nav-link">
                            <FiShoppingCart size={30} color="gray" />
                        </span>
                    </li>
                </ul>
            </nav>
        </div>
    )
};

export default TopMenu