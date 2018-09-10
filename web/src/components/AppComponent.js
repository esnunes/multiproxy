import React from 'react';
import { AppProvider, Page, FooterHelp, Link } from '@shopify/polaris';

import AppList from './AppList';
import EnvList from './EnvList';
import Config from './Config';
import About from './About';

const pageToComponent = {
  AppList,
  EnvList,
  Config,
  About,
};

const InvalidPage = () => <div>Invalid page</div>;

const App = props => {
  const { page, onMenu } = props;
  const Content = pageToComponent[page] || InvalidPage;

  return (
    <AppProvider>
      <Page
        separator
        title="Multiproxy"
        secondaryActions={[
          { content: 'Applications', onAction: () => onMenu('AppList') },
          { content: 'Config', onAction: () => onMenu('Config') },
          { content: 'About', onAction: () => onMenu('About') },
        ]}
      >
        <Content />
        <FooterHelp>
          Learn more about{' '}
          <Link external url="https://github.com/esnunes/multiproxy">
            Multiproxy
          </Link>.
        </FooterHelp>
      </Page>
    </AppProvider>
  );
};

App.propTypes = {};

export default App;
