import Tile from "./Tile"

const BoardContainerStyle = "row-span-5 col-span-5 grid grid-rows-3 gap-2 w-full h-full"
const BoardRowStyle = "grid grid-cols-3 gap-1"

export default function BoardLayout({ board }) {
  const mapBoard = (board) => {
    if (!board) return
    return board.tiles.map((arr, i) => (
      <div className={BoardRowStyle}>
        {arr.map((tile, j) => (
          <Tile indexI={i} indexJ={j} tile={tile}/> 
       ))}
      </div>
    ));
  };

  return (
    <div className={BoardContainerStyle}>
      {mapBoard(board)}
    </div>
  );
}
