import {useDispatch, useSelector} from "react-redux";
import React, {useEffect} from "react";
import {setCountdown} from "../redux/actions";

export default function Timer() {
    const dispatch = useDispatch();
    const countdown = useSelector(state => state.game.countdown);
    const time = useSelector(state => state.game.time);
    const started = useSelector(state => state.game.started);
    const winner = useSelector(state => state.game.winner);
    const isActive = started && winner.length <= 0
    useEffect(() => {
        let interval = null;
        if (isActive) {
            interval = setInterval(() => {
                if (countdown > 0) dispatch(setCountdown(countdown - 1));
            }, 1000);
        } else if (!isActive || countdown === 0) {
            clearInterval(interval);
        }
        return () => clearInterval(interval);
    });
    return <div className="standard-txt boldest-txt inverse">time: { isActive ? countdown : time}</div>
}