import {
    SET_BLIND,
    SET_BOARD, SET_CHANGE, SET_DARK,
    SET_GAMEID, SET_HAND, SET_JOINED,
    SET_PAGE,
    SET_PLAYERS, SET_STARTED, SET_TEAM, SET_TEAMS, SET_TIME, SET_TIMER, SET_TURN, SET_WINNER,
} from "./actionTypes";

// Options
export const setPlayers = (players) => ({
    type: SET_PLAYERS,
    players: players
});
export const setTimer = (timer) => ({
    type: SET_TIMER,
    timer: timer
});
export const setChange = (change) => ({
    type: SET_CHANGE,
    change: change
});

// Settings
export const setDark = (dark) => ({
    type: SET_DARK,
    dark: dark
});
export const setBlind = (blind) => ({
    type: SET_BLIND,
    blind: blind
});

// Site
export const setPage = (page) => ({
    type: SET_PAGE,
    page: page
});
export const setJoined = (joined) => ({
    type: SET_JOINED,
    joined: joined
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