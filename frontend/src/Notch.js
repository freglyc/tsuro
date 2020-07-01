import React from "react";

export default function Notch(props) {
    let path;
    const notch = props.notch;
    const color = props.color;
    switch (notch) {
        case "A":
            path = <circle fill={color} cx="25" cy="4" r="4"/>;
            break;
        case "B":
            path = <circle fill={color} cx="50" cy="4" r="4"/>;
            break;
        case "C":
            path = <circle fill={color} cx="71" cy="25" r="4"/>;
            break;
        case "D":
            path = <circle fill={color} cx="71" cy="50" r="4"/>;
            break;
        case "E":
            path = <circle fill={color} cx="50" cy="71" r="4"/>;
            break;
        case "F":
            path = <circle fill={color} cx="25" cy="71" r="4"/>;
            break;
        case "G":
            path = <circle fill={color} cx="4" cy="50" r="4"/>;
            break;
        case "H":
            path = <circle fill={color} cx="4" cy="25" r="4"/>;
            break;
        default:
            path = <path/>
    }
    return path
}