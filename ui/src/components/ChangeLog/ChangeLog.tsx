
import './ChangeLog.sass';
import { CollapseButton } from '../../useCommonStyles';

export interface Props {

}

export const ChangeLog: React.FC<Props> = (props: Props) => {
    return (
        <div className='chagnelogContainer'>
            <div className="endpointChangelog">
                <span className='title'>Endpoints Changelog</span>
                <CollapseButton onClick={() => { }} >Collapse All</CollapseButton>
            </div>
        </div>
    )
}