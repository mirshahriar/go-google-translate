#!/usr/bin/env python

import sys
import subprocess
from os.path import expandvars

ROOT = expandvars('$GOPATH') + '/src/github.com/aerokite/go-google-translate/cmd'

def call(cmd, stdin=None, cwd=ROOT):
    print('$ ' + cmd)
    subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)

def fmt():
    call('goimports -w ../.')
    call('gofmt -s -w ../.')

def compile():
    call('go install ./...')

def build():
    call('CGO_ENABED=0 go build -o build/gopret ROOT')

def default():
    fmt()
    compile()

if __name__ == "__main__":
    if len(sys.argv) > 1:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        default()
