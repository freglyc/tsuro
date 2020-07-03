import React, {useContext, useEffect} from "react";
import {useDispatch, useSelector} from "react-redux";
import {WebSocketContext} from "../websocket/websocket";
import Tile from "./Tile";
import Timer from "./Timer";
import {setPage, setTeam} from "../redux/actions";
import {DropSpace} from "./DropSpace";
import favicon from "../resources/favicons/favicon.ico"
import red from "../resources/favicons/red.ico"
import blue from "../resources/favicons/blue.ico"
import green from "../resources/favicons/green.ico"
import yellow from "../resources/favicons/yellow.ico"
import orange from "../resources/favicons/orange.ico"
import pink from "../resources/favicons/pink.ico"
import purple from "../resources/favicons/purple.ico"
import turquoise from "../resources/favicons/turquoise.ico"

export default function GamePage() {
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);

    const gameID = useSelector(state => state.site.gameID);
    const team = useSelector(state => state.game.team);
    const turn = useSelector(state => state.game.turn);
    const winner = useSelector(state => state.game.winner);
    const teams = useSelector(state => state.game.teams);
    const board = useSelector(state => state.game.board);
    const time = useSelector(state => state.game.time);
    const players = useSelector(state => state.options.players);
    const change = useSelector(state => state.options.change);

    let hand = []
    teams.forEach(t => { if (t.color === team) hand = t.hand.tiles })

    let rotateMsg = { "gameID": gameID, "kind":"rotateRight", "team":team, "idx":-1, "row":-1, "col":-1,
        "players":players, "size":6, "time":-1, "change": change }
    let resetMsg = { "gameID": gameID, "kind":"reset", "team":"Neutral", "idx":-1, "row":-1, "col":-1,
        "players":players, "size":6, "time":-1, "change": change }

    useEffect(() => {
        if (winner.length !== 0) {
            document.getElementById("favicon").setAttribute("href", favicon);
        } else {
            let favi;
            if (turn === "Red") favi = red;
            else if (turn === "Yellow") favi = yellow;
            else if (turn === "Green") favi = green;
            else if (turn === "Blue") favi = blue;
            else if (turn === "Orange") favi = orange;
            else if (turn === "Purple") favi = purple;
            else if (turn === "Pink") favi = pink;
            else if (turn === "Turquoise") favi = turquoise;
            else favi = favicon;
            document.getElementById("favicon").setAttribute("href", favi);
        }
    })

    const turnColor = winner.length !== 0 ? winner[0].toLowerCase() : turn.toLowerCase();
    return (
        <div className="flexbox flex-column flex-center full-height">
            <div className="flexbox flex-column flex-center full-width">
                <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>

                <div className="flexbox full-width justify-center">
                    <div className="flexbox flex-column team-btns small-padding-right">
                        {teams.map(t => {
                            let background = t.color === team ? t.color.toLowerCase() + "-background" : ""
                            return <div key={t.color + "-btn"}
                                className={"team-btn " + background + " " + t.color.toLowerCase() + "-border pointer"}
                                onClick={e => {
                                    e.stopPropagation();
                                    if (team !== "Neutral" && !change) return
                                    dispatch(setTeam(t.color))
                                }}/>
                        })}
                    </div>

                    <div className="game-width">
                        <p className="flex-self-start small-txt lighter-txt inverse medium-padding-top">Share this link with friends: <a className="inverse" href={"https://" +  window.location.host + "/" + gameID}>{"https://" +  window.location.host + "/" + gameID}</a></p>
                        <hr className="full-width dark"/>

                        <div className="flexbox flex-column">
                            <p className={turnColor + " standard-txt boldest-txt flex-self-end"}>{
                                winner.length !== 0 ? winner.map(w => w.toLowerCase()) + " wins!" : turn.toLowerCase() + " turn"
                            }</p>
                        </div>

                        <div className="full-width">
                            <div className="flexbox flex-column flex-center smallest-margin-top medium-padding board">
                                <div className="full-width">{
                                    board.map((row, idx1) => {
                                        return (
                                            <div key={"row-" + idx1} className="full-width flexbox space-between">
                                                {row.map((tile, idx2) => {
                                                    return <DropSpace key={ "space" + idx1 + "," + idx2 } droppable={tile.edges === null} row={idx1} col={idx2}>
                                                        <Tile key={"tile" + idx1 + "," + idx2} edges={tile.edges} paths={tile.paths} row={idx1} col={idx2}/>
                                                    </DropSpace>
                                                })}
                                            </div>
                                        )
                                    })
                                }</div>
                            </div>
                        </div>

                        <div className="flexbox space-between full-width small-padding-top">
                            { time > 0 ? <Timer/> : <div/> }
                            <div className="flexbox flex-center">
                                <div className="flexbox flex-center smallest-padding-right">
                                    <button className="fas fa-cog inverse gear" onClick={(e) => {
                                        e.preventDefault();
                                        dispatch(setPage("SETTINGS"));
                                    }}/>
                                </div>
                                <div className="flexbox flex-center small-padding-right">
                                    <button onClick={(e) => {
                                        e.preventDefault();
                                        dispatch(setPage("RULES"));
                                    }} className="resetBtn smallest-txt bolder-txt">rules</button>
                                </div>
                                <div className="flexbox flex-center">
                                    <button onClick={e => {
                                        e.preventDefault();
                                        ws.sendMessage(resetMsg);
                                    }} className="resetBtn smallest-txt bolder-txt">new game</button>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className="flexbox flex-column small-padding-left hand">
                        <div className="empty hide-overflow">
                            { hand.length > 0 ?
                                <div className="pointer"
                                     onClick={e => {
                                        e.preventDefault();
                                        rotateMsg["idx"] = 0;
                                        ws.sendMessage(rotateMsg);
                                     }}>
                                    <Tile idx={0} edges={hand[0].edges} paths={hand[0].paths} row={-1} col={-1}/>
                                </div> : null}
                        </div>
                        <div className="empty small-margin-top hide-overflow">
                            { hand.length > 1 ?
                                <div className="pointer"
                                     onClick={e => {
                                         e.preventDefault();
                                         rotateMsg["idx"] = 1;
                                         ws.sendMessage(rotateMsg);
                                     }}>
                                    <Tile idx={1} edges={hand[1].edges} paths={hand[1].paths} row={-1} col={-1}/>
                                </div> : null}
                        </div>
                        <div className="empty small-margin-top hide-overflow">
                            { hand.length > 2 ?
                                <div className="pointer"
                                     onClick={e => {
                                         e.preventDefault();
                                         rotateMsg["idx"] = 2;
                                         ws.sendMessage(rotateMsg);
                                     }}>
                                    <Tile idx={2} edges={hand[2].edges} paths={hand[2].paths} row={-1} col={-1}/>
                                </div> : null}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
