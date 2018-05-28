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
  "admin": "/_multiproxy",
  "cookie": "multiproxy",
  "upstreams": [
    { "key": "branch-a", "addr": "http://localhost:8001" },
    { "key": "branch-b", "addr": "http://localhost:8002" }
  ],
  "broadcast": ["/webhooks"]
}
```

