import { Accordion, AccordionSummary, AccordionDetails, Grid } from "@mui/material";
import { useContext, useMemo, useEffect } from "react";
import { CollapsedContext } from "../../CollapsedContext";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathListItem.sass';

export interface PathListItemProps {
    change: any
    showChangeType?: string
}

export const PathListItem: React.FC<PathListItemProps> = ({ change, showChangeType = "" }) => {
    const changeVal = change.Value
    const { accordions, setAccordions } = useContext(CollapsedContext);
    const filteredChanges = useMemo(() => {
        const changes = showChangeType ? changeVal?.Paths.filter((path: any) => path.Changelog.type === showChangeType) : changeVal?.Paths
        setAccordions((prev) => {
            let newArr = [...prev, changes.map((path: any) => { return { isCollpased: true, id: JSON.stringify(path) } })].flat()
            return newArr
        })
        return changes
    }, [changeVal?.Paths, setAccordions, showChangeType])

    const onClick = (id: string) => {
        setAccordions((prev) => {
            let newArr = [...prev]
            let accordion = newArr.find(x => x.id === id)
            if (accordion)
                accordion.isCollpased = !accordion?.isCollpased
            return newArr
        })
    }

    useEffect(() => {
        setAccordions(prev => [...prev, { isCollpased: true, id: JSON.stringify(change) }])
    }, [change, setAccordions])

    const isExpand = (path: any) => {
        const acc = accordions?.find(x => x.id === JSON.stringify(path))
        if (acc)
            return !acc.isCollpased
        return false
    }

    const getToTypeColor = (type: string) => {
        switch (type) {
            case "create":
                return "create"
            case "update":
                return "create"
            case "delete":
                return "delete"
            default:
                return "info"
        }
    }

    const getFromTypeColor = (type: string) => {
        switch (type) {
            case "create":
                return "create"
            case "update":
                return "delete"
            case "delete":
                return "delete"
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
                        <span className='pathPrefix'>{change.Value.Key}</span>&nbsp;
                        <span className='pathName'>{change.Key}</span>
                    </div>
                    <div>
                        <span className='change total'>Changes: {changeVal.TotalChanges}</span>&nbsp;
                        {changeVal.CreatedChanges > 0 && <span className='change create'>Created: {changeVal.CreatedChanges}</span>}
                        {changeVal.UpdatedChanges > 0 && <span className='change update'>Updated: {changeVal.UpdatedChanges}</span>}
                        {changeVal.DeletedChanges > 0 && <span className='change delete'>Deleted: {changeVal.DeletedChanges}</span>}
                    </div>
                </div>
            </AccordionSummary>
            <AccordionDetails>
                {filteredChanges?.map((path: any) => {
                        <Accordion key={JSON.stringify(path)} expanded={(() => isExpand(path))()}>
                            <AccordionSummary
                                onClick={() => onClick(JSON.stringify(path))}
                                expandIcon={<ExpandMoreIcon />}
                                aria-controls="panel2a-content">
                                <div>
                                    <span className={`operation ${path.Operation}`}>{path.Operation}</span>&nbsp;
                                    <span className='pathName'>{path.Changelog.path.join(" ")}</span>
                                </div>
                            </AccordionSummary>
                            <AccordionDetails>
                                <span>Path:</span>
                                {path.Changelog.path.slice(1).map((path: any, index: number) =>
                                    <div key={`${path + index}`} style={{ paddingLeft: `${(index + 1 * 0.4)}em` }}>{path}</div>)
                                }
                                <div style={{ marginTop: "10px" }}>
                                    <Grid container spacing={2}>
                                        {path?.Changelog?.from && <Grid item md>
                                            <div>From:</div>
                                            <pre className={`${getFromTypeColor(path.Changelog.type)}`} style={{ whiteSpace: "pre-wrap" }}>
                                                {JSON.stringify(path.Changelog.from)}
                                            </pre>
                                        </Grid>}
                                        {path?.Changelog?.to && <Grid item md>
                                            <div>To:</div>
                                            <pre className={`${getToTypeColor(path.Changelog.type)}`} style={{ whiteSpace: "pre-wrap" }}>
                                                {JSON.stringify(path.Changelog.to)}
                                            </pre>
                                        </Grid>}
                                    </Grid>
                                </div>
                            </AccordionDetails>
                        </Accordion>
                })}
            </AccordionDetails >
        </Accordion >
    )
}