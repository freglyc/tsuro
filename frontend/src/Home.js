import React, {useContext, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {WebSocketContext} from "./websocket/websocket";
import {setChange, setGameID, setPage, setPlayers, setTimer} from "./redux/actions";

export default function HomePage() {
    const [advanced, setAdvanced] = useState(false)
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);

    const gameID = useSelector(state => state.site.gameID);
    let timer = useSelector(state => state.options.timer);
    let change = useSelector(state => state.options.change);

    return (
        <div className="flexbox flex-column flex-center full-height">
            <div className="flexbox flex-column flex-center half-width">
                <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>
                <p className="standard-txt lighter-txt gray large-padding-top">
                    Play 2-8 player Tsuro against friends on one or more devices.
                    To create a game or join an existing one, enter a game ID and click 'Go'.
                </p>
                <form className="flexbox large-padding-top full-width" onSubmit={ console.log("TODO") }>
                    <input className="input" autoFocus type="text" value={gameID}
                           onChange={(e) => dispatch(setGameID(e.target.value))}/>
                    <button className="goBtn" onClick={ console.log("TODO") }>Go</button>
                </form>
                <div className="flexbox flex-self-end small-padding-top">
                    <div className="flexbox flex-center small-padding-right">
                        <button className="fas fa-cog dark gear" onClick={(e) => {
                            e.preventDefault();
                            dispatch(setPage("SETTINGS"))
                        }}/>
                    </div>
                    <div className="flexbox flex-center">
                        <label className="small-padding-right standard-txt boldest-txt blue" htmlFor="players">PLAYERS</label>
                        <select className="small-txt boldest-txt select" name="players" id="players"
                                onChange={(e) => dispatch(setPlayers(parseInt(e.target.value)))}>
                            <option value="2">2</option>
                            <option value="3">3</option>
                            <option value="3">4</option>
                            <option value="3">5</option>
                            <option value="3">6</option>
                            <option value="3">7</option>
                            <option value="3">8</option>
                        </select>
                    </div>
                </div>
                <div className="flexbox flex-self-end small-padding-top">
                    <button onClick={(e) => {
                        e.preventDefault();
                        setAdvanced(!advanced);
                    }}>advanced options</button>
                </div>
                { advanced ?
                    <div className="full-width large-padding-top">
                        <div className="flexbox space-between full-width">
                            <div>
                                <h2 className="standard-txt boldest-txt dark">TIMER</h2>
                                <p className="small-txt gray">enable a 20 second timer</p>
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
                                <h2 className="standard-txt boldest-txt dark">TEAM CHANGE</h2>
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