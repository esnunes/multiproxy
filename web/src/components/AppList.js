import React, { Component } from 'react';
import { connect } from 'react-redux';

import { changePage, loadSelectedEnv } from '../actions';
import AppListComponent from './AppListComponent';

export class AppList extends Component {
  componentDidMount() {
    const { apps, loadSelectedEnv } = this.props;

    apps
      .filter(a => a.env.status === 'unknown')
      .forEach(a => loadSelectedEnv(a.id));
  }

  render() {
    return <AppListComponent {...this.props} />;
  }
}

const appListWithSelectedEnv = state => {
  const envs = state.envs;

  const apps = state.apps.map(a => {
    const env = envs[a.id] || { status: 'unknown' };

    return Object.assign({}, a, {
      env: { selected: env.selected, status: env.status },
    });
  });

  return apps;
};

const mapStateToProps = state => ({
  apps: appListWithSelectedEnv(state),
});

const mapDispatchToProps = dispatch => ({
  loadSelectedEnv: appId => dispatch(loadSelectedEnv(appId)),
  onAppSelected: appId => dispatch(changePage('EnvList', { appId })),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AppList);
