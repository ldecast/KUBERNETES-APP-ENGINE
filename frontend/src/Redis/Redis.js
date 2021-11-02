import { useState, useEffect } from 'react';
import "./Redis.css";
import Header from '../Header/Header';
import Loading from '../Loading/Loading';
import socket from "../socket";

function Redis() {

    const [state, setState] = useState({
        last_request: 0,
        last_gameid: "",
        last_gamename: "",
        last_winner: "",
        last_players: 0,
        last_worker: "",
        isLoaded: false
    });

    const getReports = async () => {
        fetch(process.env.REACT_APP_API_URL_REDIS)
            .then(response => response.json())
            .then(data => {
                // console.log(data);
                setState(data);
            });
    };

    useEffect(() => {
        getReports();
        return () => {
            socket.off('redis-report');
        }
    }, [])

    useEffect(() => {
        socket.on('redis-report', () => {
            console.log('redis reports socket ok');
            getReports();
        });
    }, [])

    return (
        <div className="redis" >
            <Header title={"Redis reports"} redis={true} />
            {
                state.isLoaded
                    ?
                    (
                        <div className="redis__content">

                            <div className="cards__container">
                                <div className="report__card" id="card-1">
                                    <h2>REQUESTS NUMBER</h2>
                                    <p>{state.last_request}</p>
                                </div>
                                <div className="report__card" id="card-2">
                                    <h2>LAST GAME ID</h2>
                                    <p>{state.last_gameid}</p>
                                </div>
                                <div className="report__card" id="card-3">
                                    <h2>LAST GAME NAME</h2>
                                    <p>{state.last_gamename}</p>
                                </div>
                                <div className="report__card" id="card-4">
                                    <h2>LAST WINNER</h2>
                                    <p>{state.last_winner}</p>
                                </div>
                                <div className="report__card" id="card-5">
                                    <h2>LAST PLAYERS</h2>
                                    <p>{state.last_players}</p>
                                </div>
                                <div className="report__card" id="card-6">
                                    <h2>LAST WORKER INSERTION</h2>
                                    <p>{state.last_worker}</p>
                                </div>
                            </div>

                        </div>
                    )
                    :
                    <Loading />
            }
        </div >
    );
}

export default Redis;