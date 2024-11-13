import Tile from "./Tile"

export default function BoardLayout({ board }) {
  const mapBoard = (board) => {
    return board.tiles.map((arr, i) => (
      <div className="grid grid-cols-3 gap-1">
        {arr.map((tile, j) => (
          <Tile indexI={i} indexJ={j} tile={tile}/> 
       ))}
      </div>
    ));
  };

  return (
    <div className="grid grid-rows-3 gap-2 w-full h-full">
      {mapBoard(board)}
    </div>
  );
}
