# Templating

Currently the templates are helm charts. They can be rendered using `helm template` and also `pkg.go.dev/text/template`.

To render a template you can use helm, but you can also use our bootstrapper cli.

To use helm first create a values.yaml:

```yaml
# values.yaml
openmcpOperator:
    image: ghcr.io/openmcp-project/images/openmcp-operator
    tag: v0.9.1

platformClusterKubeconfigSecretName: "platform-kubeconfig"
```

```bash
> helm template --output-dir output --values values.yaml github.com/openmcp-project/gitops-templates/v0.0.1/openmcp
```

## Parameters

### FluxCD

resources files:

```yaml
git:
    repoUrl: "" # The url to the github gitops repository
    mainBranch: "" # The main branch of the gitops repository (most of the time, set it to 'main')
```

When rendering the `overlays` files, the following values are used:

```yaml
targetPath: "" # The path were the fluxcd resources are lying relative to the overlays
gitRepoEnvBranch: "" # The branch for this environment to look at
envPathFluxSystem: "" # The path were the env overlays are located at from the root of the git repo
```

### OpenMCP

```yaml
openMCPOperator:
    image: "" # the image of the openmcp operator to use
    tag: "" # the tag of the image of the openmcp operator you want to use for deployment

onboardingClusterKubeconfigSecretName: "" # the secret name for the onboarding cluster; must be located in the openmcp-system namespace
platformClusterKubeconfigSecretName: "" # the secret name for the platform cluster; must be located in the openmcp-system namespace
```
