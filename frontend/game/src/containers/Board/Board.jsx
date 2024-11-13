import { useEffect, useState } from "react"
import { initTicTacToeGame } from "../../services/services"
import BoardLayout from "./components/BoardLayout"
export default function Board(){
  const [board,setBoard] = useState([])
  const [error,setError] = useState(null)
  const [loading,setLoading] = useState(true)
  useEffect(()=>{
    const data = async () => {
      try {
        const response = await initTicTacToeGame("http://localhost:8000")
        setBoard(response.Board)
      } catch (err) {
        setError(err.message)
      } finally {
        setLoading(false)
      }
    }
    data()
  },[])
 
  if (error) return (<div>{error}</div>)
  if (loading) return (<div>{loading}</div>)

  return(
    <>
      <BoardLayout board={board}/>
    </>
  )
}
