import {SET_CHANGE, SET_PLAYERS, SET_TIMER} from "../actionTypes";

const initialState = {
    players: 2,
    timer: false,
    change: true
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case SET_PLAYERS:
            return { ...state, page: action.players }
        case SET_TIMER:
            return { ...state, timer: action.timer }
        case SET_CHANGE:
            return { ...state, change: action.change }
        default:
            return state
    }
}