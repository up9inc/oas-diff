import './Header.sass';
import mizuIcon from "../../assets/mizu-icon.svg";
import logo from "../../assets/logo.svg";
import byUp9 from "../../assets/byUp9-img.svg";

export interface Props {
    dateGenerated: string;
}

export const Header: React.FC<Props> = ({
    dateGenerated
}) => {
    return (
        <header className='header'>
            <div className='leftSide'>
                <div className='logoWrapper'>
                    <img src={mizuIcon} className="mizuIcon" alt="logo icon"></img>
                    <div className='logoIcon'>
                        <img src={logo} alt="logo" className='mizuLogo'></img>
                        <img src={byUp9} alt="by Up9" className='byUp9Icon'></img>
                    </div>
                </div>
                <div className='title'>
                    <h1>OAS-diff
                        <span className='subTitle'>&nbsp;Report</span>
                    </h1>

                </div>
            </div>
            <div className='rightSide'>
                <span className='dateGenerated'>
                    {dateGenerated}{"18 Feb 2020"}
                </span>
            </div>
        </header>
    )
}