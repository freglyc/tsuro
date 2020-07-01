import React from "react";
import Edge from "./Edge";
import {useSelector} from "react-redux";
import Notch from "./Notch";

export default function Tile(props) {

    // TESTS
    // const edges = [["A", "B"], ["C", "D"], ["E", "F"], ["G", "H"]];
    // const paths = {"Red": ["A", "B"]};

    const edges = props.edges;
    const paths = props.paths;
    const row = props.row;
    const col = props.col;
    const teams = useSelector(state => state.game.teams);

    let notch = null;
    for (let i = 0; i < teams.length; i++) {
        let team = teams[i];
        if (team.token.row === row && team.token.col === col) {
            notch = <Notch notch={team.token.notch} color={getColor(team.color)}/>
        }
    }

    return (
        <div className={edges !== null ? "dark-background" : ""}>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 75 75">
                {edges ? edges.map(edge => {
                    let color = "#979797";
                    if (paths) {
                        for (const [key, val] of Object.entries(paths)) {
                            if (JSON.stringify(edge) === JSON.stringify(val)) {
                                color = getColor(key);
                            }
                        }
                    }
                    return <Edge edge={edge} color={color}/>
                }) : null }
                {notch}
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