import { useRef } from 'react'
import '../../App.css'
import AppLayout from "./components/AppLayout.jsx"
import { useGameStore } from '../../services/stores.js'
import { initTicTacToeGame, killGame } from '../../services/services.js'

function App() {
  const boardState = useGameStore(state => state.boardState)
  const setBoardState = useGameStore(state => state.setBoardState)
  const gameUUID = useGameStore(state => state.gameUUID)
  const setGameUUID = useGameStore(state => state.setGameUUID)
  const setSelectedTile = useGameStore(state => state.setSelectedTile)

  const resetGameState = useRef(() => {
    setBoardState(null)
    setSelectedTile(null)
    setGameUUID(null)
  })

  const handleSubmit = async () => {
    if (boardState) {
      resetGameState.current()
      try {
        await killGame("http://localhost:8000", gameUUID)
        const response = await initTicTacToeGame("http://localhost:8000")
        setGameUUID(response.ID)
        setBoardState(response.Board)
      } catch (err) {
        console.log(err.Payload)
      }
    }
  }

  return (
    <>
      <AppLayout onSubmit={handleSubmit} />
    </>
  )
}
export default App
