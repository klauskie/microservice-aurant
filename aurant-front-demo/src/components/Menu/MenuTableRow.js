import React from 'react'
import './Menu.css'; 


const MenuTableRow = ({ itemID, name, description, price, isAvailable, isDisplayable, itemCallback }) => {

    const onItemClicked = () => {
        let bundle = {
            itemID: itemID,
            name: name,
            description: description,
            price: price
        }
        itemCallback(bundle)
    }

    return (
        <li className={`${isAvailable ? '' : 'disable'} ${isDisplayable ? '' : 'hide'}`}>
            <div className="grid-container item-container" onClick={onItemClicked}>
                <div className="name-block">{name}</div>
                <div className="price-block">${price}</div>
                <div className="description-block">{description.slice(0, 50)}</div>
            </div>
        </li>
    )
};

export default MenuTableRow