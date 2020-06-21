import {TOGGLE_BLIND, TOGGLE_DARK} from "../actionTypes";

const initialState = {
    dark: false,
    blind: false
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case TOGGLE_DARK:
            return { ...state, dark: !state.dark }
        case TOGGLE_BLIND:
            return { ...state, blind: !state.blind }
        default:
            return state
    }
}