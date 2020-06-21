import {
    SET_BOARD,
    SET_GAMEID, SET_HAND,
    SET_PAGE,
    SET_PLAYERS, SET_STARTED, SET_TEAM, SET_TEAMS, SET_TIME, SET_TURN, SET_WINNER,
    TOGGLE_BLIND,
    TOGGLE_CHANGE,
    TOGGLE_DARK,
    TOGGLE_JOINED,
    TOGGLE_TIMER
} from "./actionTypes";

// Options
export const setPlayers = (players) => ({
    type: SET_PLAYERS,
    players: players
});
export const toggleTimer = () => ({
    type: TOGGLE_TIMER
});
export const toggleChange = () => ({
    type: TOGGLE_CHANGE
});

// Settings
export const toggleDark = () => ({
    type: TOGGLE_DARK
});
export const toggleBlind = () => ({
    type: TOGGLE_BLIND
});

// Site
export const setPage = (page) => ({
    type: SET_PAGE,
    page: page
});
export const toggleJoined = () => ({
    type: TOGGLE_JOINED
});
export const setGameID = (gameID) => ({
    type: SET_GAMEID,
    gameID: gameID
});

// Game
export const setTeam = (team) => ({
    type: SET_TEAM,
    team: team
});
export const setTeams = (teams) => ({
    type: SET_TEAMS,
    teams: teams
});
export const setBoard = (board) => ({
    type: SET_BOARD,
    board: board
});
export const setTurn = (turn) => ({
    type: SET_TURN,
    turn: turn
});
export const setWinner = (winner) => ({
    type: SET_WINNER,
    winner: winner
});
export const setHand = (hand) => ({
    type: SET_HAND,
    hand: hand
});
export const setStarted = (started) => ({
    type: SET_STARTED,
    started: started
});
export const setTime = (time) => ({
    type: SET_TIME,
    time: time
});