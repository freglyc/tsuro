import * as React from "react";
import GamePage from "./gamepage/Game";
import HomePage from "./Home";
import SettingsPage, {Settings} from "./Settings";
import {setBlind, setDark, setGameID, setJoined, setPage} from "./redux/actions";
import RulesPage from "./Rules";
import {WebSocketContext} from "./websocket/websocket";
import {useContext, useEffect} from "react";
import {useDispatch, useSelector} from "react-redux";

export default function App() {
  const dispatch = useDispatch();
  const ws = useContext(WebSocketContext);
  const page = useSelector(state => state.site.page);
  const joined = useSelector(state => state.site.joined);
  const change = useSelector(state => state.options.change);

  // Load settings
  let settings = Settings.load();
  if (settings.dark) {
    dispatch(setDark(settings.dark));
    document.body.setAttribute('data-theme', 'dark');
  } else document.body.removeAttribute('data-theme');
  if (settings.blind) dispatch(setBlind(settings.blind));

  // Select page to render
  let render = <HomePage/>;
  if (page === "SETTINGS") render = <SettingsPage/>;
  else if (page === "RULES") render = <RulesPage/>;
  else if (page === "GAME") render = <GamePage/>;

  // Connect to game if url has gameID
  useEffect(() => {
      if (document.location.pathname !== "/" && !joined) {
        const gameID = document.location.pathname.slice(1);
        dispatch(setGameID(gameID))
        let data = { "gameID": gameID, "kind":"join", "team":"Neutral", "idx":-1, "row":-1, "col":-1,
          "players":2, "size":6, "time":-1, "change": change };
        ws.sendMessage(data);
        dispatch(setJoined(true))
        dispatch(setPage("GAME"))
      }
  });
  return (
      <div>
        <div className={"quibbble-banner"}>
          <a className={"quibbble"} href={"https://www.quibbble.com"} target="_blank">Quibbble.com</a>&nbsp;out now. Play Tsuro and more all for free!
        </div>
        { render }
      </div>)
}