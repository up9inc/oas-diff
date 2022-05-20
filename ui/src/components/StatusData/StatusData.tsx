import './StatusData.sass';

export interface IIndexable {
    [key: string]: any;
}

const dictionary: IIndexable = {
    baseFile: "Base File",
    secondFile: "Second File",
    executionTime: "Execution Time",
    totalPathChanges: "Total Path Changes",
    flags: "Flags",
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
                            {dictionary[key]}
                        </span>
                        <span className='itemData'>
                            {props[key]}
                        </span>
                    </div>
                })}
            </div>
        </div>
    )
}