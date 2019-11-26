# envdecode



## How to load env from file

1. direnv

https://direnv.net

https://github.com/direnv/direnv

```
[[ $(which direnv) ]] && eval "$(direnv hook bash)" || echo "Install direnv"

echo ${FOO-nope}

echo export FOO=foo > .envrc
.envrc is not allowed

direnv allow .

echo ${FOO-nope}


```



## Load up .env
```
env - sh -c env
PWD=/Users/ivostoyanov/src
SHLVL=1
_=/usr/bin/env
```

```
set -o allexport # or set -o a
[[ -f .env ]] && source .env
set +o allexport
```

