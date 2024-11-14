import React from "react";
import { useRef } from "react";

const ProjectContainerStyle = 
  "bg-dark-blue text-soft-blue p-6 m-4 rounded-lg shadow-[inset_0_2px_8px_rgba(0,0,0,0.6)] border border-soft-blue";
const ProjectTitleContainer = 
  "font-bold text-4xl text-start border-b-4 border-bright-blue mb-4";
const ProjectImgStyle = 
  "w-full h-auto object-cover rounded-lg";
const ProjectGridStyle = 
  "grid grid-cols-1 lg:grid-cols-2 gap-4 items-center";
const ProjectInfoContainerStyle = 
  "flex flex-col justify-center text-start p-4";
const ProjectInfoTextStyle = 
  "text-soft-blue leading-relaxed";

export default function Project({name, img, technologies, description}) {
  const imgRef = useRef(img)

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
          </div>
        </div>
      </div>
    </>
  )
}
