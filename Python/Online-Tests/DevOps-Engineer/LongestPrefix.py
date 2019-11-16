from pip._vendor.distlib.compat import raw_input


def longPref(a):
    # Sort the array to bring the shortest element at pos 0
    # We dont need to run the loop beyond the length of shortest element of the list
    a = sorted(a, key=len)

    longestPrefix = 0

    # Outer loop is to iterate through each character of the shortest element
    # Inner loop is used to compare each element in the list with outer loops prefix value
    # For e.g. Shortest element is 'App'
    # Loop 1 : comparing if each element have 'A'
    # Loop 2 : comparing if each element have 'Ap'
    # Loop 3 : comparing if each element have 'App'
    for i in range(len(a[0])):
        # Need flag value to stop direct break from the loop
        # Incrementing longestPrefix has to happen in outer loop only once the prefix matches all elements
        flag = False
        for j in range(len(a)):

            if (a[0][0:(i + 1)] == a[j][0:(i + 1)]):
                continue
            else:
                flag = True
                break
        if flag == False:
            longestPrefix += 1
        else:
            break

    if longestPrefix != 0:
        return (a[0][0:longestPrefix])
    else:
        return ("No matching prefix")


if __name__ == '__main__':
    # Read list as an input from STDIN
    a = raw_input("Enter elements of list here(separated by a space): \n").split(' ')

    # Print Output
    print(longPref(a))
