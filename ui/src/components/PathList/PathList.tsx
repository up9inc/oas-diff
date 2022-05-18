import { Accordion, AccordionDetails, AccordionSummary, FormControl, InputLabel, MenuItem, Select, Typography, TextField, SelectChangeEvent } from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import './PathList.sass';
import { useCommonStyles } from '../../useCommonStyles';
import { getData } from '../../DataService';
import React, { ChangeEvent, useState } from 'react';

const data: any = getData()
export const PathListItem: React.FC<Props> = ({ }) => {
    return (
        <Accordion>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content"
                id="panel2a-header"
            >
                {/* <Typography>[PATHS] /customers/&#123;customerId&#125;addresses</Typography> */}
                <div className='accordionTitle'>
                    <div className='path'>
                        <span className='pathPrefix'>PATHS</span>&nbsp;
                        <span className='pathName'>/customers/&#123;customerId&#125;addresses</span>
                    </div>
                    <div className='changes'>
                        <span className='change'>Changes: 2</span>&nbsp;
                        <span className='change created'>Created: 2</span>
                    </div>
                </div>
            </AccordionSummary>
            <AccordionDetails>
                <Typography>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
                    malesuada lacus ex, sit amet blandit leo lobortis eget.
                </Typography>
            </AccordionDetails>
        </Accordion>
    )
}

export interface Props {

}

export const PathList: React.FC<Props> = (props: Props) => {
    const commonClasses = useCommonStyles()
    const services: any = data.changelog.paths.map((x: any) => x.path[0])
    const uniqueServices = Array.from(new Set(services))
    const types = data.changelog.paths.map((x: any) => x.type)
    const uniqueTypes = Array.from(new Set(types))

    const [service, setService] = useState('')
    const [type, setType] = useState('')
    const [path, setPath] = useState('')

    const onChange = (setFunc: Function) => (event: SelectChangeEvent<string> | React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => setFunc(event.target.value)

    return (
        <div className='pathListContainer'>
            <div className="innerContainer">
                <div className="title">
                    PATHS LIST
                </div>
                <div className="filters">
                    <FormControl size="small" sx={{ m: 1, minWidth: 120 }} className={`${commonClasses.select}`}>
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
                    </FormControl>
                    <FormControl>
                        <TextField id="outlined-basic" label="Path" variant="outlined" size="small" value={path} onChange={onChange(setPath)} />
                    </FormControl>
                    <FormControl size="small" sx={{ m: 1, minWidth: 150 }} className={`${commonClasses.select}`}>
                        <InputLabel id="demo-simple-select-label">Change Type</InputLabel>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            label="Change Type"
                            value={type}
                            onChange={onChange(setType)}
                        >
                            {uniqueTypes.map((type: any) => <MenuItem key={type} value={type}>{type}</MenuItem>)}
                        </Select>
                    </FormControl>

                </div>
                <div className='pathList'>
                    <PathListItem></PathListItem>
                </div>
            </div>
        </div>
    )
}