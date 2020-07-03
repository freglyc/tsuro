import React, { createContext } from 'react'
import { useDispatch } from "react-redux";
import {
    setBoard,
    setChange,
    setCountdown,
    setHand,
    setStarted,
    setTeams,
    setTime,
    setTurn,
    setWinner
} from "../redux/actions";

const WebSocketContext = createContext(null)
export { WebSocketContext }

export default ({ children }) => {
    let server = "localhost:8080"
    let socket;
    let ws;

    const dispatch = useDispatch();
    const sendMessage = (msg) => {
        setTimeout(() => {
            socket.send(JSON.stringify(msg))
        }, 250)
    }
    if (!socket) {
        socket = new WebSocket('ws://' + server + '/ws');
        socket.onopen = () => {}
        socket.onmessage = (msg) => {
            const message = JSON.parse(msg.data)
            dispatch(setTeams(message.game.teams))
            dispatch(setBoard(message.game.board))
            dispatch(setTurn(message.game.turn))
            dispatch(setWinner(message.game.winner))
            dispatch(setHand(message.game.hand))
            dispatch(setStarted(message.game.started))
            dispatch(setTime(message.game.time))
            dispatch(setCountdown(message.game.countdown))
            dispatch(setChange(message.game.change))
        }
        socket.onclose = () => {}
        ws = {
            socket: socket,
            sendMessage
        }
    }
    return (
        <WebSocketContext.Provider value={ws}>
            {children}
        </WebSocketContext.Provider>
    )
}

