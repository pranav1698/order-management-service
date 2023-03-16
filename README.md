# Order Management Service

This is a Golang based micro-service to add, update and get orders

To run the application, clone the repo in your local system:
```
https://github.com/pranav1698/order-management-service.git
```

To build binary executable, run:
```
go build
```

To run the binary executble, run:
```
./order-management-service
```

The service will now run on localhost port 8080,

### Create New Order
To create a new order, open a new terminal and run the following curl command:

```
curl --request POST   --url http://localhost:8080/create  --header 'content-type: application/json'   --data '{"id": "abcdef-123456","status": "PENDING_INVOICE","items": [{"id": "123456 ","description": "a product description","price": 12.40,"quantity": 1}],"total": 12.40,"currencyUnit": "USD"}'
```

You will get the following JSON response, representing the current orders:

```
{
    "id":"abcdef-123456",
    "status":"PENDING_INVOICE",
    "items":[
        {
            "id":"123456 ",
            "description":"a product description","price":12.4,
            "quantity":1
        }],
    "total":12.4,
    "currencyUnit":"USD"
}

```

### Update Status of current order
To update the status of current order, you can run the following command:

```
curl --request PUT   --url http://localhost:8080/update  --header 'content-type: application/json'   --data '{"id": "abcdef-123456","status": "INVOICE_RECIEVED"}'
```

You will get the following JSON response, representing the updated order:

```
[
    {
        "id":"abcdef-123456",
        "status":"INVOICE_RECIEVED",
        "items":[
            {
                "id":"123456 ",
                "description":"a product description","price":12.4,
                "quantity":1
            }
        ],
        "total":12.4,
        "currencyUnit":"USD"
    }
]

```

### Get All Orders
To get all orders present currently, run the following command:

```
curl --request GET --url http://localhost:8080/orders
```

Response:

```
[
    {
        "id":"abcdef-123456",
        "status":"INVOICE_RECIEVED",
        "items":[
            {
                "id":"123456 ",
                "description":"a product description","price":12.4,
                "quantity":1
            }
        ],
        "total":12.4,
        "currencyUnit":"USD"
    },
    {
        "id":"abcdef-12345786",
        "status":"PENDING_INVOICE",
        "items":[
            {
                "id":"123456 ",
                "description":"a product description","price":12.4,
                "quantity":1
            }
        ],
        "total":12.4,
        "currencyUnit":"USD"
    }
]

```

### Get Order filtered by order fields
To get orders filtered by order keys, run the following command

```
curl --request GET --url http://localhost:8080/orders/id/abcdef-123456
```

Response:

```
[
    {
        "id":"abcdef-123456",
        "status":"INVOICE_RECIEVED",
        "items":
        [
            {
                "id":"123456 ",
                "description":"a product description","price":12.4,
                "quantity":1
            }
        ],
        "total":12.4,
        "currencyUnit":"USD"
    }
]

```

You can get orders by other filters like status, currency unit and total by changing the above URL.

### Running the application inside docker container

Building docker image for the above application using Dockerfile present in the project,

```
docker build -t order-management-service .
```

To verify that our image exists on out machine run,

```
docker images
```

Output:

```
REPOSITORY                 TAG       IMAGE ID       CREATED              SIZE
order-management-service   latest    91d7093dc66a   About a minute ago   238MB
```

To run this newly created image, we can use the following command:

```
docker run -p 8080:8080 -it order-management-service
```

Now, you can run the above curl requests in another terminal and get the same output