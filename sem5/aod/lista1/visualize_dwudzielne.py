import os
import re
import pandas as pd
import matplotlib.pyplot as plt

def parse_bipartite_file(filepath):
    """
    Parsuje pojedynczy plik wyjściowy i zwraca słownik z informacjami.
    """
    data = {
        'Graph': os.path.basename(filepath),
        'IsBipartite': False,
        'V0_size': 0,
        'V1_size': 0,
        'Time_us': 0
    }

    with open(filepath, 'r', encoding='utf-8') as file:
        lines = file.readlines()

    for line in lines:
        line = line.strip()
        if line.startswith("Graf jest dwudzielny"):
            data['IsBipartite'] = True
        elif line.startswith("Podzbiór V0:"):
            v0 = line.split(":")[1].strip().split()
            data['V0_size'] = len(v0)
        elif line.startswith("Podzbiór V1:"):
            v1 = line.split(":")[1].strip().split()
            data['V1_size'] = len(v1)
        elif line.startswith("Czas wykonania algorytmu sprawdzania dwudzielności:"):
            time_match = re.findall(r"([\d\.]+)(ms|µs)", line)
            if time_match:
                value, unit = time_match[0]
                value = float(value)
                if unit == 'ms':
                    data['Time_us'] = value * 1000  # Konwersja ms na µs
                elif unit == 'µs':
                    data['Time_us'] = value

    return data

def main():
    # Ścieżka do katalogu z plikami wyjściowymi
    input_dir = './results/results_dwudzielne'
    output_dir = './visualizations/bipartite'

    # Tworzenie katalogu na wizualizacje, jeśli nie istnieje
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    # Listowanie wszystkich plików w katalogu input_dir
    files = [f for f in os.listdir(input_dir) if f.endswith('_bipartite.txt')]

    # Lista do przechowywania danych
    data_list = []

    for file in files:
        filepath = os.path.join(input_dir, file)
        data = parse_bipartite_file(filepath)
        data_list.append(data)

    # Tworzenie DataFrame
    df = pd.DataFrame(data_list)

    # Wykres 1: Liczba grafów dwudzielnych vs. niedwudzielnych
    bipartite_counts = df['IsBipartite'].value_counts()
    plt.figure(figsize=(6,4))
    bipartite_counts.plot(kind='bar', color=['skyblue', 'salmon'])
    plt.title('Liczba Grafów Dwudzielnych vs. Niedwudzielnych')
    plt.xlabel('Czy jest dwudzielny')
    plt.ylabel('Liczba Grafów')
    plt.xticks([0,1], ['Nie', 'Tak'], rotation=0)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'bipartite_counts.png'))
    plt.close()

    # Wykres 2: Czas wykonania algorytmu dla każdego grafu
    plt.figure(figsize=(10,6))
    plt.bar(df['Graph'], df['Time_us'], color='green')
    plt.title('Czas Wykonania Algorytmu Sprawdzania Dwudzielności')
    plt.xlabel('Graf')
    plt.ylabel('Czas (µs)')
    plt.xticks(rotation=90)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution.png'))
    plt.close()

    # Wykres 3: Rozmiary podzbiorów V0 i V1 dla grafów dwudzielnych
    bipartite_df = df[df['IsBipartite']]
    if not bipartite_df.empty:
        plt.figure(figsize=(10,6))
        width = 0.35
        plt.bar(bipartite_df['Graph'], bipartite_df['V0_size'], width, label='V0', color='orange')
        plt.bar(bipartite_df['Graph'], bipartite_df['V1_size'], width, bottom=bipartite_df['V0_size'], label='V1', color='purple')
        plt.title('Rozmiary Podzbiorów V0 i V1 dla Grafów Dwudzielnych')
        plt.xlabel('Graf')
        plt.ylabel('Liczba Wierzchołków')
        plt.legend()
        plt.xticks(rotation=90)
        plt.tight_layout()
        plt.savefig(os.path.join(output_dir, 'partition_sizes.png'))
        plt.close()
    else:
        print("Brak grafów dwudzielnych do wizualizacji rozmiarów podzbiorów.")

    print(f"Wizualizacje zostały zapisane w katalogu: {output_dir}")

if __name__ == "__main__":
    main()
