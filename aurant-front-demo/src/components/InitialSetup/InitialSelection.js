import React from 'react'
import { Link } from 'react-router-dom';

const InitialSelection = () => {
    return (
        <div>
            <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
                <ul className="navbar-nav mr-auto">
                    <li className="nav-item">
                        <Link to='/create-party/74d760b9-83cc-4baa-bdbb-9e07debb58e1' className="nav-link">Create Party</Link>
                    </li>
                    <li className="nav-item">
                        <Link to='/join-party/74d760b9-83cc-4baa-bdbb-9e07debb58e1' className="nav-link">Join Party</Link>
                    </li>
                    <li className="nav-item">
                        <Link to='/menu' className="nav-link">Menu</Link>
                    </li>
                </ul>
            </nav>
        </div>
    )
}

export default InitialSelection;