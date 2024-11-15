import React from 'react';

const ServerDownContainer = 
  "row-span-5 col-span-5 grid grid-rows-3 gap-2 w-full h-full";
const ServerDownMessage = 
  "row-span-1 text-2xl font-bold mb-4 text-bright-blue";

export default function ServerDown({ message }) {
  return (
    <div className={ServerDownContainer}>
      <h2 className={ServerDownMessage}>Server is Down</h2>
      <p>{message}</p>
    </div>
  );
}

