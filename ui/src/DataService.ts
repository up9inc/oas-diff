
const dataFromWindow = (window as { [key: string]: any })["reportData"] as Object;
!dataFromWindow && console.error("No Data Found on Window.")

class DataService {
    private windowData: any;

    constructor(data: any) {
        this.windowData = data
    }

    getData() {
        return this.windowData?.data
    }

    getStatus() {
        return this.windowData?.status
    }

    getTotalChanged() {
        return this.windowData?.data?.reduce((previousValue: number, current: number, index: number, array: Array<any>) => {
            return previousValue + array[index].Value.TotalChanges
        }, 0)
    }
}

export const dataService = new DataService(dataFromWindow)