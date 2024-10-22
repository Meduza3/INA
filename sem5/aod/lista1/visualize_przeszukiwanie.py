import os
import re
import pandas as pd
import matplotlib.pyplot as plt

import os
import re
import pandas as pd
import matplotlib.pyplot as plt

def parse_traversal_file(filepath):
    """
    Parsuje pojedynczy plik wyjściowy i zwraca listę słowników z informacjami.
    Każdy słownik reprezentuje jedno uruchomienie algorytmu (BFS lub DFS).
    """
    traversal_data = []
    print(f"Przetwarzanie pliku: {filepath}")
    try:
        with open(filepath, 'r', encoding='utf-8') as file:
            lines = file.readlines()
    except Exception as e:
        print(f"Błąd podczas otwierania pliku {filepath}: {e}")
        return traversal_data

    current_traversal = {}
    i = 0
    while i < len(lines):
        line = lines[i].strip()
        
        # Rozpoczęcie nowego bloku algorytmu
        if line.startswith("Algorytm"):
            # Resetowanie danych dla nowego algorytmu
            current_traversal = {
                'Algorithm': '',
                'Graph': '',
                'VisitOrder': [],
                'ParentTree': {},
                'Time_us': 0,
                'VisitOrder_length': 0
            }
            # Wyciągnięcie typu algorytmu i nazwy pliku
            match = re.match(r"Algorytm (\w+) dla pliku (.+)", line)
            if match:
                current_traversal['Algorithm'] = match.group(1)
                current_traversal['Graph'] = match.group(2)
                print(f"  Znaleziono algorytm: {current_traversal['Algorithm']}, graf: {current_traversal['Graph']}")
            else:
                print(f"  Nie udało się dopasować algorytmu i grafu w linii: {line}")
            i += 1
            continue
        
        # Kolejność odwiedzania wierzchołków
        if line.startswith("Kolejność odwiedzania wierzchołków:"):
            if i + 1 < len(lines):
                visit_order = lines[i + 1].strip().split()
                current_traversal['VisitOrder'] = visit_order
                current_traversal['VisitOrder_length'] = len(visit_order)
                print(f"  Kolejność odwiedzania: {visit_order}")
                i += 2
                continue
        
        # Drzewo przeszukiwania
        if line.startswith("Drzewo przeszukiwania"):
            # Czytanie kolejnych linii z drzewem przeszukiwania
            i += 1
            while i < len(lines):
                parent_line = lines[i].strip()
                if parent_line == '':
                    break
                parent_match = re.match(r"(\d+):\s*(\w+)", parent_line)
                if parent_match:
                    node = int(parent_match.group(1))
                    parent = parent_match.group(2)
                    if parent.lower() == 'nil':
                        current_traversal['ParentTree'][node] = None
                    else:
                        try:
                            current_traversal['ParentTree'][node] = int(parent)
                        except ValueError:
                            current_traversal['ParentTree'][node] = parent  # Jeśli parent nie jest liczbą
                    print(f"    Wierzchołek: {node}, Rodzic: {current_traversal['ParentTree'][node]}")
                else:
                    # Jeśli linia nie pasuje do formatu wierzchołek: rodzic, przerywamy blok drzewa
                    print(f"    Nie udało się dopasować wiersza drzewa przeszukiwania: {parent_line}")
                    break  # Przerywamy przetwarzanie drzewa
                i += 1
            continue
        
        # Czas wykonania algorytmu
        if line.startswith("Czas wykonania"):
            time_match = re.findall(r"Czas wykonania \w+:\s*([\d\.]+)(ns|µs|ms)", line)
            if time_match:
                value, unit = time_match[0]
                value = float(value)
                if unit == 'ms':
                    time_us = value * 1000  # Konwersja ms na µs
                elif unit == 'µs':
                    time_us = value
                elif unit == 'ns':
                    time_us = value / 1000  # Konwersja ns na µs
                else:
                    time_us = 0  # Nieznana jednostka
                current_traversal['Time_us'] = time_us
                print(f"  Czas wykonania: {time_us} µs")
            else:
                print(f"  Nie udało się dopasować czasu w linii: {line}")
            # Dodanie zebranych danych do listy
            traversal_data.append(current_traversal)
            print(f"  Dodano dane: {current_traversal}")
            i += 1
            continue
        
        i += 1  # Przejście do następnej linii

    return traversal_data


def create_individual_traversal_plots(df, output_dir):
    """
    Tworzy osobne wykresy liniowe dla kolejności odwiedzania wierzchołków każdego grafu.
    """
    # Tworzenie katalogu na indywidualne wykresy, jeśli nie istnieje
    individual_dir = os.path.join(output_dir, 'individual_traversals')
    if not os.path.exists(individual_dir):
        os.makedirs(individual_dir)

    for index, row in df.iterrows():
        graph = row['Graph']
        algorithm = row['Algorithm']
        visit_order = row['VisitOrder']
        visit_order_length = row['VisitOrder_length']

        # Przygotowanie danych do wykresu
        try:
            vertices = list(map(int, visit_order))
        except ValueError:
            print(f"  Błąd konwersji wierzchołków dla grafu {graph}, algorytm {algorithm}.")
            vertices = []
        indices = list(range(1, len(vertices) + 1))

        plt.figure(figsize=(10,6))
        plt.plot(indices, vertices, marker='o', linestyle='-', color='blue')
        plt.title(f'Kolejność Odwiedzania Wierzchołków dla {algorithm} - Graf {graph}')
        plt.xlabel('Krok')
        plt.ylabel('Wierzchołek')
        plt.grid(True)
        plt.tight_layout()

        # Nazwa pliku wykresu
        plot_filename = f'traversal_{algorithm}_{os.path.splitext(graph)[0]}.png'
        plt.savefig(os.path.join(individual_dir, plot_filename))
        plt.close()
        print(f"  Zapisano indywidualny wykres: {plot_filename}")

