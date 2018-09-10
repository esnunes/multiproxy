import { combineReducers } from 'redux';
import config from './config';
import page from './page';
import apps from './apps';
import envs from './envs';

export default combineReducers({
  config,
  page,
  apps,
  envs,
});
