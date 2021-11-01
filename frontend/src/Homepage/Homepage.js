import { useEffect, useState } from 'react';
import "./Homepage.css";
import Header from '../Header/Header';
import Loading from '../Loading/Loading';
import socket from "../socket";

function Homepage() {

  const [state, setState] = useState({
    last_games: [],
    top10_players: []
  });
  const [isLoaded, setLoaded] = useState(false);

  const getData = () => {
    fetch(process.env.REACT_APP_API_URL_HOME)
      .then(response => response.json())
      .then(data => {
        // console.log(data);
        setState(data);
        setLoaded(true);
      });
  };

  useEffect(() => {
    getData();
  }, [])

  useEffect(() => {
    socket.on('log-inserted', () => {
      console.log('homepage socket ok');
      getData();
    });
  }, [])


  return (
    <div className="homepage" >
      <Header title={"USAC Squid Game"} />
      {
        isLoaded
          ?
          (
            <div className="home__content">
              <h3>Last 10 games</h3>
              <div className="home__table">
                <table>
                  <tbody>
                    <tr>
                      <th>&nbsp;&nbsp;Game#</th>
                      <th>Player#</th>
                      <th>Game Name</th>
                    </tr>
                    {
                      state.last_games.map((game) => {
                        return (
                          <tr key={game._id}>
                            <td>&nbsp;&nbsp; {game.gameid}</td>
                            <td>&nbsp;{game.winner}</td>
                            <td>{game.gamename}</td>
                          </tr>
                        );
                      })
                    }
                  </tbody>
                </table>
              </div>

              <h3>Top 10 Players</h3>
              <div className="home__table">
                <table>
                  <tbody>
                    <tr>
                      <th>&nbsp;&nbsp;Player#</th>
                      <th>Wins</th>
                    </tr>
                    {
                      state.top10_players.map((game) => {
                        return (
                          <tr key={game._id}>
                            <td>&nbsp; &nbsp; {game._id}</td>
                            <td>&nbsp;{game.wins}</td>
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

export default Homepage;