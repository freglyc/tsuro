import React from "react";
import {useDispatch} from "react-redux";

export default function SettingsPage() {
    const dispatch = useDispatch();
    return (<div>SETTINGS</div>)
}

// Saves settings to local storage
export class Settings {
    // Load settings from location storage if there
    static load() { return JSON.parse(window.localStorage.getItem('settings')) || {}; }
    // Save settings to local storage
    static save(settings) { window.localStorage.setItem('settings', JSON.stringify(settings)); }
}