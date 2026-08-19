# ============================================================
# CELL 31: EXECUTION TIME COMPARISON
# ============================================================

plt.figure(figsize=(10, 6))

plt.plot(
    execution_comparison_df["Qubits"],
    execution_comparison_df["Median_Time_sec_Original"],
    marker="o",
    label="Original Circuit"
)

plt.plot(
    execution_comparison_df["Qubits"],
    execution_comparison_df["Median_Time_sec_Optimized"],
    marker="s",
    label="Feedback Optimized Circuit"
)

plt.xlabel("Number of Qubits")
plt.ylabel("Median Execution Time (s)")
plt.title(
    "Original and Feedback Optimized Circuit Performance"
)

plt.legend()
plt.grid(True)
plt.tight_layout()

plt.show()
