# atlasgo

Context: https://github.com/ariga/atlas/issues/2696

Run the following:

```sh
atlas migrate --env local diff
```

We get this output:

```
Error: data.hcl_schema.app: decoding body: atlas.hcl:2,11-18: Call to unknown function; There is no function named "fileset"., and 1 other diagnostic(s)
```

This is with version:

```sh
> atlas version
atlas version v0.27.1-4af8d72-canary
https://github.com/ariga/atlas/releases/latest
```
