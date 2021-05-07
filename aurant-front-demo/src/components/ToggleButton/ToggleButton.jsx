import React, { useState } from 'react'
import './ToggleButton.css'

const ToggleButton = ( {callback} ) => {

    const [toggle, setToggle] = useState(true);

    const triggerToggle = () => {
        setToggle( !toggle )
        callback(toggle)
    }

    return(
        <div onClick={triggerToggle} className={`wrg-toggle ${toggle ? 'wrg-toggle--checked' : ''}`}>
            <div className={`wrg-toggle-container ${toggle ? 'wrg-toggle-container-checked' : ''}`}>
                <div className="wrg-toggle-check">
                    <span></span>
                </div>
                <div className="wrg-toggle-uncheck">
                    <span></span>
                </div>
            </div>
            <div className="wrg-toggle-circle"></div>
            <input className="wrg-toggle-input" name="flexSwitchCheckChecked" type="checkbox" aria-label="Toggle Button" />
            <label className="wrg-toggle-label" for="flexSwitchCheckChecked">All</label>
        </div>
    )
}

export default ToggleButton;