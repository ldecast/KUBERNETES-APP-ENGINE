import { useState, useEffect } from 'react';
import "./Charts.css";
import Header from '../Header/Header';
import Loading from '../Loading/Loading';
import Bars from './Bars';
import Pie from './Pie';
import socket from "../socket";

function Charts() {

  const [state, setState] = useState({
    top3_games: [],
    inserts_workers: [],
    isLoaded: false
  });

  const getData = () => {
    fetch(process.env.REACT_APP_API_URL + "/analytics")
      .then(response => response.json())
      .then(data => {
        // console.log(data);
        setState(data)
      });
  };

  useEffect(() => {
    getData();
    return () => {
      socket.off('log-inserted');
    }
  }, [])

  useEffect(() => {
    socket.on('log-inserted', () => {
      console.log('charts socket ok');
      getData();
    });
  }, [])


  return (
    <div className="charts" >
      <Header title={"Analytics"} />
      {
        state.isLoaded
          ?
          (
            <div className="flex__container">

              <div className="vrs__container">
                <h2>WORKERS COMPARISON</h2>
                <div className="vrs__chart">
                  <Bars
                    workers={state.inserts_workers}
                  />
                </div>
              </div>

              <div className="top__container">
                <h2>TOP 3 GAMES</h2>
                <div className="top__chart">
                  <Pie
                    top_3={state.top3_games}
                  />
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

export default Charts;