import React, {useContext, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {WebSocketContext} from "./websocket/websocket";
import {setChange, setGameID, setJoined, setPage, setPlayers, setTimer} from "./redux/actions";

export default function HomePage() {
    const [advanced, setAdvanced] = useState(false)
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);
    let gameID = useSelector(state => state.site.gameID);
    let players = useSelector(state => state.options.players);
    let timer = useSelector(state => state.options.timer);
    let change = useSelector(state => state.options.change);

    function handleClick(e) {
        e.preventDefault();
        if (gameID.includes(" ") || gameID.length < 3) return
        let data = { "gameID": gameID, "kind":"join", "team":"Neutral", "idx":-1, "row":-1, "col":-1,
            "players": players, "size":6, "time": timer ? 20 : -1, "change": change };
        ws.sendMessage(data);
        dispatch(setJoined(true));
        dispatch(setPage("GAME"));
        window.history.pushState(null, '', '/' + gameID);
    }
    return (
        <div className="flexbox flex-column flex-center full-height">
            <div className="flexbox flex-column flex-center half-width">
                <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>
                <p className="standard-txt lighter-txt gray large-padding-top">
                    Play 2-8 player Tsuro against friends on one or more devices.
                    To create a game or join an existing one, enter a game ID and click 'Go'.
                </p>
                <form className="flexbox large-padding-top full-width" onSubmit={handleClick}>
                    <input className="input" autoFocus type="text" value={gameID}
                           onChange={(e) => dispatch(setGameID(e.target.value))}/>
                    <button className="goBtn" onClick={handleClick}>Go</button>
                </form>
                <div className="flexbox flex-self-end small-padding-top">
                    <div className="flexbox flex-center small-padding-right">
                        <button className="fas fa-cog inverse gear" onClick={(e) => {
                            e.preventDefault();
                            dispatch(setPage("SETTINGS"))
                        }}/>
                    </div>
                    <div className="flexbox flex-center">
                        <label className="small-padding-right standard-txt boldest-txt blue" htmlFor="players">PLAYERS</label>
                        <select className="small-txt boldest-txt select inverse" name="players" id="players"
                                onChange={(e) => {
                                        players = parseInt(e.target.value);
                                        dispatch(setPlayers(players))
                                }}>
                            <option value="2">2</option>
                            <option value="3">3</option>
                            <option value="4">4</option>
                            <option value="5">5</option>
                            <option value="6">6</option>
                            <option value="7">7</option>
                            <option value="8">8</option>
                        </select>
                    </div>
                </div>
                <div className="flexbox flex-self-end small-padding-top">
                    <button className="inverse advanced pointer" onClick={(e) => {
                        e.preventDefault();
                        setAdvanced(!advanced);
                    }}>{advanced ? "hide advanced options" : "show advanced options"}</button>
                </div>
                { advanced ?
                    <div className="full-width large-padding-top">
                        <div className="flexbox space-between full-width">
                            <div>
                                <h2 className="standard-txt boldest-txt inverse">TIMER</h2>
                                <p className="small-txt gray">enable a 20 second turn timer</p>
                            </div>
                            <label className="switch">
                                <input type="checkbox" defaultChecked={timer} onChange={(e) => {
                                    e.stopPropagation();
                                    timer = !timer;
                                    dispatch(setTimer(timer));
                                }}/>
                                <span className="slider round"/>
                            </label>
                        </div>
                        <div className="flexbox space-between full-width medium-padding-top">
                            <div>
                                <h2 className="standard-txt boldest-txt inverse">TEAM CHANGE</h2>
                                <p className="small-txt gray">allow players to changing teams after first selecting one</p>
                            </div>
                            <label className="switch">
                                <input type="checkbox" defaultChecked={change} onChange={(e) => {
                                    e.stopPropagation();
                                    change = !change;
                                    dispatch(setChange(change));
                                }}/>
                                <span className="slider round"/>
                            </label>
                        </div>
                    </div> : null
                }
            </div>
            <div className="absolute bottom">
                <p className="small-txt lighter-txt gray">Keep the developer <a target="_blank" rel="noopener noreferrer" className="gray" href="https://www.buymeacoffee.com/cfregly">caffeinated</a></p>
            </div>
        </div>
    )
}