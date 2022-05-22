import './Header.sass';
import { ReactComponent as LogoSvg } from "../../assets/Mizu-logo.svg";

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
                    <LogoSvg />
                </div>
                <div className='title'>
                    <h1>OAS-diff
                        <span className='subTitle'>Report</span>
                    </h1>
                </div>
            </div>
            <div className='rightSide'>
                <span className='dateGenerated'>
                    {new Date(dateGenerated).toLocaleTimeString('en-us', {
                        month: "long", year: "numeric", day: "numeric", hour: '2-digit', minute: '2-digit', hour12: false
                    })}
                </span>
            </div>
        </header>
    )
}
