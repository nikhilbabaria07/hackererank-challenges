#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the minimumAbsoluteDifference function below.
def minimumAbsoluteDifference(arr):

    arrLen = len(arr)
    arr.sort()
    minAbsDiff = abs(arr[1] - arr[0])

    for i in range(1,arrLen):
        absDiff = abs(arr[i] - arr[i-1])

        if absDiff < minAbsDiff:
            minAbsDiff = absDiff

    return minAbsDiff

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    result = minimumAbsoluteDifference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
