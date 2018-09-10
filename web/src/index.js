import '@shopify/polaris/styles.css';
import './index.css';

import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import createStore from './store';
import * as services from './services';
import { loadConfig } from './actions';

import App from './components/App';

export function init(opts = {}) {
  const store = createStore(services);
  if (opts.debug) store.subscribe(() => console.log('state', store.getState()));

  store.dispatch(loadConfig(opts.config));

  ReactDOM.render(
    <Provider store={store}>
      <App />
    </Provider>,
    document.getElementById('root')
  );
}
