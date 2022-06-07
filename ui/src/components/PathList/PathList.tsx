import { FormControl, InputLabel, MenuItem, Select, TextField, SelectChangeEvent } from '@mui/material';
import './PathList.sass';
import React, { useMemo, useState } from 'react';
import PathListItem from './PathListItem';
import { DataItem, Path } from '../../interfaces';
import { ChangeTypeEnum } from '../../consts';
import useDebounce from '../../hooks/useDebounce';

export interface Props {
    changeList: DataItem[]
}

const PathList: React.FC<Props> = ({ changeList }) => {
    const [type, setType] = useState('')
    const [path, setPath] = useState('')
    const debouncedSearchTerm = useDebounce(path, 200);
    const onChangeTypeChange = (event) => { setType(event.target.value); }
    const onPathchange = (event) => { setPath(event.target.value); }

    const filteredChanges = useMemo(() => {
        let listAfterFilters = changeList?.filter((change: DataItem) => change?.key.toLowerCase().includes(debouncedSearchTerm.toLowerCase()))
        if (type)
            listAfterFilters = listAfterFilters?.filter((change: DataItem) => change?.value?.path.some((path: Path) => path.changelog.type === type))
        return listAfterFilters
    }, [changeList, debouncedSearchTerm, type])

    return (
        <div className='pathListContainer'>
            <div className="innerContainer">
                <div className="title">
                    PATHS LIST
                </div>
                <div className="filters">
                    <FormControl>
                        <TextField id="outlined-basic" label="Path" variant="outlined" size="small" value={path} onChange={onPathchange} />
                    </FormControl>
                    <div className='seperatorLine'></div>
                    <FormControl size='small' sx={{ minWidth: 150 }} >
                        <InputLabel>Change Type</InputLabel>
                        <Select
                            label="Change Type"
                            value={type}
                            onChange={onChangeTypeChange}
                            sx={{
                                margin: "0px !important",
                                width: "250px"
                            }}
                        >
                            <MenuItem key={"All"} value={""}>All</MenuItem>
                            <MenuItem key={ChangeTypeEnum.Created} value={ChangeTypeEnum.Created}>Create</MenuItem>
                            <MenuItem key={ChangeTypeEnum.Updated} value={ChangeTypeEnum.Updated}>Update</MenuItem>
                            <MenuItem key={ChangeTypeEnum.Deleted} value={ChangeTypeEnum.Deleted}>Delete</MenuItem>
                        </Select>
                    </FormControl>
                </div>
                <div className='changeLogList'>
                    {filteredChanges?.map((change: DataItem, index: number) =>
                        <PathListItem key={index} changeLogItem={change} showChangeType={type} />
                    )}
                </div>
            </div>
        </div>
    )
}

export default React.memo(PathList)
