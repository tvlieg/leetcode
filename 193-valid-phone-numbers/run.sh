#! /usr/bin/env bash
awk '/^(...-|\(...\) )...-....$/{print $0}' file.txt
