#!/usr/bin/env bash

#export PS1=${ret_status} %{$fg[cyan]%}%c%{$reset_color%} $(git_prompt_info)
#set -e

#echo 🔥 ✋ 🛑  💣
### echo v.2.3.19.0
#echo 🔥 ✋ 🛑  💣


export PS1="\T \W$ "
export HISTCONTROL=erasedups
export HISTSIZE=10000
export HISTFILESIZE=10000
# dont save commands starting with space 
export HISTIGNORE="[ \t]*:pwd:ls:ll:h:a:rm"
#export HISTIGNORE="rm *:h:a"

#MYIP=$(ifconfig | grep 'inet addr:'| grep -v '127.0.0.1' | tail -1 | cut -d: -f2 | awk '{ print $1}')
#export PS1=$MYIP" \W $"

#echo 🔥

alias a='alias'
alias e='echo'
alias h='history'
alias s='source'
alias c=clear
alias t=tree
alias gr=egrep
alias del='rm'
#alias subl='/Applications/Sublime\ Text.app/Contents/SharedSupport/Bin/subl'
#alias m=multipass
#alias ep='subl ~/.profile; source ~/.profile'
#alias ep='subl ~/.bash_profile; source ~/.bash_profile'
#alias erc='subl ~/.bashrc; source ~/.bashrc'
#alias ezrc='subl ~/.zshrc; source ~/.zshrc'

# Easier navigation: .., ..., ...., ....., ~ and -
alias ..="cd .."
alias ...="cd ../.."
alias ....="cd ../../.."
alias .....="cd ../../../.."
alias ~="cd ~" # `cd` is probably faster to type though
alias -- -="cd -"
alias cex='chmod +x'

# Shortcuts
alias dl="cd ~/Downloads"
alias dt="cd ~/Desktop"
#alias p="cd ~/projects"

# for locked files on mac
a zap='sudo chflags -R noschg,nohidden'

# Detect which `ls` flavor is in use
if ls --color > /dev/null 2>&1; then # GNU `ls`
	# shellcheck disable=SC2034
	colorflag="--color"
	export LS_COLORS='no=00:fi=00:di=01;31:ln=01;36:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.gz=01;31:*.bz2=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.avi=01;35:*.fli=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.ogg=01;35:*.mp3=01;35:*.wav=01;35:'
else # macOS `ls`
	# shellcheck disable=SC2034
	colorflag="-G"
	export LSCOLORS='BxBxhxDxfxhxhxhxhxcxcx'
fi

####################
###  🔥 LS 🔥  ###
####################

# List all files colorized in long format
alias l='ls -lF $colorflag'

# List all files colorized in long format, excluding . and ..
alias lsall='ls -lAF $colorflag'

# List only directories
alias lsdir='ls -lF $colorflag | grep --color=never "^d" '

# Always use color output for `ls`
alias ls='command ls $colorflag'

alias ls1='ls -F1'
alias lst='ls -FLlhtr'
alias ll='ls -Floghtr'
alias lat='ls -FLalhtr'

####################
###  🔥 GIT 🔥  ###
####################

alias clone='git clone'
alias amend='git commit --amend -m'
alias orphan='git checkout --orphan'

alias gita='git add -A '
alias gitb='git branch '
# delete remote branch
alias gitdrb='git push origin --delete'

alias gitc='git commit -a -m '
alias gitconf='git config --list --show-origin'
alias gitck='git checkout'
alias gitp='git push; git push --tags'
alias gitl='git pull'
alias gitls='git ls-files'
alias gitmr="git merge --strategy-option theirs"
alias gitml="git merge --strategy-option ours"
alias gits='git status'
# alias gitr='git rebase -i'
# alias gitrc='git rebase --continue'
# alias gitra='git rebase --abort'
# alias gitre='git rebase --edit-todo'
# squash last N commits before push e.g. gitsq HEAD~5
alias gitsq='echo "to squash last N commits - append HEAD~N" && git reset --soft'
alias gitresetDEVELOP='git reset --hard origin/develop'

alias wip='git commit -a -m wip'
alias wipp='git commit -a -m wip && git push'

#alias gitclean='git clean -fxd'
# alias gitckr='git checkout --recurse-submodules --remote'
# alias gitcl='git clone --recurse-submodules'
# alias gitlr='git submodule update --recursive --remote'
# alias gitsub='git submodule update --init --recursive'
# alias gitf='git fetch --recurse-submodules'
# alias gitpull='for d in */ ; do  pushd $d;    git pull; popd; done'
#alias gitsz='git count-objects -v'
# alias gitmu='git submodule update'
#alias gitpull='for d in */ ; do  pushd $d;    git pull; popd; done'

alias gitlog='git log --graph --decorate --oneline'

alias gitlog='git log --graph --full-history --all --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'

