import React, { createContext } from 'react'
import { useDispatch } from "react-redux";
import {setBoard, setHand, setStarted, setTeams, setTime, setTurn, setWinner} from "../redux/actions";

const WebSocketContext = createContext(null)
export { WebSocketContext }

export default ({ children }) => {
    let server = "localhost:8080"
    let socket;
    let ws;
    let connected = false;

    const dispatch = useDispatch();
    const sendMessage = (msg) => {
        setTimeout(() => {
            socket.send(JSON.stringify(msg))
        }, 250)
        // if (socket && connected) {
        //     socket.send(JSON.stringify(msg))
        //     console.log("MSG SENT")
        // }
    }
    if (!socket) {
        socket = new WebSocket('ws://' + server + '/ws');
        socket.onopen = () => { connected = true; }
        socket.onmessage = (msg) => {
            console.log("RECEIVED")
            const message = JSON.parse(msg.data)
            dispatch(setTeams(message.game.teams))
            dispatch(setBoard(message.game.board))
            dispatch(setTurn(message.game.turn))
            dispatch(setWinner(message.game.winner))
            dispatch(setHand(message.game.hand))
            dispatch(setStarted(message.game.started))
            dispatch(setTime(message.game.time))
            console.log(message);
        }
        socket.onclose = () => { connected = false; }
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

