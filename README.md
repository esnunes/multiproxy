# multiproxy: Multicast HTTP Reverse Proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/esnunes/multiproxy)](https://goreportcard.com/report/github.com/esnunes/multiproxy)

## Tasks

- [ ] Cover unicast pkg with tests;
- [ ] Add logger to broadcast handler;
- [ ] Develop admin handler;
- [ ] Use buffers to broadcast body;
- [ ] Integrate cobra; 
- [ ] Add support to modify host binding;
- [ ] Modify app to access dynamic config;
- [ ] Add multicast configuration per addr (e.g. response status code, require
  all succeeded, ...);

## Configuration

```json
{
  "admin": "http://admin.localhost:8080/",
  "cookie": "multiproxy",
  "apps": [
    {
      "id": 1,
      "name": "My app 1",
      "description": "This is my fancy app 1",
      "addr": "http://app1.localhost:8080",
      "broadcast": ["webhooks/"],
      "envs": [
        { "name": "Environment 1", "key": "env1", "upstream": "http://localhost:8001" }
      ]
    }
  ]
}
```

