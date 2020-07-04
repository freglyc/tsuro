import {adjectives, nouns} from "../../words";
import {SET_CONNECTED, SET_GAMEID, SET_JOINED, SET_PAGE} from "../actionTypes";

const initialState = {
    page: "HOME",
    joined: false,
    gameID: adjectives[Math.floor(Math.random() * 50)] + "-" + nouns[Math.floor(Math.random() * 50)],
    connected: false,
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case SET_PAGE:
            return { ...state, page: action.page }
        case SET_JOINED:
            return { ...state, joined: action.joined }
        case SET_GAMEID:
            return { ...state, gameID: action.gameID }
        case SET_CONNECTED:
            return { ...state, connected: action.connected }
        default:
            return state
    }
}