import React from 'react';
import { Link } from 'react-router-dom';
import './404.css'

class NotFoundPage extends React.Component {
    render() {
        return (<div className="notfound">
            <h1 id="title">404</h1>
            <h1 id="text">Not Found</h1>
            <p>
                <Link to="/home">Go to Home</Link>
            </p>
        </div>);
    }
} export default NotFoundPage;