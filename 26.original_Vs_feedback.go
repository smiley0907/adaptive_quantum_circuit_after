# ============================================================
# CELL 26: ORIGINAL VS FEEDBACK OPTIMIZED
# ============================================================

execution_comparison_df = pd.merge(
    original_df[
        [
            "Qubits",
            "Median_Time_sec",
            "Mean_Time_sec",
            "Std_Time_sec"
        ]
    ],
    optimized_df[
        [
            "Qubits",
            "Median_Time_sec",
            "Mean_Time_sec",
            "Std_Time_sec"
        ]
    ],
    on="Qubits",
    suffixes=(
        "_Original",
        "_Optimized"
    )
)

display(execution_comparison_df)
