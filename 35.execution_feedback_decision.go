# ============================================================
# CELL 35: EXECUTION FEEDBACK DECISION
# ============================================================

feedback_decisions = []

for _, row in final_results_df.iterrows():

    if row["Execution_Time_Improvement_%"] > 0:

        decision = "RETAIN OPTIMIZED CIRCUIT"

    else:

        decision = "RETAIN ORIGINAL CIRCUIT"

    feedback_decisions.append({
        "Qubits": int(row["Qubits"]),
        "Execution_Time_Improvement_%":
            row["Execution_Time_Improvement_%"],
        "Decision": decision
    })

feedback_decision_df = pd.DataFrame(
    feedback_decisions
)

display(feedback_decision_df)
