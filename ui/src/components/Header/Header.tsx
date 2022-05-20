import './Header.sass';
import logo from "../../assets/Mizu-logo.svg";

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
                    <img src={logo} className="" alt="logo icon"></img>
                </div>
                <div className='title'>
                    <h1>OAS-diff
                        <span className='subTitle'>&nbsp;Report</span>
                    </h1>
                </div>
            </div>
            <div className='rightSide'>
                <span className='dateGenerated'>
                    {dateGenerated}
                </span>
            </div>
        </header>
    )
}