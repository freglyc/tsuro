import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import store from "./redux/store";
import { Provider } from 'react-redux'
import WebSocketProvider from './websocket/websocket'
import { DndProvider } from 'react-dnd-multi-backend';
import MultiBackend from 'react-dnd-multi-backend';
import HTML5toTouch from 'react-dnd-multi-backend/dist/esm/HTML5toTouch';

ReactDOM.render(
  <Provider store={store}>
      <WebSocketProvider>
          <DndProvider backend={MultiBackend} options={HTML5toTouch}>
              <App/>
          </DndProvider>
      </WebSocketProvider>
  </Provider>,
  document.getElementById('root')
);
