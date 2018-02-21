genplugin
---

`genplugin` is a command-line tool to generate Docker plugins for testing.

## Usage

```
$ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION=1.36 hinshun/genplugin some/plugin
$ docker plugin ls
ID                  NAME                 DESCRIPTION         ENABLED
1c9d60be44a4        some/plugin:latest                       false
```
