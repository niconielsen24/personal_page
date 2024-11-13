import React from "react";
import { Popover, PopoverButton, PopoverPanel } from "@headlessui/react";
import Board from "../../Board/Board";

export default function AppLayout({ onInput, name }) {
  const PopoverButtonStyle = "px-4 py-2 font-semibold text-white bg-gray-800 rounded-md shadow-lg transition-all hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2";
  const PopoverPanelStyle = "absolute z-10 mt-2 p-4 w-48 bg-gray-900 rounded-lg shadow-lg border border-indigo-500 text-white animate-fade-in transition ease-out duration-300 transform opacity-0 scale-95 neon-shadow"
  const PopoverInputStyle = "w-full px-3 py-2 bg-gray-800 text-white placeholder-gray-400 border border-gray-700 rounded-md shadow-inner focus:outline-none focus:ring-2 focus:ring-indigo-400"

  return (
    <>
      {name &&
        <h1 className="text-3xl font-bold animate-fadeIn" >
          Hola {name} !
        </h1>
      }
      <Popover className="relative">
        <PopoverButton className={PopoverButtonStyle}>
          Jugar
        </PopoverButton>
        <PopoverPanel className={PopoverPanelStyle}>
          <div>
            <input
              placeholder="Enter Name"
              className={PopoverInputStyle}
              onInput={onInput}
            />
          </div>
        </PopoverPanel>
      </Popover>
      <div className="p-4 w-[400px] h-[400px]">
        <Board />
      </div>

    </>
  );
}
