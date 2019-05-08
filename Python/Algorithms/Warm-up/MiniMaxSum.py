#!/bin/python3


# Complete the miniMaxSum function below.
def miniMaxSum(arr):
    maxSum = 0;
    minSum = 0;
    arrLen = len(arr);

    arr.sort()

    for i in range(arrLen):
        if (i == arrLen - 1):
            minSum += arr[i];
            continue;
        elif (i == 0):
            maxSum += arr[i];
            continue;
        else:
            minSum += arr[i];
            maxSum += arr[i];

    print(maxSum, minSum)


if __name__ == '__main__':
    arr = list(map(int, input().rstrip().split()))

    miniMaxSum(arr)
