#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the diagonalDifference function below.
def diagonalDifference(arr):

    arrLen = len(arr);
    lDiag = 0;
    rDiag = 0;

    for i in range(0,arrLen,1):
        lDiag += arr[i][i];

    for i in range(0,arrLen,1):
        for j in range(arrLen - 1 - i,-1,-1):
            rDiag += arr[i][j];
            break;

    return abs(lDiag-rDiag)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = []

    for _ in range(n):
        arr.append(list(map(int, input().rstrip().split())))

    result = diagonalDifference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
