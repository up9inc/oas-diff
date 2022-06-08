import { Accordion, AccordionSummary, AccordionDetails } from "@mui/material"
import { Path } from "../../interfaces"
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import SyntaxHighlighter from "../SyntaxHighlighter"
import { ChangeTypeEnum, infoClass } from "../../consts";
import { useState } from "react";


export interface PathDisplayProps {
    path: Path
}

export const PathDisplay: React.FC<PathDisplayProps> = ({ path }) => {
    const [isExpanded, setIsExpanded] = useState(false)
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
    return (<>
        <Accordion expanded={isExpanded}>
            <AccordionSummary
                onClick={() => { setIsExpanded(!isExpanded) }}
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel2a-content">
                <div className="singleLine">
                    <span className={`operation ${path.operation}`}>{path.operation}</span>
                    <span className='pathName'>{path.changelog?.path?.join(" ")}</span>
                </div>
            </AccordionSummary>
            <AccordionDetails>
                {isExpanded && <><span>Path:</span>
                    {path.changelog?.path?.slice(1).map((path: string, index: number) =>
                        <div key={index} style={{ paddingLeft: `${(index + 1 * 0.4)}em` }}>{path}</div>)
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
                    </div></>}

            </AccordionDetails>
        </Accordion>
    </>)
}
