#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'GetDistance' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. BOOLEAN isClosest
#  2. INTEGER_ARRAY sortedArray
#

def GetDistance(isClosest, sortedArray):
    arrLen=sortedArray.__len__
    print("clostest")
    if isClosest == "True":
        i=1
        closest=sortedArray[1]-sortedArray[0]
        while(i<arrLen-1):
            diff=(sortedArray[i+1]-sortedArray[i])
            if diff < closest:
                closest=diff
            i=i+1
        return closest
    else:
        print("Furtherst")
        return (sortedArray[arrLen-1]-sortedArray[0])


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    isClosest = int(input().strip()) != 0

    sortedArray_count = int(input().strip())

    sortedArray = []

    for _ in range(sortedArray_count):
        sortedArray_item = int(input().strip())
        sortedArray.append(sortedArray_item)

    result = GetDistance(isClosest, sortedArray)

    fptr.write(str(result) + '\n')

    fptr.close()
