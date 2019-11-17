if __name__ == '__main__':
    n = int(input())
    arr = list(map(int, input().split()))
    arr.sort()

    # print(list(set(arr))[::-1][1])

    z = max(arr)

    while max(arr) == z:
        arr.remove(max(arr))

    print(max(arr))