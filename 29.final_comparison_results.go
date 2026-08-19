# ============================================================
# CELL 29: FINAL EXPERIMENTAL COMPARISON
# ============================================================

final_results_df = pd.merge(
    final_structure_df,
    execution_comparison_df[
        [
            "Qubits",
            "Median_Time_sec_Original",
            "Median_Time_sec_Optimized",
            "Execution_Time_Improvement_%"
        ]
    ],
    on="Qubits"
)

display(final_results_df)
