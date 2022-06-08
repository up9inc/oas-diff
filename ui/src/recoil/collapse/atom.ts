import { atom } from "recoil"

export interface IAccordion {
    isCollapsed: boolean,
    id: string
}

const collapseItemsList = atom({
    key: "collapseItemsList",
    default: [] as IAccordion[]
})

const collapseSubItemsList = atom({
    key: "collapseSubItemsList",
    default: [] as IAccordion[]
})

export { collapseItemsList, collapseSubItemsList };


