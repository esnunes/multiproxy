import { delay } from 'redux-saga';
import { call, put, getContext, select } from 'redux-saga/effects';

import {
  SELECT_ENV_REQUESTED,
  SELECT_ENV_SUCCEEDED,
  SELECT_ENV_FAILED,
} from '../actions';

const appUrl = (state, appId) => {
  const app = state.config.apps.find(a => a.id === appId);
  if (!app) throw new Error('invalid app');

  return app.addr;
};

export default function* selectEnv(action) {
  const { appId, key } = action.payload;

  try {
    const services = yield getContext('services');

    yield put({ type: SELECT_ENV_REQUESTED, payload: { appId } });

    const url = yield select(appUrl, appId);
    yield call(services.selectEnv, url, key);

    yield put({
      type: SELECT_ENV_SUCCEEDED,
      payload: { appId, selected: key },
    });
  } catch (error) {
    yield put({ type: SELECT_ENV_FAILED, payload: { appId, error } });
  }
}
