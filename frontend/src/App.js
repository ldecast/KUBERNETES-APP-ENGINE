import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import Sidebar from "./Sidebar/Sidebar";
import Homepage from './Homepage/Homepage';
import NotFoundPage from './NotFound/404';
import GamerStats from './GamerStats/GamerStats';
import Logs from './Logs/Logs';
import Charts from './Charts/Charts';
import Redis from './Redis/Redis';
import { useState } from 'react';

function App() {

  const [state, setState] = useState(0);

  function ClickRequest() {
    // console.log(state)
    setState(state + 1);
  }

  return (
    <div className="app" key={state}>
      <Router>
        <Sidebar click={ClickRequest} />
        <Switch>
          <Route exact path={["/", "/home"]}>
            <Homepage />
          </Route>
          <Route exact path={"/stats"}>
            <GamerStats />
          </Route>
          <Route exact path={"/charts"}>
            <Charts />
          </Route>
          <Route exact path={"/transactions"}>
            <Logs />
          </Route>
          <Route exact path={"/redisReports"}>
            <Redis />
          </Route>
          <Route component={NotFoundPage} />
        </Switch>
      </Router >
    </div>
  );
}

export default App;
