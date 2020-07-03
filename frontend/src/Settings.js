import React from "react";
import {useDispatch, useSelector} from "react-redux";
import {setBlind, setDark, setPage} from "./redux/actions";

export default function SettingsPage() {
    const dispatch = useDispatch();
    const joined = useSelector(state => state.site.joined);
    let dark = useSelector(state => state.settings.dark);
    let blind = useSelector(state => state.settings.blind);
    return (
        <div className="flexbox flex-column flex-center full-height">
            <button className="absolute exit" onClick={(e) => {
                e.preventDefault();
                dispatch(setPage(joined ? "GAME" : "HOME"))
            }}/>
            <div className="flexbox flex-column flex-center half-width">
                <div className="flexbox flex-column flex-center">
                    <h1 className="title-txt large-padding-top"><a className="red remove-hyperlink" href={'http://' + window.location.host}>Tsuro</a></h1>
                    <h1 className="small-txt boldest-txt flex-self-end inverse">SETTINGS</h1>
                </div>
                <div className="full-width large-padding-top">
                    <div className="flexbox space-between full-width">
                        <div>
                            <h2 className="standard-txt boldest-txt inverse">DARK MODE</h2>
                            <p className="small-txt gray">darken the mood and may conserve battery life</p>
                        </div>
                        <label className="switch">
                            <input type="checkbox" defaultChecked={dark} onChange={(e) => {
                                e.stopPropagation();
                                dark = !dark;
                                dispatch(setDark(dark));
                                Settings.save({"dark": dark, "blind": blind});
                                if (dark) document.body.setAttribute('data-theme', 'dark');
                                else document.body.removeAttribute('data-theme')
                            }}/>
                            <span className="slider round"/>
                        </label>
                    </div>
                    <div className="flexbox space-between full-width medium-padding-top">
                        <div>
                            <h2 className="standard-txt boldest-txt inverse">*WIP* COLOR BLIND MODE</h2>
                            <p className="small-txt gray">add patterns to colors to distinguish teams</p>
                        </div>
                        <label className="switch">
                            <input disabled={true} type="checkbox" defaultChecked={blind} onChange={(e) => {
                                e.stopPropagation();
                                blind = !blind;
                                dispatch(setBlind(blind))
                                Settings.save({"dark": dark, "blind": blind});
                            }}/>
                            <span className="slider round"/>
                        </label>
                    </div>
                </div>
            </div>
            <div className="absolute bottom">
                <p className="small-txt lighter-txt gray">Keep the developer <a target="_blank" rel="noopener noreferrer" className="gray" href="https://www.buymeacoffee.com/cfregly">caffeinated</a></p>
            </div>
        </div>
    )
}

// Saves settings to local storage
export class Settings {
    // Load settings from location storage if there
    static load() { return JSON.parse(window.localStorage.getItem('settings')) || {}; }
    // Save settings to local storage
    static save(settings) { window.localStorage.setItem('settings', JSON.stringify(settings)); }
}