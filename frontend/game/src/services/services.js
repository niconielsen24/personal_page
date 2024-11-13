export async function tryGoServer(){
  const response = await fetch("http://localhost:8000",{
    method : "GET",
    headers : {
      "Content-Type" : "application/json"
    }
  });

  if (!response.ok){
    throw new Error("walter frank ? se rompio xd");
  }

  return response.json();
}

export async function initTicTacToeGame(ip_addr) {
  const Addr = ip_addr;
  const response = await fetch(`${Addr}/init`,{
    method : "POST",
    headers : {
      "Content-Type" : "application/json"
    }
  });

  if (!response.ok){
    throw new Error(response.status, response.statusText)
  }
  
  return response.json();
}
