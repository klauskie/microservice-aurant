import React from 'react'
import CategoryTableRow from './CategoryTableRow';
import MenuTableRow from './MenuTableRow';

const MenuTable = ({ list, itemCallback }) => {
    return (
        <div className="container">
            <ul>
                {
                    list.map((categoryItem, index) => {
                        return (
                            <div key={index}>
                                <CategoryTableRow categoryId={categoryItem.categoryId} name={categoryItem.name} key={index} />
                                {
                                    categoryItem.menuItems.map((menuItem, index) => (
                                        <MenuTableRow 
                                            itemID={menuItem.itemId} 
                                            name={menuItem.name} 
                                            description={menuItem.description} 
                                            price={menuItem.price} 
                                            isAvailable={menuItem.isAvailable} 
                                            isDisplayable={menuItem.isDisplayable} 
                                            itemCallback={itemCallback} 
                                            key={index} 
                                        />
                                    ))
                                }
                            </div>     
                        )   
                    })
                }
            </ul>

        </div>
    )
}

export default MenuTable;