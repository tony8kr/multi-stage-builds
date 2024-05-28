# Multi-stage builds
In order to build an image using a specific Dockerfile you need to use **-f**
```
docker build -t http-app-go:multi-stage-alpine -f Dockerfile-multi-stage-alpine .
```
