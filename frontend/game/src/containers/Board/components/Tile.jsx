import { useCallback } from "react"
import { useGameStore } from "../../../services/stores"

const TileStyle =
  "row-span-1 col-span-1 flex items-center justify-center w-full h-full border border-gray-700 p-2 bg-gray-800 shadow-lg rounded-lg hover:border-indigo-400 transition-all duration-300";
const OplayerTileStyle =
  "text-blue-500 font-bold md:text-3xl lg:text-5xl animate-fadeIn";
const XplayerTileStyle =
  "text-red-500 font-bold md:text-3xl lg:text-5xl animate-fadeIn";

const STYLE_MAP = Object.freeze({
  X: XplayerTileStyle,
  O: OplayerTileStyle,
  Empty: "",
})

export default function Tile({ indexI, indexJ, tile }) {
  const setSelectedTile = useGameStore(state => state.setSelectedTile)
  const handleClick = useCallback(() => {
    const t = {
      player: tile,
      pos: { x: indexI, y: indexJ }
    }
    setSelectedTile(t)
  }, [])

  const char = String.fromCharCode(tile)
  const mapTileColor = (t) => STYLE_MAP[t] || ""

  return (
    <>
      <div
        onClick={handleClick}
        key={`tile-${indexI}-${indexJ}`}
        className={`${TileStyle}`}
      >
        <p className={`${mapTileColor(char)}`}>{String.fromCharCode(tile)}</p>
      </div>
    </>
  )
}
