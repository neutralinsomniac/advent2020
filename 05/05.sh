#!/bin/sh
# did this in bed on my phone. programming with a keyboard is no bueno
cat input |sed -e 's/F/0/g;s/L/0/g;s/R/1/g;s/B/1/g' |sort -n |python -c 'import fileinput;l=[int(n,2) for n in [line.strip() for line in fileinput.input()]]; print(l[-1], [x+1 for i, x in enumerate(l[:-1]) if l[i] != l[i+1]-1])'
