import { useState } from "react"
import { client } from "../libs/axios"

export const Home = () => {
  const [response, setResponse] = useState("")
  const get = async () => {
    client.get("/articles").then((response) => {
      setResponse(response.data)
    })
  }

  return (
    <div className="mx-4 my-2">
      <div className="text-3xl">Home page.</div>
      <hr className="my-2" />
      <button onClick={get} className="border border-black px-1">
        Get
      </button>
      <div>
        Response:<br />
        {response}
      </div>
    </div>
  )
}