#alias gitlog='git log --graph --decorate --oneline'
#alias gitlog='git log --graph --full-history --all --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'
# alias gitlog='git log --graph --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'
# alias gitlogfull='git log --graph --full-history --all --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'
# alias gitd1='git diff HEAD^ HEAD'
# alias gitd2='git diff HEAD^^ HEAD'
# alias gitd3='git diff HEAD^^^ HEAD'
# alias gitl1w='git log --oneline --since=1.weeks'
# alias gitl1m='git log --oneline --since=1.months'
# alias gitl3m='git log --oneline --since=3.months'
# alias gitll='git log --oneline HEAD^..HEAD'
# alias gitll2='git log --oneline HEAD^^..HEAD'

# PATCH
#  git diff 79937af   29deb34  -- internal/generator/generator.go /internal/generator/ops.go
# 
# git push origin master --force-with-lease
# alias gitlog='git log --graph --oneline'
# alias gitlogp='git log --pretty="%h - %an, %ar : %s"'
# alias gitbis='git bisect start'
# alias gitbisr='git bisect reset'
# alias gitbisb='git bisect bad'
# alias gitbisg='git bisect good'
# alias gitsh='git show HEAD'
# alias gitref='git for-each-ref'

#git config --global core.editor "subl -n -w"

# Always enable colored `egrep` output
# Note: `GREP_OPTIONS="--color=auto"` is deprecated, hence the alias usage.
#alias grep='grep --color=auto'

# Enable aliases to be sudo’ed
alias sudo='sudo '

# always use unidiff
alias diff='diff -u'
alias k9='kill -9'

# alias v='vagrant'
# alias vs='vagrant status'

alias pssh='ps aux|grep ssh'
alias gign='vi .gitignore'

export PATH=/usr/local/bin:$HOME/tools:$PATH

alias sba='source ~/.bash_aliases'
alias pu='lsof -i '

# remove apple quarantine extended attribs
#alias jailbreak='sudo xattr -r -d com.apple.quarantine'
#alias wait='wait-for-it.sh'

alias npmi='npm install'

alias d=docker

alias dclr='docker rm $(docker ps -a -f status=exited -q)'
alias dim='docker images'
alias dk='docker kill'
alias dps='docker ps -a'
alias prune='docker system prune -f'
alias dl='docker logs'
alias dk='docker kill'
alias dimd='docker image rm -f'

# alias d_c_clean='printf "\n>>> Deleting stopped containers\n\n" && docker rm $(docker ps -alias -q)'
# alias d_c_kill='docker kill $(docker ps -q)'
# alias d_clean='dockercleanc || true && dockercleani'h
# alias d_i_clean='printf "\n>>> Deleting untagged images\n\n" && docker rmi $(docker images -q -f dangling=true)'
# alias dc=docker-compose
# alias ddel='docker rmi -f '
# alias dex=ssh_docker
# alias dh='docker history '
# alias di='docker inspect'
# alias dim='docker images|less'
# alias dk='docker kill'
# alias dl='docker logs'
# alias dr='docker run --rm -p8000:8000 -p8001:8001 -p8443:8443 -p8444:8444 '

#export GOPATH=$HOME/go
#export GO111MODULE=on

#export PATH=$GOPATH/bin:$PATH:/usr/local/go/bin
#alias nomod='export GO111MODULE=off'
#alias mod='export GO111MODULE=on'


# URL-encode strings
alias urlencode='python -c "import sys, urllib as ul; print ul.quote_plus(sys.argv[1]);"'
alias p='python'
alias gor="go run -gcflags='-m -l'  ."

# alias got='go test ./...'
# alias gocc='go clean -cache'
# alias gb='go build -ldflags "-s -w" -o /tmp/t; ls -alh /tmp/t'
# alias gw=goweight
# alias gmw='go mod why'
# alias armgo='CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -v -alias -tags netgo -ldflags "-w -extldflags '\''-static'\'' " '
# alias xgo='CGO_ENABLED=1 CC=aarch64-buildroot-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build '
# alias skms='cd ~/.ssh; cp id_rsa_t.pub id_rsa.pub; cp id_rsa_t id_rsa; '
# alias sksd='cd ~/.ssh; cp id_rsa_sd.pub id_rsa.pub; cp id_rsa_sd id_rsa; '

alias ve=virtualenv
alias eh='sudo vi /etc/hosts'
alias prot='chmod 0400'

function MD() {
  echo creating dir "$1"
  mkdir -p "$1"
  cd "$1" || return
  pwd
}

function DEL_T() {
  echo geleting git tag "$1"
  git tag -d "$1"
  git push --delete origin "$1"
  git tag
}

