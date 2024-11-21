# atlasgo

Context: https://github.com/ariga/atlas/issues/3156

The only difference between the working and not_working dirs is the file name `a.hcl` vs `z.hcl`.

By "working" I mean that we get a good error message for the faulty "string" type instead of "text".

Run the following from each directory:

```sh
atlas migrate --env local diff
```