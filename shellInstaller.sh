#!/bin/sh

echo `go`

if [[ $? != 0 ]]; then
    echo "Go is not installed on your system!"
    echo "Do you want I install it?[Y/n]"
    read $RESULT
    if [[ $RESULT == "n" ]] || [[ $RESULT == "N" ]]; then
        echo "So, You must install Go in the moment . . ."
    else
        echo "Installing GoLang . . ."
        echo `pkg`, `install`, `golang`

        if [[ $? != 0 ]]; then
            echo "There is a problem! Plaese take action yourself."
        else
            echo -e "Install Successful!\nYou can run main.go!"
        fi
    fi
fi
