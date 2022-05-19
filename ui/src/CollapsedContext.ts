import React from "react";

type ObjectDescriptor<D, M> = {
    data?: D;
    methods?: M & ThisType<D & M>; // Type of 'this' in methods is D & M
};

function makeObject<D, M>(desc: ObjectDescriptor<D, M>): D & M {
    let data: object = desc.data || {};
    let methods: object = desc.methods || {};
    return { ...data, ...methods } as D & M;
}

const a = makeObject({
    // data: {
    //     collapsed: false,
    //     accordions: []
    // },
    // methods: {
    //     moveBy(accordion: any) {
    //         (this.accordions as []{}) .push(accordion)
    //     }
    // }
})

export const CollapsedContext = React.createContext(a); 