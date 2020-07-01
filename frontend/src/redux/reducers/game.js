import {
    SET_BOARD,
    SET_HAND,
    SET_STARTED,
    SET_TEAM,
    SET_TEAMS, SET_TIME,
    SET_TURN,
    SET_WINNER,
} from "../actionTypes";

const initialState = {
    team: "Neutral",
    teams: {},
    board: [],
    turn: "Neutral",
    winner: [],
    hand: [],
    started: false,
    time: -1
}

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case SET_TEAM:
            return { ...state, team: action.team }
        case SET_TEAMS:
            return { ...state, teams: action.teams }
        case SET_BOARD:
            return { ...state, board: action.board }
        case SET_TURN:
            return { ...state, turn: action.turn }
        case SET_WINNER:
            return { ...state, winner: action.winner }
        case SET_HAND:
            return { ...state, hand: action.hand }
        case SET_STARTED:
            return { ...state, started: action.started }
        case SET_TIME:
            return { ...state, time: action.time }
        default:
            return state
    }
}