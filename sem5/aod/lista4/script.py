import pandas as pd
import matplotlib.pyplot as plt

# Load data from CSV
csv_file = "data.csv"  # Replace with the path to your CSV file
data = pd.read_csv(csv_file)

# Plot graph for 'val'
plt.figure(figsize=(8, 6))
plt.plot(data['k'], data['val'], marker='o', color='blue')
plt.title('Średnia wielkość maksymalnego przepływu w zależności od k')
plt.xlabel('k')
plt.ylabel('val')
plt.grid(True)
plt.savefig('max_flow_vs_k.png')
plt.close()

# Plot graph for 'time'
plt.figure(figsize=(8, 6))
plt.plot(data['k'], data['time'], marker='o', color='green')
plt.title('Czas wykonania algorytmu w zależności od k')
plt.xlabel('k')
plt.ylabel('Czas (s)')
plt.grid(True)
plt.savefig('execution_time_vs_k.png')
plt.close()

# Plot graph for 'bfs'
plt.figure(figsize=(8, 6))
plt.plot(data['k'], data['bfs'], marker='o', color='red')
plt.title('Liczba ścieżek powiększających w BFS w zależności od k')
plt.xlabel('k')
plt.ylabel('Liczba ścieżek powiększającyc BFS')
plt.grid(True)
plt.savefig('bfs_calls_vs_k.png')
plt.close()
