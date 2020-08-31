## lb-proxy-mesh

### Linkerd Introduction

> Linkerd is a service mesh for Kubernetes. It makes running services easier and safer by giving you runtime debugging, observability, reliability, and security—all without requiring any changes to your code.

```sh
.
├── images-2.8.1.sh     // download images for linkerd in CHINA (2.8.1).
├── images.sh           // download images for linkerd in CHINA (2.7.1).
├── linkerd-2.8.1.yaml  // linkerd (2.8.1) k8s apply config file; get from `linkerd install > linkerd-2.8.1.yaml`.
├── linkerd.yaml        // linkerd (2.7.1) k8s apply config file; get from `linkerd install > linkerd.yaml`.
├── client.yaml         // client k8s apply config YAML file. 
└── server.yaml         // server k8s apply config YAML file.
```

for now, `2.8.1` is stable for `Linkerd`, so there is no need to use `2.7.1`.

### References

https://linkerd.io/2/getting-started/#step-1-install-the-cli