import React, { useRef } from "react";
import Tech from "./Tech";
import { GithubIcon } from "./Icons";
import { TECH_MAP } from "./Icons";

const ProjectContainerStyle =
  "bg-dark-blue text-soft-blue p-6 m-4 rounded-lg shadow-[inset_0_2px_8px_rgba(0,0,0,0.6)] border border-soft-blue";
const ProjectTitleContainer =
  "font-bold text-4xl text-start border-b-4 border-bright-blue mb-4";
const ProjectImgStyle =
  "w-full h-auto object-cover rounded-lg";
const ProjectGridStyle =
  "grid grid-cols-1 lg:grid-cols-2 gap-4 items-center";
const ProjectInfoContainerStyle =
  "flex flex-col justify-center text-start p-4 bg-gray-700 rounded-lg";
const ProjectInfoTextStyle =
  "text-white text-xl leading-relaxed";
const TechAndRepoInfoStyle =
  "mt-2 p-2 items-center"
const TechHeaderStyle =
  "text-2xl text-white font-bold"
const TechContainerStyle =
  "flex flex-row items-center justify-center"
const RepoLinkContainerStyle =
  "flex flex-row items-center justify-start bg-gray-300 p-2 rounded-lg"

export default function Project({ name, img, technologies, description, repo_url }) {
  const imgRef = useRef(img)
  const urlRef = useRef(repo_url)

  const mapTechnologies = (tech_array) => {
    if (!Array.isArray(tech_array)) return

    return tech_array.map((tech, _) => {
      return <Tech tech={TECH_MAP[tech] || tech} />
    })
  }

  return (
    <>
      <div className={ProjectContainerStyle}>
        <h3 className={ProjectTitleContainer}>{name}</h3>
        <div className={ProjectGridStyle}>
          <img className={ProjectImgStyle} ref={imgRef}></img>
          <div className={ProjectInfoContainerStyle}>
            <p className={ProjectInfoTextStyle}>
              {description}
            </p>
            <div className={TechAndRepoInfoStyle}>
              <h4 className={TechHeaderStyle}>Technologies used</h4>
              <div className={TechContainerStyle}>
                {mapTechnologies(technologies)}
              </div>
            </div>
            <div className={RepoLinkContainerStyle}>
              <div>
                <GithubIcon />
              </div>
              <a className="p-2 font-bold" href={urlRef.current} target="blank">{name}</a>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}
