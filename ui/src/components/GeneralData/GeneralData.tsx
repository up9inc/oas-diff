import { getStatus } from '../../DataService';
import './GeneralData.sass';

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

const status = getStatus()

export const GeneralData: React.FC<Props> = (props: Props) => {
    return (
        <div className='generalData'>
            <div className='details'>
                {Object.entries(props).map(([key, val]) => {
                    return <div className='item'>
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