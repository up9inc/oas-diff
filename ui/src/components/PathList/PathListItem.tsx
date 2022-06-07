import { Accordion, AccordionSummary, AccordionDetails } from "@mui/material";
import { useContext, useMemo, useState, useCallback, useRef } from "react";
import { CollapsedContext } from "../../CollapsedContext";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathListItem.sass';
import { DataItem, Path } from "../../interfaces";
import { ChangeTypeEnum } from "../../consts";
import { PathDisplay } from "./PathDisplay";
import React from "react";

const getKeyByValue = (value: string) => {
    const indexOfS = Object.values(ChangeTypeEnum).indexOf(value as unknown as ChangeTypeEnum);
    return Object.keys(ChangeTypeEnum)[indexOfS];
}

export interface PathListItemProps {
    changeLogItem: DataItem
    showChangeType?: string
    style?: any
}

const PathListItem: React.FC<PathListItemProps> = ({ changeLogItem, showChangeType = "" }) => {
    const changeVal = changeLogItem.value
    changeLogItem.isExpanded = false
    const { accordions, setAccordions } = useContext(CollapsedContext);
    const [isExpanded, setIsExpanded] = useState(false)
    const changes = useMemo(() => {
        return changeVal?.path
    }, [changeVal?.path])

    const filteredChanges = useMemo(() => {
        return changes.filter((path) => path.changelog.type.indexOf(showChangeType) >= 0)
    }, [changes, showChangeType])

    // useEffect(() => {
    //     // const changes = showChangeType ? changeVal?.path.filter((path: Path) => path.changelog.type === showChangeType) : changeVal?.path
    //     // setAccordions((prev) => {
    //     //     return [...prev, changes.map((path: Path) => { return { isCollpased: true, id: JSON.stringify(path) } })].flat()
    //     // })
    // }, [changeVal?.path, setAccordions, showChangeType])

    const onClick = () => {
        setIsExpanded(!isExpanded)
    }

    // useEffect(() => {
    //     //setAccordions(prev => [...prev, { isCollpased: true, id: JSON.stringify(changeLogItem) }])
    // }, [changeLogItem, setAccordions])


    return (
        <Accordion expanded={isExpanded}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content" onClick={() => onClick()}>
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
                    {filteredChanges?.map((path: Path, index) => {
                        return <PathDisplay path={path} key={index} />
                    })}
                </>}
            </AccordionDetails>
        </Accordion >
    )
}

export default React.memo(PathListItem)
