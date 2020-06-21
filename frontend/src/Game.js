import React, {useContext} from "react";
import {useDispatch, useSelector} from "react-redux";
import {WebSocketContext} from "./websocket/websocket";

export default function GamePage() {
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);

    const team = useSelector(state => state.game.team);
    const turn = useSelector(state => state.game.turn);
    const winner = useSelector(state => state.game.winner);
    const teams = useSelector(state => state.game.teams);
    const board = useSelector(state => state.game.board);

    return (<div> {team} </div>)
}
