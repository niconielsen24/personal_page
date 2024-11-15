export default function Tech({ tech }) {
  return (
    <div className="tech-container p-2 flex flex-col items-center justify-center">
      <div className="tech-image w-16 h-16 flex items-center justify-center">
        {tech}
      </div>
      <div className="tech-name text-center mt-2 text-lg font-semibold">{tech.name}</div>
    </div>
  );
}

