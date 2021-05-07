import React from 'react'
import './Order.css'; 


const OrderItem = ({ item }) => {

    const calcPrice = () => {
        let price = parseFloat(item.metadata.quantity) * parseFloat(item.catalog_item.price);
        return price.toFixed(2)
    }

    return (
        <li>
            <div className="grid-order-container order-container">
                <div className="A-block">{item.catalog_item.name}</div>
                <div className="C-block">{item.catalog_item.description}</div>
                <div className="D-block"><span className="italic second-color">Instructions:</span> {item.metadata.instructions}</div>
                <div className="B-block">
                    <span>
                        {item.metadata.quantity} x ${item.catalog_item.price} = <span className="attention-color">${calcPrice()}</span>
                    </span> 
                </div>
            </div>
        </li>
    )
};

export default OrderItem