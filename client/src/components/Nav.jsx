import React from "react";

export default function Nav({addCard, setAddCard}){
    return(
      <nav className="fixed-top">
         <p className="nav-text" onClick={()=>setAddCard(true)}> Create a new card</p>
    </nav>  
    )
}
