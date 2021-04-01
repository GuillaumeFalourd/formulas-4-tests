#!/bin/sh
runFormula() {
  echo "***********************************************************"
  echo "* Running with credentials                                 "
  echo "* USERID: $CREDENTIAL_AWS_CALLERUSERID "
  echo "* ARN: $CREDENTIAL_AWS_CALLERARN "
  echo "* ACCOUNT: $CREDENTIAL_AWS_CALLERACCOUNT "
  echo "***********************************************************"
  echo "Running command [$RIT_COMMAND] using AWS CLI "

  AWS_PAGER="" aws $RIT_COMMAND
}
