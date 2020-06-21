import { combineReducers } from "redux";
import site from "./site";
import options from "./options";
import settings from "./settings";
import game from "./game";


export default combineReducers({ site, options, settings, game })