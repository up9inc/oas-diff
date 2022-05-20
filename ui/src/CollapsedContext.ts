import React from "react";

export interface IAccordion {
    isCollpased: boolean,
    id: string
}

const state = {
    accordions: ([] as IAccordion[]),
    setAccordions: {} as React.Dispatch<React.SetStateAction<IAccordion[]>>
}

export const CollapsedContext = React.createContext(state); 