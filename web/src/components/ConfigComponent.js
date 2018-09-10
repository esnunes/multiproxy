import React from 'react';
import {
  Layout,
  Card,
  Stack,
  Link,
  ResourceList,
  TextStyle,
} from '@shopify/polaris';

const Config = ({ config }) => (
  <Layout>
    <Layout.AnnotatedSection
      title="Config"
      description="Multiproxy configuration."
    >
      <Stack vertical distribution="fill">
        <Card sectioned title="General">
          <dl>
            <dt>Admin URL</dt>
            <dd>
              <Link external url={config.admin}>
                {config.admin}
              </Link>
            </dd>

            <dt>Cookie name</dt>
            <dd>{config.cookie}</dd>
          </dl>
        </Card>

        <Card title="Applications">
          {config.apps.map(app => (
            <Card.Section key={app.id} title={app.name}>
              <dl>
                <dt>ID</dt>
                <dd>{app.id || '-'}</dd>

                <dt>Description</dt>
                <dd>{app.description || '-'}</dd>

                <dt>URL</dt>
                <dd>
                  <Link external url={app.addr}>
                    {app.addr || '-'}
                  </Link>
                </dd>

                {app.broadcast.length && <dt>Broadcast</dt>}
                {app.broadcast.map((b, i) => <dd key={i}>{b || '-'}</dd>)}
              </dl>

              <Card sectioned subdued>
                {app.envs.map(env => (
                  <Stack key={env.key} alignment="center" wrap={false}>
                    <Stack.Item fill>
                      <div>
                        <h3>
                          <TextStyle variation="strong">{env.name}</TextStyle>
                        </h3>
                        <dl>
                          <dt>Key</dt>
                          <dd>{env.key || '-'}</dd>

                          <dt>Name</dt>
                          <dd>{env.name || '-'}</dd>

                          <dt>Upstream</dt>
                          <dd>
                            <Link external url={env.upstream}>
                              {env.upstream}
                            </Link>
                          </dd>
                        </dl>
                      </div>
                    </Stack.Item>
                  </Stack>
                ))}
              </Card>
            </Card.Section>
          ))}
        </Card>
      </Stack>
    </Layout.AnnotatedSection>
  </Layout>
);

export default Config;
