import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import os

def load_data(file_name):
    df = pd.read_csv(f"data/{file_name}.txt", delimiter='\t', header=None, names=['n', 'comp', 'swaps'])
    return df.groupby('n').mean().reset_index()
        
grouped_hybrid = load_data("hybrid")
grouped_insertion = load_data("insertion")
grouped_quick = load_data("quick")
grouped_merge = load_data("merge")
grouped_custom = load_data("custom")
grouped_dualpivot = load_data("dualpivot")

plt.figure(figsize=(10, 6))
plt.plot(grouped_hybrid['n'], grouped_hybrid['comp'], color='red', label=f"hybrid")
plt.plot(grouped_insertion['n'], grouped_insertion['comp'], color='blue', label=f"insertion")
plt.plot(grouped_quick['n'], grouped_quick['comp'], color='green', label=f"quick")
plt.plot(grouped_merge['n'], grouped_merge['comp'], color='orange', label=f"merge")
plt.plot(grouped_custom['n'], grouped_custom['comp'], color='purple', label=f"custom")
plt.plot(grouped_dualpivot['n'], grouped_dualpivot['comp'], color='pink', label=f"dual pivot")

plt.title(f"Comparisons")
plt.xlabel('Input Size')
plt.ylabel('Count')
plt.grid(True)
plt.legend()
plot_filename = os.path.join('plots', f"Comparisons.png")
plt.savefig(plot_filename)
plt.close()

plt.figure(figsize=(10, 6))
plt.plot(grouped_hybrid['n'], grouped_hybrid['swaps'], color='red', label=f"hybrid")
plt.plot(grouped_insertion['n'], grouped_insertion['swaps'], color='blue', label=f"insertion")
plt.plot(grouped_quick['n'], grouped_quick['swaps'], color='green', label=f"quick")
plt.plot(grouped_merge['n'], grouped_merge['swaps'], color='orange', label=f"merge")
plt.plot(grouped_custom['n'], grouped_custom['swaps'], color='purple', label=f"custom")
plt.plot(grouped_dualpivot['n'], grouped_dualpivot['swaps'], color='pink', label=f"dual pivot")

plt.title(f"Swaps")
plt.xlabel('Input Size')
plt.ylabel('Count')
plt.grid(True)
plt.legend()
plot_filename = os.path.join('plots', f"Swaps.png")
plt.savefig(plot_filename)
plt.close()

plt.figure(figsize=(10, 6))
plt.plot(grouped_hybrid['n'], (grouped_hybrid['comp']/grouped_hybrid['n']), color='red', label=f"hybrid")
plt.plot(grouped_insertion['n'], (grouped_insertion['comp']/grouped_insertion['n']), color='blue', label=f"insertion")
plt.plot(grouped_quick['n'], (grouped_quick['comp']/grouped_quick['n']), color='green', label=f"quick")
plt.plot(grouped_merge['n'], (grouped_merge['comp']/grouped_merge['n']), color='orange', label=f"merge")
plt.plot(grouped_custom['n'], (grouped_custom['comp']/grouped_custom['n']), color='purple', label=f"custom")
plt.plot(grouped_dualpivot['n'], (grouped_dualpivot['comp']/grouped_dualpivot['n']), color='pink', label=f"dual pivot")

plt.title(f"Comparisons / n")
plt.xlabel('Input Size')
plt.ylabel('Count')
plt.grid(True)
plt.legend()
plot_filename = os.path.join('plots', f"Comparisonsovern.png")
plt.savefig(plot_filename)
plt.close()

plt.figure(figsize=(10, 6))
plt.plot(grouped_hybrid['n'], (grouped_hybrid['swaps']/grouped_hybrid['n']), color='red', label=f"hybrid")
plt.plot(grouped_insertion['n'], (grouped_insertion['swaps']/grouped_insertion['n']), color='blue', label=f"insertion")
plt.plot(grouped_quick['n'], (grouped_quick['swaps']/grouped_quick['n']), color='green', label=f"quick")
plt.plot(grouped_merge['n'], (grouped_merge['swaps']/grouped_merge['n']), color='orange', label=f"merge")
plt.plot(grouped_custom['n'], (grouped_custom['swaps']/grouped_custom['n']), color='purple', label=f"custom")
plt.plot(grouped_dualpivot['n'], (grouped_dualpivot['swaps']/grouped_dualpivot['n']), color='pink', label=f"dual pivot")

plt.title(f"Swaps / n")
plt.xlabel('Input Size')
plt.ylabel('Count')
plt.grid(True)
plt.legend()
plot_filename = os.path.join('plots', f"Swapsovern.png")
plt.savefig(plot_filename)
plt.close()

