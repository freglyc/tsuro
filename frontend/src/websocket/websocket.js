import React, { createContext } from 'react'
import {useDispatch, useSelector} from "react-redux";
import {
    setBoard,
    setChange,
    setConnected,
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
    let server = "fregly-tsuro.herokuapp.com"
    let socket;
    let ws;
    let connected = useSelector(state => state.site.connected);
    let queue = [] // holds queue of backlogged messages to be sent if not connected

    const dispatch = useDispatch();
    const sendMessage = (msg) => {
        if (connected) socket.send(JSON.stringify(msg))
        else if (!queue.includes(msg)) queue.push(msg)
    }
    if (!socket) {
        socket = new WebSocket('wss://' + server + '/ws');
        socket.onopen = () => {
            connected = true
            dispatch(setConnected(connected))
            for (let i = 0; i < queue.length; i++) sendMessage(queue[i])
            queue = []
        }
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
        socket.onclose = () => {
            connected = false
            dispatch(setConnected(connected))
        }
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

