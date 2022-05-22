import { FormControl, InputLabel, MenuItem, Select, TextField, SelectChangeEvent } from '@mui/material';
import './PathList.sass';
import React, { useMemo, useState } from 'react';
import { PathListItem } from './PathListItem';
import { DataItem, Path } from '../../interfaces';
import { ChangeTypeEnum } from '../../consts';

export interface Props {
    changeList: DataItem[]
}

export const PathList: React.FC<Props> = ({ changeList }) => {
    const [type, setType] = useState('')
    const [path, setPath] = useState('')

    const onPathFilterChange = (setFunc: Function) => (event: SelectChangeEvent<string> | React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => setFunc(event.target.value)

    const filteredChanges = useMemo(() => {
        let relevantList = changeList
        if (type)
            relevantList = changeList?.filter((change: DataItem) => change?.value?.path.some((path: Path) => path.changelog.type === type))

        return relevantList?.filter((change: DataItem) => change?.key.toLowerCase().includes(path.toLowerCase()))
    }, [changeList, path, type])

    return (
        <div className='pathListContainer'>
            <div className="innerContainer">
                <div className="title">
                    PATHS LIST
                </div>
                <div className="filters">
                    <FormControl>
                        <TextField id="outlined-basic" label="Path" variant="outlined" size="small" value={path} onChange={onPathFilterChange(setPath)} />
                    </FormControl>
                    <div className='seperatorLine'></div>
                    <FormControl size='small' sx={{ minWidth: 150 }} >
                        <InputLabel>Change Type</InputLabel>
                        <Select
                            label="Change Type"
                            value={type}
                            onChange={onPathFilterChange(setType)}
                            sx={{
                                margin: "0px !important",
                                width: "250px"
                            }}
                        >
                            <MenuItem key={"None"} value={""}>None</MenuItem>
                            <MenuItem key={"created"} value={ChangeTypeEnum.Created}>Create</MenuItem>
                            <MenuItem key={"updated"} value={ChangeTypeEnum.Updated}>Update</MenuItem>
                            <MenuItem key={"deleted"} value={ChangeTypeEnum.Deleted}>Delete</MenuItem>
                        </Select>
                    </FormControl>
                </div>
                <div className='changeLogList'>
                    {filteredChanges?.map((change: DataItem, index: number) => <div key={"changeLogItem" + index} className='changeLogItem'>
                        <PathListItem change={change} showChangeType={type} /></div>)
                    }
                </div>
            </div>
        </div>
    )
}
