#!/usr/bin/env bash

set -e

echo -e "\e[1;94mBuilding CBM...\e[0m"
go build -o cbm main.go
echo -e "\e[1;94mMoving CBM to /usr/bin\e[0m"
sudo mv cbm /usr/bin/
