#!/usr/bin/env bash

echo -e "\e[1;94mBuilding CBM...\e[0m"
go build -o cbm main.go
if [[ $? != 0 ]]; then echo "CodeBaseManager Couldn't build"; exit 1; fi
echo -e "\e[1;94mMoving CBM to /usr/bin\e[0m"
sudo mv cbm /usr/bin/

USER_HOME=""
CBM_DIRNAME=".cbm"
CBM_GLOBAL_PATH=""
CBM_PORT=""
CBM_BACKEND_DIR="backend"
CBM_GLOBAL_BACKEND_DIR=""
CBM_API_ENV_FILENAME=".env"
CBM_API_ENV_PATH=""

function validateHome() {
  while true; do
    echo -e "\e[1;94mPlease validate your HOME directory: $HOME (Yy/Nn)\e[0m"
    read -r yn
    case $yn in
        [Yy]* ) USER_HOME="$HOME";
          CBM_GLOBAL_PATH="$USER_HOME/$CBM_DIRNAME"
          mkdir -p "$USER_HOME/$CBM_DIRNAME";
          break;;
        [Nn]* ) exit;;
        * ) echo -e "\e[1;94mPlease answer yes or no.\e[0m";;
    esac
done
}

function installBackend() {
  CBM_GLOBAL_BACKEND_DIR="$CBM_GLOBAL_PATH/$CBM_BACKEND_DIR"
  CBM_API_ENV_PATH="$CBM_GLOBAL_PATH/$CBM_BACKEND_DIR/$CBM_API_ENV_FILENAME"

  rm -rfd "$CBM_GLOBAL_BACKEND_DIR"
  cp -r "$CBM_BACKEND_DIR" "$CBM_GLOBAL_BACKEND_DIR"
  echo "PORT=$CBM_PORT" > "$CBM_API_ENV_PATH"

  cd "$CBM_GLOBAL_BACKEND_DIR" && make api-clean-stop && make api-start
}

function cbmPort() {

  while true; do
    if [[ $CBM_PORT != "" ]]; then echo "\n$CBM_PORT didn't work" ;fi
    echo -e "\e[1;94mPlease enter a valid port on which Cbm Backend could run:\e[0m"
    read -r CBM_PORT
    echo -e "\e[1;94mLaunching CodeBaseManager Backend ...\n\n\e[0m"

    if ! installBackend ;then continue ; else break ; fi
done
echo -e "\e[1;94mApi Running ... \n\e[0m"
}

validateHome
cbmPort

echo -e "\e[1;94mCodebase manager Successfully Installed !!\n\n\e[0m"
