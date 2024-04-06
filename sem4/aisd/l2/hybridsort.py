liczba_porownan = 0
liczba_zastapien = 0

def main():
    input_data = input().split()  # Odczytaj dane jako linie tekstu ze standardowego wejścia
    length = int(input_data[0])
    table = list(map(int, input_data[1:]))
    boundary = 100

    print("TABLICA WEJSCIOWA")
    print(table)

    if(length <= boundary):
        insertionsort(table)
    else:
        quicksort(table, 0, len(table)-1)

    print("TABLICA WEJSCIOWA jeszcze raz")
    print(table_old)
    print("TABLICA WYJSCIOWA!:")
    print(table)
    print("LICZBA PRZESTAWIEN", liczba_zastapien)
    print("LICZBA POROWNAN", liczba_porownan)

def insertionsort(arr):
    global liczba_porownan, liczba_zastapien
    for i in range(1, len(arr)):
        print(i, arr)
        key = arr[i]
        j = i - 1
        while j >= 0 and key < arr[j]:
            liczba_porownan += 2
            arr[j + 1] = arr[j]
            liczba_zastapien += 1
            j -= 1
        arr[j + 1] = key
        liczba_zastapien += 1

def quicksort(arr, low, high):
    global liczba_porownan, liczba_zastapien
    print(arr)
    if low < high:
        liczba_porownan += 1
        pivot = partition(arr, low, high)
        quicksort(arr, low, pivot - 1)
        quicksort(arr, pivot + 1, high)

def partition(arr, low, high):
    global liczba_porownan, liczba_zastapien
    i = (low - 1)
    pivot = arr[high]
    for j in range(low, high):
        if arr[j] <= pivot:
            liczba_porownan += 1
            i = i + 1
            arr[i], arr[j] = arr[j], arr[i]
            liczba_zastapien += 2
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    liczba_zastapien += 1
    return (i+1)


if __name__ == "__main__":
    main()