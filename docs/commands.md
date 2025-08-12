```
ocm add componentversions --copy-resources --force  --create --file .out --settings settings.yaml component-constructor.yaml

ocm transfer ctf --copy-resources --enforce -f .out ghcr.io/n3rdc4ptn/ocm

ocm download resources --downloader ocm/dirtree --repo OCIRegistry::ghcr.io/n3rdc4ptn/ocm github.com/openmcp-project/gitops-templates:0.0.1 openmcp openmcpdir
```
