import { useState } from "react";
import LoanForm from "./components/LoanForm";
import BalanceSheet from "./components/BalanceSheet";
import LoanResult from "./components/LoanResult";

const App = () => {
  const [balanceSheet, setBalanceSheet] = useState([]);
  const [loanResult, setLoanResult] = useState(null);

  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;

  const fetchBalanceSheet = async (accountingProvider) => {
    try {
      const response = await fetch(`${apiBaseUrl}/balance-sheet/${accountingProvider}`);
      const data = await response.json();
      setBalanceSheet(data);
    } catch (error) {
      console.error("Error fetching balance sheet:", error);
    }
  };

  const submitApplication = async (formData) => {
    try {
      const response = await fetch(`${apiBaseUrl}/loan`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });
      const data = await response.json();
      setLoanResult(data);
    } catch (error) {
      console.error("Error submitting loan application:", error);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 py-6 flex flex-col justify-center sm:py-12">
      <div className="relative py-3 sm:max-w-xl sm:mx-auto">
        <div className="absolute inset-0 bg-gradient-to-r from-cyan-400 to-light-blue-500 shadow-lg transform -skew-y-6 sm:skew-y-0 sm:-rotate-6 sm:rounded-3xl"></div>
        <div className="relative px-4 py-10 bg-white shadow-lg sm:rounded-3xl sm:p-20">
          <LoanForm
            onFetchBalanceSheet={fetchBalanceSheet}
            onSubmitApplication={submitApplication}
          />
          {balanceSheet.length > 0 && (
            <BalanceSheet balanceSheet={balanceSheet} />
          )}
          {loanResult && <LoanResult result={loanResult} />}
        </div>
      </div>
    </div>
  );
};

export default App;
