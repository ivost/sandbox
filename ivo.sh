
export PS1="\T \W$ "
export HISTCONTROL=erasedups
export HISTSIZE=10000
export HISTFILESIZE=10000
# dont save commands starting with space 
#export HISTIGNORE="[ \t]*:pwd:ls:ll:h:a:rm"
export HISTIGNORE="rm *:h:a"

#MYIP=$(ifconfig | grep 'inet addr:'| grep -v '127.0.0.1' | tail -1 | cut -d: -f2 | awk '{ print $1}')
#export PS1=$MYIP" \W $"

alias a='alias'
alias e='echo'
alias h='history'
alias s='source'
alias c=clear
alias t=tree
alias gr=egrep
alias del='rm'
alias subl=vi

alias inst='sudo apt install'
alias upd='sudo apt update'
export EDITOR=vim
alias eba='vim ~/ivo.sh; source ~/ivo.sh'
alias sba='source ~/ivo.sh'

# Easier navigation: .., ..., ...., ....., ~ and -
alias ..="cd .."
alias ...="cd ../.."
alias ....="cd ../../.."
alias .....="cd ../../../.."
alias ~="cd ~" # `cd` is probably faster to type though
alias -- -="cd -"

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

alias wip='git commit -a -m wip && git push'
alias gitclean='git clean -fxd'

alias gitc='git commit -a -m '
alias gitconf='git config --list --show-origin'
alias gitck='git checkout'
alias gitp='git push; git push --tags'
alias gitl='git pull'
alias gitls='git ls-files'
alias gitmr="git merge --strategy-option theirs"
alias gitml="git merge --strategy-option ours"
alias gits='git status'
alias gitckr='git checkout --recurse-submodules --remote'
alias gitcl='git clone --recurse-submodules'
alias gitlr='git submodule update --recursive --remote'
alias gitsub='git submodule update --init --recursive'
alias gitf='git fetch --recurse-submodules'
alias gitpull='for d in */ ; do  pushd $d;    git pull; popd; done'
alias gitlog='git log --graph --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'
alias gitlogfull='git log --graph --full-history --all --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'
alias gitd1='git diff HEAD^ HEAD'
alias gitd2='git diff HEAD^^ HEAD'
alias gitd3='git diff HEAD^^^ HEAD'
alias gitl1w='git log --oneline --since=1.weeks'
alias gitl1m='git log --oneline --since=1.months'
alias gitl3m='git log --oneline --since=3.months'
alias gitll='git log --oneline HEAD^..HEAD'
alias gitll2='git log --oneline HEAD^^..HEAD'

a mks='make sandbox'
a mki='make inspect'
a mkm='make maintain'
a mc='machinectl'
a mkf='make sandbox.refresh'
a cds='cd /mnt/sdcard/shining_software'
a cdr='cd /mnt/sdcard/shining_software/src/dep/roc_services'
a cdrc='cd /mnt/sdcard/shining_software/src/dep/roc_services/cmd/roc/cmd'
a cdsb='cd ~/sandbox'
a fresh='git submodule update --init --recursive; ./little-firmware.sh'


