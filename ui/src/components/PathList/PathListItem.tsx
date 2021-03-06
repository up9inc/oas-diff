import { Accordion, AccordionSummary, AccordionDetails } from "@mui/material";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathListItem.sass';
import { DataItem, Path } from "../../interfaces";
import { ChangeTypeEnum } from "../../consts";
import { PathDisplay } from "./PathDisplay";
import React, { useState, useEffect, useMemo, useCallback } from "react";
import { useRecoilValue, useSetRecoilState } from "recoil";
import { mainAccordionsList, subAccordionsList } from "../../recoil/collapse";

export interface PathListItemProps {
    changeLogItem: DataItem
    showChangeType?: string
}

const PathListItem: React.FC<PathListItemProps> = ({ changeLogItem, showChangeType = "" }) => {
    const changeVal = changeLogItem.value
    const accordions = useRecoilValue(mainAccordionsList);
    const setSubAccordions = useSetRecoilState(subAccordionsList);
    const [isExpanded, setIsExpanded] = useState(false)

    useEffect(() => {
        const isGloballyExpanded = !accordions.find(x => x.id === JSON.stringify(changeLogItem))?.isCollapsed
        setIsExpanded(isGloballyExpanded)
    }, [accordions, changeLogItem])

    const changes = useMemo(() => {
        return changeVal?.path
    }, [changeVal?.path])

    useEffect(() => {
        const subAccordions = changeVal?.path.map((path: Path) => {
            return { isCollapsed: true, id: JSON.stringify(path) }
        })
        setSubAccordions(subAccordions)
    }, [changeVal?.path, setSubAccordions])

    const filteredChanges = useMemo(() => {
        return changes.filter((path) => path.changelog.type.indexOf(showChangeType) >= 0)
    }, [changes, showChangeType])

    const onAccordionClick = useCallback(() => {
        setIsExpanded(!isExpanded)
    }, [isExpanded])

    return (
        <Accordion expanded={isExpanded}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content" onClick={onAccordionClick}>
                <div className='accordionTitle'>
                    <div className='path'>
                        <span className='pathPrefix'>{changeLogItem.value.key}</span>
                        <span className='pathName'>{changeLogItem.key}</span>
                    </div>
                    <div>
                        <span className='change total'>Changes: {changeVal.totalChanges}</span>
                        {changeVal.createdChanges > 0 && <span className='change create'>Created: {changeVal.createdChanges}</span>}
                        {changeVal.updatedChanges > 0 && <span className='change update'>Updated: {changeVal.updatedChanges}</span>}
                        {changeVal.deletedChanges > 0 && <span className='change delete'>Deleted: {changeVal.deletedChanges}</span>}
                    </div>
                </div>
            </AccordionSummary>
            <AccordionDetails>
                {isExpanded && <>
                    {Object.keys(ChangeTypeEnum).map((changeType) => {
                        const changeOfType = filteredChanges.filter(x => x.changelog.type === ChangeTypeEnum[changeType])
                        return changeOfType.length > 0 && <div key={ChangeTypeEnum[changeType]}>
                            <div className={`${ChangeTypeEnum[changeType]} changeCategory`} >{
                                Object.keys(ChangeTypeEnum).find(key => changeType === key)}
                            </div>
                            {changeOfType?.map((path: Path, index) => {
                                return <PathDisplay path={path} key={index} />
                            })}
                        </div>
                    })}
                </>
                }
            </AccordionDetails>
        </Accordion >
    )
}

export default React.memo(PathListItem)
