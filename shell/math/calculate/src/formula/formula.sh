#!/bin/sh
runFormula() {
  if [ "sum" = "$RIT_OPERATION" ]; then
    echo "rit shell math sum numbers --rit_number_one=$RIT_NUMBER_ONE --rit_number_two=$RIT_NUMBER_TWO"
  elif [ "multiplication" = "$RIT_OPERATION" ]; then
    echo "rit shell math multiply numbers --rit_number_one=$RIT_NUMBER_ONE --rit_number_two=$RIT_NUMBER_TWO"
  else
    echo "Unexpected operation type: $RIT_OPERATION"
  fi
}
