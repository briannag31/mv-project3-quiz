// import { useState, useEffect } from "react"
// import New from "./New"
// import Index from "./Index"

export default function Home ({cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, setViewCards, setViewDecks}){

    return(
        <>
            <h1 style={{color: "white", fontSize: "70px", marginTop: "90px", marginBottom: "50px",textAlign: "center"}}>Welcome! What would you like to do today?</h1>
            {/* <button onClick={()=>setViewDecks(true)}>View All Study Decks</button> */}
            <div style={{justifyContent: "center", display: "flex"}}>
                <button className="home-button" onClick={()=>setViewCards(true)}>View My Study Cards</button>
                <button className="home-button" onClick={()=>setAddCard(true)}>Create A New Study Card</button>
            </div>
           
        </>   
    )
}