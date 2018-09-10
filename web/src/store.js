import { createStore, applyMiddleware } from 'redux';
import createSagaMiddleware from 'redux-saga';

import reducers from './reducers';
import sagas from './sagas';

export default function(services, state = {}) {
  const sagaMiddleware = createSagaMiddleware({
    context: { services },
  });

  const store = createStore(reducers, state, applyMiddleware(sagaMiddleware));

  sagaMiddleware.run(sagas);

  if (module.hot) {
    module.hot.accept('./reducers', () => {
      const next = require('./reducers').default;
      store.replaceReducer(next);
    });
  }

  return store;
}
