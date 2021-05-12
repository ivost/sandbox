#!/usr/bin/env bash

#export PS1=${ret_status} %{$fg[cyan]%}%c%{$reset_color%} $(git_prompt_info)
#set -e

#echo 🔥 ✋ 🛑  💣
### echo v.7.19.21.0
#echo 🔥 ✋ 🛑  💣

export PS1="\T \W$ "
export HISTCONTROL=erasedups
export HISTSIZE=10000
export HISTFILESIZE=10000
# dont save commands starting with space 
#export HISTIGNORE="[ \t]*:pwd:ls:ll:h:a:rm"
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

alias clone='git clone --recursive -j8'
alias clone1='git clone --recurse-submodules -j8 --depth 1'
alias gitsubu='git submodule update --init --recursive'

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
alias gitsq='echo "to squash last N commits - append HEAD~N" && git reset --soft'
alias gitresetDEVELOP='git reset --hard origin/develop'
alias gitli='git lfs install'

alias wip='git commit -a -m wip'
alias wipp='git commit -a -m wip && git push'

alias gitlog='git log --graph --decorate --oneline'
alias gitlog='git log --graph --full-history --all --color --pretty=format:"%x1b[31m%h%x09%x1b[32m%d%x1b[0m%x20%s"'

# Enable aliases to be sudo’ed
#alias sudo='sudo '

# always use unidiff
alias diff='diff -u'
alias k9='kill -9'

alias pssh='ps aux|grep ssh'
alias gign='vi .gitignore'

export PATH=/usr/local/bin:$HOME/tools:$PATH
export PATH=$HOME/tools/platform-tools:$PATH
export PATH=$HOME/.bin:$PATH

# export PATH=/Users/ivo/Library/Python/3.9/bin:$PATH
#export PATH=$HOME/go/bin:$PATH
#export GOPATH=$HOME/go


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

