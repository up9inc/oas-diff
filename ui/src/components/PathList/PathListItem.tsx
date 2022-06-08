import { Accordion, AccordionSummary, AccordionDetails } from "@mui/material";
import { useMemo, useState } from "react";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathListItem.sass';
import { DataItem, Path } from "../../interfaces";
import { PathDisplay } from "./PathDisplay";
import React from "react";
import { ChangeTypeEnum } from "../../consts";

export interface PathListItemProps {
    changeLogItem: DataItem
    showChangeType?: string
}

const PathListItem: React.FC<PathListItemProps> = ({ changeLogItem, showChangeType = "" }) => {
    const changeVal = changeLogItem.value
    const [isExpanded, setIsExpanded] = useState(false)
    const changes = useMemo(() => {
        return changeVal?.path
    }, [changeVal?.path])

    const filteredChanges = useMemo(() => {
        return changes.filter((path) => path.changelog.type.indexOf(showChangeType) >= 0)
    }, [changes, showChangeType])

    const onAccordionClick = () => {
        setIsExpanded(!isExpanded)
    }

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
