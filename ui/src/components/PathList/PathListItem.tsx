import { Accordion, AccordionSummary, AccordionDetails } from "@mui/material";
import { useContext, useMemo, useEffect } from "react";
import { CollapsedContext } from "../../CollapsedContext";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathListItem.sass';
import { DataItem, Path } from "../../interfaces";
import { infoClass, ChangeTypeEnum } from "../../consts";
import { SyntaxHighlighter } from "../SyntaxHighlighter";

export interface PathListItemProps {
    changeLogItem: DataItem
    showChangeType?: string
}

export const PathListItem: React.FC<PathListItemProps> = ({ changeLogItem, showChangeType = "" }) => {
    const changeVal = changeLogItem.value
    const { accordions, setAccordions } = useContext(CollapsedContext);
    const filteredChanges = useMemo(() => {
        return showChangeType ? changeVal?.path.filter((path: Path) => path.changelog.type === showChangeType) : changeVal?.path
    }, [changeVal?.path, showChangeType])

    useEffect(() => {
        const changes = showChangeType ? changeVal?.path.filter((path: Path) => path.changelog.type === showChangeType) : changeVal?.path
        setAccordions((prev) => {
            return [...prev, changes.map((path: Path) => { return { isCollpased: true, id: JSON.stringify(path) } })].flat()
        })
    }, [changeVal?.path, setAccordions, showChangeType])

    const onClick = (id: string) => {
        setAccordions((prev) => {
            const newArr = [...prev]
            const accordion = newArr.find(x => x.id === id)
            if (accordion)
                accordion.isCollpased = !accordion?.isCollpased
            return newArr
        })
    }

    useEffect(() => {
        setAccordions(prev => [...prev, { isCollpased: true, id: JSON.stringify(changeLogItem) }])
    }, [changeLogItem, setAccordions])

    const isExpand = (path: Path) => {
        const acc = accordions?.find(x => x.id === JSON.stringify(path))
        return acc ? !acc.isCollpased : false
    }

    const getToTypeColor = (type: string) => {
        switch (type) {
            case ChangeTypeEnum.Created:
                return ChangeTypeEnum.Created
            case ChangeTypeEnum.Updated:
                return ChangeTypeEnum.Created
            case ChangeTypeEnum.Deleted:
                return ChangeTypeEnum.Deleted
            default:
                return infoClass
        }
    }

    const getFromTypeColor = (type: string) => {
        switch (type) {
            case ChangeTypeEnum.Created:
                return ChangeTypeEnum.Created
            case ChangeTypeEnum.Updated:
                return ChangeTypeEnum.Deleted
            case ChangeTypeEnum.Deleted:
                return ChangeTypeEnum.Deleted
            default:
                return infoClass
        }
    }

    return (
        <Accordion expanded={!accordions.find(x => x.id === JSON.stringify(changeLogItem))?.isCollpased}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content" onClick={() => onClick(JSON.stringify(changeLogItem))}>
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
                {Object.keys(ChangeTypeEnum).map((changeType) => {
                    const changeOfType = filteredChanges.filter(x => x.changelog.type === ChangeTypeEnum[changeType])
                    return changeOfType.length > 0 && <div key={ChangeTypeEnum[changeType]}>
                        <div className={`${ChangeTypeEnum[changeType]} changeCategory`} >{Object.keys(ChangeTypeEnum).find(key => changeType === key)}</div>
                        {changeOfType?.map((path: Path) => {
                            return (<Accordion key={JSON.stringify(path)} expanded={(() => isExpand(path))()}>
                                <AccordionSummary
                                    onClick={() => onClick(JSON.stringify(path))}
                                    expandIcon={<ExpandMoreIcon />}
                                    aria-controls="panel2a-content">
                                    <div>
                                        <span className={`operation ${path.operation}`}>{path.operation}</span>
                                        <span className='pathName'>{path.changelog?.path?.join(" ")}</span>
                                    </div>
                                </AccordionSummary>
                                <AccordionDetails>
                                    <span>Path:</span>
                                    {path.changelog?.path?.slice(1).map((path: string, index: number) =>
                                        <div key={`${path + index}`} style={{ paddingLeft: `${(index + 1 * 0.4)}em` }}>{path}</div>)
                                    }
                                    <div style={{ marginTop: "10px" }} className="diffContainer">

                                        {path?.changelog?.from && <div style={{ flex: 1, width: "100%" }}>
                                            <div>From:</div>
                                            <SyntaxHighlighter
                                                code={JSON.stringify(path.changelog.from)}
                                                language="json"
                                                className={`${getFromTypeColor(path.changelog.type)}`}
                                            />
                                        </div>}
                                        {path?.changelog?.to && <div style={{ flex: 1, width: "100%" }}>
                                            <div>To:</div>
                                            <SyntaxHighlighter
                                                code={JSON.stringify(path.changelog.to)}
                                                language="json"
                                                className={`${getToTypeColor(path.changelog.type)}`}
                                            />
                                        </div>}
                                    </div>
                                </AccordionDetails>
                            </Accordion>)
                        })}
                    </div>
                }
                )}


            </AccordionDetails >
        </Accordion >
    )
}
