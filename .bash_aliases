#!/usr/bin/env bash

#export PS1=${ret_status} %{$fg[cyan]%}%c%{$reset_color%} $(git_prompt_info)
#set -e

#echo 🔥 ✋ 🛑  💣
### echo v.2.4.7.0
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
alias ivo-git='git config --global user.email "ivostoy@gmail.com"; git config --global user.name "ivo";  git config --global push.default simple'

alias gitc='git commit -a -m '
alias gitconf='git config --list --show-origin'
alias gitck='git checkout'
alias gitp='git push; git push --tags'
alias gitl='git pull'
alias gitls='git ls-files'
alias gitmr="git merge --strategy-option theirs"
alias gitml="git merge --strategy-option ours"
alias gits='git status'
alias gitsq='echo "to squash last N commits - append HEAD~N" && git reset --soft'
alias gitresetDEVELOP='git reset --hard origin/develop'


alias wip='git commit -a -m wip; git push'
alias wipp='git commit -a -m wip && git push'

alias gitlog='git log --graph --decorate --oneline'
alias gitlog='git log --graph --full-history --all --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'

# Enable aliases to be sudo’ed
alias sudo='sudo '

# always use unidiff
alias diff='diff -u'
alias k9='kill -9'

alias pssh='ps aux|grep ssh'
alias gign='vi .gitignore'

export PATH=/usr/local/bin:$HOME/tools:$PATH

alias sba='source ~/.bash_aliases'
alias pu='lsof -i '
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
alias eh='sudo vi /etc/hosts'
alias prot='chmod 0400'

function MD() {
  echo creating dir "$1"
  mkdir -p "$1"
  cd "$1" || return
  pwd
}

function LEN() {
  V=$1
  echo length of "$V": ${#V}
}

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
  
function setup_git() {
  git config --global user.email ivostoy@gmail.com
  git config --global user.name ivo
  git config --global core.editor "subl -n -w"
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

export EDITOR=subl
alias inst='sudo apt install'
alias upd='sudo apt update'
alias eba='nano ~/.bash_aliases; source ~/.bash_aliases'

## aliases depending on OS
### MAC
if [[ "$OSTYPE" =~ ^darwin ]]; then 
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
  alias subl='/Applications/Sublime\ Text.app/Contents/SharedSupport/Bin/subl'
  #alias m=multipass

	#x=$(brew --prefix)/etc/bash_completion
	# shellcheck disable=SC1090
	#[[ -f "$x" ]] && source "$x"
fi
### END MAC

#curl -i -X POST -H "$A" $URL  -d '{"value": "'$L'"}'

# colorizer
[[ -s "/usr/local/etc/grc.bashrc" ]] && source /usr/local/etc/grc.bashrc

# if you have ssh problems
alias sshv='ssh -vvv -o LogLevel=DEBUG3'

#a mb='make extraclean && make api && make generate'
alias mode="stat -f '%A %a %N' "

###############
# Coral board #
###############
# ping bored-kid.local
# 10.0.1.194
export CORAL=10.0.1.194
a sh-coral='mdt shell'
alias N='date +''%s'''

# default file creation mask
umask 0022

alias mm='make menuconfig'

# brew cask install adoptopenjdk13

if [[ -f /usr/libexec/java ]]; then
  export JAVA_HOME=`/usr/libexec/java_home -v 13.0`
  export PATH=$JAVA_HOME/bin:$PATH

  export MAVEN_HOME=$HOME/tools/apache-maven-3.6.3
  export PATH=$MAVEN_HOME/bin:$PATH

  #echo $JAVA_HOME
  #java -version
  #mvn -version
fi
a cdj='cd ~/github/sandbox/java'



