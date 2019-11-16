from pip._vendor.distlib.compat import raw_input


def arrangeArrays(a):
    a_len = len(a)

    if a_len == 0:
        return "Array consists no elements"

    # Sorting the array to properly fetch the first and last element of the array
    a.sort()

    # Initialized an empty array to be returned as a result
    new_arr = []

    # The loop is desingned to work like this
    # for e.g. A list with 6 elements
    # Loop 1: i =0 j = 5
    # Loop 2: i =1 j = 4
    # Loop 3: i =2 j = 3

    # for e.g. A list with 7 elements
    # Loop 1: i =0 j = 6
    # Loop 2: i =1 j = 5
    # Loop 3: i =2 j = 4
    # Loop 4: i =3 j = 3
    for i in range(0, a_len, 1):
        for j in range(a_len - 1 - i, i - 1, -1):
            if i != j:
                # Appending the last element which is also the biggest element.
                new_arr.append(a[j])
                # Appending the first element which is also the smallest element.
                new_arr.append(a[i])
            else:
                # Its the middle element(if number of elements in the list is odd) and can use either 'i' or 'j'.
                new_arr.append(a[j])
            break

    return new_arr


if __name__ == '__main__':
    # Read list as an input from STDIN
    str_a = raw_input("Enter elements of list here(separated by a space): \n").split(' ')
    a = [int(num) for num in str_a]

    # Print Output
    print(arrangeArrays(a))
