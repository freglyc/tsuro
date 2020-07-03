import React from "react";

export default function Edge(props) {
    let path;
    const edge = JSON.stringify(props.edge);
    const color = props.color;
    switch (true) {
        case edge === JSON.stringify(["A", "B"]) || edge === JSON.stringify(["B", "A"]):
            path = <path fill={color} d="M37.5,19.54c-11.07,0-14-12.78-14-19.54h3c0,.67.15,16.54,11,16.54S48.5.67,48.5,0h3C51.5,6.76,48.57,19.54,37.5,19.54Z"/>
            break
        case edge === JSON.stringify(["A", "C"]) || edge === JSON.stringify(["C", "A"]):
            path = <path fill={color} d="M75,26.51C24.08,26.51,23.5,1.09,23.5,0h3c0,.06.13,5.88,6,11.63C38,17.05,49.89,23.51,75,23.51Z"/>
            break
        case edge === JSON.stringify(["A", "D"]) || edge === JSON.stringify(["D", "A"]):
            path = <path fill={color} d="M75,51.5c-26.59,0-39.2-14-45.09-25.83A63.14,63.14,0,0,1,23.5,0h3a61.37,61.37,0,0,0,6.14,24.42c8,16,22.28,24.08,42.36,24.08Z"/>
            break
        case edge === JSON.stringify(["A", "E"]) || edge === JSON.stringify(["E", "A"]):
            path = <path fill={color} d="M48.5,75h3c0-25.67-3.9-29.56-12.46-38.08l-1-1C26.5,24.38,26.5,23.92,26.5,0h-3c0,24.17,0,25.62,12.44,38.06l1,1C45.35,47.43,48.5,50.57,48.5,75Z"/>
            break
        case edge === JSON.stringify(["A", "F"]) || edge === JSON.stringify(["F", "A"]):
            path = <rect fill={color} x="23.5" width="3" height="75"/>
            break
        case edge === JSON.stringify(["A", "G"]) || edge === JSON.stringify(["G", "A"]):
            path = <path fill={color} d="M26.51,0C26.51,50.92,1.09,51.5,0,51.5v-3c.06,0,5.88-.13,11.63-6C17.05,37,23.51,25.11,23.51,0Z"/>
            break
        case edge === JSON.stringify(["A", "H"]) || edge === JSON.stringify(["H", "A"]):
            path = <path fill={color} d="M26.5,0C26.5,26.2.27,26.5,0,26.5v-3c1,0,23.5-.33,23.5-23.5Z"/>
            break
        case edge === JSON.stringify(["B", "C"]) || edge === JSON.stringify(["C", "B"]):
            path = <path fill={color} d="M75,26.5C48.8,26.5,48.5.27,48.5,0h3c0,1,.33,23.5,23.5,23.5Z"/>
            break
        case edge === JSON.stringify(["B", "D"]) || edge === JSON.stringify(["D", "B"]):
            path = <path fill={color} d="M48.49,0c0,50.92,25.42,51.5,26.5,51.5v-3c-.06,0-5.88-.13-11.63-6C58,37,51.49,25.11,51.49,0Z"/>
            break
        case edge === JSON.stringify(["B", "E"]) || edge === JSON.stringify(["E", "B"]):
            path = <rect fill={color} x="48.5" width="3" height="75"/>
            break
        case edge === JSON.stringify(["B", "F"]) || edge === JSON.stringify(["F", "B"]):
            path = <path fill={color} d="M26.5,75h-3c0-25.67,3.9-29.56,12.46-38.08l1-1C48.5,24.38,48.5,23.92,48.5,0h3c0,24.17,0,25.62-12.44,38.06l-1,1C29.65,47.43,26.5,50.57,26.5,75Z"/>
            break
        case edge === JSON.stringify(["B", "G"]) || edge === JSON.stringify(["G", "B"]):
            path = <path fill={color} d="M51.5,0c0,26.59-14,39.2-25.83,45.09A63.14,63.14,0,0,1,0,51.5v-3a61.37,61.37,0,0,0,24.42-6.14C40.4,34.33,48.5,20.08,48.5,0Z"/>
            break
        case edge === JSON.stringify(["B", "H"]) || edge === JSON.stringify(["H", "B"]):
            path = <path fill={color} d="M0,26.51C50.92,26.51,51.5,1.09,51.5,0h-3c0,.06-.13,5.88-6,11.63C37,17.05,25.11,23.51,0,23.51Z"/>
            break
        case edge === JSON.stringify(["C", "D"]) || edge === JSON.stringify(["D", "C"]):
            path = <path fill={color} d="M55.46,37.5c0-11.07,12.78-14,19.54-14v3c-.67,0-16.54.15-16.54,11s15.87,11,16.54,11v3C68.24,51.5,55.46,48.57,55.46,37.5Z"/>
            break
        case edge === JSON.stringify(["C", "E"]) || edge === JSON.stringify(["E", "C"]):
            path = <path fill={color} d="M48.49,75c0-50.92,25.42-51.5,26.5-51.5v3c-.06,0-5.88.13-11.63,6C58,38,51.49,49.89,51.49,75Z"/>
            break
        case edge === JSON.stringify(["C", "F"]) || edge === JSON.stringify(["F", "C"]):
            path = <path fill={color} d="M23.5,75c0-26.59,14-39.2,25.83-45.09A63.14,63.14,0,0,1,75,23.5v3a61.37,61.37,0,0,0-24.42,6.14C34.6,40.67,26.5,54.92,26.5,75Z"/>
            break
        case edge === JSON.stringify(["C", "G"]) || edge === JSON.stringify(["G", "C"]):
            path = <path fill={color} d="M0,48.5v3c25.67,0,29.56-3.9,38.08-12.46l1-1C50.62,26.5,51.08,26.5,75,26.5v-3c-24.17,0-25.62,0-38.06,12.44l-1,1C27.57,45.35,24.43,48.5,0,48.5Z"/>
            break
        case edge === JSON.stringify(["C", "H"]) || edge === JSON.stringify(["H", "C"]):
            path = <rect fill={color} y="23.5" width="75" height="3"/>
            break
        case edge === JSON.stringify(["D", "E"]) || edge === JSON.stringify(["E", "D"]):
            path = <path fill={color} d="M48.5,75c0-26.2,26.23-26.5,26.5-26.5v3c-1,0-23.5.33-23.5,23.5Z"/>
            break
        case edge === JSON.stringify(["D", "F"]) || edge === JSON.stringify(["F", "D"]):
            path = <path fill={color} d="M75,48.49c-50.92,0-51.5,25.42-51.5,26.5h3c0-.06.13-5.88,6-11.63C38,58,49.89,51.49,75,51.49Z"/>
            break
        case edge === JSON.stringify(["D", "G"]) || edge === JSON.stringify(["G", "D"]):
            path = <rect fill={color} y="48.5" width="75" height="3"/>
            break
        case edge === JSON.stringify(["D", "H"]) || edge === JSON.stringify(["H", "D"]):
            path = <path fill={color} d="M0,26.5v-3c25.67,0,29.56,3.9,38.08,12.46l1,1C50.62,48.5,51.08,48.5,75,48.5v3c-24.17,0-25.62,0-38.06-12.44l-1-1C27.57,29.65,24.43,26.5,0,26.5Z"/>
            break
        case edge === JSON.stringify(["E", "F"]) || edge === JSON.stringify(["F", "E"]):
            path = <path fill={color} d="M37.5,55.46c11.07,0,14,12.78,14,19.54h-3c0-.67-.15-16.54-11-16.54S26.5,74.33,26.5,75h-3C23.5,68.24,26.43,55.46,37.5,55.46Z"/>
            break
        case edge === JSON.stringify(["E", "G"]) || edge === JSON.stringify(["G", "E"]):
            path = <path fill={color} d="M0,48.49c50.92,0,51.5,25.42,51.5,26.5h-3c0-.06-.13-5.88-6-11.63C37,58,25.11,51.49,0,51.49Z"/>
            break
        case edge === JSON.stringify(["E", "H"]) || edge === JSON.stringify(["H", "E"]):
            path = <path fill={color} d="M0,23.5c26.59,0,39.2,14,45.09,25.83A63.14,63.14,0,0,1,51.5,75h-3a61.37,61.37,0,0,0-6.14-24.42C34.33,34.6,20.08,26.5,0,26.5Z"/>
            break
        case edge === JSON.stringify(["F", "G"]) || edge === JSON.stringify(["G", "F"]):
            path = <path fill={color} d="M0,48.5c26.2,0,26.5,26.23,26.5,26.5h-3c0-1-.33-23.5-23.5-23.5Z"/>
            break
        case edge === JSON.stringify(["F", "H"]) || edge === JSON.stringify(["H", "F"]):
            path = <path fill={color} d="M26.51,75C26.51,24.08,1.09,23.5,0,23.5v3c.06,0,5.88.13,11.63,6C17.05,38,23.51,49.89,23.51,75Z"/>
            break
        case edge === JSON.stringify(["G", "H"]) || edge === JSON.stringify(["H", "G"]):
            path = <path fill={color} d="M19.54,37.5c0,11.07-12.78,14-19.54,14v-3c.67,0,16.54-.15,16.54-11S.67,26.5,0,26.5v-3C6.76,23.5,19.54,26.43,19.54,37.5Z"/>
            break
        default:
            path = <path/>
            break
    }
    return path
}
