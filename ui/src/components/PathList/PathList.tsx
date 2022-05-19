import { Accordion, AccordionDetails, AccordionSummary, FormControl, InputLabel, MenuItem, Select, Typography, TextField, SelectChangeEvent, Grid } from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathList.sass';
import { useCommonStyles } from '../../useCommonStyles';
import React, { useContext, useMemo, useState } from 'react';
import { CollapsedContext } from '../../CollapsedContext';

export interface PathListItemProps {
    change: any
    showChangeType?: string
}


export const PathListItem: React.FC<PathListItemProps> = ({ change, showChangeType = "" }) => {

    const changeVal = change.Value
    const filteredChanges = useMemo(() => {
        return showChangeType ? changeVal?.Paths.filter((path: any) => path.Changelog.type === showChangeType) : changeVal?.Paths
    }, [changeVal?.Paths, showChangeType])

    const { collapsed } = useContext(CollapsedContext);

    return (
        <Accordion expanded={!collapsed}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content"
                id="panel2a-header"
            >
                <div className='accordionTitle'>
                    <div className='path'>
                        <span className='pathPrefix'>{change.Value.Key}</span>&nbsp;
                        <span className='pathName'>{change.Key}</span>
                    </div>
                    <div className='changes'>
                        <span className='change total'>Changes: {changeVal.TotalChanges}</span>&nbsp;
                        {changeVal.CreatedChanges > 0 && <span className='change create'>Created: {changeVal.CreatedChanges}</span>}
                        {changeVal.UpdatedChanges > 0 && <span className='change update'>Updated: {changeVal.UpdatedChanges}</span>}
                        {changeVal.DeletedChanges > 0 && <span className='change delete'>Deleted: {changeVal.DeletedChanges}</span>}
                    </div>
                </div>
            </AccordionSummary>
            <AccordionDetails>
                {filteredChanges?.map((path: any) => {
                    return <Accordion key={JSON.stringify(path)} expanded={!collapsed}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panel2a-content"
                            id="panel2a-header">
                            <div className='path'>
                                <span className={`operation ${path.Operation}`}>{path.Operation}</span>&nbsp;
                                <span className='pathName'>{path.Changelog.path.join(" ")}</span>
                            </div>
                        </AccordionSummary>
                        <AccordionDetails>
                            <span>Path:</span>
                            {path.Changelog.path.slice(1).map((path: any, index: number) => <div key={`${path + index}`} style={{ paddingLeft: `${(index + 1 * 0.4)}em` }}>{path}</div>)}
                            <div style={{ marginTop: "10px" }}>
                                <Grid container spacing={2}>
                                    {path?.Changelog?.from && <Grid item md>
                                        <div>From:</div>
                                        <pre className={`${path.Changelog.type}`} style={{ whiteSpace: "pre-wrap" }}>
                                            {JSON.stringify(path.Changelog.from)}
                                        </pre>
                                    </Grid>}
                                    {path?.Changelog?.to && <Grid item md>
                                        <div>To:</div>
                                        <pre className={`${path.Changelog.type}`} style={{ whiteSpace: "pre-wrap" }}>
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

export interface Props {
    changeList: any
}

export const PathList: React.FC<Props> = ({ changeList }) => {
    const commonClasses = useCommonStyles()
    const services: any = changeList.map((x: any) => x.Key)
    const uniqueServices = Array.from(new Set(services))
    // const types = data.changelog.paths.map((x: any) => x.type)
    // const uniqueTypes = Array.from(new Set(types))

    const [service, setService] = useState('')
    const [type, setType] = useState('')
    const [path, setPath] = useState('')

    const onChange = (setFunc: Function) => (event: SelectChangeEvent<string> | React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => setFunc(event.target.value)


    const filteredChanges = useMemo(() => {
        let relevantList = changeList
        if (type !== "")
            relevantList = changeList?.filter((change: any) => change?.Value?.Paths.some((path: any) => path.Changelog.type === type))

        return relevantList?.filter((change: any) => change?.Key.toLowerCase().includes(path.toLowerCase()))
    }, [changeList, path, type])

    return (
        <div className='pathListContainer'>
            <div className="innerContainer">
                <div className="title">
                    PATHS LIST
                </div>
                <div className="filters">
                    {/* <FormControl size="small" sx={{ m: 1, minWidth: 120 }} className={`${commonClasses.select}`}>
                        <InputLabel id="demo-simple-select-label">Services</InputLabel>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            label="Services"
                            value={service}
                            onChange={onChange(setService)}
                        >
                            {uniqueServices.map((service: any) => <MenuItem key={service} value={service}>{service}</MenuItem>)}
                        </Select>
                    </FormControl> */}
                    <FormControl>
                        <TextField id="outlined-basic" label="Path" variant="outlined" size="small" value={path} onChange={onChange(setPath)} />
                    </FormControl>
                    <div className='seperatorLine'></div>

                    <FormControl size='small' sx={{ m: 1, minWidth: 150 }} className={`${commonClasses.select}`}>
                        <InputLabel id="demo-simple-select-label">Change Type</InputLabel>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            label="Change Type"
                            value={type}
                            onChange={onChange(setType)}

                        >
                            <MenuItem key={"None"} value={""}>None</MenuItem>
                            <MenuItem key={"Created"} value={"create"}>Create</MenuItem>
                            <MenuItem key={"updated"} value={"update"}>Update</MenuItem>
                            <MenuItem key={"delete"} value={"delete"}>Delete</MenuItem>
                            {/* {uniqueTypes.map((type: any) => <MenuItem key={type} value={type}>{type}</MenuItem>)} */}
                        </Select>
                    </FormControl>

                </div>
                <div className='changeLogList'>
                    {filteredChanges.map((change: any, index: string) => <div key={"changeLogItem" + index} className='changeLogItem'>
                        <PathListItem change={change} showChangeType={type}></PathListItem></div>)
                    }
                </div>
            </div>
        </div>
    )
}