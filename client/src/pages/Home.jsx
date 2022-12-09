// import { useState, useEffect } from "react"
// import New from "./New"
// import Index from "./Index"

export default function Home ({cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, setViewCards, setViewDecks}){

    return(
        <>
            <h1>Welcome, What would you like to do today?</h1>
            <button onClick={()=>setViewDecks(true)}>View All Study Decks</button>
            <button onClick={()=>setViewCards(true)}>View All Cards</button>
            <button onClick={()=>setAddCard(true)}>Create New Deck</button>
        </>   
    )
}