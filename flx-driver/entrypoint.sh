#!/bin/bash

if [ -d "/opt/felix/drivers_rcc/script/" ]
then
  cd /opt/felix/drivers_rcc/script/
  . drivers_flx_local start
else
  echo "No /opt/felix/drivers_rcc/ found"
fi
