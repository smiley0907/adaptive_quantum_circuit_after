# ============================================================
# CELL 29: FINAL RESULTS FOR PAPER
# ============================================================

paper_results = final_results_df[
    [
        "Qubits",
        "Gate_Count_Original",
        "Gate_Count_Optimized",
        "Circuit_Depth_Original",
        "Circuit_Depth_Optimized",
        "Gate_Reduction_%",
        "Depth_Reduction_%",
        "Median_Time_sec_Original",
        "Median_Time_sec_Optimized",
        "Execution_Time_Improvement_%"
    ]
].copy()

# Round percentages and times for presentation
paper_results["Gate_Reduction_%"] = (
    paper_results["Gate_Reduction_%"].round(2)
)

paper_results["Depth_Reduction_%"] = (
    paper_results["Depth_Reduction_%"].round(2)
)

paper_results["Median_Time_sec_Original"] = (
    paper_results["Median_Time_sec_Original"].round(6)
)

paper_results["Median_Time_sec_Optimized"] = (
    paper_results["Median_Time_sec_Optimized"].round(6)
)

paper_results["Execution_Time_Improvement_%"] = (
    paper_results["Execution_Time_Improvement_%"].round(2)
)

display(paper_results)
