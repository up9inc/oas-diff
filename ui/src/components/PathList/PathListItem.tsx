import { Accordion, AccordionSummary, AccordionDetails, Grid } from "@mui/material";
import { useContext, useMemo, useEffect } from "react";
import { CollapsedContext } from "../../CollapsedContext";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathListItem.sass';
import { DataItem, Path } from "../../interfaces";
import { createClass, deleteClass } from "../../consts";

export interface PathListItemProps {
    change: DataItem
    showChangeType?: string
}

export const PathListItem: React.FC<PathListItemProps> = ({ change, showChangeType = "" }) => {
    const changeVal = change.value
    const { accordions, setAccordions } = useContext(CollapsedContext);
    const filteredChanges = useMemo(() => {
        const changes = showChangeType ? changeVal?.path.filter((path: Path) => path.changelog.type === showChangeType) : changeVal?.path
        setAccordions((prev) => {
            return [...prev, changes.map((path: Path) => { return { isCollpased: true, id: JSON.stringify(path) } })].flat()
        })
        return changes
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
        setAccordions(prev => [...prev, { isCollpased: true, id: JSON.stringify(change) }])
    }, [change, setAccordions])

    const isExpand = (path: Path) => {
        const acc = accordions?.find(x => x.id === JSON.stringify(path))
        return acc ? !acc.isCollpased : false
    }

    const getToTypeColor = (type: string) => {
        switch (type) {
            case "create":
                return createClass
            case "update":
                return createClass
            case "delete":
                return deleteClass
            default:
                return "info"
        }
    }

    const getFromTypeColor = (type: string) => {
        switch (type) {
            case "create":
                return createClass
            case "update":
                return deleteClass
            case "delete":
                return deleteClass
            default:
                return "info"
        }
    }

    return (
        <Accordion expanded={!accordions.find(x => x.id === JSON.stringify(change))?.isCollpased}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content" onClick={() => onClick(JSON.stringify(change))}>
                <div className='accordionTitle'>
                    <div className='path'>
                        <span className='pathPrefix'>{change.value.key}</span>
                        <span className='pathName'>{change.key}</span>
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
                <div className={`${filteredChanges && filteredChanges[0]?.changelog?.type} changeCategory`}>{filteredChanges[0]?.changelog?.type + "d"}</div>
                {filteredChanges?.map((path: Path) => {
                    return (<Accordion key={JSON.stringify(path)} expanded={(() => isExpand(path))()}>
                        <AccordionSummary
                            onClick={() => onClick(JSON.stringify(path))}
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panel2a-content">
                            <div>
                                <span className={`operation ${path.operation}`}>{path.operation}</span>
                                <span className='pathName'>{path.changelog?.paths?.join(" ")}</span>
                            </div>
                        </AccordionSummary>
                        <AccordionDetails>
                            <span>Path:</span>
                            {path.changelog?.paths?.slice(1).map((path: string, index: number) =>
                                <div key={`${path + index}`} style={{ paddingLeft: `${(index + 1 * 0.4)}em` }}>{path}</div>)
                            }
                            <div style={{ marginTop: "10px" }}>
                                <Grid container spacing={2}>
                                    {path?.changelog?.from && <Grid item md>
                                        <div>From:</div>
                                        <pre className={`${getFromTypeColor(path.changelog.type)}`} style={{ whiteSpace: "pre-wrap" }}>
                                            {JSON.stringify(path.changelog.from)}
                                        </pre>
                                    </Grid>}
                                    {path?.changelog?.to && <Grid item md>
                                        <div>To:</div>
                                        <pre className={`${getToTypeColor(path.changelog.type)}`} style={{ whiteSpace: "pre-wrap" }}>
                                            {JSON.stringify(path.changelog.to)}
                                        </pre>
                                    </Grid>}
                                </Grid>
                            </div>
                        </AccordionDetails>
                    </Accordion>)
                })}
            </AccordionDetails >
        </Accordion >
    )
}
