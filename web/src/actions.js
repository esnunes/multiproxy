export const CONFIG_SUCCEEDED = 'CONFIG_SUCCEEDED';
export const loadConfig = config => ({
  type: CONFIG_SUCCEEDED,
  payload: config,
});

export const CHANGE_PAGE = 'CHANGE_PAGE';
export const changePage = (page, params = {}) => ({
  type: CHANGE_PAGE,
  payload: { page, params },
});

export const LOAD_SELECTED_ENV = 'LOAD_SELECTED_ENV';
export const loadSelectedEnv = appId => ({
  type: LOAD_SELECTED_ENV,
  payload: { appId },
});
export const SELECTED_ENV_REQUESTED = 'SELECTED_ENV_REQUESTED';
export const SELECTED_ENV_SUCCEEDED = 'SELECTED_ENV_SUCCEEDED';
export const SELECTED_ENV_FAILED = 'SELECTED_ENV_FAILED';

export const SELECT_ENV = 'SELECT_ENV';
export const selectEnv = (appId, key) => ({
  type: SELECT_ENV,
  payload: { appId, key },
});
export const SELECT_ENV_REQUESTED = 'SELECT_ENV_REQUESTED';
export const SELECT_ENV_SUCCEEDED = 'SELECT_ENV_SUCCEEDED';
export const SELECT_ENV_FAILED = 'SELECT_ENV_FAILED';
