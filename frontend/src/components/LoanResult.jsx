import React from "react";

const LoanResult = ({ result }) => {
  if (!result) {
    return <p>No loan result available.</p>;
  }

  const decision = result.decision || {}; // Ensure decision object exists

  return (
    <div>
      <h2 className="text-2xl font-semibold mt-4 mb-4">Loan Application Result</h2>
      <p className="text-lg">
        {decision.Approved ? (
          <span className="text-green-500">Approved</span>
        ) : (
          <span className="text-red-500">Not Approved</span>
        )}
      </p>
      <p className="text-sm">Message: {decision.Message}</p>
      <p className="text-sm">Pre-assessment value: {result.preAssessment}</p>
      <p className="text-sm">Approval status: {result.approvalStatus}</p>
    </div>
  );
};

export default LoanResult;
