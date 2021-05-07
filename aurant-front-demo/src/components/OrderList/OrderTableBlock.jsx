import React from 'react'
import './Order.css'; 
import OrderItem from './OrderItem';


const OrderTableBlock = ({ orderOwner }) => {

    const calcClientTotal = () => {
        let total = 0;
        for (let i = 0; i < orderOwner.order_list.length; i++) {
            let item = orderOwner.order_list[i];
            total += parseFloat(item.metadata.quantity) * parseFloat(item.catalog_item.price);
        }
        return total.toFixed(2);
    }

    return (
        <div className="vertical-filler">
            <li>
                <div className="flex-row-container">
                    <div className="h3-tag left-f">{orderOwner.client.Name}</div>
                    <div className="h3-tag right-f font-smaller-20 attention-color-red">${calcClientTotal()}</div>
                </div>
            </li>

            <ul className="flex-col-container">
                {
                    orderOwner.order_list.map((item, index) => {
                        return (
                            <div className="flex-item" key={index}>
                                <div className="item-flex-row-s">
                                    <span className="v-line"></span>
                                </div>
                                <div className="item-flex-row-l">
                                    <OrderItem item={item} key={index}/>
                                </div>
                            </div>
                        )
                    })
                }
            </ul>            
        </div>
    )
};

export default OrderTableBlock