import React from 'react';
import {
  Layout,
  Card,
  ResourceList,
  Avatar,
  TextStyle,
  Stack,
  Badge,
  Spinner,
  Banner,
  Link,
  DisplayText,
  Button,
} from '@shopify/polaris';

const EnvItem = props => {
  const { env, selected, status, onEnvSelected } = props;

  const media = (
    <Avatar
      customer={true}
      size="medium"
      initials={env.name.substring(0, 2).toUpperCase()}
      name={env.name}
    />
  );

  let selectedMarkup = null;
  switch (status) {
    case 'loading':
      selectedMarkup = <Spinner color="inkLightest" size="small" />;
      break;
    case 'loaded':
      selectedMarkup = selected ? <Badge status="info">Selected</Badge> : null;
      break;
    case 'failed':
      selectedMarkup = <Badge status="warning">Failed</Badge>;
      break;
  }

  return (
    <ResourceList.Item
      id={env.key}
      onClick={() => onEnvSelected(env.key)}
      media={media}
    >
      <Stack alignment="center" wrap={false}>
        <Stack.Item fill>
          <div>
            <h3>
              <TextStyle variation="strong">{env.name}</TextStyle>
            </h3>
          </div>
        </Stack.Item>
        <Stack.Item>
          <div>{selectedMarkup}</div>
        </Stack.Item>
      </Stack>
    </ResourceList.Item>
  );
};

const EnvList = props => {
  const { app, envs, onEnvSelected } = props;

  return (
    <Layout>
      <Layout.AnnotatedSection title={app.name} description={app.description}>
        <Stack vertical distribution="fill">
          <Stack distribution="equalSpacing">
            <DisplayText size="medium">Environments</DisplayText>
            <Button primary external disabled={envs.selected === null} url={app.addr}>Open app</Button>
          </Stack>
          {envs.status === 'failed' && (
            <Banner
              title="Could not retrieve selected environment"
              status="warning"
            >
              <p>
                This app might not be properly configured. Please check if{' '}
                <Link url={app.addr} external>
                  {app.addr}
                </Link>{' '}
                is responding.
              </p>
            </Banner>
          )}
          <Card>
            <ResourceList
              items={envs.list}
              renderItem={env => {
                return (
                  <EnvItem
                    env={env}
                    selected={env.name === envs.selected}
                    status={envs.status}
                    onEnvSelected={key => onEnvSelected(app.id, key)}
                  />
                );
              }}
            />
          </Card>
        </Stack>
      </Layout.AnnotatedSection>
    </Layout>
  );
};

export default EnvList;
