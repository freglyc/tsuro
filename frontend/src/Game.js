import React, {useContext} from "react";
import {useDispatch, useSelector} from "react-redux";
import {WebSocketContext} from "./websocket/websocket";
import Tile from "./Tile";
import {setPage} from "./redux/actions";

export default function GamePage() {
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);

    const gameID = useSelector(state => state.site.gameID);
    const team = useSelector(state => state.game.team);
    const turn = useSelector(state => state.game.turn);
    const winner = useSelector(state => state.game.winner);
    const teams = useSelector(state => state.game.teams);
    const board = useSelector(state => state.game.board);

    return (
        <div className="flexbox flex-column flex-center full-height">
            <div className="flexbox flex-column flex-center game-width">
                <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>
                <p className="flex-self-start small-txt lighter-txt inverse medium-padding-top">Share this link with friends: <a className="inverse" href={"https://" +  window.location.host + "/" + gameID}>{"https://" +  window.location.host + "/" + gameID}</a></p>
                <hr className="full-width dark"/>
                <div className="full-width">

                    <div className="flexbox flex-column">
                        <p className={turn.toLowerCase() + " standard-txt boldest-txt flex-self-end"}>{
                            winner.length !== 0 ? winner.forEach(w => w.toLowerCase() + " ") + "wins!" : turn.toLowerCase() + " turn"
                        }</p>
                    </div>

                    <div className="full-width">
                        <div className="flexbox flex-column flex-center smallest-margin-top medium-padding board">
                            <div className="full-width">{
                                board.map((row, idx1) => {
                                    return (
                                        <div key={"row-" + idx1} className="full-width flexbox space-between">
                                            {row.map((tile, idx2) => {
                                                return <div className="empty" key={idx1 + "," + idx2} onClick={ e => {
                                                    e.preventDefault();
                                                    // TODO add functionality
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
                            <div className="flexbox flex-center small-padding-right">
                                <button className="fas fa-cog inverse gear" onClick={(e) => {
                                    e.preventDefault();
                                    dispatch(setPage("SETTINGS"));
                                }}/>
                            </div>
                            <div className="flexbox flex-center">
                                <button onClick={(e) => {
                                    e.preventDefault();
                                    // TODO send reset message
                                }} className="resetBtn smallest-txt bolder-txt">new game</button>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    )
}
