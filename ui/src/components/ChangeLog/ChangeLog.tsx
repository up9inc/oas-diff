import './ChangeLog.sass';
import { CollapseButton } from '../../useCommonStyles';

export interface Props {
    onCollapseAll: () => void
}

export const ChangeLog: React.FC<Props> = ({ onCollapseAll }) => {
    return (
        <div className='chagnelogContainer'>
            <div className="endpointChangelog">
                <span className='title'>Endpoints Changelog</span>
                <CollapseButton onClick={onCollapseAll} >Collapse All</CollapseButton>
            </div>
        </div>
    )
}
