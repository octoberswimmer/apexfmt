# apexfmt

Apexfmt formats
[Apex](https://developer.salesforce.com/docs/atlas.en-us.apexcode.meta/apexcode/apex_dev_guide.htm)
code.  It uses tabs for indentation.

Given a file, it writes the formatted code to standard output by default.  The
`--write`/`-w` flag can be used to overwrite the original file(s).  The
`--list`/`-l` flag can be used to list files with formatting different from
apexfmt's.


# Usage

## CLI
```
$ apexfmt -w sfdx/main/default/classes/*.cls sfdx/main/default/triggers/*.trigger
```

## Vim

apexfmt is included as a default formatter in [vim-autoformat](https://github.com/vim-autoformat/vim-autoformat/pull/394).

Use the following settings to display parse errors and disable the default formatting if `apexfmt` fails:

```
let g:autoformat_verbosemode=1
let g:autoformat_autoindent = 0
let g:autoformat_retab = 0
let g:autoformat_remove_trailing_spaces = 0
```

# Demo

Try out apexfmt in a browser at https://apexfmt.octoberswimmer.com/.

# Thanks

apexfmt is inspired by [gofmt](https://pkg.go.dev/cmd/gofmt).

The antlr4 grammar is based on the @nawforce's
[apex-parser](https://github.com/nawforce/apex-parser) grammar, which is
originally based on the grammer used by @neowit's
[tooling-force.com](https://github.com/neowit/tooling-force.com).

apexfmt in the browser forked from [go-fmt-wasm](https://github.com/junedev/go-fmt-wasm).
