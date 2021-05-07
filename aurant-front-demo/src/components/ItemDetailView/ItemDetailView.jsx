import React, { useRef, useState } from 'react'
import { AiOutlinePlus, AiOutlineMinus } from 'react-icons/ai';
import './ItemDetailView.css';


const ItemDetailView = ({ itemBundle, orderCallback }) => {

    let quantityInput = useRef(null)
    let [quantity, setQuantity] = useState(1)
    let [instructions, setInstructions] = useState("")

    const onMinusClicked = () => {
        let currentValue = parseInt(quantityInput.current.value)
        if (currentValue > 1) {
            let newVal = currentValue - 1
            quantityInput.current.value = newVal
            setQuantity(newVal)
        }
    }

    const onPlusClicked = () => {
        let currentValue = parseInt(quantityInput.current.value)
        let newVal = currentValue + 1
        quantityInput.current.value = newVal
        setQuantity(newVal)
    }

    const prepareDataOrder = () => {
        let order = {
            item_id: itemBundle.itemID,
            instructions: instructions,
            quantity: quantity
        }
        orderCallback(order)
    }

    return (
        <div className="card-container">
            <div className="container">

                <div className="row name-block row-item ">{itemBundle.name}</div>
                <div className="row row-item ">{itemBundle.description}</div>

                <div className="row row-item ">
                    <input className="col-sm input-box-lr" onChange={(e) => {setInstructions(e.target.value)}} type="text" placeholder="Add instructions" />
                </div>

                <div className="row row-item button-group justify-content-center">
                    <span className="input-group-btn">
                        <button onClick={onMinusClicked} type="button" className="quantity-left-minus btn btn-light btn-number"  datatype="minus" data-field="">
                            <span>
                                <AiOutlineMinus />
                            </span>
                        </button>
                    </span>


                    <input id="quantity-input" ref={quantityInput} value={quantity} onChange={(e) => {}} type="number" className="form-control input-number" min={1} max={100}/>
                    
                    
                    <span className="input-group-btn">
                        <button onClick={onPlusClicked} type="button" className="quantity-right-plus btn btn-light btn-number" datatype="plus" data-field="">
                            <span>
                                <AiOutlinePlus />
                            </span>
                        </button>
                    </span>
                </div>
                
                <div className="row row-item justify-content-center">
                    <button onClick={prepareDataOrder} type="button" className="btn btn-dark">
                        Add to cart: ${itemBundle.price * parseFloat(quantity)}
                    </button>
                </div>
            </div>
        </div>
    )
};

export default ItemDetailView