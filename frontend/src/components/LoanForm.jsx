import { useState, useEffect } from "react";

const LoanForm = ({ onFetchBalanceSheet, onSubmitApplication, onAccountingProviderChange }) => {
  const [formData, setFormData] = useState({
    businessDetails: {
      name: "",
      yearEstablished: "",
    },
    loanAmount: "",
    accountingProvider: "",
  });

  const [accountingProvider, setAccountingProvider] = useState("");

  useEffect(() => {
    if (accountingProvider) {
      onFetchBalanceSheet(accountingProvider);
    }
  }, [accountingProvider]);

  const handleChange = (event) => {
    const { name, value } = event.target;
    if (name === "loanAmount") {
      setFormData({ ...formData, [name]: parseInt(value, 10) });
    } else if (name === "name" || name === "yearEstablished") {
      setFormData({
        ...formData,
        businessDetails: { ...formData.businessDetails, [name]: value },
      });
    } else {
      setFormData({ ...formData, [name]: value });
      setAccountingProvider(value);
      onAccountingProviderChange(); // Call the new function when the accounting provider is changed
    }
  };

  const handleSubmitApplication = (event) => {
    event.preventDefault();
    onSubmitApplication(formData);
  };

  return (
    <form onSubmit={handleSubmitApplication} className="space-y-4">
      <h2 className="text-2xl font-semibold">Loan Application</h2>
      <div>
        <label htmlFor="name" className="block text-sm font-medium">
          Business Name:
        </label>
        <input
          type="text"
          name="name"
          value={formData.businessDetails.name}
          onChange={handleChange}
          className="mt-1 w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
        />
      </div>
      <div>
        <label htmlFor="yearEstablished" className="block text-sm font-medium">
          Year Established:
        </label>
        <input
          type="number"
          name="yearEstablished"
          value={formData.businessDetails.yearEstablished}
          onChange={handleChange}
          className="mt-1 w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
        />
      </div>
      <div>
        <label htmlFor="loanAmount" className="block text-sm font-medium">
          Loan Amount:
        </label>
        <input
          type="number"
          name="loanAmount"
          value={formData.loanAmount}
          onChange={handleChange}
          className="mt-1 w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
        />
      </div>
      <div>
        <label
          htmlFor="accountingProvider"
          className="block text-sm font-medium"
        >
          Accounting Provider:
        </label>
        <select
          name="accountingProvider"
          value={formData.accountingProvider}
          onChange={handleChange}
          className="mt-1 w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
        >
          <option value="">Select a provider</option>
          <option value="Xero">Xero</option>
          <option value="MYOB">MYOB</option>
        </select>
      </div>
      <button
        type="submit"
        className="w-full py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Submit Application
      </button>
    </form>
  );
};

export default LoanForm;
