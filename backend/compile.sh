# Simple script to build and run the program on a default image or one passed from the user

#!/usr/bin/bash

IMAGE=./images/windypic.jpg

if [ $# -eq 0 ]; then
    echo "Warning: No arguments provided, using default image windypic.jpg"
else
    IMAGE=$1
fi

cd $HOME/Programming/go-image/backend

go build && ./goimg $IMAGE && xdg-open NEWIMAGE.jpeg &> /dev/null 
