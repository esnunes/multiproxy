import { delay } from 'redux-saga';
import { call, put, getContext, select } from 'redux-saga/effects';

import {
  SELECTED_ENV_REQUESTED,
  SELECTED_ENV_SUCCEEDED,
  SELECTED_ENV_FAILED,
} from '../actions';

const envName = (state, appId, key) => {
  const app = state.config.apps.find(a => a.id === appId);
  if (!app) return null;

  const env = app.envs.find(e => e.key === key);
  if (!env) return null;

  return env.name;
};

const appUrl = (state, appId) => {
  const app = state.config.apps.find(a => a.id === appId);
  if (!app) throw new Error('invalid app');

  return app.addr;
};

export default function* loadSelectedEnv(action) {
  const { appId } = action.payload;

  try {
    const services = yield getContext('services');

    yield put({ type: SELECTED_ENV_REQUESTED, payload: { appId } });

    const url = yield select(appUrl, appId);
    const { key } = yield call(services.loadSelectedEnv, url);

    yield put({
      type: SELECTED_ENV_SUCCEEDED,
      payload: { appId, selected: key },
    });
  } catch (error) {
    yield put({ type: SELECTED_ENV_FAILED, payload: { appId, error } });
  }
}
