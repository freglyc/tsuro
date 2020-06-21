import {SET_PLAYERS, TOGGLE_CHANGE, TOGGLE_TIMER} from "../actionTypes";

const initialState = {
    players: 2,
    timer: false,
    change: true
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case SET_PLAYERS:
            return { ...state, page: action.players }
        case TOGGLE_TIMER:
            return { ...state, timer: !state.timer }
        case TOGGLE_CHANGE:
            return { ...state, change: !state.change }
        default:
            return state
    }
}