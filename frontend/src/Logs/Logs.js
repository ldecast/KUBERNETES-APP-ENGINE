import { useState, useEffect } from 'react';
import "./Logs.css";
import Header from '../Header/Header';
import Loading from '../Loading/Loading';
import socket from "../socket";

function Logs() {

  const [state, setState] = useState({
    logs: [],
    isLoaded: false
  });

  const getLogs = async () => {
    fetch(process.env.REACT_APP_API_URL + "/transactions")
      .then(response => response.json())
      .then(data => {
        // console.log(data);
        setState({
          logs: data,
          isLoaded: true
        });
      });
  };

  useEffect(() => {
    getLogs();
    return () => {
      socket.off('log-inserted');
    }
  }, [])

  useEffect(() => {
    socket.on('log-inserted', () => {
      console.log('logs sockets ok');
      getLogs();
    });
  }, [])


  return (
    <div className="logs" >
      <Header title={"Transactions"} />
      {
        state.isLoaded
          ?
          (
            <div className="logs__content">
              <h3>Log de datos almacenados</h3>
              <div className="logs__table">
                <table>
                  <tbody>
                    <tr>
                      <th>&nbsp;&nbsp;Request#</th>
                      <th>Game#</th>
                      <th>Game Name</th>
                      <th>Winner</th>
                      <th>Players</th>
                      <th>Worker</th>
                    </tr>
                    {
                      state.logs.map((log, i) => {
                        return (
                          <tr key={log._id}>
                            <td>&nbsp;&nbsp; {i + 1}</td>
                            <td>{log.gameid}</td>
                            <td>{log.gamename}</td>
                            <td>{log.winner}</td>
                            <td>{log.players}</td>
                            <td>{log.worker}</td>
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

export default Logs;