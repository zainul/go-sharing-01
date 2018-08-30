#!/bin/bash
set -e


GometalinterVariable=(
           "aligncheck"
           "deadcode"
           "dupl"
           "errcheck"
          #  "gas"
           "goconst"
           "gocyclo"
           "goimports"
           "golint"
           "gosimple"
        #    "gotype"
           # about gotype please see this issue https://github.com/alecthomas/gometalinter/issues/206
           # and https://github.com/alecthomas/gometalinter/issues/355
           "ineffassign"
          #  "interfacer"
          #  "lll"
           "misspell"
           "safesql"
           "staticcheck"
           "structcheck"
           "unconvert"
           "unparam"
           "unused"
           "varcheck"
)


Directory=(
            "pkg/station/handler"
            "pkg/station/service"
            "pkg/station/store"
          )

arrayGometalinterVariable=${#GometalinterVariable[@]}
arrayDirectory=${#Directory[@]}


for ((k=0; k<${arrayDirectory}; k++));
do
        #cd ${Directory[$k]}
  for ((i=1; i<${arrayGometalinterVariable}; i++));
  do
        if [ "${Directory[$k]}" == "controllers" ]
          then
          if [ "${GometalinterVariable[$i]}" != "gocyclo" ] && [ "${GometalinterVariable[$i]}" != "lll" ] && [ "${GometalinterVariable[$i]}" != "dupl" ] && [ "${GometalinterVariable[$i]}" != "goconst" ]
            then
            echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
            gometalinter.v2 -j 1 --disable-all --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
          fi
        else
          echo "Currently linter (without exception)running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
          gometalinter.v2 -j 1 --disable-all --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
        fi

        sleep 1
        wait

  done
done