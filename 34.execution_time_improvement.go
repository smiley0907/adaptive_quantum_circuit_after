# ============================================================
# CELL 34: EXECUTION TIME IMPROVEMENT
# ============================================================

plt.figure(figsize=(10, 6))

plt.bar(
    final_results_df["Qubits"].astype(str),
    final_results_df["Execution_Time_Improvement_%"]
)

plt.axhline(
    0,
    linewidth=1
)

plt.xlabel("Number of Qubits")
plt.ylabel("Execution Time Improvement (%)")
plt.title(
    "Execution Feedback Optimization Improvement"
)

plt.grid(
    axis="y"
)

plt.tight_layout()
plt.show()
