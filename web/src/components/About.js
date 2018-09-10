import React from 'react';
import { EmptyState, Layout } from '@shopify/polaris';

import image from '../assets/404.svg';

const About = props => (
  <Layout sectioned>
    <EmptyState
      heading="Multiple environments, single setup"
      action={{
        content: 'Learn more',
        url: 'https://github.com/esnunes/multiproxy',
        external: true,
      }}
      image={image}
    >
      <p>
        Access different environments using a single application configuration.
      </p>
    </EmptyState>
  </Layout>
);

export default About;
