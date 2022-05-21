import { FormControl, InputLabel, MenuItem, Select, TextField, SelectChangeEvent } from '@mui/material';
import './PathList.sass';
import React, { useMemo, useState } from 'react';
import { PathListItem } from './PathListItem';

export interface Props {
    changeList: any
}

export const PathList: React.FC<Props> = ({ changeList }) => {
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
                    <FormControl>
                        <TextField id="outlined-basic" label="Path" variant="outlined" size="small" value={path} onChange={onChange(setPath)} />
                    </FormControl>
                    <div className='seperatorLine'></div>
                    <FormControl size='small' sx={{ minWidth: 150 }} >
                        <InputLabel id="demo-simple-select-label">Change Type</InputLabel>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            label="Change Type"
                            value={type}
                            onChange={onChange(setType)}
                            sx={{
                                margin: "0px !important",
                                width: "250px"
                            }}
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
                    {filteredChanges?.map((change: any, index: string) => <div key={"changeLogItem" + index} className='changeLogItem'>
                        <PathListItem change={change} showChangeType={type}></PathListItem></div>)
                    }
                </div>
            </div>
        </div>
    )
}