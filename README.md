# apexfmt

Apexfmt formats
[Apex](https://developer.salesforce.com/docs/atlas.en-us.apexcode.meta/apexcode/apex_dev_guide.htm)
code.  It uses tabs for indentation.

Given a file, it writes the formatted code to standard output by default.  The
`--write`/`-w` flag can be used to overwrite the original file(s).  The
`--list`/`-l` flag can be used to list files with formatting different from
apexfmt's.

# Usage

```
$ apexfmt -w sfdx/main/default/classes/*.cls sfdx/main/default/triggers/*.trigger
```

# Thanks

apexfmt is inspired by [gofmt](https://pkg.go.dev/cmd/gofmt).

The antlr4 grammar is based on the @nawforce's
[apex-parser](https://github.com/nawforce/apex-parser) grammar, which is
originally based on the grammer used by @neowit's
[tooling-force.com](https://github.com/neowit/tooling-force.com).
