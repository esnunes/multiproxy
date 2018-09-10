import { CHANGE_PAGE } from '../actions';

const initialState = {
  current: 'AppList',
  params: {},
};

export default function(state = initialState, action) {
  switch (action.type) {
    case CHANGE_PAGE: {
      const { page, params } = action.payload;

      return Object.assign({}, state, { params, current: page });
    }
    default:
      return state;
  }
}
