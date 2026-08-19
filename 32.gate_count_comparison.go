# ============================================================
# CELL 32: GATE COUNT COMPARISON
# ============================================================

plt.figure(figsize=(10, 6))

plt.plot(
    final_results_df["Qubits"],
    final_results_df["Gate_Count_Original"],
    marker="o",
    label="Original Circuit"
)

plt.plot(
    final_results_df["Qubits"],
    final_results_df["Gate_Count_Optimized"],
    marker="s",
    label="Feedback Optimized Circuit"
)

plt.xlabel("Number of Qubits")
plt.ylabel("Gate Count")
plt.title("Circuit Gate Count Comparison")

plt.legend()
plt.grid(True)
plt.tight_layout()

plt.show()
