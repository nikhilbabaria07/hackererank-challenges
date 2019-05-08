#!/bin/python3


# Complete the minimumBribes function below.
def minimumBribes(q):
    qLen = len(q)
    chaosFlag = False
    bribeCntr = 0

    for i in range(qLen - 1):
        flag = True
        for j in range(qLen - 1):
            if (q[j] > j + 3):
                chaosFlag = True
                break
            if (q[j] > q[j + 1]):
                temp = q[j + 1]
                q[j + 1] = q[j]
                q[j] = temp
                bribeCntr += 1
        for k in range(qLen):
            if ((k + 1) % (qLen + 1) != q[k]):
                flag = False
        if (flag == True):
            break

    if (chaosFlag == False):
        print(bribeCntr)
    else:
        print("Too chaotic")


if __name__ == '__main__':
    # t = int(input())
    #
    # for t_itr in range(t):
    #     n = int(input())
    #
    #     q = list(map(int, input().rstrip().split()))

    # q = [1, 2, 3, 4, 5, 6, 7, 10, 8, 9, 11]

    q = [1, 5, 4, 2, 3, 7, 8, 6]
    minimumBribes(q)
