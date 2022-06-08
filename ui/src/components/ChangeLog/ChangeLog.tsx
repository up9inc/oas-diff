import './ChangeLog.sass';
import { CollapseButton } from '../../useCommonStyles';
import { useSetRecoilState } from 'recoil';
import { mainAccordionsList } from '../../recoil/collapse/';


export const ChangeLog = () => {

    const setAccordions = useSetRecoilState(mainAccordionsList)

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
