import sys

def main():
    input_data = input().split()  # Odczytaj dane jako linie tekstu ze standardowego wejścia
    length = int(input_data[0])
    table = list(map(int, input_data[1:]))
    boundary = 100
    if(length <= boundary):
        insertionsort(table)
    else:
        quicksort(table)

    print(table)

def insertionsort(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key

def quicksort(arr, low, high):
    if low < high:
        pivot = partition(arr, low, high)
        quicksort(arr, low, pivot - 1)
        quicksort(arr, pivot + 1, high)

def partition(arr, low, high):
    i = (low - 1)
    pivot = arr[high]
    for j in range(low, high):
        if arr[j] <= pivot:
            i = i + 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return (i+1)


if __name__ == "__main__":
    main()