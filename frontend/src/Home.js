import React, {useContext} from "react";
import {useDispatch} from "react-redux";
import {WebSocketContext} from "./websocket/websocket";

export default function HomePage() {
    const dispatch = useDispatch();
    const ws = useContext(WebSocketContext);

    return (<div>HOME</div>)
}