import React, {useContext} from 'react'
import { useDrop } from 'react-dnd'
import { ItemTypes } from './ItemTypes'
import {WebSocketContext} from "../websocket/websocket";
import {useSelector} from "react-redux";

// Drop logic for dropping a tile on the board
export const DropSpace = ({ droppable, row, col, children }) => {
    const ws = useContext(WebSocketContext);
    const gameID = useSelector(state => state.site.gameID);
    const team = useSelector(state => state.game.team);
    const change = useSelector(state => state.options.change);
    let msg = { "gameID": gameID, "kind": "place", "team": team, "idx": -1, "row": row, "col": col,
        "players": -1, "size": -1, "time": -1, "change": change }
    const [{}, drop] = useDrop({
        accept: ItemTypes.TILE,
        canDrop: () => droppable,
        drop: (obj) => {
            msg['idx'] = obj.idx
            ws.sendMessage(msg);
        },
    })
    return (
        <div ref={drop} className={"empty hide-overflow"}>{children}</div>
    )
}
