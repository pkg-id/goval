#!/usr/bin/env bash
set -e -o pipefail
## this will retrieve all of the .go files that have been
## changed since the last commit
STAGED_GO_FILES=$(git diff --cached --diff-filter=AM --name-only -- '*.go')

## we can check to see if this is empty
if [[ $STAGED_GO_FILES == "" ]]; then
    echo "No Go Files to Update"
## otherwise we can do stuff with these changed go files
else
    if ! command -v gofmt &> /dev/null ; then
        echo "gofmt not installed or available in the PATH" >&2
        exit 1
    fi

    for file in $STAGED_GO_FILES; do
        ## format our file
        ## add any potential changes from our formatting to the
        ## commit
        gofmt -l -w $file
        git add $file
    done
fi
