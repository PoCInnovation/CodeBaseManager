#!/usr/bin/env bash

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
            [Yy]*)
                USER_HOME="$HOME"
                CBM_GLOBAL_PATH="$USER_HOME/$CBM_DIRNAME"
                mkdir -p "$USER_HOME/$CBM_DIRNAME"
                break
                ;;
            [Nn]*) exit ;;
            *) echo -e "\e[1;94mPlease answer yes or no.\e[0m" ;;
        esac
    done
}

function copyConfigs() {
    FILES=("watcher" "func-test")

    for file in ${FILES[*]}; do
        cp "$CBM_DIRNAME/$file.toml" "$CBM_GLOBAL_PATH"
    done
}

function installBackend() {
    CBM_GLOBAL_BACKEND_DIR="$CBM_GLOBAL_PATH/$CBM_BACKEND_DIR"
    CBM_API_ENV_PATH="$CBM_GLOBAL_PATH/$CBM_BACKEND_DIR/$CBM_API_ENV_FILENAME"

    echo -e "\e[1;94mRemoving directory: $CBM_GLOBAL_BACKEND_DIR\e[0m"
    rm -rfd "$CBM_GLOBAL_BACKEND_DIR"
    echo -e "\e[1;94mInstalling CodeBaseManager Backend in $CBM_GLOBAL_BACKEND_DIR\e[0m"
    cp -r "$CBM_BACKEND_DIR" "$CBM_GLOBAL_BACKEND_DIR"
}

function writePort() {
    echo "CBM_PORT=$CBM_PORT" >"$CBM_API_ENV_PATH"

    cd "$CBM_GLOBAL_BACKEND_DIR" && make api-clean-stop && make api-start || return 1
}

function getPortCBM() {
    local wd
    wd=$(pwd)

    while true; do
        if [[ $CBM_PORT != "" ]]; then echo -e "\n\e[1;94m$CBM_PORT didn't work\e[0m"; fi

        echo -e "\e[1;94mPlease enter a valid port on which Cbm Backend could run:\e[0m"
        read -r CBM_PORT
        echo -e "\e[1;94mLaunching CodeBaseManager Backend ...\n\n\e[0m"

        if ! writePort; then
            cd "$wd" || echo -e "\e[1;94mProblem with CodebaseManager installation.\e[0m" && exit
            continue
        else
            break
        fi
    done

    cd "$wd" || exit 1
    echo -e "\e[1;94mApi Running ... \n\e[0m"
}

function installCbm() {
    # Installing the CLI
    BIN_PATH="/usr/bin"
    echo -e "\e[1;94mBuilding CodeBaseManager(cli)...\e[0m"
    go build -o cbm cli/main.go
    if [[ $? != 0 ]]; then
        echo "CodeBaseManager(cli) couldn't build"
        exit 1
    fi

    echo -e "\e[1;94mMoving CodeBaseManager(cli) to $BIN_PATH\e[0m"
    sudo mv cbm $BIN_PATH

    # Installing the Watcher
    WATCHER="cbm-watcher"
    WHERE="watcher"
    echo -e "\e[1;94mBuilding $WATCHER...\e[0m"
    go build -o $WATCHER $WHERE/main.go
    if [[ $? != 0 ]]; then
        echo "$WATCHER Couldn't build"
        exit 1
    fi

    echo -e "\e[1;94mMoving $WATCHER to $BIN_PATH\e[0m"
    sudo mv $WATCHER $BIN_PATH

    # Preping home directory
    validateHome
    copyConfigs

    # Installing backend
    if ! installBackend; then
        echo -e "\e[1;94mProblem with CodebaseManager installation.\e[0m"
        exit 1
    fi

    getPortCBM
    echo -e "\e[1;94mCodebase manager Successfully Installed !!\n\n\e[0m"
}

installCbm
