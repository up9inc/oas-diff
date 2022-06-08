import './ChangeLog.sass';
import { CollapseButton } from '../../useCommonStyles';
import { useSetRecoilState } from 'recoil';
import { collapseItemsList } from '../../recoil/collapse/';

export interface Props {
    //onCollapseAll: () => void
}

export const ChangeLog: React.FC<Props> = () => {

    const setAccordions = useSetRecoilState(collapseItemsList)

    const onCollapseAll = () => {
        setAccordions((prev) => {
            return prev.map(x => { return { ...x, isCollapsed: true } })
        })
    }

    return (
        <div className='chagnelogContainer'>
            <div className="endpointChangelog">
                <span className='title'>Endpoints Changelog</span>
                <CollapseButton onClick={onCollapseAll} >Collapse All</CollapseButton>
            </div>
        </div>
    )
}
