import sys

def main():
    if len(sys.argv) != 3:
        print("Usage: python insertionsort.py LENGTH TABLE")
        sys.exit(1)
    length = int(sys.argv[1])
    table = list(map(int, sys.argv[2].split()))
    insertion_sort(table)
    print(table)


def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key


if __name__ == "__main__":
    main()


