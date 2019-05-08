#!/bin/python3


# Complete the rotLeft function below.
def rotLeft(a, d):
    aLen = len(a)
    d = d % aLen

    aRot = [None] * aLen

    for i in range(aLen):
        aRot[i] = a[(i + d) % aLen]

    return aRot


if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')
    #
    # nd = input().split()
    #
    # n = int(nd[0])
    #
    # d = int(nd[1])
    #
    # a = list(map(int, input().rstrip().split()))

    a = [1, 2, 3, 5, 7, 8]
    d = 2001

    result = rotLeft(a, d)

    # fptr.write(' '.join(map(str, result)))
    # fptr.write('\n')
    #
    # fptr.close()

    print(result)
