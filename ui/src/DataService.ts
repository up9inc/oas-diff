import { DataItem, ReaportData } from "./interfaces";

const dataFromWindow = (window as { [key: string]: any })["reportData"] as ReaportData;
!dataFromWindow && console.error("No Data Found on Window.")

class DataService {
    private windowData: ReaportData;

    constructor(data: ReaportData) {
        this.windowData = data
    }

    getData() {
        return this.windowData?.data
    }

    getStatus() {
        return this.windowData?.status
    }

    getTotalChanged() {
        return this.windowData?.data?.reduce((previousValue: number, current: DataItem) => {
            return previousValue + current.value.totalChanges
        }, 0)
    }
}

export const dataService = new DataService(dataFromWindow)
