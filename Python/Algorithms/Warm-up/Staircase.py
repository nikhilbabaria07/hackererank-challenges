#!/bin/python3


# Complete the staircase function below.
def staircase(n):
    for i in range(0, n, 1):
        for j in range(n - 1, i, -1):
            print(" ", end="");
        for k in range(i + 1, 0, -1):
            print("#", end="");
        print("")


if __name__ == '__main__':
    n = int(input())

    staircase(n)