function LEN() {
  V=$1
  echo length of "$V": ${#V}
}

# message=`git log --format=%B origin..HEAD | sort | uniq | grep -v '^$'`
# git reset --soft origin
# git commit -m "$message"
# git reset --soft HEAD~3 && 
# git commit --edit -m"$(git log --format=%B --reverse HEAD..HEAD@{1})"
#[[   -f $CF ]] && source $CF
#kubectl run nginx --image=nginx
#export PATH=$HOME/Library/Python/2.7/bin:$PATH
#export NANO=10.0.1.187
export NANO=10.0.1.171
a sh-nano="ssh ivo@$NANO"

export PI=10.0.1.188
a sh-pi="ssh pi@$PI"

# spy on redis a redis-proxy='socat -v tcp-listen:16379 tcp:127.0.0.1:6379'

# a start-pg='pg_ctl -D /usr/local/var/postgres start && brew services start postgresql'
# a start-pg='brew services start postgresql'
# a stop-pg='brew services stop postgresql'
#a stat-pg='pg_ctl -D /usr/local/var/postgres  status'
# pg_ctl -D /usr/local/var/postgres start
# a start-redis='brew services start redis'
# a stop-redis='brew services stop redis'

# CREATE ROLE ivostoyanov WITH SUPERUSER CREATEDB CREATEROLE a;
# CREATE DATABASE XXX;

##########################
###  🔥 FUNCTIONS 🔥  ###
#########################

# OS detection
function is_osx() {
  [[ "$OSTYPE" =~ ^darwin ]] || return 1
}

function is_ubuntu() {
  [[ "$(cat /etc/issue 2> /dev/null)" =~ Ubuntu ]] || return 1
}

function is_ubuntu_desktop() {
  dpkg -l ubuntu-desktop >/dev/null 2>&1 || return 1
}

function get_os() {
  for os in osx ubuntu ubuntu_desktop; do
    is_$os; [[ $? == "${1:-0}" ]] && echo $os
  done
}

# git branch
function git_br() {
	echo $( (git branch 2> /dev/null) | grep \* | cut -c3-)
}
# for iterm badge
function iterm2_print_user_vars() {
  iterm2_set_user_var gitBranch "$(git_br)"
}

function login() {

}

function bun_check() {
	openssl crl2pkcs7 -nocrl -certfile $1 | openssl pkcs7 -print_certs -text -noout
}

# git rebase remote after push
function rebase() {
  local B=$(git_br)
  local N=$1
  #echo git rebase -i origin/$B~$N $B
  git rebase -i origin/$B~$N $B
}

# git force push
function gitfp() {
  local B
  B=$(git_br)
  git push origin +"$B"
}

alias inst='sudo apt install'
alias upd='sudo apt update'
export EDITOR=subl
alias eba='subl ~/.bash_aliases; source ~/.bash_aliases'

## aliases depending on OS
### MAC
if [[ "$OSTYPE" =~ ^darwin ]]; then 
	export EDITOR=subl
	alias eba='subl ~/.bash_aliases; source ~/.bash_aliases'
	alias inst='brew install'
	alias bd='brew doctor'

	function iterm2_print_user_vars() {
	   iterm2_set_user_var gitBranch "$( (git branch 2> /dev/null) | grep '\*' | cut -c3-)"
	}

  # Recursively delete `.DS_Store` files
  alias cleanup="find . -type f -name '*.DS_Store' -ls -delete"

  # Empty the Trash on all mounted volumes and the main HDD.
  # Also, clear Apple’s System Logs to improve shell startup speed.
  # Finally, clear download history from quarantine. https://mths.be/bum
  alias emptytrash="sudo rm -rfv /Volumes/*/.Trashes; sudo rm -rfv ~/.Trash; sudo rm -rfv /private/var/log/asl/*.asl; sqlite3 ~/Library/Preferences/com.apple.LaunchServices.QuarantineEventsV* 'delete from LSQuarantineEvent'"

  # Show/hide hidden files in Finder
  alias show="defaults write com.apple.finder AppleShowAllFiles -bool true && killall Finder"
  alias hide="defaults write com.apple.finder AppleShowAllFiles -bool false && killall Finder"

  # Hide/show all desktop icons (useful when presenting)
  alias hidedesktop="defaults write com.apple.finder CreateDesktop -bool false && killall Finder"
  alias showdesktop="defaults write com.apple.finder CreateDesktop -bool true && killall Finder"
	#x=$(brew --prefix)/etc/bash_completion
	# shellcheck disable=SC1090
	#[[ -f "$x" ]] && source "$x"
fi
### END MAC


#curl -i -X POST -H "$A" $URL  -d '{"value": "'$L'"}'
#curl -i -X POST -H "$A" $URL2  -d '{"value": "'$L'"}'

# colorizer
[[ -s "/usr/local/etc/grc.bashrc" ]] && source /usr/local/etc/grc.bashrc


# if you have ssh problems
alias sshv='ssh -vvv -o LogLevel=DEBUG3'

#a mb='make extraclean && make api && make generate'
alias mode="stat -f '%A %a %N' "

export CFH=http://localhost:8888

export PATH=$PATH:~/tools/platform-tools/:~/Library/Python/3.7/bin

###############
# Coral board #
###############
# ping bored-kid.local
# 10.0.1.194
export CORAL=10.0.1.194
a sh-coral='mdt shell'

# colorizer
#[[ -s "/usr/local/etc/grc.bashrc" ]] && source /usr/local/etc/grc.bashrc

# if you have ssh problems
#alias sshv='ssh -vvv -o LogLevel=DEBUG3'
#export region=us-west-2

#alias mk=minikube
#alias emk='eval $(minikube docker-env)'
#alias mkd='minikube dashboard'
#alias sk=skaffold
#minikube start --kubernetes-version v1.16.0 --vm-driver=none/virtualbox

#a ac='asciinema'
# exit / ctrl/D
# asciinema auth
# asciinema play $1
#asciinema rec -t "My git tutorial"

# alias ecr-login='$(aws ecr get-login --region ${region} --no-include-email)'
alias tf=terraform
alias N='date +''%s'''
alias S-H='N; http $S/readyz; N'

#############################
###  🔥🔥🔥 XXX 🔥🔥🔥   ###
#############################

# a DROP='psql -c "DROP DATABASE IF EXISTS XXX;"'
# a CRDR='psql -c "CREATE DATABASE XXX;"'

alias mode="stat -f '%A %a %N' "
#alias ecr-login='$(aws ecr get-login --region ${region} --no-include-email)'
#eval $(minikube docker-env)

a mb='make build'
a mh='make help'
a mt='make test'
# integration tests
a mti='make testi'

a mr='make run'
a mf='make refresh'

a da='direnv allow'

####################
###  🔥 K8S 🔥  ###
####################
# https://kubernetes.io/docs/reference/kubectl/cheatsheet/#scaling-resources
# alias k='kubectl'
# alias kc='kubectl create'
# alias kd='kubectl describe'
# alias kg='kubectl get'
# alias kl='kubectl logs'
# alias kp='kube-prompt'
# alias ki='kubectl cluster-info'

# alias kpf='kubectl port-forward '

# alias kcg='kubectl config get-contexts'
# alias kcu='kubectl config use-context'
# alias kcud='kubectl config use-context docker-desktop'

# alias kgn='kubectl get nodes -o wide'
# alias kdn='kubectl describe  node'
# alias kgns='kubectl get namespaces'
# alias kgp='kubectl get pods '
# alias kgpw='kubectl get pods -o wide'
# alias kdp='kubectl describe pod'
# alias kd-='kubectl delete'
# alias kd-p='kubectl delete pod'
# alias kd-s='kubectl delete service'
# alias kd-d='kubectl delete deployment'

# alias kgpa='kubectl get pods --all-namespaces -o wide'
# alias kgs='kubectl get services -o wide'

# alias kgd='kubectl get deployments -o wide'
# alias kdd='kubectl describe deployment'

# alias kaf='kubectl apply -f'
# alias kdf='kubectl delete -f'
# alias kex='kubectl exec -it $1'

# alias ksd1='kubectl scale --replicas 1 deployment'
# alias ksd2='kubectl scale --replicas 2 deployment'
# alias ksd3='kubectl scale --replicas 3 deployment'

#alias kgpa='kubectl get pod -l app=$1 -o jsonpath='{.items[0].metadata.name}''

# function kg_podname() {
#     kubectl get pod -l app=$1 -o jsonpath='{.items[0].metadata.name}'
# }

# function kg_podip() {
#     kubectl get pod -l app=$1 -o jsonpath='{.items[0].status.podIP}'
# }

# alias pfv='kubectl port-forward vault-0 8200:8200'
# alias sh-uu='kubectl exec -it ubuntu-util -- /bin/bash'

#kubectl exec -it $(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}') -c ratings -- curl productpage:9080/productpage | grep -o "<title>.*</title>"

#eval "$(direnv hook bash)"

#export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
#export PATH=$PATH:/usr/local/Cellar/minikube/1.5.2/bin
#export PATH="$PATH:/Users/ivostoyanov/scripts/istio-1.3.4/bin"
# comment unless using kind
#export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
#export PATH="$PATH:/$HOME/istio/istio-1.4.0/bin"
#a i='istioctl'
#a dex='docker exec -it'
#a mks='minikube start  --kubernetes-version=1.15.6 --memory='10000mb' --cpus=4'
# docker run -d --name=logtest alpine /bin/sh -c "while true; do sleep 5; echo working...; done"

# when using minikube
#### eval $(minikube docker-env) && echo "USING MINIKUBE"
#echo 🔥🔥🔥🔥

##############
