# ============================================================
# CELL 24: FEEDBACK OPTIMIZED RESULTS
# ============================================================

optimized_df = pd.DataFrame(
    optimized_results
)

display(
    optimized_df[
        [
            "Qubits",
            "Gate_Count",
            "Circuit_Depth",
            "Median_Time_sec",
            "Mean_Time_sec",
            "Std_Time_sec"
        ]
    ]
)
