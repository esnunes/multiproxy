import 'regenerator-runtime/runtime';
import { takeEvery, takeLatest } from 'redux-saga/effects';

import { LOAD_SELECTED_ENV, SELECT_ENV } from '../actions';

import loadSelectedEnv from './loadSelectedEnv';
import selectEnv from './selectEnv';

export default function* sagas() {
  yield takeEvery(LOAD_SELECTED_ENV, loadSelectedEnv);
  yield takeLatest(SELECT_ENV, selectEnv);
}
