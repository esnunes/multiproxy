import React from 'react';
import { connect } from 'react-redux';

import { changePage } from '../actions';
import AppComponent from './AppComponent';

const mapStateToProps = state => ({
  page: state.page.current,
});

const mapDispatchToProps = dispatch => ({
  onMenu: page => dispatch(changePage(page)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AppComponent);
