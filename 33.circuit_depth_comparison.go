# ============================================================
# CELL 33: CIRCUIT DEPTH COMPARISON
# ============================================================

plt.figure(figsize=(10, 6))

plt.plot(
    final_results_df["Qubits"],
    final_results_df["Circuit_Depth_Original"],
    marker="o",
    label="Original Circuit"
)

plt.plot(
    final_results_df["Qubits"],
    final_results_df["Circuit_Depth_Optimized"],
    marker="s",
    label="Feedback Optimized Circuit"
)

plt.xlabel("Number of Qubits")
plt.ylabel("Circuit Depth")
plt.title("Circuit Depth Comparison")

plt.legend()
plt.grid(True)
plt.tight_layout()

plt.show()
