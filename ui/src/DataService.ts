const data = (window as { [key: string]: any })["reportData"] as Object;

export const getData = () => {
    return {
        "execution-status": {
            "base-file": "/home/gustavomassa/projects/up9/oas-diff/examples/simple.json",
            "second-file": "/home/gustavomassa/projects/up9/oas-diff/examples/simple2.json",
            "start-time": "May 17 10:18:37.835",
            "execution-time": "9.779455ms",
            "execution-flags": {
                "type-filter": "",
                "loose": false,
                "include-file-path": false,
                "ignore-descriptions": false,
                "ignore-examples": false
            }
        },
        "changelog": {
            "components": [],
            "externalDocs": [],
            "info": [
                {
                    "type": "update",
                    "path": [
                        "title"
                    ],
                    "from": "Simple Example",
                    "to": "simple example"
                },
                {
                    "type": "update",
                    "path": [
                        "version"
                    ],
                    "from": "1.0.0",
                    "to": "1.1.0"
                }
            ],
            "paths": [
                {
                    "type": "delete",
                    "path": [
                        "/users",
                        "get",
                        "parameters",
                        "x-accept"
                    ],
                    "identifier": {
                        "name": "x-accept"
                    },
                    "from": {
                        "name": "x-accept",
                        "in": "header",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    "to": null
                },
                {
                    "type": "delete",
                    "path": [
                        "/users",
                        "get",
                        "parameters",
                        "Id"
                    ],
                    "identifier": {
                        "name": "Id"
                    },
                    "from": {
                        "name": "Id",
                        "in": "path",
                        "required": true,
                        "style": "simple",
                        "schema": {
                            "type": "string",
                            "pattern": ".+(_|-|\\.).+"
                        },
                        "example": "some-uuid-maybe"
                    },
                    "to": null
                },
                {
                    "type": "create",
                    "path": [
                        "/users",
                        "get",
                        "parameters",
                        "x-custom"
                    ],
                    "identifier": {
                        "name": "x-custom"
                    },
                    "from": null,
                    "to": {
                        "name": "x-custom",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        }
                    }
                },
                {
                    "type": "create",
                    "path": [
                        "/users",
                        "get",
                        "responses",
                        "200"
                    ],
                    "from": null,
                    "to": {
                        "description": "A simple string response",
                        "headers": {
                            "X-Rate-Limit-Limit": {
                                "description": "The number of allowed requests in the current period",
                                "schema": {
                                    "type": "integer"
                                }
                            },
                            "X-Rate-Limit-Remaining": {
                                "description": "The number of remaining requests in the current period",
                                "schema": {
                                    "type": "integer"
                                }
                            },
                            "X-Rate-Limit-Reset": {
                                "description": "The number of seconds left in the current period",
                                "schema": {
                                    "type": "integer"
                                }
                            }
                        },
                        "content": {
                            "text/plain": {
                                "schema": {
                                    "type": "string",
                                    "example": "whoa!"
                                }
                            }
                        }
                    }
                },
                {
                    "type": "update",
                    "path": [
                        "/users",
                        "parameters",
                        "id",
                        "example"
                    ],
                    "identifier": {
                        "name": "id"
                    },
                    "from": "custom uuid",
                    "to": "some example"
                },
                {
                    "type": "delete",
                    "path": [
                        "/users",
                        "parameters",
                        "accept"
                    ],
                    "identifier": {
                        "name": "accept"
                    },
                    "from": {
                        "name": "accept",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "to": null
                }
            ],
            "security": [],
            "servers": [
                {
                    "type": "delete",
                    "path": [
                        "https://test.com"
                    ],
                    "identifier": {
                        "url": "https://test.com"
                    },
                    "from": {
                        "url": "https://test.com",
                        "description": "some description"
                    },
                    "to": null
                },
                {
                    "type": "delete",
                    "path": [
                        "http://refael.shipping.sock-shop"
                    ],
                    "identifier": {
                        "url": "http://refael.shipping.sock-shop"
                    },
                    "from": {
                        "url": "http://refael.shipping.sock-shop",
                        "description": "refael up9-demo-link all"
                    },
                    "to": null
                },
                {
                    "type": "create",
                    "path": [
                        "HTTP://REFAEL.SHIPPING.SOCK-SHOP"
                    ],
                    "identifier": {
                        "url": "HTTP://REFAEL.SHIPPING.SOCK-SHOP"
                    },
                    "from": null,
                    "to": {
                        "url": "HTTP://REFAEL.SHIPPING.SOCK-SHOP",
                        "description": "REFAEL UP9-DEMO-LINK ALLa"
                    }
                },
                {
                    "type": "create",
                    "path": [
                        "http://gustavo.shipping.sock-shop"
                    ],
                    "identifier": {
                        "url": "http://gustavo.shipping.sock-shop"
                    },
                    "from": null,
                    "to": {
                        "url": "http://gustavo.shipping.sock-shop",
                        "description": "gustavo up9-demo-link all"
                    }
                },
                {
                    "type": "create",
                    "path": [
                        "https://test2.com"
                    ],
                    "identifier": {
                        "url": "https://test2.com"
                    },
                    "from": null,
                    "to": {
                        "url": "https://test2.com",
                        "description": "some description 2"
                    }
                }
            ],
            "tags": [],
            "webhooks": []
        }
    };
}