# K8S YAML Example

From terminal, go to `examples/k8s-yaml` directory:

```bash
cd examples/k8s-yaml
```

## Generate Deployment & Config from template file

Run `subenv` againts `deployment.tmpl.yml` using `test.env` file as the environment file source:

```bash
ENV=development subenv -e test.env deployment.tmpl.yml > deployment.result.yml
```

## Generate Secret from template file

`Kubernetes` secret value is encoded using `base64`. `subenv` have `encoder` (`-c`) option for us to use, to generate
secret simply put `-c base64` argument.

```bash
subenv -e test.env -c base64 secret.tmpl.yml > deployment.result.yml
```