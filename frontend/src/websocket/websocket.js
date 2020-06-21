import React, { createContext } from 'react'
import { useDispatch } from "react-redux";
import {setTeam} from "../redux/actions";

const WebSocketContext = createContext(null)
export { WebSocketContext }


// type Message struct {
//     GameID string `json:"gameID"`
//     Kind   string `json:"kind"`
//     Team   int    `json:"team"`
//     Idx    int    `json:"idx"`
//     Row    int    `json:"row"`
//     Col    int    `json:"col"`
//
//     tsuro.Options
// }

// type Options struct {
//     Players int `json:"players"` // number of players
//     Size    int `json:"size"`    // width and height of the board
//     Time    int `json:"time"`    // timer length, -1 means no timer
// }

export default ({ children }) => {
    let server = "localhost:8080"
    let socket;
    let ws;

    const dispatch = useDispatch();

    const sendMessage = (msg) => {

    }

    if (!socket) {
        let socket = new WebSocket('wss://' + server + '/ws');
        socket.onmessage = (message) => {
            // team: "Neutral",
            //     teams: {},
            // board: [],
            //     turn: "Neutral",
            //     winner: "Neutral",
            //     hand: [],
            //     started: false,
            //     time: -1
            dispatch(setTeam(message.team))
        }
    }

    // const sendMessage = (roomId, message) => {
    //     const payload = {
    //         roomId: roomId,
    //         data: message
    //     }
    //     socket.emit("event://send-message", JSON.stringify(payload));
    //     dispatch(updateChatLog(payload));
    // }
    //
    // if (!socket) {
    //     socket = io.connect(WS_BASE)
    //
    //     socket.on("event://get-message", (msg) => {
    //         const payload = JSON.parse(msg);
    //         dispatch(updateChatLog(payload));
    //     })
    //
    //     ws = {
    //         socket: socket,
    //         sendMessage
    //     }
    // }

    return (
        <WebSocketContext.Provider value={ws}>
            {children}
        </WebSocketContext.Provider>
    )
}

