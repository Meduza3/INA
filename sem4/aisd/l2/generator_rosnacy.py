import random
import sys

def main():
    liczba_kluczy = int(sys.argv[1])

    min_klucz = 0
    max_klucz = 2*liczba_kluczy - 1
    losowe_klucze = [random.randint(min_klucz, max_klucz) for _ in range(liczba_kluczy)]
    sorted(losowe_klucze)
    print(liczba_kluczy, ' '.join(map(str, losowe_klucze)))


if __name__ == "__main__":
    main()