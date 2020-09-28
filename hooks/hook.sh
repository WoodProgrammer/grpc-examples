#!/bin/sh

get_current_branch(){
  export BRANCH=$(git branch --show-current)
}

check_remote_master(){
  git log ${BRANCH} --pretty=format:'%h' > /tmp/local_branch
  git fetch --all
  git log origin/master --pretty=format:'%h' > /tmp/remote_master

  while IFS= read line; do
    if grep -q $line "/tmp/local_branch"; then
      echo "Exist"
    else
      echo "Hey One missing commit from remote master ${line}"
    fi

  done < /tmp/remote_master
}

get_current_branch
check_remote_master

