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
# The information for the git repo
git:
    repoUrl: ""
    mainBranch: ""
    # Flux GitRepository spec.provider. "generic" (default) uses secretRef auth; "github" enables GitHub App auth.
    provider: "generic"

# Image pull secrets to be added to all deployments
imagePullSecrets: []
    # - name: my-registry-secret

# Image replacement variables
# You can either specify a prefix which is put in front of xxx-controller or directly specify the images
images:
    prefix: "ghcr.io/openmcp-project/fluxcd"
    sourceController:
        image: "ghcr.io/fluxcd/source-controller"
        tag: "latest" # optional
        digest: "" # optional
    notificationController:
        image: "ghcr.io/fluxcd/notification-controller"
    kustomizeController:
        image: "ghcr.io/fluxcd/kustomize-controller"
    helmController:
        image: "ghcr.io/fluxcd/helm-controller"
    imageReflectorController:
        image: "ghcr.io/fluxcd/image-reflector-controller"
    imageAutomationController:
        image: "ghcr.io/fluxcd/image-automation-controller"
```

When rendering the `overlays` files, the following values are used:

```yaml
# Path from the overlays folder to the resources folder of fluxcd (e.g. ../../../resources/fluxcd)
fluxCDResourcesPath: ""
# Path to the env fluxCD folder (e.g. envs/%ENV%/fluxcd)
fluxCDEnvPath: ""
# branch of the env (e.g. dev)
gitRepoEnvBranch: ""
```

### OpenMCP

```yaml
openMCPResourcesPath: "" # The path were the fluxcd resources are lying relative to the overlays (e.g. ../../../resources/openmcp)
openMCPOperator:
    image: "" # the image of the openmcp operator to use
    tag: "" # the tag of the image of the openmcp operator you want to use for deployment
```
