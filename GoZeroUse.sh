#! /bin/bash

OutType=$1

if [ -z "$1" ]
then
    echo "useage: ./GoZeroUse [operation] [type]"
    echo "example1: ./GoZeroUse create"
    echo "example2: ./GoZeroUse test sql"
    exit 1
fi

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
    go build Baseinterface-Api.go
    # 需要使用go build生成的exe文件来执行，这样os.Executable()获取到的才是正确的路径
    ./Baseinterface-Api -f etc/BaseInterface-Api.yaml  
    # go run来运行，会将可执行文件默认放到/tmp/go-build...
    # 需要配置GOTMPDIR=""来改变go run生成可执行文件的位置
    # go run Baseinterface-Api.go -f etc/BaseInterface-Api.yaml
    cd -
fi

if [[ $OutType == "test" ]];then 
    case $2 in
    "minio") echo "test minio"
    ;;
    "mysql") echo "test mysql"
    ;;
    *) echo "useage: ./GoZeroUse test sql"
    ;;
    esac
fi
