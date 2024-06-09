import pandas as pd
import matplotlib.pyplot as plt

# Load the data
file_path = "lambda.txt"
data = pd.read_csv(file_path, sep=" ", header=None, names=["accuracy", "lambda"])

# Create the plot
plt.figure(figsize=(10, 6))
plt.plot(data["accuracy"], data["lambda"], marker='.', linestyle='', color='b')
plt.xlabel("Точність, %")
plt.ylabel("λ")
plt.grid(True)
plt.savefig('lambda-plot.png', dpi=300)
