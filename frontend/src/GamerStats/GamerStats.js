import { useState, useEffect } from 'react';
import "./GamerStats.css";
import Header from '../Header/Header';
import Loading from '../Loading/Loading';
import socket from "../socket";

function GamerStats() {

  const [player, setPlayer] = useState("-");
  const [stats, setStats] = useState([]);
  const [players, setPlayers] = useState([]);
  const [isLoaded, setLoaded] = useState(false);

  useEffect(() => {
    const getPlayers = () => {
      fetch(process.env.REACT_APP_API_URL_STATS)
        .then(response => response.json())
        .then(data => {
          // console.log(data);
          setPlayers(data);
          setLoaded(true);
        });
    }
    getPlayers();
    socket.on('log-inserted', () => {
      console.log('gamerStats socket ok');
      getPlayers();
    });
    return () => {
      socket.off('log-inserted');
    }
  }, [])

  useEffect(() => {
    const getStats = () => {
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          player: player
        })
      };
      fetch(process.env.REACT_APP_API_URL_STATS, requestOptions)
        .then(response => response.json())
        .then(data => {
          // console.log(data);
          setStats(data);
          setLoaded(true);
        });
    }
    getStats();
  }, [player])


  function changePlayer(val) {
    setLoaded(false);
    setPlayer(val);
  }

  return (
    <div className="stats" >
      <Header title={"Realtime Gamer Stats"} />
      {
        isLoaded
          ?
          (
            <div className="gamer__content">
              <select id="player__select" defaultValue={player} onChange={(e) => changePlayer(e.target.value)}>
                <option value="-" disabled>Choose a player</option>
                {
                  players.map((player) => {
                    return (
                      <option key={player} value={player}>{player}</option>
                    );
                  })
                }
              </select>
              <div className="gamer__table">
                <table>
                  <tbody>
                    <tr>
                      <th>&nbsp;&nbsp;Game#</th>
                      <th>Game Name</th>
                      <th>State</th>
                    </tr>
                    {
                      stats.map((stat) => {
                        return (
                          <tr key={stat._id} >
                            <td>&nbsp;&nbsp; {stat.gameid}</td>
                            <td>&nbsp;{stat.gamename}</td>
                            <td>{"Win"}</td>
                          </tr>
                        );
                      })
                    }
                  </tbody>
                </table>
              </div>

            </div>
          )
          :
          <Loading />
      }
    </div >
  );
}

export default GamerStats;