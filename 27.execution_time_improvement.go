# ============================================================
# CELL 27: EXECUTION TIME IMPROVEMENT
# ============================================================

execution_comparison_df[
    "Execution_Time_Improvement_%"
] = (
    (
        execution_comparison_df["Median_Time_sec_Original"]
        -
        execution_comparison_df["Median_Time_sec_Optimized"]
    )
    /
    execution_comparison_df["Median_Time_sec_Original"]
) * 100

display(execution_comparison_df)
