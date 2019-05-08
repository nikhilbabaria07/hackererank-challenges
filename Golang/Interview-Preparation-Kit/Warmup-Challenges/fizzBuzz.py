#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'netVirtaCheck' function below.
#
# The function is expected to return a STRING.
# The function accepts INTEGER_ARRAY inputArray as parameter.
#

def fizzBuzz(n):
    if (n%4==0):
        if (n%5==0):
            return "netVirta"
        else:
            return "net"
    elif (n%5==0):
        return "Virta"
    else:
        return str(n)


def addArrVal(inputArray):
    cntr = 0
    for i in inputArray:
        cntr = cntr + i
    return cntr


def netVirtaCheck(inputArray):
    arrLen = inputArray.__len__;
    arrSum = addArrVal(inputArray);
    arrAvg = int(round(arrSum/arrLen))
    return fizzBuzz(arrAvg)


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    inputArray_count = int(input().strip())

    inputArray = []

    for _ in range(inputArray_count):
        inputArray_item = int(input().strip())
        inputArray.append(inputArray_item)

    result = netVirtaCheck(inputArray)

    fptr.write(result + '\n')

    fptr.close()
