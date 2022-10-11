#!/usr/bin/env python
import sys


for line in sys.stdin:
    data, width = line.split(",")
    width = int(width)
    print("{:<30}{:=<{width}}".format(data, "", width=width))
