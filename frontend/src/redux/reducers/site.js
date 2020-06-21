import {adjectives, nouns} from "../../words";
import {SET_GAMEID, SET_PAGE, TOGGLE_JOINED} from "../actionTypes";

const initialState = {
    page: "HOME",
    joined: false,
    gameID: adjectives[Math.floor(Math.random() * 50)] + "-" + nouns[Math.floor(Math.random() * 50)],
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case SET_PAGE:
            return { ...state, page: action.page }
        case TOGGLE_JOINED:
            return { ...state, joined: !state.joined }
        case SET_GAMEID:
            return { ...state, gameID: action.gameID }
        default:
            return state
    }
}