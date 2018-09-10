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
} from '@shopify/polaris';

const AppItem = props => {
  const { app, onAppSelected } = props;

  const media = (
    <Avatar
      customer={true}
      size="medium"
      initials={app.name.substring(0, 2).toUpperCase()}
      name={app.name}
    />
  );

  let envMarkup = null;
  switch (app.env.status) {
    case 'loading':
      envMarkup = <Spinner color="inkLightest" size="small" />;
      break;
    case 'loaded':
      envMarkup = app.env.selected ? (
        <Badge status="info">{app.env.selected}</Badge>
      ) : (
        <Badge status="default">Not Definied</Badge>
      );
      break;
    case 'failed':
      envMarkup = <Badge status="warning">Failed</Badge>;
      break;
    default:
      envMarkup = <Badge status="attention">Unknown</Badge>;
  }

  return (
    <ResourceList.Item
      id={app.id}
      onClick={() => onAppSelected(app.id)}
      media={media}
    >
      <Stack alignment="center" wrap={false}>
        <Stack.Item fill>
          <div>
            <h3>
              <TextStyle variation="strong">{app.name}</TextStyle>
            </h3>
            <div>{app.description}</div>
          </div>
        </Stack.Item>
        <Stack.Item>
          <div>{envMarkup}</div>
        </Stack.Item>
      </Stack>
    </ResourceList.Item>
  );
};

const AppList = props => {
  const { apps, onAppSelected } = props;

  return (
    <Layout>
      <Layout.AnnotatedSection
        title="Applications"
        description="List of applications available."
      >
        <Card>
          <ResourceList
            items={apps}
            renderItem={app => {
              return <AppItem app={app} onAppSelected={onAppSelected} />;
            }}
          />
        </Card>
      </Layout.AnnotatedSection>
    </Layout>
  );
};

export default AppList;
