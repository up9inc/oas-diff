import { atom } from "recoil"

export interface IAccordion {
    isCollapsed: boolean,
    id: string
}

const mainAccordionsList = atom({
    key: "mainAccordionsList",
    default: [] as IAccordion[]
})

const subAccordionsList = atom({
    key: "subAccordionsList",
    default: [] as IAccordion[]
})

export { mainAccordionsList, subAccordionsList };


