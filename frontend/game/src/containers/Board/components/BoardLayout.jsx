export default function BoardLayout({ board }) {
  const mapBoard = (board) => {
    return board.tiles.map((arr, i) => (
      <div key={`row-${i}`} className="grid grid-cols-3 gap-1">
        {arr.map((tile, j) => (
          <div
            key={`tile-${i}-${j}`}
            className="row-span-1 col-span-1 flex items-center justify-center border border-gray-300 p-2"
          >
            {String.fromCharCode(tile)}
          </div>
        ))}
      </div>
    ));
  };

  return (
    <div className="grid grid-rows-3 gap-2">
      {mapBoard(board)}
    </div>
  );
}
