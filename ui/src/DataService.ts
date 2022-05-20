
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

export const dataService = new DataService(data)