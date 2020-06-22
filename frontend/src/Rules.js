import React from "react";
import {useDispatch, useSelector} from "react-redux";
import {setPage} from "./redux/actions";

export default function RulesPage() {
    const dispatch = useDispatch();
    const joined = useSelector(state => state.site.joined)
    return (
        <div className="flexbox flex-column flex-center full-height">
            <button className="absolute exit" onClick={(e) => {
                e.preventDefault();
                dispatch(setPage(joined ? "GAME" : "HOME"))
            }}/>
            <div className="flexbox flex-column flex-center half-width">
                <div className="flexbox flex-column flex-center">
                    <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>
                    <h1 className="small-txt boldest-txt flex-self-end dark">RULES</h1>
                </div>
                <div className="flexbox flex-column flex-self-start large-padding-top dark standard-txt">
                    <div>
                        Tsuro is a 2-8 player board game where players place tiles on the board to build paths beginning
                        at the edges and travel around the interior. The goal of the game is to keep your path from
                        leading off the board. The last player(s) standing wins.
                    </div>
                    <div className="small-padding-top">
                        <span className="boldest-txt">SELECTING A TEAM:</span> Upon joining a game, select a team by clicking the colored circles to the left of the screen.
                    </div>
                    <div className="small-padding-top">
                        <span className="boldest-txt">ROTATING A TILE:</span> At any time you may rotate the tiles in your hand by clicking them.
                    </div>
                    <div className="small-padding-top">
                        <span className="boldest-txt">PLACING A TILE:</span> On your turn place a tile by dragging and dropping your chosen tile onto the highlighted space on the board.
                    </div>
                </div>
            </div>
        </div>
    )
}