def main():
    # Ścieżka do katalogu z plikami wyjściowymi
    input_dir = './results/results_przeszukiwanie'
    output_dir = './visualizations/przeszukiwanie'

    # Tworzenie katalogu na wizualizacje, jeśli nie istnieje
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    # Tworzenie katalogu na indywidualne wykresy
    individual_dir = os.path.join(output_dir, 'individual_traversals')
    if not os.path.exists(individual_dir):
        os.makedirs(individual_dir)

    # Listowanie wszystkich plików w katalogu input_dir
    files = [f for f in os.listdir(input_dir) if f.endswith('.txt')]

    print(f"Znaleziono {len(files)} plików w katalogu {input_dir}:")
    for f in files:
        print(f"- {f}")

    if not files:
        print("Brak plików .txt w katalogu input_dir. Upewnij się, że ścieżka jest poprawna.")
        return

    # Lista do przechowywania wszystkich danych
    all_traversals = []

    for file in files:
        filepath = os.path.join(input_dir, file)
        traversal_runs = parse_traversal_file(filepath)
        all_traversals.extend(traversal_runs)

    # Tworzenie DataFrame
    df = pd.DataFrame(all_traversals)

    print("\nPrzykładowe dane z DataFrame:")
    print(df.head())

    # Sprawdzenie, czy dane zostały poprawnie wczytane
    if df.empty:
        print("Brak danych do wizualizacji. Upewnij się, że pliki wyjściowe zawierają odpowiednie dane.")
        return

    # Konwersja czasu wykonania do float (jeśli nie jest)
    df['Time_us'] = pd.to_numeric(df['Time_us'], errors='coerce').fillna(0)

    # Wykres 1: Liczba uruchomień BFS vs. DFS
    algo_counts = df['Algorithm'].value_counts()
    plt.figure(figsize=(6,4))
    algo_counts.plot(kind='bar', color=['skyblue', 'salmon'])
    plt.title('Liczba Uruchomień Algorytmów BFS vs. DFS')
    plt.xlabel('Algorytm')
    plt.ylabel('Liczba Uruchomień')
    plt.xticks(rotation=0)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'algorithm_counts.png'))
    plt.close()
    print("Zapisano wykres: algorithm_counts.png")

    # Wykres 2: Czas wykonania algorytmu dla każdego grafu, rozróżniając BFS i DFS
    plt.figure(figsize=(12,6))
    algorithms = df['Algorithm'].unique()
    colors = {'BFS': 'green', 'DFS': 'purple'}
    for algo in algorithms:
        subset = df[df['Algorithm'] == algo]
        plt.bar(subset['Graph'], subset['Time_us'], label=algo, color=colors.get(algo, 'gray'))
    plt.title('Czas Wykonania Algorytmów BFS i DFS')
    plt.xlabel('Graf')
    plt.ylabel('Czas (µs)')
    plt.legend()
    plt.xticks(rotation=90)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution.png'))
    plt.close()
    print("Zapisano wykres: time_execution.png")

    # Wykres 3: Histogram czasu wykonania algorytmów BFS i DFS
    plt.figure(figsize=(8,6))
    for algo in algorithms:
        subset = df[df['Algorithm'] == algo]
        plt.hist(subset['Time_us'], bins=20, alpha=0.5, label=algo, color=colors.get(algo, 'gray'))
    plt.title('Rozkład Czasu Wykonania Algorytmów BFS i DFS')
    plt.xlabel('Czas (µs)')
    plt.ylabel('Liczba Uruchomień')
    plt.legend()
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'time_execution_histogram.png'))
    plt.close()
    print("Zapisano wykres: time_execution_histogram.png")

    # Wykres 4: Średni czas wykonania algorytmów BFS i DFS
    avg_times = df.groupby('Algorithm')['Time_us'].mean()
    plt.figure(figsize=(6,4))
    avg_times.plot(kind='bar', color=[colors.get(algo, 'gray') for algo in avg_times.index])
    plt.title('Średni Czas Wykonania Algorytmów BFS i DFS')
    plt.xlabel('Algorytm')
    plt.ylabel('Średni Czas (µs)')
    plt.xticks(rotation=0)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'average_time_execution.png'))
    plt.close()
    print("Zapisano wykres: average_time_execution.png")

    # Wykres 5: Długość kolejności odwiedzania wierzchołków dla BFS i DFS
    plt.figure(figsize=(12,6))
    for algo in algorithms:
        subset = df[df['Algorithm'] == algo]
        plt.bar(subset['Graph'], subset['VisitOrder_length'], label=algo, color=colors.get(algo, 'gray'))
    plt.title('Długość Kolejności Odwiedzania Wierzchołków dla BFS i DFS')
    plt.xlabel('Graf')
    plt.ylabel('Długość Kolejności')
    plt.legend()
    plt.xticks(rotation=90)
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, 'visit_order_length.png'))
    plt.close()
    print("Zapisano wykres: visit_order_length.png")

    # Nowa Wizualizacja: Osobne Wykresy dla Kolejności Odwiedzania Wierzchołków każdego Grafu
    create_individual_traversal_plots(df, output_dir)
    print("Zapisano osobne wykresy kolejności odwiedzania wierzchołków w katalogu: individual_traversals")

    print(f"\nWizualizacje zostały zapisane w katalogu: {output_dir}")


if __name__ == "__main__":
    main()
