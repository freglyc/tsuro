import React, {useContext} from "react";
import {useDispatch, useSelector} from "react-redux";
import {WebSocketContext} from "./websocket/websocket";
import Tile from "./Tile";
import {setPage, setTeam} from "./redux/actions";

export default function GamePage() {
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);

    const gameID = useSelector(state => state.site.gameID);
    const team = useSelector(state => state.game.team);
    const turn = useSelector(state => state.game.turn);
    const winner = useSelector(state => state.game.winner);
    const teams = useSelector(state => state.game.teams);
    const board = useSelector(state => state.game.board);

    let hand = []
    teams.forEach(t => { if (t.color === team) hand = t.hand.tiles })

    let placeMsg = {
        "gameID": gameID,
        "kind":"place",
        "team":team,
        "idx":-1,
        "row":-1,
        "col":-1,
        "players":-1,
        "size":-1,
        "time":-1
    }
    let rotateMsg = {
        "gameID": gameID,
        "kind":"rotateRight",
        "team":team,
        "idx":-1,
        "row":-1,
        "col":-1,
        "players":-1,
        "size":-1,
        "time":-1
    }
    let resetMsg = {
        "gameID": gameID,
        "kind":"reset",
        "team":"Neutral",
        "idx":-1,
        "row":-1,
        "col":-1,
        "players":-1,
        "size":-1,
        "time":-1
    }

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
                                                    return <div className="empty" key={idx1 + "," + idx2}
                                                                onDragOver={ e => {e.preventDefault();}}
                                                                onDrop={ e => {
                                                                    const id = e.dataTransfer.getData('text');
                                                                    placeMsg['row'] = idx1
                                                                    placeMsg['col'] = idx2
                                                                    if (id === "hand0") placeMsg['idx'] = 0
                                                                    else if (id === "hand1") placeMsg['idx'] = 1
                                                                    else if (id === "hand2") placeMsg['idx'] = 2
                                                                    ws.sendMessage(placeMsg);
                                                                    e.dataTransfer.clearData();
                                                                }}>
                                                        <Tile key={"tile" + idx1 + "," + idx2} edges={tile.edges} paths={tile.paths} row={idx1} col={idx2}/>
                                                    </div>
                                                })}
                                            </div>
                                        )
                                    })
                                }</div>
                            </div>
                        </div>

                        <div className="flexbox space-between full-width small-padding-top">
                            {
                                // this.props.started && this.props.timer ?
                                    // <Timer time={this.props.time} currentTime={this.props.currentTime} turn={this.props.turn} winner={this.props.winner}/> :
                                    // this.props.timer ?
                                    //     <div className="standard-txt boldest-txt dark">time: {this.props.time}</div> :
                                    //     <div/>
                                <div/>
                            }
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
                                <div id="hand0" className="pointer"
                                     onClick={e => {
                                        e.preventDefault();
                                        rotateMsg["idx"] = 0;
                                        ws.sendMessage(rotateMsg);
                                     }}
                                     draggable={true}
                                     onDragStart={e => e.dataTransfer.setData('text/plain', e.target.id)}>
                                    <Tile edges={hand[0].edges} paths={hand[0].paths} row={-1} col={-1}/>
                                </div> : null}
                        </div>
                        <div className="empty small-margin-top hide-overflow">
                            { hand.length > 1 ?
                                <div id="hand1" className="pointer"
                                     onClick={e => {
                                         e.preventDefault();
                                         rotateMsg["idx"] = 1;
                                         ws.sendMessage(rotateMsg);
                                     }}
                                     draggable={true}
                                     onDragStart={e => e.dataTransfer.setData('text/plain', e.target.id)}>
                                    <Tile edges={hand[1].edges} paths={hand[1].paths} row={-1} col={-1}/></div> : null}
                        </div>
                        <div className="empty small-margin-top hide-overflow">
                            { hand.length > 2 ?
                                <div id="hand2" className="pointer"
                                     onClick={e => {
                                         e.preventDefault();
                                         rotateMsg["idx"] = 2;
                                         ws.sendMessage(rotateMsg);
                                     }}
                                     draggable={true}
                                     onDragStart={e => e.dataTransfer.setData('text/plain', e.target.id)}>
                                <Tile edges={hand[2].edges} paths={hand[2].paths} row={-1} col={-1}/></div> : null}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
