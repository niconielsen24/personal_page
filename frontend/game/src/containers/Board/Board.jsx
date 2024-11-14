import { useEffect, useState } from "react"
import { initTicTacToeGame, makeMove } from "../../services/services"
import BoardLayout from "./components/BoardLayout"
import { useGameStore } from "../../services/stores"

export default function Board() {
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(true)
  const gameUUID = useGameStore(state => state.gameUUID)
  const setGameUUID = useGameStore(state => state.setGameUUID)
  const selectedTile = useGameStore(state => state.selectedTile) 
  const boardState = useGameStore(state => state.boardState)
  const setBoardState = useGameStore(state => state.setBoardState)

  useEffect(() => {
    const fetchGame = async () => {
      try {
        const response = await initTicTacToeGame("http://localhost:8000")
        setBoardState(response.Board)
        if(!gameUUID) setGameUUID(response.ID)
      } catch (err) {
        setError(err.message)
      } finally {
        setLoading(false)
      }
    }
    fetchGame()
  }, [])

  useEffect(() => {
    if (!selectedTile) return
    if (!gameUUID) return
    const fetchMove = async () => {
      try {
        const response = await makeMove("http://localhost:8000", selectedTile.pos, gameUUID)
        setBoardState(response.Board)
      } catch (err) {
        console.log(err)
      }
    }
    fetchMove()
 }, [selectedTile, gameUUID])

  if (error) return (<div>{error}</div>)
  if (loading) return (<div>{loading}</div>)

  return (
    <>
      <BoardLayout board={boardState} />
    </>
  )
}
