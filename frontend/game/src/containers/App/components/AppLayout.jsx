import React, { useState, useEffect } from "react";
import { Button } from "@headlessui/react";
import Board from "../../Board/Board";
import Header from "./Header";
import Project from "../../../components/Project";

const BoardContainerStyle =
  "p-4 grid grid-rows-6 grid-cols-5 justify-items-start bg-custom-navy rounded-lg sm:w-[200px] sm:h-[200px] lg:w-[400px] lg:h-[400px] border border-custom-blue shadow-inset-light";
const HeaderContainerStyle =
  "fixed top-0 left-0 w-full bg-custom-blue text-custom-cyan shadow-md z-50 p-4 font-bold text-xl flex justify-between items-center";
const TicTacToeButtonContainerStyle =
  "p-3 row-span-1 col-span-5 text-custom-navy";

export default function AppLayout({ onSubmit }) {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    fetch('../../../src/data/projects.json') 
      .then((response) => response.json())
      .then((data) => setProjects(data.projects))
      .catch((error) => console.error('Error loading projects:', error));
  }, []);

  return (
    <>
      <div className="flex mt-20 flex-col items-center justify-center">
        <div className={HeaderContainerStyle}>
          <Header />
        </div>

        <div className="">
          {projects.map((project) => (
            <Project
              key={project.id}
              name={project.name}
              description={project.description}
              img={project.image_url}
              technologies={project.technologies}
              repo_url={project.repo_url}
            />
          ))}
        </div>

        <div className={BoardContainerStyle}>
          <div className={TicTacToeButtonContainerStyle}>
            <Button onClick={onSubmit}>Jugar</Button>
          </div>
          <Board />
        </div>
      </div>
    </>
  );
}

