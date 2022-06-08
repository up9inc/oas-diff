import { FormControl, InputLabel, MenuItem, Select, TextField } from '@mui/material';
import './PathList.sass';
import React, { useMemo, useState, useEffect } from 'react';
import PathListItem from './PathListItem';
import { DataItem, Path } from '../../interfaces';
import { ChangeTypeEnum } from '../../consts';
import useDebounce from '../../hooks/useDebounce';
import { useSetRecoilState } from 'recoil';
import { collapseItemsList } from '../../recoil/collapse';

export interface Props {
    changeList: DataItem[]
}

const PathList: React.FC<Props> = ({ changeList }) => {
    const [typeFilter, setTypeFilter] = useState('')
    const [pathFilter, setPathFilter] = useState('')
    const debouncedSearchTerm = useDebounce(pathFilter, 200);
    const setAccordions = useSetRecoilState(collapseItemsList);
    const onTypeChange = (event) => { setTypeFilter(event.target.value); }
    const onPathChange = (event) => { setPathFilter(event.target.value); }

    const filteredListItems = useMemo(() => {
        let listAfterFilters = changeList?.filter((change: DataItem) => change?.key.toLowerCase().includes(debouncedSearchTerm.toLowerCase()))
        if (typeFilter)
            listAfterFilters = listAfterFilters?.filter((change: DataItem) => change?.value?.path.some((path: Path) => path.changelog.type === typeFilter))
        return listAfterFilters
    }, [changeList, debouncedSearchTerm, typeFilter])

    useEffect(() => {
        const accordions = changeList.map(change => {
            return { isCollapsed: true, id: JSON.stringify(change) }
        })
        setAccordions(accordions)
    }, [changeList, setAccordions])

    return (
        <div className='pathListContainer'>
            <div className="innerContainer">
                <div className="title">
                    PATHS LIST
                </div>
                <div className="filters">
                    <FormControl>
                        <TextField id="outlined-basic" label="Path" variant="outlined" size="small" value={pathFilter} onChange={onPathChange} />
                    </FormControl>
                    <div className='seperatorLine'></div>
                    <FormControl size='small' sx={{ minWidth: 150 }} >
                        <InputLabel>Change Type</InputLabel>
                        <Select
                            label="Change Type"
                            value={typeFilter}
                            onChange={onTypeChange}
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
                    {filteredListItems?.map((change: DataItem, index: number) =>
                        <PathListItem key={index} changeLogItem={change} showChangeType={typeFilter} />
                    )}
                </div>
            </div>
        </div>
    )
}

export default React.memo(PathList)
