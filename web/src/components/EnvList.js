import React from 'react';
import { connect } from 'react-redux';

import { selectEnv } from '../actions';
import EnvListComponent from './EnvListComponent';

const mapStateToProps = state => ({
  app: state.apps.find(a => a.id === state.page.params.appId),
  envs: state.envs[state.page.params.appId],
});

const mapDispatchToProps = dispatch => ({
  onEnvSelected: (appId, key) => dispatch(selectEnv(appId, key)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(EnvListComponent);
