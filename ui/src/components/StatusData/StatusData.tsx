import './StatusData.sass';

export interface IIndexable {
    [key: string]: any;
}

const TimeExecutionComponent = (ms: number) => <div className='timeContainer'>
    <span className='time'>{ms}</span>
    <span className='measure'>ms</span>
</div>

const propsToPresentationDictionary: IIndexable = {
    baseFile: { name: "Base File" },
    secondFile: { name: "Second File" },
    executionTime: { name: "Execution Time", component: (val: any) => TimeExecutionComponent(val) },
    totalPathChanges: { name: "Total Path Changes", component: (val: any) => <span className='singleNumberCard'>{val}</span> },
    flags: { name: "Flags", component: (val: any) => <span className='singleNumberCard'>{val}</span> },
}

export interface Props extends IIndexable {
    baseFile: string;
    secondFile: string;
    executionTime: string;
    totalPathChanges: number;
    flags: number
}

export const StatusData: React.FC<Props> = (props: Props) => {
    return (
        <div className='generalData'>
            <div className='details'>
                {Object.entries(props).map(([key, val]) => {
                    return <div key={key} className='item'>
                        <span className='itemTitle'>
                            {propsToPresentationDictionary[key].name}
                        </span>
                        <div className='itemData' title={`${val}`}>
                            {propsToPresentationDictionary[key]?.component ? propsToPresentationDictionary[key].component(val) : val}
                        </div>
                    </div>
                })}
            </div>
        </div>
    )
}
