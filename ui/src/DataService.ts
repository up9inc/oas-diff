const data = (window as { [key: string]: any })["reportData"] as Object;

export const getData = () => {
    return { "data": [{ "Key": "/orders", "Value": 
            { "Key": "paths", "Paths":
             [{ "Endpoint": "/orders", "Operation": "post", "Changelog": { "type": "delete", "path": ["/orders", "post"], "from": { "post": { "summary": "/orders", "description": "Mizu observed 123301 entries (13905 failed), at 1.065 hits/s, average response time is 0.018 seconds", "operationId": "ceece857-1bd2-433c-8dfe-f95460ca5d01", "requestBody": { "description": "Generic request body", "content": { "application/json": { "example": { "address": "[REDACTED]", "card": "http://user/cards/57a98d98e4b00679b4a830b1", "customer": "http://user/customers/57a98d98e4b00679b4a830b2", "items": "http://carts/carts/57a98d98e4b00679b4a830b2/items" } } }, "required": true }, "responses": { "201": { "description": "Successful call with status 201", "content": { "application/json": { "example": { "address": { "city": "Glasgow", "country": "[REDACTED]", "id": null, "number": "246", "postcode": "G67 3DL", "street": "Whitelees Road" }, "card": { "ccv": "958", "expires": "08/19", "id": null, "longNum": "5544154011345918" }, "customer": { "addresses": "[REDACTED]", "cards": [], "firstName": "[REDACTED]", "id": null, "lastName": "[REDACTED]", "username": "[REDACTED]" }, "customerId": "57a98d98e4b00679b4a830b2", "date": "2022-05-15T14:06:23.168+0000", "id": "628108dff7006000071971df", "items": [{ "id": "628108df31f3780007258e6c", "itemId": "819e1fbf-8b7e-4f6d-811f-693534916a8b", "quantity": 1, "unitPrice": 14 }], "shipment": { "id": "74d1de4c-d0d1-4b75-9b34-729b0c5d3e79", "name": "57a98d98e4b00679b4a830b2" }, "total": 18.99 } } } }, "406": { "description": "Failed call with status 406", "content": { "application/json": { "example": { "error": "Not Acceptable", "exception": "works.weave.socks.orders.controllers.OrdersController$PaymentDeclinedException", "message": "Payment declined: amount exceeds 100.00", "path": "/orders", "status": 406, "timestamp": 1652623585235 } } } } } } }, "to": null } }], "TotalChanges": 1, "CreatedChanges": 0, "UpdatedChanges": 0, "DeletedChanges": 1 } }], "status": { "base-file": "/Users/leon/dev/oas-diff/swagger.json", "second-file": "/Users/leon/dev/oas-diff/swagger copy.json", "start-time": "May 18 15:37:57.395", "execution-time": "50.6345ms", "execution-flags": { "type-filter": "", "loose": false, "include-file-path": false, "ignore-descriptions": false, "ignore-examples": false } } }
}

export const getStatus = (): any => getData().status