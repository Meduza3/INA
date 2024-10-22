import os
import re
import pandas as pd
import matplotlib.pyplot as plt

def parse_scc_file(filepath):
    """
    Parsuje pojedynczy plik wyjściowy i zwraca słownik z informacjami.
    """
    data = {
        'Graph': os.path.basename(filepath),
        'Num_SCC': 0,
        'SCC_Sizes': [],
        'Time_us': 0
    }

    with open(filepath, 'r', encoding='utf-8') as file:
        lines = file.readlines()

    for line in lines:
        line = line.strip()
        if line.startswith("Liczba silnie spójnych składowych:"):
            match = re.findall(r"Liczba silnie spójnych składowych:\s*(\d+)", line)
            if match:
                data['Num_SCC'] = int(match[0])
        elif line.startswith("Składowa"):
            # Wyciąganie liczby wierzchołków w składowej
            match = re.findall(r"Składowa \d+:\s*(\d+)\s*wierzchołków?", line)
            if match:
                scc_size = int(match[0])
                data['SCC_Sizes'].append(scc_size)
        elif line.startswith("Czas wykonania algorytmu SCC:"):
            # Wyciąganie czasu i konwersja do µs
            time_match = re.findall(r"Czas wykonania algorytmu SCC:\s*([\d\.]+)(ms|µs)", line)
            if time_match:
                value, unit = time_match[0]
                value = float(value)
                if unit == 'ms':
                    time_us = value * 1000  # Konwersja ms na µs
                elif unit == 'µs':
                    time_us = value
                else:
                    time_us = 0  # Nieznana jednostka
                data['Time_us'] = time_us

    return data

def create_individual_scc_plots(df, output_dir):
    """
    Tworzy osobne wykresy słupkowe dla rozmiarów SCC każdego grafu.
    """
    # Tworzenie katalogu na indywidualne wykresy, jeśli nie istnieje
    individual_dir = os.path.join(output_dir, 'scc_individual')
    if not os.path.exists(individual_dir):
        os.makedirs(individual_dir)

    for index, row in df.iterrows():
        graph = row['Graph']
        scc_sizes = row['SCC_Sizes']
        num_scc = row['Num_SCC']

        # Przygotowanie danych do wykresu
        scc_labels = [f'Składowa {i+1}' for i in range(num_scc)]
        sizes = scc_sizes

        plt.figure(figsize=(8,6))
        plt.bar(scc_labels, sizes, color='teal')
        plt.title(f'Rozmiary Silnie Spójnych Składowych dla Grafu {graph}')
        plt.xlabel('Silnie Spójna Składowa')
        plt.ylabel('Liczba Wierzchołków')
        plt.tight_layout()

        # Nazwa pliku wykresu
        plot_filename = f'scc_sizes_{os.path.splitext(graph)[0]}.png'
        plt.savefig(os.path.join(individual_dir, plot_filename))
        plt.close()

def main():
    # Ścieżka do katalogu z plikami wyjściowymi
    input_dir = './results/results_skladowe'
    output_dir = './visualizations/skladowe'

    # Tworzenie katalogu na wizualizacje, jeśli nie istnieje
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    # Listowanie wszystkich plików w katalogu input_dir
    files = [f for f in os.listdir(input_dir) if f.endswith('.txt')]

    # Lista do przechowywania wszystkich danych
    all_scc_data = []

    for file in files:
        filepath = os.path.join(input_dir, file)
        scc_data = parse_scc_file(filepath)
        all_scc_data.append(scc_data)

    # Tworzenie DataFrame
    df = pd.DataFrame(all_scc_data)

    # Sprawdzenie, czy dane zostały poprawnie wczytane
    if df.empty:
        print("Brak danych do wizualizacji. Upewnij się, że katalog z plikami wyjściowymi nie jest pusty.")
        return

    # Konwersja czasu wykonania do float (jeśli nie jest)
    df['Time_us'] = pd.to_numeric(df['Time_us'], errors='coerce').fillna(0)

    # Wykres 1: Czas wykonania algorytmu SCC dla każdego grafu
    plt.figure(figsize=(12,6))
    plt.bar(df['Graph'], df['Time_us'], color='green')
    plt.title('Czas Wykonania Algorytmu SCC dla Każdego Grafu')
    plt.xlabel('Graf')
    plt.ylabel('Czas (µs)')
    plt.xticks(rotation=90)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution.png'))
    plt.close()

    # Przygotowanie danych do wykresu 2: Rozmiary SCC
    # Tworzymy listę słowników, gdzie każdy słownik to jedna SCC
    scc_records = []
    for index, row in df.iterrows():
        graph = row['Graph']
        for i, size in enumerate(row['SCC_Sizes'], start=1):
            scc_records.append({'Graph': graph, 'SCC_ID': f'Składowa {i}', 'Size': size})

    scc_df = pd.DataFrame(scc_records)

    # Wykres 2: Rozmiary silnie spójnych składowych dla każdego grafu (stacked bar)
    plt.figure(figsize=(12,6))
    # Pivotujemy dane tak, aby każda składowa miała swoją kolumnę
    pivot_df = scc_df.pivot(index='Graph', columns='SCC_ID', values='Size').fillna(0)
    pivot_df.plot(kind='bar', stacked=True, figsize=(12,6))
    plt.title('Rozmiary Silnie Spójnych Składowych dla Każdego Grafu')
    plt.xlabel('Graf')
    plt.ylabel('Liczba Wierzchołków')
    plt.legend(title='Silnie Spójna Składowa', bbox_to_anchor=(1.05, 1), loc='upper left')
    plt.xticks(rotation=90)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'scc_sizes.png'))
    plt.close()

    # Opcjonalnie: Wykres 3 - Histogram czasu wykonania
    plt.figure(figsize=(8,6))
    plt.hist(df['Time_us'], bins=20, color='skyblue', edgecolor='black')
    plt.title('Rozkład Czasu Wykonania Algorytmu SCC')
    plt.xlabel('Czas (µs)')
    plt.ylabel('Liczba Grafów')
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution_histogram.png'))
    plt.close()

    # Nowa Wizualizacja: Osobne Wykresy dla Rozmiarów SCC każdego Grafu
    create_individual_scc_plots(df, output_dir)

    print(f"Wizualizacje zostały zapisane w katalogu: {output_dir}")
    print(f"Indywidualne wykresy rozmiarów SCC zostały zapisane w katalogu: {os.path.join(output_dir, 'scc_individual')}")

if __name__ == "__main__":
    main()
