import React from "react";

export default function Nav({addCard, setAddCard, setViewCards}){
    return(
      <nav className="fixed-top">
             <img className="d-flex flex-row  owl-logo" src="https://i.imgur.com/de3nbRq.png" alt="app logo of an owl" onClick={()=>{setAddCard(false); setViewCards(false)}}/>
         <p className="nav-text" onClick={()=>setAddCard(true)}> Create a new card</p>
    </nav>  
    )
}
