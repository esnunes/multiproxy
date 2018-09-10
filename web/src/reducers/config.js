import { CONFIG_SUCCEEDED } from '../actions';

const initialState = {};

export default function(state = initialState, action) {
  switch (action.type) {
    case CONFIG_SUCCEEDED: {
      return action.payload;
    }
    default:
      return state;
  }
}
