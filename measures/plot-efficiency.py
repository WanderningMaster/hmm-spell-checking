import pandas as pd
import matplotlib.pyplot as plt

# Load the data
file_path = "speed_sync.txt"
data = pd.read_csv(file_path, sep=" ", header=None, names=["time", "variants"])

# Create the plot
plt.figure(figsize=(10, 6))
plt.plot(data["time"], data["variants"], marker='.', linestyle='', color='b', label='Синхронне виконання')

file_path = "speed.txt"
data = pd.read_csv(file_path, sep=" ", header=None, names=["time", "variants"])
plt.plot(data["time"], data["variants"], marker='.', linestyle='', color='r', label='Паралельне виконання')

plt.xlabel("Час виконання, мс")
plt.ylabel("Максимальна кількість варіантів")
plt.grid(True)
plt.savefig('time-sync-compare-plot.png', dpi=300)
