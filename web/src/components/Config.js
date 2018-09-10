import React from 'react';
import { connect } from 'react-redux';

import { changePage } from '../actions';
import ConfigComponent from './ConfigComponent';

const mapStateToProps = state => ({
  config: state.config,
});

const mapDispatchToProps = dispatch => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ConfigComponent);
