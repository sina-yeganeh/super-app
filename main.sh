#!/bin/sh
# SinaYeaneh0-0

# File name variable:
FILENAME[0]="Calualtor.go"
FILENAME[1]="guessNumber.go"
FILENAME[2]="Rock_Cuter_Paper.go"
FILENAME[3]="Shutdown.go"
FILENAME[4]="timeTool.go"
FILENAME[5]="banner.go"

function banner() {
    cd modules
    clear
    go run ${FILENAME[5]}
    if [[ $? != 0 ]]; then
        echo "Some think went wrong!"
    else
        cd ..
    fi
}

function root() {
    echo -e "
[1] Calualtor
[2] Guess Number
[3] Rock, Cuter, Paper
[4] Time Tool
[5] Shutdown
[6] About Me
[7] My GitHub
[0] Exit"

    read -p "Please enter a option number: " $OPTION
    export $OPTION
}

function runtools(TOOL) {
    cd modules
    go run TOOL
}

ABOUTME="I am SinaYegaeneh0-0 from Iran!\nI love book and programming :)"
MYGITHUB="My username is: SinaYeaneh0-0 And my name is Sina Yeganeh Faal"

function checkandrun() {
    case $OPTION in
        1 ) runtools ${FILENAME[0]}
            ;;
        2 ) runtools ${FILENAME[1]}
            ;;
        3 ) runtools ${FILENAME[2]}
            ;;
        4 ) runtools ${FILENAME[3]}
            ;;
        5 ) runtools ${FILENAME[4]}
            ;;
        6 ) printf $ABOUTME
            ;;
        7 ) printf $MYGITHUB
            ;;
        0 ) break
    esac
}

for (( i = 0; ; i++ )); do
    banner
    checkandrun
    read -p "Do you want to use agian?[Y/n]" $USE
    if [[ $? == 0 ]]; then
        echo "Some think went wrong! Please try again. . ."
    elif [[ $USE == "Y" || "y" || "" ]]; then
        continue
    else
        break
    fi
done
