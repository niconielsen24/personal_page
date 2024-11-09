import { useState } from 'react'
import '../../App.css'
import AppLayout from "./components/AppLayout.jsx"
import { tryGoServer } from "../../services/services.js"

function App() {
  const [message, setMessage] = useState(null)
  const [name, setName] = useState(null)

  const handleClick = async () => {
    try {
      const response = await tryGoServer();
      console.log(response);
      setMessage(response.message);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <>
      <AppLayout onClick={handleClick} onInput={(e) => setName(e.target.value)} name={name} message={message}/>
    </>
  )
}
export default App
