# Catalog Service

## TODO

- Support API KEY
- Support Redis Cache


## Endpoints

### Item

> GET all items: **api/item**

> GET all vendor's items: **api/item/:restaurantID**

> POST create new item: **api/item**

> PUT update item: **api/item**

> POST get item definition map: **api/item-definition**


### Category

> GET all categories: **api/category**

> GET by ID: **api/category/:uuid**

> GET all categories by vendor: **api/category/restaurant/:restaurantID**

> POST create category: **api/category**

> PUT update category: **api/category**

> PUT add item to category: **api/category/add/:itemID**


### Restaurant

> GET all restaurants: **api/restaurant**

> GET by ID: **api/restaurant/:uuid**

> GET all categories by vendor: **api/restaurant/categories/:restaurantID**

> POST create restaurant: **api/restaurant**

> PUT update restaurant: **api/restaurant**

> POST import data by csv: **api/restaurant/upload-csv/:restaurantUUID**