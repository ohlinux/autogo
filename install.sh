#!/usr/bin/env bash

if [ ! -f install.sh ]; then
    echo 'install.sh must be run within its container folder' 1>&2
    exit 1
fi

AUTOGO_CMD="/usr/bin/autogo"
CURDIR=`pwd`
export GOPATH="$CURDIR"

cat << EOF > $AUTOGO_CMD
#!/usr/bin/env bash
export AUTOGO_CMD="$AUTOGO_CMD"
export AUTOGO_ROOT="$CURDIR"
$CURDIR/bin/autogo
EOF

chmod +x $AUTOGO_CMD

gofmt -tabs=false -tabwidth=4 -w src

go install autogo

echo 'Install finished!!!'
