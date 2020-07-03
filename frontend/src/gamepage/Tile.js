import React from "react";
import Edge from "./Edge";
import {useSelector} from "react-redux";
import Notch from "./Notch";
import {ItemTypes} from "./ItemTypes";
import {useDrag} from "react-dnd";

export default function Tile({ edges, paths, row, col, idx }) {
    const [, drag] = useDrag({ item: { type: ItemTypes.TILE, idx: idx } })
    const teams = useSelector(state => state.game.teams);
    let notches = [];
    for (let i = 0; i < teams.length; i++) {
        let team = teams[i];
        if (team.token.row === row && team.token.col === col) {
            notches.push(<Notch key={team.color + "notch"} notch={team.token.notch} color={getColor(team.color)}/>);
        }
    }
    return (
        <div ref={drag} className={edges !== null ? "dark-background" : ""}>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 75 75">
                {edges ? edges.map(edge => {
                    let color = "#979797";
                    if (paths) {
                        for (const [key, val] of Object.entries(paths)) {
                            for (let i = 0; i < val.length; i ++) {
                                let e = val[i]
                                if ((edge[0] === e[0] && edge[1] === e[1]) ||
                                    (edge[1] === e[0] && edge[0] === e[1])) {
                                    color = getColor(key);
                                }
                            }
                        }
                    }
                    return <Edge key={row + "," + col + edge} edge={edge} color={color}/>
                }) : null }
                {notches.map(n => n)}
            </svg>
        </div>
    )
}

function getColor(find) {
    let color;
    switch (find) {
        case "Red":
            color = "#E53338";
            break;
        case "Blue":
            color = "#1C9BE6";
            break;
        case "Green":
            color = "#24D650";
            break;
        case "Yellow":
            color = "#F9EA36";
            break;
        case "Orange":
            color = "#FA9D33";
            break;
        case "Purple":
            color = "#C564EF";
            break;
        case "Pink":
            color = "#FF4FB1";
            break;
        case "Turquoise":
            color = "#44D7B6";
            break;
        default:
            color = "#979797";
    }
    return color;
}