# Provider Zitadel

`provider-zitadel` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
[Zitadel](https://zitadel.com/) API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/vhdirk/provider-zitadel):

```sh
up ctp provider install xpkg.upbound.io/vhdirk/provider-zitadel:v0.2.0
```

Alternatively, you can use declarative installation:

```sh
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-zitadel
spec:
  package: xpkg.upbound.io/vhdirk/provider-zitadel:v0.2.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/vhdirk/crossplane-provider-zitadel).

## Developing

Run code-generation pipeline:

```sh
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```sh
make run
```

Build, push, and install:

```sh
make all
```

Build binary:

```sh
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/vhdirk/crossplane-provider-zitadel/issues).
