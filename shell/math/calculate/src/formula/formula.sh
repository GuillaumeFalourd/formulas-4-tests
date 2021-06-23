#!/bin/sh
runFormula() {
  if [ "sum" = $OPERATION]; then
    echo "rit shell math sum numbers --rit_number_one=$RIT_NUMBER_ONE --rit_number_two=$RIT_NUMBER_TWO"
  elif [ "multiplication" = $OPERATION]; then
    echo "rit shell math multiply numbers --rit_number_one=$RIT_NUMBER_ONE --rit_number_two=$RIT_NUMBER_TWO"
  else
    echo "Unexpected operation type: $OPERATION"
  fi
}
