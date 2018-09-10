import { CONFIG_SUCCEEDED } from '../actions';

const initialState = [];

export default function(state = initialState, action) {
  switch (action.type) {
    case CONFIG_SUCCEEDED:
      return action.payload.apps.map(a => {
        return {
          id: a.id,
          name: a.name,
          description: a.description,
          addr: a.addr,
        };
      });
      break;
    default:
      return state;
  }
}
