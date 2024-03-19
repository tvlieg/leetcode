#! /usr/bin/env bash
< words.txt tr -s ' ' '\n' | sort | uniq -c | sort -n -r | tr -s ' ' | awk '{print $2 " " $1}'
