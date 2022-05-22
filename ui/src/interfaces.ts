type JSONValue =
    | string
    | number
    | boolean
    | JSONObject

interface JSONObject {
    [x: string]: JSONValue;
}

export interface ChangeLog {
    type: string
    path: string[]
    from: JSONObject
    to: JSONObject
}

export interface Path {
    endpoint: string,
    operation: string,
    changelog: ChangeLog
}

export interface DataItemValue {
    key: string
    path: Path[]
    totalChanges: number,
    createdChanges: number,
    updatedChanges: number,
    deletedChanges: number
}

export interface DataItem {
    key: string
    value: DataItemValue
}

export interface Status {
    baseFile: string
    secondFile: string
    executionTime: string
    executionFlags: {
        typeFilter: string
        loose: boolean
        includeFilePath: boolean
        ignoreDescriptions: boolean
        ignoreExamples: boolean
    }
    startTime: string
}

export interface ReaportData {
    data: DataItem[]
    status: Status
}
