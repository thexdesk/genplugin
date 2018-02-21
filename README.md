genplugin
---

`genplugin` is a command-line tool to generate Docker plugins for testing.

## Usage

An example using `hinshun/genplugin` to generate a scratch plugin to push and
pull from a private registry. Plugins expect to connect to a unix socket to
figure out what privileges to grant the plugin, but since it does nothing it
is expected to hang for a bit, don't worry.

```sh
$ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION=1.36 hinshun/genplugin 172.17.0.1/admin/plugin
$ docker plugin push 172.17.0.1/admin/plugin
The push refers to repository [172.17.0.1/admin/plugin]
38370c72c6aa: Pushed
latest: digest: sha256:877688936bb2b953f0f3bcc0399589ca045ad8055574837ac4af9b7fcd4a5d97 size: 514
$ docker plugin rm 172.17.0.1/admin/plugin
172.17.0.1/admin/plugin
$ docker plugin install 172.17.0.1/admin/plugin
latest: Pulling from admin/plugin
fa3cce888389: Download complete
Digest: sha256:877688936bb2b953f0f3bcc0399589ca045ad8055574837ac4af9b7fcd4a5d97
Status: Downloaded newer image for 172.17.0.1/admin/plugin:latest
Error response from daemon: dial unix /run/docker/plugins/0df9e568dae0d613ae104a33cf1b8922d6993ad4e5256d435c6051cb3daf9d70: connect: connection refused
```
