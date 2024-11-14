import React from "react";

const HeaderStyle = 
  "flex justify-between items-center px-6 py-4 bg-custom-blue text-custom-cyan";
const LogoStyle = 
  "text-2xl font-bold tracking-wide text-custom-yellow";
const NavStyle = 
  "flex gap-6 ml-auto";
const NavLinkStyle = 
  "text-lg font-medium text-custom-cyan hover:text-custom-blue-light transition-colors duration-300";

export default function Header() {
  return (
    <header className={HeaderStyle}>
      <div className={LogoStyle}>Nico Nielsen</div>
      <nav className={NavStyle}>
        <a href="#about" className={NavLinkStyle}>About</a>
        <a href="#projects" className={NavLinkStyle}>Projects</a>
        <a href="#contact" className={NavLinkStyle}>Contact</a>
      </nav>
    </header>
  );
}

