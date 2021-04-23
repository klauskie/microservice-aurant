# Party API

## Headers

Every request must contain the header key, <i>token</i> and a valid token string provided at login.

> token : <valid_token_id>

> Content-Type: application/json

## Endpoints

### Party

> GET party: **party-api/party/:partyTag**

> POST new party: **party-api/party**

> PUT join party: **party-api/party/:partyTag**

> GET party members: **party-api/party-clients/:partyTag**

> GET party status: **party-api/party-status/:partyTag**

> PUT update status: **party-api/party-status/:partyTag**

> DELETE kick member from party: **party-api/party/:partyTag/kick/:<client-id>**

> DELETE remove party: **party-api/party/:partyTag**


### Order

> POST create order by member: **party-api/order/:partyTag**

> GET order by member: **party-api/order/:partyTag**

> GET complete party order: **party-api/party-order/:partyTag**

> POST send prepare request to order-api: **party-api/prepare-order/:partyTag**


## Errors codes:

HTTP 400: incomplete or invalid parameters provided.

HTTP 404: object not found.

HTTP 401: unauthorized. Missing token.


## Request structure

> GET: party-api/party/:partyTag

> PUT: party-api/party/:partyTag

> Response: Party object. Status: 200

    {
        "message": "GET party",
        "party": {
            "tag": <string>,
            "client_list": [
                {
                    "Id": <client_id:string>,
                    "Name": <string>
                }
            ],
            "restaurant_id": <string>,
            "host": {
                "Id": <client_id:string>,
                "Name": <string>
            },
            "order_map": {
                <client_id:string>: [
                    {
                        "item_id": <item_id:string>,
                        "instructions": <string>,
                        "quantity": <number>,
                        "Owner": {
                            "Id": <client_id:string>,
                            "Name": <string>
                        }
                    }
                ]
            },
            "is_ok": <bool>
        },
        "party-tag": <string>
    }


___


> POST: party-api/party

> Response: Party object. Status: 201

    {
        "message": "Party created successfully",
        "party": {
            "tag": <string>,
            "client_list": [
                {
                    "Id": <client_id:string>,
                    "Name": <string>
                }
            ],
            "restaurant_id": <string>,
            "host": {
                "Id": <client_id:string>,
                "Name": <string>
            },
            "order_map": {},
            "is_ok": <bool>
        },
        "party-tag": <string>
    }

___


> GET: party-api/party-clients/:partyTag

> Response: Client list object. Status: 200

    {
        "clients": [
            {
                "Id": <client:id:string>,
                "Name": <string>
            },
            {
                "Id": <client:id:string>,
                "Name": <string>
            }
        ],
        "message": "GET party",
        "party-tag": <string>
    }


___


> PUT: party-api/party-status/:partyTag

> Provide the following structure in body:

    {
        "ready": true
    }

> Response. Status: 202

    {
        "message": "Party status updated",
        "party-tag": "NOCW",
        "status": true
    }

___


> POST: party-api/order/:partyTag

> Provide the following structure in body:

    {
        "item_id": "222",
        "instructions": "no cheese",
        "quantity": 2
    }

> Response: Order object from member. Status: 202

    {
        "Orders": [
            {
                "item_id": "222",
                "instructions": "no cheese",
                "quantity": 2,
                "Owner": {
                    "Id": <client_id:string>,
                    "Name": <string>
                }
            }
        ],
        "message": "Order added",
        "party-tag": "NOCW"
    }


___

GET: party-api/order/:partyTag

> Response: Order object from member. Status: 202

    {
        "Orders": [
            {
                "item_id": "222",
                "instructions": "no cheese",
                "quantity": 2,
                "Owner": {
                    "Id": <client_id:string>,
                    "Name": <string>
                }
            }
        ],
        "message": "Order added",
        "party-tag": "NOCW"
    }


___

> GET: party-api/party-order/:partyTag

> Response: Order list grouped by member. Status: 200

    {
        "Orders": [
            {
                "client": {
                    "Id": <client_id:string>,
                    "Name": <string>
                },
                "order_list": [
                    {
                        "item_id": "222",
                        "instructions": "no cheese",
                        "quantity": 2,
                        "Owner": {
                            "Id": <client_id:string>,
                            "Name": <string>
                        }
                    },
                    {
                        "item_id": "333",
                        "instructions": "napkins",
                        "quantity": 1,
                        "Owner": {
                            "Id": <client_id:string>,
                            "Name": <string>
                        }
                    }
                ]
            }
        ],
        "message": "Complete Order fetched",
        "party-tag": <string>
    }

___

> POST: party-api/prepare-order/:partyTag

> Caller must be HOST

> Response: Party object. Status: 201

    {
        "message": "Orders pushed",
        "party": {
            "tag": <string>,
            "client_list": [
                {
                    "Id": <client_id:string>,
                    "Name": <string>
                }
            ],
            "restaurant_id": <string>,
            "host": {
                "Id": <client_id:string>,
                "Name": <string>
            },
            "order_map": {
                <client_id:string>: [
                    {
                        "item_id": <item_id:string>,
                        "instructions": <string>,
                        "quantity": <number>,
                        "Owner": {
                            "Id": <client_id:string>,
                            "Name": <string>
                        }
                    }
                ]
            },
            "is_ok": <bool>
        },
        "party-tag": <string>
    }