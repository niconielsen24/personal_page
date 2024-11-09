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
