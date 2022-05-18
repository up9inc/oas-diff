
import { Button } from '@material-ui/core';
import './ChangeLog.sass';
import { useCommonStyles } from '../../useCommonStyles';

export interface Props {

}

export const ChangeLog: React.FC<Props> = (props: Props) => {
    const commonClasses = useCommonStyles()
    return (
        <div className='chagnelogContainer'>
            <div className="endpointChangelog">
                <span className='title'>Endpoints Changelog</span>
                <Button className={`${commonClasses.button}`} onClick={() => { }} >Collapse All</Button>
            </div>
        </div>
    )
}