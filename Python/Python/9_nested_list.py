if __name__ == '__main__':

    n = int(input())

    arr = [ [input(), float(input())] for _ in range(n) ]

    # for _ in range(int(input())):
    #     name = input()
    #     score = float(input())
    #     arr.append( [ name, score ] )

    min_val = min( [ i[1] for i in arr ] )
    # print(min_val)

    # print( [ i for i in arr if i[1] == min_val ][0] )
    while min_val == min( [ i[1] for i in arr ] ):
        arr.remove( [ i for i in arr if i[1] == min_val ][0] )

    # print(arr)
    # print( [ i[0] for i in arr ] )

    new_min_val = min( [ i[1] for i in arr ] )
    # print(new_min_val)

    res = [ i[0] for i in arr if new_min_val == i[1] ]
    # print(res)
    
    res.sort()
    # print(res)

    for i in res:
        print(i)