function scp1() {
  scp $1 root@192.168.1.1:
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

export EDITOR=gedit
alias inst='sudo apt install'
alias upd='sudo apt update'
alias eba='subl ~/.bash_aliases; source ~/.bash_aliases'

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
#[[ -s "/usr/local/etc/grc.bashrc" ]] && source /usr/local/etc/grc.bashrc

# if you have ssh problems
alias sshv='ssh -vvv -o LogLevel=DEBUG3'

#a mb='make extraclean && make api && make generate'
alias mode="stat -f '%A %a %N' "
alias mm='make menuconfig'

# export NVM_DIR="$HOME/.nvm"
# [ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
# [ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && . "/usr/local/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion

a vlc='/Applications/VLC.app/Contents/MacOS/VLC'

export OPENSSL_DIR=/usr/local/opt/openssl
export OPENSSL_ROOT_DIR=/usr/local/opt/openssl

# export OPENSSL_DIR=/usr/local/Cellar/openssl@1.1/1.1.1g
# export OPENSSL_ROOT_DIR=/usr/local/Cellar/openssl@1.1/1.1.1g

#If you need to have openssl first in your PATH run:
#  echo 'export PATH="/usr/local/opt/openssl/bin:$PATH"' >> ~/.profile

#export PATH="/usr/local/opt/openssl/bin:$PATH"
#export PATH="/usr/local/Cellar/mosquitto/1.6.12/bin:$PATH"
#alias mqstart='mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf'
alias mqstart='brew services start mosquitto'
alias mqstop='brew services stop mosquitto'
alias mqs='mosquitto_sub -t top'
alias mqp='mosquitto_pub -t top'
alias mqs1='mosquitto_sub -h 192.168.1.1 -t top'

#For compilers to find openssl you may need to set:
export LDFLAGS="-L/usr/local/opt/openssl/lib"
export CPPFLAGS="-I/usr/local/opt/openssl/include"

#For pkg-config to find openssl you may need to set:
export PKG_CONFIG_PATH="/usr/local/opt/openssl/lib/pkgconfig"

export I=192.168.1.1
alias sshj='ssh root@$I'

a sshu='ssh -i "~/.ssh/ivo-keypair-2020.pem" ubuntu@ec2-18-189-20-67.us-east-2.compute.amazonaws.com'
complete -C '/usr/local/bin/aws_completer' aws

a gb='go build  -ldflags "-s -w"; ls -altrh'

a gba='GOOS=linux GOARCH=arm  CGO_ENABLED=0 go  build -ldflags "-s -w" ; ls -altrh'
a gba1='GOOS=linux GOARCH=arm  go  build; ls -altrh'

# export PATH="$HOME/.jenv/bin:$PATH"
# eval "$(jenv init -)"

#export PATH="$HOME/tools/apache-maven-3.6.3/bin:$PATH"

if type brew &>/dev/null; then
  HOMEBREW_PREFIX="$(brew --prefix)"
  if [[ -r "${HOMEBREW_PREFIX}/etc/profile.d/bash_completion.sh" ]]; then
    source "${HOMEBREW_PREFIX}/etc/profile.d/bash_completion.sh"
  else
    for COMPLETION in "${HOMEBREW_PREFIX}/etc/bash_completion.d/"*; do
      [[ -r "$COMPLETION" ]] && source "$COMPLETION"
    done
  fi
fi

###############
# Coral board #
###############
# ping bored-kid.local
# 10.0.1.194
export CORAL=192.168.4.144
a mdt-coral='mdt shell'
a sh-coral='ssh mendel@$CORAL'

export U_C="mendel@$CORAL"

alias N='date +''%s'''

# default file creation mask
#umask 0022

# AI-CAM
export AC=192.168.0.20
# pw is S...1
a sh-ac='ssh ivo@$AC'

a p=python3

# eval "$(pyenv init -)"
# eval "$(pyenv virtualenv-init -)"

export A1=51.143.89.221
a az1='ssh -i ~/.ssh/ivo-ubuntu-1_key.pem ivo@$A1'

#export R3=10.0.1.20
export R3=10.0.1.4
a r3='ssh pi@$R3'

a sa1='ssh -i ~/.ssh/ivostoy-897440107178keypair.pem ubuntu@52.40.243.181'

# export R4=10.0.1.194
# a r4='ssh pi@$R4'


export KE='/Users/ivo/go/src/github.com/kubeedge'
a cdke='cd $KE'

#a ss-pi='ssh pi@10.0.1.170'

a k='kubectl --insecure-skip-tls-verify'

export LEN=192.168.4.164
export PI4e=192.168.4.119
export PI4=192.168.4.55
export ODY=192.168.4.124

a sh-pi4='ssh ivo@$PI4'

export UB=10.1.2.143

export U1=52.250.11.246
export U2=52.229.50.232

export AU1=AzureUser@$U1
export AU2=azureuser@$U2

a ub='ssh ivo.stoyanov@$UB'
a u1='ssh -i ~/.ssh/ivo-ubuntu-1_key.pem  $AU1'
a u2='ssh -i ~/.ssh/ivo-ubuntu-1_key.pem  $AU2'
export UUB="ivo.stoyanov@$UB"

# jetson nano
export JN=192.168.4.147
export UJ=ivo@$JN
a sh-jn='ssh $UJ'

export V3=192.168.4.157
export UK=khadas@$V3
a sh-k='ssh $UK'
a sh-k='ssh $V3'

export NUC=192.168.4.165
export UN=ivo@$NUC
a nuc='ssh $UN'

export TB=192.168.5.173
export UTB=ivo@$TB

a scp2='scp -p222'

a tb='ssh -p222 $UTB'

# scp to
a tb-to='scp $1 $UTB:'
# scp from
a tb-from='scp -p222 $UTB:$1 .'

a rpi3a='ssh pi@rpi3a.local'
a m=multipass
a ody='ssh ivo@odyssey.local'
a sh-len='ssh $LEN'
# scp to
a ub-ct='scp $1 ivo.stoyanov@$UB:'
# scp from
a ub-cg='scp ivo.stoyanov@$UB:$1 .'

#  export IP=$(multipass info faasd --format json| jq '.info.faasd.ipv4[0]' | tr -d '\"')

a pyc='open -na "PyCharm.app"'
a cd-ov='cd ~/github/openvino-playground'

#export OV="/opt/intel/openvino"
#source $OV/bin/setupvars.sh

a pipi='pip install -r requirements.txt'
a pipu='python -m pip install --upgrade pip'

a nb='jupyter notebook'

### /opt/intel/openvino_2021/deployment_tools/inference_engine/lib/intel64/libinference_engine*.so /lib/x86_64-linux-gnu/ -v

# export OV="$HOME/intel/openvino_2021"
# a ovi='source $OV/bin/setupvars.sh'

# export OV_DT=$OV/deployment_tools
# export OV_IE=$OV/inference_engine

# export DYLD_LIBRARY_PATH=$OV/opencv/lib:$OV_DT/ngraph/lib:$OV_DT/inference_engine/lib/intel64:$OV_DT/inference_engine/external/tbb/lib
# export LD_LIBRARY_PATH=$DYLD_LIBRARY_PATH

# echo $DYLD_LIBRARY_PATH
# echo $LD_LIBRARY_PATH

alias caz='conda activate zipline'
alias preq='pip freeze > requirements.txt'
alias dai-env='pyenv local depthai'

#export X=/usr/local/gcc-arm-none-eabi-9-2019-q4-major
#export PATH=$PATH:$X/bin

# git clone --depth 1 https://github.com/tensorflow/tensorflow
# pyenv virtualenv 3.7.9 tf2

# export LDFLAGS="-L/usr/local/opt/tcl-tk/lib"
# export DYLDFLAGS="-L/usr/local/opt/tcl-tk/lib"
# export CPPFLAGS="-I/usr/local/opt/tcl-tk/include"
# export PKG_CONFIG_PATH="/usr/local/opt/tcl-tk/lib/pkgconfig"

export MP="$HOME/github/myriad-playground"

# export PYTHON_PATH=$MP/insg:$MP/insg/common:$MP/insg/oak:$PYTHON_PATH
# echo $PYTHON_PATH

export PATH=$MP/insg/common:$PATH
a cdm='cd $MP'
a cdo='cd $MP/insg/oak'

a pips='pipenv shell'

a ys='yarn start'
a yb='yarn run build'
a dc='docker compose'
a dea='source deactivate'

# brew info bison

export PATH="/usr/local/opt/bison/bin:$PATH"

export LDFLAGS="-L/usr/local/opt/bison/lib"


# export PATH="$HOME/.jenv/bin:$PATH"
# eval "$(jenv init -)"
