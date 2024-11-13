import { useCallback } from "react"
import { useGameStore } from "../../../services/stores"

export default function Tile({ indexI, indexJ, tile }) {
  const setSelectedTile = useGameStore(state => state.setSelectedTile)
  const handleClick = useCallback(() => {
    const t = { 
      player : tile,
      pos : { x : indexI, y : indexJ}
    }
    setSelectedTile(t)
  },[])

  return (
    <>
      <div
        onClick={handleClick}
        key={`tile-${indexI}-${indexJ}`}
        className="
          row-span-1 col-span-1 flex items-center justify-center w-full h-full border border-gray-300 p-2
          "
      >
        {String.fromCharCode(tile)}
      </div>
    </>
  )
}
