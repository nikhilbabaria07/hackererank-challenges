if __name__ == '__main__':

    n = int(input())
    s = set(map(int, input().split()))
    N = int(input())

    res = 0

    for i in range(N):
        opr = input().split()

        if opr[0] == "pop":
            s.pop()
        elif opr[0] == "discard":
            s.discard(int(opr[1]))
        elif opr[0] == "remove":
            s.remove(int(opr[1]))

    for i in s:
        res += i

    print(res)

