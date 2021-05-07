import React from 'react'
import './Menu.css'; 


const CategoryTableRow = ({ categoryId, name }) => {
    return (
        <li>
            <div>
                <div className="h2-tag">{name}</div>
            </div>
        </li>
    )
};

export default CategoryTableRow