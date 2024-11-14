export async function tryGoServer() {
  const response = await fetch("http://localhost:8000", {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  });

  if (!response.ok) {
    throw new Error("walter frank ? se rompio xd");
  }

  return response.json();
}

export async function initTicTacToeGame(ip_addr) {
  const Addr = encodeURI(ip_addr);
  const response = await fetch(`${Addr}/init`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    }
  });

  if (!response.ok) {
    throw new Error(`HTTP error! Status: ${response.status} ${response.statusText}`)
  }

  return response.json();
}

export async function killGame(ip_addr, game_id) {
  const Addr = encodeURI(ip_addr)
  const response = await fetch(`${Addr}/killGame`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      game_id: game_id,
    })
  })

  if (!response.ok) {
    throw new Error(`HTTP error! Status: ${response.status} ${response.statusText}`)
  }
}

export async function makeMove(ip_addr, pos, game_uuid) {
  const Addr = encodeURI(ip_addr);
  const response = await fetch(`${Addr}/makeMove`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      game_id: game_uuid,
      position: pos
    })
  })
  if (!response.ok) {
    throw new Error(`HTTP error! Status: ${response.status} ${response.statusText}`)
  }

  return response.json()
}
