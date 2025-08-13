# Usage

The gitops templates are built into one ocm component which is called `gitops-templates`.

You can fetch it from the ocm repo: `ghcr.io/openmcp-project/components`.
The component name is: `github.com/openmcp-project/gitops-templates`.

To get a list of all resources, run the ocm cli:

```bash
> ocm get resources --repo OCIRegistry::ghcr.io/openmcp-project/components github.com/openmcp-project/gitops-templates:0.0.1

NAME             VERSION IDENTITY TYPE       RELATION
fluxcd           v0.0.1           fileSystem local
gitops-templates v0.0.1           blob       external
openmcp          v0.0.1           fileSystem local
openmcp-operator v0.0.1           ociImage   local
```

The templates itself are fileSystem types. You can fetch these using:

```bash
# Just openmcp
> ocm download resources --downloader ocm/dirtree --repo OCIRegistry::ghcr.io/openmcp-project/components github.com/openmcp-project/gitops-templates:v0.0.1 openmcp

# Openmcp and fluxcd
> ocm download resources --downloader ocm/dirtree --repo OCIRegistry::ghcr.io/openmcp-project/components github.com/openmcp-project/gitops-templates:v0.0.1 openmcp fluxcd
```

If you specify multiple templates, they will be put into a directory structure in the directory you are currently in:

```txt
github.com/
  openmcp-project/
    gitops-templates/
      v0.0.1/
        openmcp/
        fluxcd/
```

## Rendering a template

For detailed instructions on rendering templates, see the [Templating guide](./Templating.md).
