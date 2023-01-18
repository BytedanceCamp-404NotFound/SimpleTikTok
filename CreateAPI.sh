#! /bin/bash

echo ==============create api file=================
cd BaseInterface
goctl api go -api api/BaseInterface.api -dir . -style GoZero
cd -