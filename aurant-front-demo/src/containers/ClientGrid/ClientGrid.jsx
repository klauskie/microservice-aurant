import React, { useEffect, useState } from 'react'
import Client from '../Clients/Client';
import ErrorBoundary from '../ErrorBoundary/ErrorBoundary';
import EmptyClient from './EmptyClient';
import './ClientGrid.css'

const ClientGrid = () => {

    const [list, setList] = useState([<Client id={0}/>])

    useEffect(() => {
        const eList = [<EmptyClient addToGridCallback={() => appendClient()} />]
        setList([...list, ...eList])
    }, [])

    const appendClient = () => {
        console.log('appendClient called')

        const nId = list.length

        setList(old => [...old.slice(0, old.length-1), <Client id={nId} />, old[old.length-1]])
    }

    return (
        <div className="client-grid">
            {
                list.map((client, index) => {
                    return (
                        <div className="item" key={index} >
                            <ErrorBoundary>
                                {client}
                            </ErrorBoundary>
                        </div>
                    )
                })
            }
        </div>
    )
}

export default ClientGrid;