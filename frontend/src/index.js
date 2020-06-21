import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import store from "./redux/store";
import { Provider } from 'react-redux'
import WebSocketProvider from './websocket/websocket'

ReactDOM.render(
  <Provider store={store}>
      <WebSocketProvider>
        <App/>
      </WebSocketProvider>
  </Provider>,
  document.getElementById('root')
);
