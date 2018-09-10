import {
  CONFIG_SUCCEEDED,
  SELECTED_ENV_REQUESTED,
  SELECTED_ENV_SUCCEEDED,
  SELECTED_ENV_FAILED,
  SELECT_ENV_REQUESTED,
  SELECT_ENV_SUCCEEDED,
  SELECT_ENV_FAILED,
} from '../actions';

const Status = {
  Unkonwn: 'unknown',
  Loading: 'loading',
  Loaded: 'loaded',
  Failed: 'failed',
};

const envName = (state, appId, key) => {
  const env = state[appId].list.find(e => e.key === key);
  if (!env) return null;

  return env.name;
};

const initialState = {};

export default function(state = initialState, action) {
  switch (action.type) {
    case CONFIG_SUCCEEDED: {
      const state = {};

      action.payload.apps.forEach(a => {
        state[a.id] = {
          list: a.envs.map(e => Object.assign({}, e)),
          selected: null,
          status: Status.Unkonwn,
        };
      });

      return state;
    }

    case SELECT_ENV_REQUESTED:
    case SELECTED_ENV_REQUESTED: {
      const appId = action.payload.appId;

      const newState = Object.assign({}, state);

      newState[appId] = Object.assign({}, newState[appId], {
        selected: null,
        status: Status.Loading,
      });

      return newState;
    }

    case SELECT_ENV_SUCCEEDED:
    case SELECTED_ENV_SUCCEEDED: {
      const { appId, selected } = action.payload;

      const newState = Object.assign({}, state);

      newState[appId] = Object.assign({}, newState[appId], {
        selected: envName(state, appId, selected) || null,
        status: Status.Loaded,
      });

      return newState;
    }

    case SELECT_ENV_FAILED:
    case SELECTED_ENV_FAILED: {
      const { appId, selected } = action.payload;

      const newState = Object.assign({}, state);

      newState[appId] = Object.assign({}, newState[appId], {
        status: Status.Failed,
      });

      return newState;
    }
    default:
      return state;
  }
}
