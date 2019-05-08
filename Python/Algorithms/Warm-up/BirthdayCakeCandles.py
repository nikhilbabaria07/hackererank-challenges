#!/bin/python3

import os


# Complete the birthdayCakeCandles function below.
def birthdayCakeCandles(ar):
    ar.sort();
    arLen = len(ar);
    maxVal = ar[arLen - 1];
    tallestCandles = 0;

    for i in range(arLen):
        if ar[i] == maxVal:
            tallestCandles += 1

    return tallestCandles


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    ar_count = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = birthdayCakeCandles(ar)

    fptr.write(str(result) + '\n')

    fptr.close()
