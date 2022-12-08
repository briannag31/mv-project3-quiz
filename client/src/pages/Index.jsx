import React from "react";
export default function Index({cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, fetchCards}){
    const url = "http://localhost:3001/api"

    const deleteCard = async (id) =>{
        const res = await fetch(`${url}/deleteCards/${id}`, {
          method: "DELETE",
          header: {
            "Content-Type": "application/x-www-form-urlencoded"  
          },
        })
        const data = await res.json()
        fetchCards()
      }
  return (  <>
    <button onClick={()=>setAddCard(true)}>Add a New Card</button>
  <div>
   { cards.map((individualCard, idx) =>{
    let border = "solid red"
    if( individualCard.known === true){
        border = "solid green"
    }
      return <div style={{border: border}} className={"card"}>
        <h1 key={idx}>{individualCard.term}</h1>
        <h6>{individualCard.collection}</h6>
        <button type="button" className="btn btn-danger"  onClick={()=> { const id = (individualCard._id); deleteCard(id)}}>Delete</button>
      </div> 
    })}</div>
    </>)
}