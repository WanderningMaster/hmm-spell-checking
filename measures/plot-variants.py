import pandas as pd
import matplotlib.pyplot as plt

# Load the data
file_path = "variants.txt"
data = pd.read_csv(file_path, sep=" ", header=None, names=["accuracy", "variants"])

# Create the plot
plt.figure(figsize=(10, 6))
plt.plot(data["accuracy"], data["variants"], marker='.', linestyle='', color='b')
plt.xlabel("Точність, %")
plt.ylabel("Максимальна кількість варіантів")
plt.grid(True)
plt.savefig('variants-plot.png', dpi=300)
