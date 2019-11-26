#!/bin/bash
# append -x to debug

######
##
## ROC API Load test script
## requires working installation of vegeta https://github.com/tsenart/vegeta
##
######

CONFIG_FILE="./stress.config"
DEVICE_LIST="/tmp/devices.tmp"

function parse_args() {
  echo "Parsing command line args " "$@" >&2
  while getopts "hc:d:e:" opt; do
    case $opt in
    c)
      CONN=$OPTARG
      ;;
    d)
      MAX_DEV=$OPTARG
      ;;
    e)
      ENV=$OPTARG
      ;;
    h)
      print_help
      exit 0
      ;;
    r)
      RATE=$OPTARG
      ;;
    t)
      DUR=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
    esac
  done

  BASE="https://api.$ENV.roc.braintld.com"
}

function print_help() {
  echo "Usage: roc_load_test.sh [-c connections] [-d maxdevices] [-e environment] [-r QPS] [-t duration]" >&2
  echo "Config file: $CONFIG_FILE" >&2
}

function read_config() {
  [[ -z "$CONFIG_FILE" ]] && CONFIG_FILE="stress.config"

  echo "Reading configuration $CONFIG_FILE" >&2
  source $CONFIG_FILE
}

function print_config() {
  printf "Environment (-e, ENV):\t %s \n" "$ENV" >&2
  printf "Number of connections (-c, CONN):\t %s \n" "$CONN" >&2
  printf "Max number of devices (-d, MAX_DEV):\t %s \n" "$MAX_DEV" >&2
  printf "Request rate (QPS) (-r, RATE):\t %s \n" "$RATE" >&2
  printf "Test duration (-t, DUR):\t %s \n" "$DUR" >&2
}

# write device ids to temp file
# max $1 devices
function write_dev_ids() {
  set -x
  roc devices --env $ENV list --limit $MAX_DEV --output json | jq '.[] | .id' -r >"$DEVICE_LIST"
  set +x
}

function hit_devices() {
  COUNT=$(wc -l "$DEVICE_LIST" | awk '{ print $1 }')
  rm "$REPORT"
  echo "$COUNT devices in $DEVICE_LIST" >&2
  while read -r line
  do
    echo "$line" >&2
    attack "$line"
  done < "$DEVICE_LIST"
  echo "report written to $REPORT" >&2
  cat "$REPORT"
}

# execute in background
function attack() {
  URL=$BASE/v1/devices/$1/status
  #set -x
  jq -ncM '{method: "GET", url: "'"$URL"'",header: {"Authorization":["Bearer '"$TOKEN"'"]}}' |
    vegeta attack -format=json -duration="$DUR" -rate="$RATE" -connections="$CONN" |
    vegeta report >> "$REPORT" &
  #set +x
}

read_config
parse_args "$@"
print_config
write_dev_ids
hit_devices
