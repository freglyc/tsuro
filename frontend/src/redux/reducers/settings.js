import {SET_BLIND, SET_DARK} from "../actionTypes";

const initialState = {
    dark: false,
    blind: false
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case SET_DARK:
            return { ...state, dark: action.dark }
        case SET_BLIND:
            return { ...state, blind: action.blind }
        default:
            return state
    }
}