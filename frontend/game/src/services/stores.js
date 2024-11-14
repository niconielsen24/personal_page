import { create } from "zustand"

export const useGameStore = create(set => ({
  gameUUID: null,
  selectedTile: null,
  lastTile: null,
  boardState: null,
  setGameUUID: (new_uuid) => set({ gameUUID: new_uuid }),
  setSelectedTile: (new_tile) => set({ selectedTile: new_tile}),
  setLastTile: (new_last_tile) => set({ lastTile: new_last_tile}),
  setBoardState: (new_board_state) => set({ boardState: new_board_state}),
}))
