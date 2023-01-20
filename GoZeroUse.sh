#! /bin/bash

OutType=$1

if [[ $OutType == "create" ]];then 
    goctl api go -api api/BaseInterface.api -dir BaseInterface/ -style GoZero
fi

if [[ $OutType == "build" ]];then 
    cd BaseInterface
    go build Baseinterface-Api.go
    cd -
fi

if [[ $OutType == "run" ]];then 
    cd BaseInterface
    go run Baseinterface-Api.go -f etc/BaseInterface-Api.yaml
    cd -
fi