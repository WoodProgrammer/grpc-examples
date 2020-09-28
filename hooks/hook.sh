#!/bin/sh

get_current_branch(){
  export BRANCH=$(git branch --show-current)
}

check_remote_master(){
  git log calculator --pretty=format:'%h' > /tmp/local_branch
  git fetch --all
  git log origin/master --pretty=format:'%h' > /tmp/remote_master
  DIFF_REMOTE_MASTER=$(comm -23 /tmp/remote_master /tmp/local_branch)

}

notify_developer(){
  echo "Hello you are far away from the remote master ${DIFF_REMOTE_MASTER}"
}

#get_current_branch
check_remote_master
notify_developer
