// Connect4 App Page

import * as React from "react";
import GamePage from "./Game";
import HomePage from "./Home";
import SettingsPage, {Settings} from "./Settings";
import {setGameID, setPage, toggleBlind, toggleDark, toggleJoined} from "./redux/actions";
import RulesPage from "./Rules";
import {WebSocketContext} from "./websocket/websocket";
import {useContext, useEffect} from "react";
import {useDispatch, useSelector} from "react-redux";

export default function App() {
  const dispatch = useDispatch();
  const ws = useContext(WebSocketContext);

  const page = useSelector(state => state.site.page);

  // Load settings
  let settings = Settings.load();
  if (settings.dark) {
    dispatch(toggleDark())
    document.body.setAttribute('data-theme', 'dark');
  } else document.body.removeAttribute('data-theme')
  if (settings.blind) dispatch(toggleBlind())

  // Select page to render
  let render = <HomePage/>;
  if (page === "SETTINGS") render = <SettingsPage/>
  else if (page === "RULES") render = <RulesPage/>
  else if (page === "GAME") render = <GamePage/>

  // Connect to game if url has gameID
  useEffect(() => {
      // Set game if in one
      if (document.location.pathname !== "/") {
        const gameID = document.location.pathname.slice(1)
        dispatch(setGameID(gameID))
        let data = {
          "gameID": gameID,
          "kind":"join",
          "team":"Neutral",
          "idx":-1,
          "row":-1,
          "col":-1,
          "players":2,
          "size":6,
          "time":-1
        };
        ws.sendMessage(data);
        dispatch(toggleJoined())
        dispatch(setPage("GAME"))
      }
  })

  return (<div> { render } </div>)
}