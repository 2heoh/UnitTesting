#!/usr/bin/env bash

# Copyright 2011 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This script creates a complete godoc app in $APPDIR.
# It copies the cmd/godoc and src/go/... sources from GOROOT,
# synthesizes an app.yaml file, and creates the .zip, index, and
# configuration files.
#
# If an argument is provided it is assumed to be the app-engine godoc directory.
# Without an argument, $APPDIR is used instead. If GOROOT is not set, "go env"
# is consulted to find the $GOROOT.
#
# The script creates a .zip file representing the $GOROOT file system
# and computes the correspondig search index files. These files are then
# copied to $APPDIR. A corresponding godoc configuration file is created
# in $APPDIR/appconfig.go.

ZIPFILE=godoc.zip
INDEXFILE=godoc.index
SPLITFILES=index.split.
GODOC=golang.org/x/tools/cmd/godoc
CONFIGFILE=$GODOC/appconfig.go

error() {
	echo "error: $1"
	exit 2
}

getArgs() {
	if [ -z $APPENGINE_SDK ]; then
		error "APPENGINE_SDK environment variable not set"
	fi
	if [ ! -x $APPENGINE_SDK/go ]; then
		error "couldn't find go comment in $APPENGINE_SDK"
	fi
	if [ -z $GOROOT ]; then
		GOROOT=$(go env GOROOT)
		echo "GOROOT not set explicitly, using go env value instead"
	fi
	if [ -z $APPDIR ]; then
		if [ $# == 0 ]; then
			error "APPDIR not set, and no argument provided"
		fi
		APPDIR=$1
		echo "APPDIR not set, using argument instead"
	fi
	
	# safety checks
	if [ ! -d $GOROOT ]; then
		error "$GOROOT is not a directory"
	fi
	if [ -e $APPDIR ]; then
		error "$APPDIR exists; check and remove it before trying again"
	fi

	# reporting
	echo "GOROOT = $GOROOT"
	echo "APPDIR = $APPDIR"
}

fetchGodoc() {
	echo "*** Fetching godoc (if not already in GOPATH)"
	unset GOBIN
	go=$APPENGINE_SDK/go
	$go get -d -tags appengine $GODOC
	mkdir -p $APPDIR/$GODOC
	cp $(find $($go list -f '{{.Dir}}' $GODOC) -type f -depth 1) $APPDIR/$GODOC/
}

makeAppYaml() {
	echo "*** make $APPDIR/app.yaml"
	cat > $APPDIR/app.yaml <<EOF
application: godoc
version: 1
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
EOF
}

makeZipfile() {
	echo "*** make $APPDIR/$ZIPFILE"
	zip -q -r $APPDIR/$ZIPFILE $GOROOT/*
}

makeIndexfile() {
	echo "*** make $APPDIR/$INDEXFILE"
	GOPATH= godoc -write_index -index_files=$APPDIR/$INDEXFILE -zip=$APPDIR/$ZIPFILE
}

splitIndexfile() {
	echo "*** split $APPDIR/$INDEXFILE"
	split -b8m $APPDIR/$INDEXFILE $APPDIR/$SPLITFILES
}

makeConfigfile() {
	echo "*** make $APPDIR/$CONFIGFILE"
	cat > $APPDIR/$CONFIGFILE <<EOF
package main

// GENERATED FILE - DO NOT MODIFY BY HAND.
// (generated by $GOROOT/src/cmd/godoc/setup-godoc-app.bash)

const (
	// .zip filename
	zipFilename = "$ZIPFILE"

	// goroot directory in .zip file
	zipGoroot = "$GOROOT"

	// glob pattern describing search index files
	// (if empty, the index is built at run-time)
	indexFilenames = "$SPLITFILES*"
)
EOF
}

getArgs "$@"
set -e
mkdir $APPDIR
fetchGodoc
makeAppYaml
makeZipfile
makeIndexfile
splitIndexfile
makeConfigfile

echo "*** setup complete"
