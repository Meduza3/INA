import os
import re
import pandas as pd
import matplotlib.pyplot as plt

def parse_topo_file(filepath):
    """
    Parsuje pojedynczy plik wyjściowy i zwraca słownik z informacjami.
    """
    data = {
        'Graph': os.path.basename(filepath),
        'IsAcyclic': False,
        'TopoOrder': None,
        'TopoLength': 0,
        'Time_us': 0
    }

    with open(filepath, 'r', encoding='utf-8') as file:
        lines = file.readlines()

    for i, line in enumerate(lines):
        line = line.strip()
        if line.startswith("Graf jest acykliczny"):
            data['IsAcyclic'] = True
        elif line.startswith("Porządek topologiczny wierzchołków:"):
            # Odczyt następnej linii z porządkiem topologicznym
            if i + 1 < len(lines):
                topo_order = lines[i + 1].strip().split()
                data['TopoOrder'] = topo_order
                data['TopoLength'] = len(topo_order)
        elif line.startswith("Czas wykonania programu:"):
            time_match = re.findall(r"([\d\.]+)(ms|µs)", line)
            if time_match:
                value, unit = time_match[0]
                value = float(value)
                if unit == 'ms':
                    data['Time_us'] = value * 1000  # Konwersja ms na µs
                elif unit == 'µs':
                    data['Time_us'] = value
            else:
                # Handle different time units if necessary
                data['Time_us'] = 0

    return data

def main():
    # Ścieżka do katalogu z plikami wyjściowymi
    input_dir = './results/results_topo'
    output_dir = './visualizations/visualizations_topo'

    # Tworzenie katalogu na wizualizacje, jeśli nie istnieje
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    # Listowanie wszystkich plików w katalogu input_dir
    files = [f for f in os.listdir(input_dir) if f.endswith('_topo.txt')]

    # Lista do przechowywania danych
    data_list = []

    for file in files:
        filepath = os.path.join(input_dir, file)
        data = parse_topo_file(filepath)
        data_list.append(data)

    # Tworzenie DataFrame
    df = pd.DataFrame(data_list)

    # Wykres 1: Liczba grafów acyklicznych vs. zawierających cykle
    acyclic_counts = df['IsAcyclic'].value_counts()
    plt.figure(figsize=(6,4))
    acyclic_counts.plot(kind='bar', color=['salmon', 'skyblue'])
    plt.title('Liczba Grafów Acyklicznych vs. Zawierających Cykle')
    plt.xlabel('Czy graf jest acykliczny')
    plt.ylabel('Liczba Grafów')
    plt.xticks([0,1], ['Nie', 'Tak'], rotation=0)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'acyclic_counts.png'))
    plt.close()

    # Wykres 2: Czas wykonania algorytmu dla każdego grafu
    plt.figure(figsize=(12,6))
    plt.bar(df['Graph'], df['Time_us'], color='green')
    plt.title('Czas Wykonania Algorytmu Sortowania Topologicznego')
    plt.xlabel('Graf')
    plt.ylabel('Czas (µs)')
    plt.xticks(rotation=90)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution.png'))
    plt.close()

    # Wykres 3: Długość porządku topologicznego dla grafów acyklicznych z n <= 200
    topo_df = df[(df['IsAcyclic']) & (df['TopoLength'] > 0)]
    if not topo_df.empty:
        plt.figure(figsize=(12,6))
        plt.bar(topo_df['Graph'], topo_df['TopoLength'], color='purple')
        plt.title('Długość Porządku Topologicznego dla Grafów Acyklicznych (n ≤ 200)')
        plt.xlabel('Graf')
        plt.ylabel('Długość Porządku Topologicznego')
        plt.xticks(rotation=90)
        plt.tight_layout()
        plt.savefig(os.path.join(output_dir, 'topo_length.png'))
        plt.close()
    else:
        print("Brak grafów acyklicznych z porządkiem topologicznym do wizualizacji.")

    # Opcjonalnie: Histogram czasu wykonania
    plt.figure(figsize=(8,6))
    plt.hist(df['Time_us'], bins=20, color='skyblue', edgecolor='black')
    plt.title('Rozkład Czasu Wykonania Algorytmu Sortowania Topologicznego')
    plt.xlabel('Czas (µs)')
    plt.ylabel('Liczba Grafów')
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution_histogram.png'))
    plt.close()

    print(f"Wizualizacje zostały zapisane w katalogu: {output_dir}")

if __name__ == "__main__":
    main()