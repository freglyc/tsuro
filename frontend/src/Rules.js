import React from "react";
import {useDispatch, useSelector} from "react-redux";
import {setPage} from "./redux/actions";

export default function RulesPage() {
    const dispatch = useDispatch();
    const joined = useSelector(state => state.site.joined);
    return (
        <div className="flexbox flex-column flex-center full-height">
            <button className="absolute exit" onClick={(e) => {
                e.preventDefault();
                dispatch(setPage(joined ? "GAME" : "HOME"));
            }}/>
            <div className="flexbox flex-column flex-center half-width">
                <div className="flexbox flex-column flex-center">
                    <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>
                    <h1 className="small-txt boldest-txt flex-self-end inverse">RULES</h1>
                </div>
                <div className="flexbox flex-column flex-self-start large-padding-top inverse standard-txt">
                    <div>
                        <span className="boldest-txt">OVERVIEW: </span>
                        Tsuro is a 2-8 player board game where players place tiles on the board to build paths beginning
                        at the edges and travel around the interior. The goal of the game is to keep your path from
                        leading off the board and from running into other players. The last player(s) standing wins.
                    </div>
                    <div className="small-padding-top">
                        <span className="boldest-txt">SELECTING A TEAM:</span> Upon joining a game, select a team by clicking a colored circle to the left of the board.
                    </div>
                    <div className="small-padding-top">
                        <span className="boldest-txt">ROTATING A TILE:</span> At any time you may rotate the tiles in your hand by clicking them.
                    </div>
                    <div className="small-padding-top">
                        <span className="boldest-txt">PLACING A TILE:</span> On your turn place a tile by dragging and dropping one from your hand onto the space that is directly adjacent to your token.
                    </div>
                </div>
            </div>
            <div className="absolute bottom">
                <p className="small-txt lighter-txt gray">Keep the developer <a target="_blank" rel="noopener noreferrer" className="gray" href="https://www.buymeacoffee.com/cfregly">caffeinated</a></p>
            </div>
        </div>
    )
}