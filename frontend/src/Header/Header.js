import { useHistory } from "react-router-dom";
import './Header.css';

function Header(props) {
    let history = useHistory();
    return (
        <div className="__header">
            <h2>{props.title}</h2>
            <select id="db_selector" defaultValue={props.redis ? "2" : "1"} onChange={(e) => {
                if (e.target.value === "2") {
                    history.push('/redisReports')
                }
            }}>
                <option disabled={props.redis ? true : false} value="1">Mongo DB</option>
                <option value="2">Redis</option>
            </select>
        </div>
    );
}

export default Header;