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
// import UserStats from './UserStats/UserStats';

function App() {
  return (
    <div className="app">
      <Router>
        <Sidebar />
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
