# scrub - maintain copyright headers in source files

- scrub check enumerates given source tree and returns count of files without headers
- scrub run prepends configured header to files w/o headers

scrub help

```
Usage:
  scrub [flags]
  scrub [command]

Available Commands:
  check       check
  config      Config
  help        Help
  run         Run scrub to check/add copyright text to source code
  version     Version

Flags:
      --color string              Use color when printing; can be 'always', 'auto', or 'never' (default "auto")
  -j, --concurrency int           Concurrency (default NumCPU) (default 12)
  -c, --config PATH               Read config from file path PATH (default "config.yaml")
      --cpu-profile-path string   Path to CPU profile output file
  -h, --help                      help for scrub
      --mem-profile-path string   Path to memory profile output file
      --no-config                 Don't read config
      --trace-path string         Path to trace output file
  -v, --verbose                   verbose output
      --version                   Print version
```

Configuration is stored in config file config.yaml and expected by default in the current directory

It can be specified via --config flag

```
################################
# scrub configuration parameters
################################

# input configuration options
input:
  root: test/testdata
  # directory with header file - if missing root will be used
  header-dir:
  # header file name
  header: header.txt

  # skip directories (regex, . is auto-escaped)
  # no need to specify: .git, vendor, testdata, third-party
  skip-dirs:
    - "subdir"

  # skip files (regex, * is auto-escaped)
  skip-files:
    - "*_gen.go"

  # patterns to check in the beginning of file
  # only 2X size of the header is searched
  # strings are converted to lowercase - so the patterns below are not case-sensitive
  # .* is regex for 0 or more characters
  skip:
    - ".*Brain Corp.*"
    - ".*copyright.*"

# output configuration options
output:


```
