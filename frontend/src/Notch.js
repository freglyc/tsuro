import React from "react";

export default function Notch(props) {
    let path;
    const notch = props.notch;
    const color = props.color;
    switch (notch) {
        case "A":
            path = <circle fill={color} cx="25" cy="5" r="5"/>;
            break;
        case "B":
            path = <circle fill={color} cx="50" cy="5" r="5"/>;
            break;
        case "C":
            path = <circle fill={color} cx="70" cy="25" r="5"/>;
            break;
        case "D":
            path = <circle fill={color} cx="70" cy="50" r="5"/>;
            break;
        case "E":
            path = <circle fill={color} cx="50" cy="70" r="5"/>;
            break;
        case "F":
            path = <circle fill={color} cx="25" cy="70" r="5"/>;
            break;
        case "G":
            path = <circle fill={color} cx="5" cy="50" r="5"/>;
            break;
        case "H":
            path = <circle fill={color} cx="5" cy="25" r="5"/>;
            break;
        default:
            path = <path/>
    }
    return path
}