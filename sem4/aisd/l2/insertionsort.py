liczba_porownan = 0
liczba_zastapien = 0

def main():
    input_data = input().split()  # Odczytaj dane jako linie tekstu ze standardowego wejścia
    length = int(input_data[0])
    table = list(map(int, input_data[1:]))
    print("TABLICA WEJSCIOWA")
    print(table)
    table_old = list(map(int, input_data[1:]))
    insertionsort(table)
    
    
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


if __name__ == "__main__":
    main()


