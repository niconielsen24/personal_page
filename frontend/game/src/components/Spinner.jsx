import React from 'react';

const SpinnerStyle = 
  "w-16 h-16 border-4 border-t-bright-blue border-dark-blue rounded-full animate-spin";

export default function Spinner() {
  return (
    <div className="flex items-center justify-center h-full">
      <div className={SpinnerStyle}></div>
    </div>
  );
}

