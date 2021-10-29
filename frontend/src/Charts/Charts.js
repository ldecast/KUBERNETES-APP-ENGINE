import React, { Component } from "react";
import "./Charts.css";
import Header from '../Header/Header';
import Loading from '../Loading/Loading';
import Bars from './Bars';
import Pie from './Pie';
import socket from "../socket";

class Charts extends Component {

  constructor() {
    super();
    this.state = {
      top3_games: [],
      inserts_workers: [],
      isLoaded: false,
    };
  }

  getData() {
    this.setState({
      isLoaded: false
    });
    fetch(process.env.REACT_APP_API_URL_ANALYTICS)
      .then(response => response.json())
      .then(data => {
        // console.log(data);
        this.setState(data)
      });
  };

  componentDidMount() {
    this.getData();

    socket.on('log-inserted', () => {
      console.log('charts socket ok');
      this.getData();
    });
  }

  render() {
    return (
      <div className="charts" >
        <Header title={"Analytics"} />
        {
          this.state.isLoaded
            ?
            (
              <div className="flex__container">

                <div className="vrs__container">
                  <h2>WORKERS COMPARISON</h2>
                  <div className="vrs__chart">
                    {console.log(this.state.inserts_workers)}
                    <Bars
                      workers={this.state.inserts_workers}
                    />
                  </div>
                </div>

                <div className="top__container">
                  <h2>TOP 3 GAMES</h2>
                  <div className="top__chart">
                    <Pie
                      top_3={this.state.top3_games}
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

}

export default Charts;