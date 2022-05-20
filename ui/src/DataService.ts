
const dataBase = {
    "data": [{
        "Key": "/orders", "Value": {
            "Key": "paths", "Paths":
                [{
                    "Endpoint": "/orders", "Operation": "post", "Changelog":
                    {
                        "type": "update", "path": ["/orders", "post", "description"]
                        , "from": "Mizu observed 5001 entries (468 failed), at 0.434 hits/s, average response time is 0.003 seconds",
                        "to": "Leon"
                    }
                }, { "Endpoint": "/orders", "Operation": "post", "Changelog": { "type": "update", "path": ["/orders", "post", "responses", "201", "description"], "from": "Successful call with status 201", "to": "Successful call with status 201 asdasd" } }], "TotalChanges": 2, "CreatedChanges": 0, "UpdatedChanges": 2, "DeletedChanges": 0
        }
    }, { "Key": "/login", "Value": { "Key": "paths", "Paths": [{ "Endpoint": "/login", "Operation": "get", "Changelog": { "type": "delete", "path": ["/login", "get"], "from": { "get": { "summary": "/login", "description": "Mizu observed 5001 entries (0 failed), at 0.435 hits/s, average response time is 0.002 seconds", "operationId": "7965bcf9-e94a-41c2-9955-c8b1ac1f57ac", "responses": { "200": { "description": "Successful call with status 200", "content": { "text/html": { "example": "Cookie is set" } } } } } }, "to": null } }], "TotalChanges": 1, "CreatedChanges": 0, "UpdatedChanges": 0, "DeletedChanges": 1 } }, { "Key": "/loginFor", "Value": { "Key": "paths", "Paths": [{ "Endpoint": "/loginFor", "Operation": "get", "Changelog": { "type": "create", "path": ["/loginFor", "get"], "from": null, "to": { "get": { "summary": "/login", "description": "Mizu observed 5001 entries (0 failed), at 0.435 hits/s, average response time is 0.002 seconds", "operationId": "7965bcf9-e94a-41c2-9955-c8b1ac1f57ac", "responses": { "200": { "description": "Successful call with status 200", "content": { "text/html": { "example": "Cookie is set" } } } } } } } }], "TotalChanges": 1, "CreatedChanges": 1, "UpdatedChanges": 0, "DeletedChanges": 0 } }, { "Key": "/catalogue", "Value": { "Key": "paths", "Paths": [{ "Endpoint": "/catalogue", "Operation": "get", "Changelog": { "type": "delete", "path": ["/catalogue", "get"], "from": { "get": { "summary": "/catalogue", "description": "Mizu observed 9999 entries (0 failed), at 0.218 hits/s, average response time is 0.052 seconds", "operationId": "33fcf7ed-87ae-4817-82f2-c593ba2c86e6", "responses": { "200": { "description": "Successful call with status 200", "content": { "": { "example": [{ "count": 1, "description": "Socks fit for a Messiah. You too can experience walking in water with these special edition beauties. Each hole is lovingly proggled to leave smooth edges. The only sock approved by a higher power.", "id": "03fef6ac-1896-4ce8-bd69-b798f85c6e0b", "imageUrl": ["/catalogue/images/holy_1.jpeg", "/catalogue/images/holy_2.jpeg"], "name": "Holy", "price": 99.99, "tag": ["magic", "action"] }, { "count": 438, "description": "proident occaecat irure et excepteur labore minim nisi amet irure", "id": "3395a43e-2d88-40de-b95f-e00e1502085b", "imageUrl": ["/catalogue/images/colourful_socks.jpg", "/catalogue/images/colourful_socks.jpg"], "name": "Colourful", "price": 18, "tag": ["blue", "brown"] }, { "count": 820, "description": "Ready for action. Engineers: be ready to smash that next bug! Be ready, with these super-action-sport-masterpieces. This particular engineer was chased away from the office with a stick.", "id": "510a0d7e-8e83-4193-b483-e27e09ddc34d", "imageUrl": ["/catalogue/images/puma_1.jpeg", "/catalogue/images/puma_2.jpeg"], "name": "SuperSport XL", "price": 15, "tag": ["black", "sport", "formal"] }, { "count": 738, "description": "A mature sock, crossed, with an air of nonchalance.", "id": "808a2de1-1aaa-4c25-a9b9-6612e8f29a38", "imageUrl": ["/catalogue/images/cross_1.jpeg", "/catalogue/images/cross_2.jpeg"], "name": "Crossed", "price": 17.32, "tag": ["red", "formal", "blue", "action"] }, { "count": 808, "description": "enim officia aliqua excepteur esse deserunt quis aliquip nostrud anim", "id": "819e1fbf-8b7e-4f6d-811f-693534916a8b", "imageUrl": ["/catalogue/images/WAT.jpg", "/catalogue/images/WAT2.jpg"], "name": "Figueroa", "price": 14, "tag": ["formal", "blue", "green"] }, { "count": 175, "description": "consequat amet cupidatat minim laborum tempor elit ex consequat in", "id": "837ab141-399e-4c1f-9abc-bace40296bac", "imageUrl": ["/catalogue/images/catsocks.jpg", "/catalogue/images/catsocks2.jpg"], "name": "Cat socks", "price": 15, "tag": ["green", "brown", "formal"] }, { "count": 115, "description": "For all those leg lovers out there. A perfect example of a swivel chair trained calf. Meticulously trained on a diet of sitting and Pina Coladas. Phwarr...", "id": "a0a4f044-b040-410d-8ead-4de0446aec7e", "imageUrl": ["/catalogue/images/bit_of_leg_1.jpeg", "/catalogue/images/bit_of_leg_2.jpeg"], "name": "Nerd leg", "price": 7.99, "tag": ["skin", "blue"] }, { "count": 801, "description": "We were not paid to sell this sock. It's just a bit geeky.", "id": "d3588630-ad8e-49df-bbd7-3167f7efb246", "imageUrl": ["/catalogue/images/youtube_1.jpeg", "/catalogue/images/youtube_2.jpeg"], "name": "YouTube.sock", "price": 10.99, "tag": ["geek", "formal"] }, { "count": 127, "description": "Keep it simple.", "id": "zzz4f044-b040-410d-8ead-4de0446aec7e", "imageUrl": ["/catalogue/images/classic.jpg", "/catalogue/images/classic2.jpg"], "name": "Classic", "price": 12, "tag": ["green", "brown"] }] } } } } } }, "to": null } }], "TotalChanges": 1, "CreatedChanges": 0, "UpdatedChanges": 0, "DeletedChanges": 1 } }], "status": { "baseFile": "/Users/leon/dev/oas-diff/swagger.json", "secondFile": "/Users/leon/dev/oas-diff/swagger copy.json", "startTime": "May 19 12:09:39.561", "executionTime": "62.968542ms", "executionFlags": { "typeFilter": "", "loose": false, "includeFilePath": false, "ignoreDescriptions": false, "ignoreExamples": false } }
}


const data = (window as { [key: string]: any })["reportData"] as Object;

class DataService {
    private data: any;

    constructor(data: any) {
        this.data = data
    }

    getData() {
        return this.data.data
    }

    getStatus() {
        return this.data.status
    }

    getTotalChanged() {
        return this.data.data.reduce((previousValue: number, current: number, index: number, array: Array<any>) => {
            return previousValue + array[index].Value.TotalChanges
        }, 0)
    }
}

export const dataService = new DataService(dataBase)