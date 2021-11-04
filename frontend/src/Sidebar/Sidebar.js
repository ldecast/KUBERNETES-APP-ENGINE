import React, { useState } from "react";
import "./Sidebar.css";
import HomeIcon from "@material-ui/icons/Home";
import PersonIcon from "@material-ui/icons/Person";
import EqualizerIcon from '@material-ui/icons/Equalizer';
import ReceiptIcon from '@material-ui/icons/Receipt';
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";
import DehazeIcon from '@material-ui/icons/Dehaze';
import RedisIcon from '@material-ui/icons/Timeline';
import { Link } from "react-router-dom";

function Sidebar(props) {

  const [class_name, setState] = useState({
    root: "root__sidebar",
    side: "sidebar",
    dehaze: "dehaze"
  });

  function SetVisible() {
    let root = (class_name.root === 'root__sidebar' ? 'root__hidden' : 'root__sidebar');
    let side = (class_name.side === 'sidebar' ? 'sidebar__hidden' : 'sidebar');
    let dehaze = (class_name.dehaze === 'dehaze' ? 'dehaze__hidden' : 'dehaze');
    let tmp = {
      root: root,
      side: side,
      dehaze: dehaze
    };
    setState(tmp);
  }

  return (
    <div className={class_name.root}>
      <div className={class_name.dehaze} onClick={SetVisible}>
        <DehazeIcon />
      </div >
      <div className={class_name.side}>
        <Link className={`sidebarOption`} to={`/home`} onClick={props.click} >
          <HomeIcon />
          {"Homepage"}
        </Link>
        <Link className={`sidebarOption`} to={`/stats`} onClick={props.click}>
          <PersonIcon />
          {"Gamer stats"}
        </Link>
        <Link className={`sidebarOption`} to={`/charts`} onClick={props.click}>
          <EqualizerIcon />
          {"Charts"}
        </Link>
        <Link className={`sidebarOption`} to={`/transactions`} onClick={props.click}>
          <ReceiptIcon />
          {"Transactions"}
        </Link>
        <Link className={`sidebarOption`} to={`/redisReports`} onClick={props.click}>
          <RedisIcon />
          {"Redis Reports"}
        </Link>
        <a className={`sidebarOption`} href={`https://github.com/sergioarmgpl/operating-systems-usac-course/blob/master/lang/en/projects/project1v4/project1.md`} target="_blank" rel="noreferrer">
          <MoreHorizIcon />
          {"More"}
        </a>
        <div className="__Copyright">
          <p>Sistemas Operativos 1&nbsp;·&nbsp;Grupo 7</p>
          <p>USAC &copy; 2S 2021&nbsp;·&nbsp;Proyecto 2</p>
        </div>
      </div>
    </div>
  );
}

export default Sidebar;